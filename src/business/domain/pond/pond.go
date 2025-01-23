package pond

import (
	"context"
	"fmt"
	"time"

	"github.com/adiatma85/new-go-template/src/business/entity"

	"github.com/adiatma85/own-go-sdk/codes"
	"github.com/adiatma85/own-go-sdk/errors"
	"github.com/adiatma85/own-go-sdk/log"
	"github.com/adiatma85/own-go-sdk/parser"
	"github.com/adiatma85/own-go-sdk/query"
	"github.com/adiatma85/own-go-sdk/redis"
	"github.com/adiatma85/own-go-sdk/sql"
	"github.com/adiatma85/own-go-sdk/timelib"
)

const (
	cacheExpirationTime = 30 * time.Minute
)

type Interface interface {
	GetList(ctx context.Context, params entity.PondParam) ([]entity.Pond, map[int64]entity.Pond, *entity.Pagination, error)
}

type InitParam struct {
	Log     log.Interface
	Db      sql.Interface
	Redis   redis.Interface
	Json    parser.JSONInterface
	Timelib timelib.Interface
}

type pond struct {
	log     log.Interface
	db      sql.Interface
	redis   redis.Interface
	json    parser.JSONInterface
	timelib timelib.Interface
}

func Init(param InitParam) Interface {
	pm := &pond{
		log:     param.Log,
		db:      param.Db,
		redis:   param.Redis,
		json:    param.Json,
		timelib: param.Timelib,
	}

	return pm
}

func (p *pond) GetList(ctx context.Context, params entity.PondParam) ([]entity.Pond, map[int64]entity.Pond, *entity.Pagination, error) {
	p.log.Debug(ctx, fmt.Sprintf("getting pond list with params: %+v", params))
	ponds := []entity.Pond{}
	mapPonds := make(map[int64]entity.Pond)

	if !params.BypassCache {
		cacheResult, cacheMap, cachePagination, err := p.getCacheList(ctx, params)
		switch {
		case errors.Is(err, redis.Nil):
			p.log.Info(ctx, fmt.Sprintf(entity.ErrorRedisNil, err.Error()))
		case err != nil:
			p.log.Error(ctx, fmt.Sprintf(entity.ErrorRedis, err.Error()))
		default:
			return cacheResult, cacheMap, &cachePagination, nil
		}
	}

	qb := query.NewSQLQueryBuilder(p.db, "param", "db", &params.QueryOption).AddAliasPrefix("p", &params)

	queryExt, queryArgs, countExt, countArgs, err := qb.Build(&params)
	if err != nil {
		return ponds, mapPonds, nil, errors.NewWithCode(codes.CodeSQLBuilder, err.Error())
	}

	rows, err := p.db.Follower().Query(ctx, "rPondList", readPondList+queryExt, queryArgs...)
	if err != nil && !errors.Is(err, sql.ErrNotFound) {
		return ponds, mapPonds, nil, errors.NewWithCode(codes.CodeSQLRead, err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		pond := entity.Pond{}
		if err := rows.StructScan(&pond); err != nil {
			p.log.Error(ctx, errors.NewWithCode(codes.CodeSQLRowScan, err.Error()))
			continue
		}
		pond = p.calculateDoc(pond)
		pond.Category = pond.GetPondType()
		ponds = append(ponds, pond)
		mapPonds[pond.ID] = pond
	}

	pg := entity.Pagination{
		CurrentPage:     params.Page,
		CurrentElements: int64(len(ponds)),
	}
	if len(ponds) > 0 && !params.QueryOption.DisableLimit && params.IncludePagination {
		if err := p.db.Follower().Get(ctx, "cPondList", readPondCount+countExt, &pg.TotalElements, countArgs...); err != nil {
			return ponds, mapPonds, nil, errors.NewWithCode(codes.CodeSQLRead, err.Error())
		}
	}

	pg.ProcessPagination(params.Limit)

	if err := p.upsertCacheList(ctx, params, ponds, mapPonds, pg, cacheExpirationTime); err != nil {
		p.log.Error(ctx, err)
	}

	return ponds, mapPonds, &pg, nil
}

func (p *pond) calculateDoc(pond entity.Pond) entity.Pond {
	if !pond.LatestMetricDate.IsNullOrZero() {
		pond.CurrentDOC = int64(p.timelib.SubstractTime(pond.LatestMetricDate.Time, pond.CurrentCycleDocStartDate.Time).Hours() / 24)
	}
	return pond
}
