package parameter_column_index

import (
	"context"
	"fmt"
	"time"

	"github.com/adiatma85/new-go-template/src/business/entity"
	"github.com/adiatma85/own-go-sdk/codes"
	"github.com/adiatma85/own-go-sdk/errors"
)

const (
	getParameterColumnIndexByKey           = "aquahero:parameter_column_index:get:%s"
	getParameterColumnIndexByQueryKey      = "aquahero:parameter_column_index:get:q:%s"
	getParameterColumnIndexByPaginationKey = "aquahero:parameter_column_index:get:p:%s"
	deleteParameterColumnIndexKeysPattern  = "aquahero:parameter_column_index*"
)

func (p *parameterColumnIndex) upsertCache(ctx context.Context, key string, parameterIndex entity.ParameterColumnIndex, ttl time.Duration) error {
	marshalledParameterIndex, err := p.json.Marshal(parameterIndex)
	if err != nil {
		return errors.NewWithCode(codes.CodeCacheMarshal, err.Error())
	}

	err = p.redis.SetEX(ctx, key, string(marshalledParameterIndex), ttl)
	if err != nil {
		return errors.NewWithCode(codes.CodeCacheSetSimpleKey, err.Error())
	}

	return nil
}

func (p *parameterColumnIndex) getCache(ctx context.Context, key string) (entity.ParameterColumnIndex, error) {
	parameterIndex := entity.ParameterColumnIndex{}

	marshalledParameterIndex, err := p.redis.Get(ctx, key)
	if err != nil {
		return parameterIndex, err
	}

	err = p.json.Unmarshal([]byte(marshalledParameterIndex), &parameterIndex)
	if err != nil {
		return parameterIndex, errors.NewWithCode(codes.CodeCacheUnmarshal, err.Error())
	}

	return parameterIndex, nil
}

func (p *parameterColumnIndex) upsertCacheList(ctx context.Context, param entity.ParameterColumnIndexParam, parameterIndexes []entity.ParameterColumnIndex, pg entity.Pagination, ttl time.Duration) error {
	keyValue, err := p.json.Marshal(param)
	if err != nil {
		return errors.NewWithCode(codes.CodeCacheMarshal, err.Error())
	}

	// Set parameter column indexes to cache
	marshalledParameterIndexes, err := p.json.Marshal(parameterIndexes)
	if err != nil {
		return errors.NewWithCode(codes.CodeCacheMarshal, err.Error())
	}
	err = p.redis.SetEX(ctx, fmt.Sprintf(getParameterColumnIndexByQueryKey, string(keyValue)), string(marshalledParameterIndexes), ttl)
	if err != nil {
		return errors.NewWithCode(codes.CodeCacheSetSimpleKey, err.Error())
	}

	// Set pagination to cache
	marshalledPagination, err := p.json.Marshal(pg)
	if err != nil {
		return errors.NewWithCode(codes.CodeCacheMarshal, err.Error())
	}

	err = p.redis.SetEX(ctx, fmt.Sprintf(getParameterColumnIndexByPaginationKey, string(keyValue)), string(marshalledPagination), ttl)
	if err != nil {
		return errors.NewWithCode(codes.CodeCacheSetSimpleKey, err.Error())
	}

	return nil
}

func (p *parameterColumnIndex) getCacheList(ctx context.Context, param entity.ParameterColumnIndexParam) ([]entity.ParameterColumnIndex, entity.Pagination, error) {
	var (
		parameterIndexes = []entity.ParameterColumnIndex{}
		pg               = entity.Pagination{}
	)

	keyValue, err := p.json.Marshal(param)
	if err != nil {
		return parameterIndexes, pg, errors.NewWithCode(codes.CodeCacheMarshal, err.Error())
	}

	// Get parameter indexes from redis
	marshalledParameterIndexes, err := p.redis.Get(ctx, fmt.Sprintf(getParameterColumnIndexByQueryKey, string(keyValue)))
	if err != nil {
		return parameterIndexes, pg, err
	}

	err = p.json.Unmarshal([]byte(marshalledParameterIndexes), &parameterIndexes)
	if err != nil {
		return parameterIndexes, pg, errors.NewWithCode(codes.CodeCacheUnmarshal, err.Error())
	}

	// Get pagination from redis
	marshalledPagination, err := p.redis.Get(ctx, fmt.Sprintf(getParameterColumnIndexByPaginationKey, string(keyValue)))
	if err != nil {
		return parameterIndexes, pg, err
	}

	err = p.json.Unmarshal([]byte(marshalledPagination), &pg)
	if err != nil {
		return parameterIndexes, pg, errors.NewWithCode(codes.CodeCacheUnmarshal, err.Error())
	}

	return parameterIndexes, pg, nil
}

func (p *parameterColumnIndex) deleteCache(ctx context.Context) error {
	err := p.redis.Del(ctx, deleteParameterColumnIndexKeysPattern)
	if err != nil {
		return err
	}

	return nil
}
