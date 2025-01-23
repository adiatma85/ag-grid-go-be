package pond

import (
	"context"
	"fmt"
	"time"

	"github.com/adiatma85/new-go-template/src/business/entity"
	"github.com/adiatma85/own-go-sdk/codes"
	"github.com/adiatma85/own-go-sdk/errors"
)

const (
	getPondByIdKey              = `aquahero:pond:get:%s`
	getPondByQueryKey           = `aquahero:pond:get:q:%s`
	getPondByPaginationQueryKey = `aquahero:pond:get:pq:%s`
	getPondListByMapQueryKey    = `aquahero:pond:get:mq:%s`
	deletePondPattern           = `aquahero:pond:*`
	deleteModulePattern         = `aquahero:module:*`
)

func (p *pond) getCacheList(ctx context.Context, params entity.PondParam) ([]entity.Pond, map[int64]entity.Pond, entity.Pagination, error) {
	var (
		result     []entity.Pond
		resultMap  map[int64]entity.Pond
		pagination entity.Pagination
	)

	key, err := p.json.Marshal(params)
	if err != nil {
		return result, resultMap, pagination, errors.NewWithCode(codes.CodeCacheMarshal, err.Error())
	}

	pondsRedis, err := p.redis.Get(ctx, fmt.Sprintf(getPondByQueryKey, string(key)))
	if err != nil {
		return result, resultMap, pagination, err
	}

	data := []byte(pondsRedis)
	if err := p.json.Unmarshal(data, &result); err != nil {
		return result, resultMap, pagination, errors.NewWithCode(codes.CodeCacheUnmarshal, err.Error())
	}

	mapPondRedis, err := p.redis.Get(ctx, fmt.Sprintf(getPondListByMapQueryKey, string(key)))
	if err != nil {
		return result, resultMap, pagination, err
	}

	dataMapPond := []byte(mapPondRedis)
	if err := p.json.Unmarshal(dataMapPond, &resultMap); err != nil {
		return result, resultMap, pagination, err
	}

	paginationRedis, err := p.redis.Get(ctx, fmt.Sprintf(getPondByPaginationQueryKey, string(key)))
	if err != nil {
		return result, resultMap, pagination, err
	}

	dataPagination := []byte(paginationRedis)
	if err := p.json.Unmarshal(dataPagination, &pagination); err != nil {
		return result, resultMap, pagination, errors.NewWithCode(codes.CodeCacheUnmarshal, err.Error())
	}

	return result, resultMap, pagination, nil
}

func (p *pond) upsertCacheList(ctx context.Context, params entity.PondParam, value []entity.Pond, valueMap map[int64]entity.Pond, valuePagination entity.Pagination, expTime time.Duration) error {
	key, err := p.json.Marshal(params)
	if err != nil {
		return errors.NewWithCode(codes.CodeCacheMarshal, err.Error())
	}

	ponds, err := p.json.Marshal(value)
	if err != nil {
		return errors.NewWithCode(codes.CodeCacheMarshal, err.Error())
	}

	if err := p.redis.SetEX(ctx, fmt.Sprintf(getPondByQueryKey, string(key)), string(ponds), expTime); err != nil {
		return errors.NewWithCode(codes.CodeCacheSetSimpleKey, err.Error())
	}

	mapPonds, err := p.json.Marshal(valueMap)
	if err != nil {
		return errors.NewWithCode(codes.CodeCacheMarshal, err.Error())
	}

	if err := p.redis.SetEX(ctx, fmt.Sprintf(getPondListByMapQueryKey, string(key)), string(mapPonds), expTime); err != nil {
		return errors.NewWithCode(codes.CodeCacheSetSimpleKey, err.Error())
	}

	pagination, err := p.json.Marshal(valuePagination)
	if err != nil {
		return errors.NewWithCode(codes.CodeCacheMarshal, err.Error())
	}

	if err := p.redis.SetEX(ctx, fmt.Sprintf(getPondByPaginationQueryKey, string(key)), string(pagination), expTime); err != nil {
		return errors.NewWithCode(codes.CodeCacheSetSimpleKey, err.Error())
	}

	return nil
}
