package parameter_column_index

import (
	"context"
	"fmt"
	"time"

	"github.com/adiatma85/new-go-template/src/business/entity"
	"github.com/adiatma85/own-go-sdk/errors"
	"github.com/adiatma85/own-go-sdk/log"
	"github.com/adiatma85/own-go-sdk/parser"
	"github.com/adiatma85/own-go-sdk/redis"
	"github.com/adiatma85/own-go-sdk/sql"
)

const (
	cacheExpirationTime time.Duration = 24 * 30 * time.Hour
)

type Interface interface {
	Get(ctx context.Context, param entity.ParameterColumnIndexParam) (entity.ParameterColumnIndex, error)
	GetList(ctx context.Context, param entity.ParameterColumnIndexParam) ([]entity.ParameterColumnIndex, *entity.Pagination, error)
	Create(ctx context.Context, inputParam entity.ParameterColumnIndexInputParam) (entity.ParameterColumnIndex, error)
	Update(ctx context.Context, updateParam entity.ParameterColumnIndexUpdateParam, selectParam entity.ParameterColumnIndexParam) error
}

type InitParam struct {
	Log   log.Interface
	Db    sql.Interface
	Redis redis.Interface
	Json  parser.JSONInterface
}

type parameterColumnIndex struct {
	log   log.Interface
	db    sql.Interface
	redis redis.Interface
	json  parser.JSONInterface
}

func Init(param InitParam) Interface {
	return &parameterColumnIndex{
		log:   param.Log,
		db:    param.Db,
		redis: param.Redis,
		json:  param.Json,
	}
}

func (p *parameterColumnIndex) Get(ctx context.Context, param entity.ParameterColumnIndexParam) (entity.ParameterColumnIndex, error) {
	parameterColumnIndex := entity.ParameterColumnIndex{}

	marshalledParam, err := p.json.Marshal(param)
	if err != nil {
		return parameterColumnIndex, err
	}

	if !param.BypassCache {
		parameterColumnIndex, err = p.getCache(ctx, fmt.Sprintf(getParameterColumnIndexByKey, string(marshalledParam)))
		switch {
		case errors.Is(err, redis.Nil):
			p.log.Warn(ctx, fmt.Sprintf(entity.ErrorRedisNil, err.Error()))
		case err != nil:
			p.log.Warn(ctx, fmt.Sprintf(entity.ErrorRedis, err.Error()))
		default:
			return parameterColumnIndex, nil
		}
	}

	parameterColumnIndex, err = p.getSQL(ctx, param)
	if err != nil {
		return parameterColumnIndex, err
	}

	err = p.upsertCache(ctx, fmt.Sprintf(getParameterColumnIndexByKey, string(marshalledParam)), parameterColumnIndex, cacheExpirationTime)
	if err != nil {
		p.log.Error(ctx, fmt.Sprintf(entity.ErrorRedis, err.Error()))
	}

	return parameterColumnIndex, nil
}

func (p *parameterColumnIndex) GetList(ctx context.Context, param entity.ParameterColumnIndexParam) ([]entity.ParameterColumnIndex, *entity.Pagination, error) {
	if !param.BypassCache {
		parameterColumnIndexList, pg, err := p.getCacheList(ctx, param)
		switch {
		case errors.Is(err, redis.Nil):
			p.log.Warn(ctx, fmt.Sprintf(entity.ErrorRedisNil, err.Error()))
		case err != nil:
			p.log.Warn(ctx, fmt.Sprintf(entity.ErrorRedis, err.Error()))
		default:
			return parameterColumnIndexList, &pg, nil
		}
	}

	parameterColumnIndexList, pg, err := p.getListSQL(ctx, param)
	if err != nil {
		return parameterColumnIndexList, pg, err
	}

	err = p.upsertCacheList(ctx, param, parameterColumnIndexList, *pg, cacheExpirationTime)
	if err != nil {
		p.log.Error(ctx, fmt.Sprintf(entity.ErrorRedis, err.Error()))
	}

	return parameterColumnIndexList, pg, nil
}

func (p *parameterColumnIndex) Create(ctx context.Context, inputParam entity.ParameterColumnIndexInputParam) (entity.ParameterColumnIndex, error) {
	var err error

	parameterColumnIndex, err := p.createSQL(ctx, inputParam)
	if err != nil {
		return parameterColumnIndex, err
	}

	err = p.deleteCache(ctx)
	if err != nil {
		p.log.Error(ctx, fmt.Sprintf(entity.ErrorRedis, err.Error()))
	}

	return parameterColumnIndex, nil
}

func (p *parameterColumnIndex) Update(ctx context.Context, updateParam entity.ParameterColumnIndexUpdateParam, selectParam entity.ParameterColumnIndexParam) error {
	var err error

	err = p.updateSQL(ctx, updateParam, selectParam)
	if err != nil {
		return err
	}

	err = p.deleteCache(ctx)
	if err != nil {
		p.log.Error(ctx, fmt.Sprintf(entity.ErrorRedis, err.Error()))
	}

	return nil
}
