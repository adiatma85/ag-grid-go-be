package parameter

import (
	"context"
	"fmt"
	"time"

	"github.com/adiatma85/new-go-template/src/business/entity"
	"github.com/adiatma85/own-go-sdk/codes"
	"github.com/adiatma85/own-go-sdk/errors"
)

const (
	getParameterByIdKey              = `aquahero:parameter:get:%s`
	getParameterByQueryKey           = `aquahero:parameter:get:q:%s`
	getParameterByPaginationQueryKey = `aquahero:parameter:get:pq:%s`
	getParameterByMapQueryKey        = `aquahero:parameter:get:mq:%s`
	getParameterByMapQueryNameKey    = `aquahero:parameter:get:mqn:%s` // map parameter by parameter_name. Will be used in the Store Pond Metric logic for parsing
	deleteParameterKeysPattern       = `aquahero:parameter*`
)

func (f *parameter) upsertCacheByID(ctx context.Context, marshalledParams []byte, parameter entity.Parameter, expTime time.Duration) error {
	parameterByIDKey := fmt.Sprintf(getParameterByIdKey, marshalledParams)

	rawJSON, err := f.json.Marshal(parameter)
	if err != nil {
		return errors.NewWithCode(codes.CodeCacheMarshal, "upsertCacheByID")
	}

	return f.redis.SetEX(ctx, parameterByIDKey, string(rawJSON), expTime)
}

func (f *parameter) getCacheByID(ctx context.Context, marshalledParams []byte) (entity.Parameter, error) {
	var result entity.Parameter
	parameterByIDKey := fmt.Sprintf(getParameterByIdKey, marshalledParams)
	parameter, err := f.redis.Get(ctx, parameterByIDKey)
	if err != nil {
		return result, err
	}

	data := []byte(parameter)
	if err := f.json.Unmarshal(data, &result); err != nil {
		return result, errors.NewWithCode(codes.CodeCacheUnmarshal, "getCacheByID")
	}

	return result, nil
}

func (f *parameter) upsertCacheByQuery(ctx context.Context, params entity.ParameterParam, parameters []entity.Parameter, pagination entity.Pagination, parameterMap map[int64]entity.Parameter, parameterMapByName map[string]entity.Parameter, expTime time.Duration) error {
	rawKey, err := f.json.Marshal(params)
	if err != nil {
		return errors.NewWithCode(codes.CodeCacheMarshal, "upsertCacheByQuery")
	}

	// build key
	key := fmt.Sprintf(getParameterByQueryKey, string(rawKey))
	paginationKey := fmt.Sprintf(getParameterByPaginationQueryKey, string(rawKey))
	mapKey := fmt.Sprintf(getParameterByMapQueryKey, string(rawKey))
	mapParameterNameKey := fmt.Sprintf(getParameterByMapQueryNameKey, string(rawKey))

	rawJSON, err := f.json.Marshal(parameters)
	if err != nil {
		return errors.NewWithCode(codes.CodeCacheMarshal, "upsertCacheByQuery")
	}

	return f.upsertCache(ctx, rawJSON, key, pagination, paginationKey, parameterMap, mapKey, parameterMapByName, mapParameterNameKey, expTime, "upsertCacheByQuery")
}

func (f *parameter) getCacheByQuery(ctx context.Context, params entity.ParameterParam) ([]entity.Parameter, map[int64]entity.Parameter, map[string]entity.Parameter, entity.Pagination, error) {
	var (
		results         []entity.Parameter
		resultMap       map[int64]entity.Parameter
		resultMapByName map[string]entity.Parameter
		pagination      entity.Pagination
	)

	// serialize query config bin param to string
	rawKey, err := f.json.Marshal(params)
	if err != nil {
		return results, resultMap, resultMapByName, pagination, errors.NewWithCode(codes.CodeCacheMarshal, "getCacheByQuery")
	}

	// build key
	key := fmt.Sprintf(getParameterByQueryKey, string(rawKey))
	paginationKey := fmt.Sprintf(getParameterByPaginationQueryKey, string(rawKey))
	mapKey := fmt.Sprintf(getParameterByMapQueryKey, string(rawKey))
	mapParameterNameKey := fmt.Sprintf(getParameterByMapQueryNameKey, string(rawKey))

	decJSON, decMapJSONm, decMapJSONName, pagination, err := f.getCache(ctx, key, mapKey, mapParameterNameKey, paginationKey, "getCacheByQuery")
	if err != nil {
		return results, resultMap, resultMapByName, pagination, err
	}

	// unmarshaling returned byte
	if err := f.json.Unmarshal(decJSON, &results); err != nil {
		return results, resultMap, resultMapByName, pagination, errors.NewWithCode(codes.CodeCacheUnmarshal, "getCacheByQuery")
	}

	// unmarshaling returned map byte
	if err := f.json.Unmarshal(decMapJSONm, &resultMap); err != nil {
		return results, resultMap, resultMapByName, pagination, errors.NewWithCode(codes.CodeCacheUnmarshal, "getCacheByQuery")
	}

	// unmarshalling returned map parameter name byte
	if err := f.json.Unmarshal(decMapJSONName, &resultMapByName); err != nil {
		return results, resultMap, resultMapByName, pagination, errors.NewWithCode(codes.CodeCacheUnmarshal, "getCacheByQuery")
	}

	return results, resultMap, resultMapByName, pagination, nil
}

func (f *parameter) upsertCache(ctx context.Context, rawJSON []byte, key string, pagination entity.Pagination, paginationKey string, parameterMap map[int64]entity.Parameter, mapKey string, parameterMapByName map[string]entity.Parameter, mapByNameKey string, expTime time.Duration, funcName string) error {
	// set key expiration
	if err := f.redis.SetEX(ctx, key, string(rawJSON), expTime); err != nil {
		return errors.NewWithCode(codes.CodeCacheSetSimpleKey, funcName)
	}

	// Pagination
	rawJSON, err := f.json.Marshal(pagination)
	if err != nil {
		return errors.NewWithCode(codes.CodeCacheMarshal, funcName)
	}

	if err = f.redis.SetEX(ctx, paginationKey, string(rawJSON), expTime); err != nil {
		return errors.NewWithCode(codes.CodeCacheSetSimpleKey, funcName)
	}

	// Map by ID
	rawJSON, err = f.json.Marshal(parameterMap)
	if err != nil {
		return errors.NewWithCode(codes.CodeCacheMarshal, funcName)
	}

	if err = f.redis.SetEX(ctx, mapKey, string(rawJSON), expTime); err != nil {
		return errors.NewWithCode(codes.CodeCacheSetSimpleKey, funcName)
	}

	// Map By Parameter Name
	rawJSON, err = f.json.Marshal(parameterMapByName)
	if err != nil {
		return errors.NewWithCode(codes.CodeCacheMarshal, funcName)
	}

	if err = f.redis.SetEX(ctx, mapByNameKey, string(rawJSON), expTime); err != nil {
		return errors.NewWithCode(codes.CodeCacheSetSimpleKey, funcName)
	}

	return nil
}

func (f *parameter) getCache(ctx context.Context, key, mapKey, mapByNameKey, paginationKey, funcName string) ([]byte, []byte, []byte, entity.Pagination, error) {
	var (
		pagination entity.Pagination
		decJSON    []byte
	)
	// fetch parameters
	resultRaw, err := f.redis.Get(ctx, key)
	if err != nil {
		return decJSON, decJSON, decJSON, pagination, err
	}

	// fetch map
	mapRaw, err := f.redis.Get(ctx, mapKey)
	if err != nil {
		return decJSON, decJSON, decJSON, pagination, err
	}

	// fetch map by parameter name
	mapByNameRaw, err := f.redis.Get(ctx, mapByNameKey)
	if err != nil {
		return decJSON, decJSON, decJSON, pagination, err
	}

	// fetch pagination
	paginationRaw, err := f.redis.Get(ctx, paginationKey)
	if err != nil {
		return decJSON, decJSON, decJSON, pagination, err
	}

	// unmarshaling returned byte
	if err := f.json.Unmarshal([]byte(paginationRaw), &pagination); err != nil {
		return decJSON, decJSON, decJSON, pagination, errors.NewWithCode(codes.CodeCacheUnmarshal, funcName)
	}

	return []byte(resultRaw), []byte(mapRaw), []byte(mapByNameRaw), pagination, nil
}

func (f *parameter) deleteParameterCache(ctx context.Context) error {
	if err := f.redis.Del(ctx, deleteParameterKeysPattern); err != nil {
		return errors.NewWithCode(codes.CodeCacheDeleteSimpleKey, "delete cache by keys pattern")
	}
	return nil
}
