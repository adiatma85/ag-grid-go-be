package aggrid

import (
	"context"
	"fmt"

	"github.com/adiatma85/new-go-template/src/business/entity"
	"github.com/adiatma85/own-go-sdk/log"
	"github.com/adiatma85/own-go-sdk/mongo"
	"github.com/adiatma85/own-go-sdk/parser"
	"github.com/adiatma85/own-go-sdk/redis"
	"go.mongodb.org/mongo-driver/bson"
)

// TODO: Add Unit Tests
type Interface interface {
	Create(ctx context.Context, insertParam bson.M) error
	Get(ctx context.Context, params entity.AgGridMetricParam) (entity.AgGridMetric, error)
	GetList(ctx context.Context, params entity.AgGridMetricParam) ([]entity.AgGridMetric, *entity.Pagination, error)
	Update(ctx context.Context, params entity.AgGridMetricParam, insertParam bson.M) error
	Upsert(ctx context.Context, params entity.AgGridMetricParam, insertParam bson.M) error
}

type InitParam struct {
	Log   log.Interface
	Redis redis.Interface
	Json  parser.JSONInterface
	Mongo mongo.Interface
}

type pondMetricV2 struct {
	log   log.Interface
	redis redis.Interface
	json  parser.JSONInterface
	mongo mongo.Interface
}

func Init(param InitParam) Interface {
	pm := &pondMetricV2{
		log:   param.Log,
		redis: param.Redis,
		json:  param.Json,
		mongo: param.Mongo,
	}

	return pm
}

func (pm *pondMetricV2) Create(ctx context.Context, insertParam bson.M) error {
	pm.log.Debug(ctx, fmt.Sprintf("inserting pond metric v2 with params: %v", insertParam))

	var err error

	_, err = pm.mongo.Leader().InsertOne(ctx, entity.AggridCollectionName, "insert one pond metric v2", insertParam)
	if err != nil {
		return err
	}

	return nil
}

func (pm *pondMetricV2) Get(ctx context.Context, params entity.AgGridMetricParam) (entity.AgGridMetric, error) {
	pm.log.Debug(ctx, fmt.Sprintf("getting pond metric v2 with params: %v", params))
	result := entity.AgGridMetric{}

	res, err := pm.mongo.Follower().FindOne(ctx, entity.AggridCollectionName, "find one pond metric v2", params)
	if err != nil {
		return result, err
	}

	if err = res.Decode(&result); err != nil {
		return result, err
	}

	return result, nil
}

func (pm *pondMetricV2) GetList(ctx context.Context, params entity.AgGridMetricParam) ([]entity.AgGridMetric, *entity.Pagination, error) {
	pm.log.Debug(ctx, fmt.Sprintf("getting pond metric v2 with params: %v", params))
	var (
		totalDocument int64
		results       []entity.AgGridMetric
	)

	findOptions := params.GetFindOptions(params.QueryOption)
	cursor, err := pm.mongo.Follower().Find(ctx, entity.AggridCollectionName, "find pond metric v2", params, findOptions)
	if err != nil {
		return results, &entity.Pagination{}, err
	}

	for cursor.Next(ctx) {
		item := entity.AgGridMetric{}
		if err = cursor.Decode(&item); err != nil {
			pm.log.Error(ctx, err)
			continue
		}

		results = append(results, item)
	}

	// Process pagination in here
	pg := entity.Pagination{
		CurrentPage:     params.Page,
		CurrentElements: int64(len(results)),
	}

	// Query limit in here must be included
	if len(results) > 0 && !params.QueryOption.DisableLimit && params.IncludePagination {
		totalDocument, err = pm.mongo.Follower().CountDocument(ctx, entity.AggridCollectionName, "count pond metric v2", params)
		if err != nil {
			return results, nil, err
		}

		pg.TotalElements = totalDocument
	}
	pg.ProcessPagination(params.Limit)

	return results, &pg, nil
}

func (pm *pondMetricV2) Update(ctx context.Context, params entity.AgGridMetricParam, insertParam bson.M) error {
	pm.log.Debug(ctx, fmt.Sprintf("update pond metric v2 with param: %v", params))

	var err error

	_, err = pm.mongo.Leader().Update(ctx, entity.AggridCollectionName, "update pond metric v2", params, insertParam)
	if err != nil {
		pm.log.Error(ctx, err)
		return err
	}

	return nil
}

func (pm *pondMetricV2) Upsert(ctx context.Context, params entity.AgGridMetricParam, insertParam bson.M) error {
	pm.log.Debug(ctx, fmt.Sprintf("upsert pond metric v2 with param: %v", params))

	var err error

	// TODO: Specify the Error Code on sdk/codes
	_, err = pm.mongo.Leader().Upsert(ctx, entity.AggridCollectionName, "upsert pond metric v2", params, insertParam)
	if err != nil {
		pm.log.Error(ctx, err)
		return err
	}

	return nil
}
