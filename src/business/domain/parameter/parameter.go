package parameter

import (
	"context"
	"fmt"
	"time"

	"github.com/adiatma85/new-go-template/src/business/entity"

	"github.com/adiatma85/own-go-sdk/appcontext"
	"github.com/adiatma85/own-go-sdk/codes"
	"github.com/adiatma85/own-go-sdk/errors"
	"github.com/adiatma85/own-go-sdk/log"
	"github.com/adiatma85/own-go-sdk/parser"
	"github.com/adiatma85/own-go-sdk/redis"
	"github.com/adiatma85/own-go-sdk/sql"
)

const (
	cacheExpirationTime = 24 * 30 * time.Hour
)

type Interface interface {
	Get(ctx context.Context, params entity.ParameterParam) (entity.Parameter, error)
	GetList(ctx context.Context, params entity.ParameterParam) ([]entity.Parameter, map[int64]entity.Parameter, map[string]entity.Parameter, *entity.Pagination, error)
	Create(ctx context.Context, parameter entity.Parameter) (entity.Parameter, error)
	Update(ctx context.Context, updateParam entity.ParameterUpdateParam, selectParam entity.ParameterParam) error
}

type InitParam struct {
	Log   log.Interface
	Db    sql.Interface
	Redis redis.Interface
	Json  parser.JSONInterface
}

type parameter struct {
	log   log.Interface
	db    sql.Interface
	redis redis.Interface
	json  parser.JSONInterface
}

func Init(param InitParam) Interface {
	return &parameter{
		log:   param.Log,
		db:    param.Db,
		redis: param.Redis,
		json:  param.Json,
	}
}

func (f *parameter) Get(ctx context.Context, params entity.ParameterParam) (entity.Parameter, error) {
	f.log.Debug(ctx, fmt.Sprintf("getting parameter with params: %v", params))
	var parameter entity.Parameter

	marshalledParams, err := f.json.Marshal(params)
	if err != nil {
		f.log.Error(ctx, errors.NewWithCode(codes.CodeCacheMarshal, err.Error()))
	}

	if !params.BypassCache {
		cacheResult, err := f.getCacheByID(ctx, marshalledParams)
		switch {
		case errors.Is(err, redis.Nil):
			f.log.Info(ctx, fmt.Sprintf(entity.ErrorRedisNil, err.Error()))
		case err != nil:
			f.log.Error(ctx, fmt.Sprintf(entity.ErrorRedis, err.Error()))
		default:
			selectedLang := appcontext.GetAcceptLanguage(ctx)
			cacheResult, _ = f.languageAndParameterInput(cacheResult, selectedLang)
			return cacheResult, nil
		}
	}

	parameter, err = f.getSQL(ctx, params)
	if err != nil {
		return parameter, err
	}

	if err = f.upsertCacheByID(ctx, marshalledParams, parameter, cacheExpirationTime); err != nil {
		f.log.Error(ctx, err)
	}

	return parameter, nil
}

func (f *parameter) Create(ctx context.Context, parameter entity.Parameter) (entity.Parameter, error) {
	return f.createSQL(ctx, parameter)
}

func (f *parameter) Update(ctx context.Context, updateParam entity.ParameterUpdateParam, selectParam entity.ParameterParam) error {
	return f.updateSQL(ctx, updateParam, selectParam)
}

func (f *parameter) GetList(ctx context.Context, params entity.ParameterParam) ([]entity.Parameter, map[int64]entity.Parameter, map[string]entity.Parameter, *entity.Pagination, error) {
	if !params.BypassCache {
		cacheResult, mapCache, mapByParameterNameCache, paginationCahce, err := f.getCacheByQuery(ctx, params)
		switch {
		case errors.Is(err, redis.Nil):
			f.log.Info(ctx, fmt.Sprintf(entity.ErrorRedisNil, err.Error()))
		case err != nil:
			f.log.Error(ctx, fmt.Sprintf(entity.ErrorRedis, err.Error()))
		default:
			cacheResult, mapCache = f.translateParametersData(ctx, cacheResult, mapCache)
			return cacheResult, mapCache, mapByParameterNameCache, &paginationCahce, nil
		}
	}

	result, resultMap, resultMapByName, pagination, err := f.getListSQL(ctx, params)
	if err != nil {
		return result, resultMap, resultMapByName, pagination, err
	}

	for i := range result {
		if result[i].Meta.Valid {
			metaParameter := entity.MetaParameter{}
			err := f.json.Unmarshal([]byte(result[i].Meta.String), &metaParameter)
			if err != nil {
				return result, resultMap, resultMapByName, pagination, err
			}
			validation := entity.ValidationParameter{
				IsRequired:       metaParameter.IsRequired,
				ErrorMessage:     metaParameter.ErrorMessage,
				ErrorDescription: metaParameter.ErrorDescription,
			}
			result[i].Validation = validation
			result[i].Autofill = metaParameter.IsAutofill
			result[i].Tooltip = metaParameter.Tooltip
			resultMap[result[i].ID] = result[i]
		}
	}

	if err = f.upsertCacheByQuery(ctx, params, result, *pagination, resultMap, resultMapByName, cacheExpirationTime); err != nil {
		f.log.Error(ctx, err)
	}

	return result, resultMap, resultMapByName, pagination, nil
}

func (f *parameter) translateParametersData(ctx context.Context, parameterList []entity.Parameter, paramMap map[int64]entity.Parameter) ([]entity.Parameter, map[int64]entity.Parameter) {

	selectedLang := appcontext.GetAcceptLanguage(ctx)

	var err error
	for i := range parameterList {
		parameterList[i], err = f.languageAndParameterInput(parameterList[i], selectedLang)
		if err != nil {
			f.log.Error(ctx, err)
			continue
		}
	}

	for i := range paramMap {
		paramMap[i], err = f.languageAndParameterInput(paramMap[i], selectedLang)
		if err != nil {
			f.log.Error(ctx, err)
			continue
		}
	}

	return parameterList, paramMap
}
