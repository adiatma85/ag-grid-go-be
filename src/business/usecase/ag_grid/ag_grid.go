package aggrid

import (
	"context"
	"fmt"

	aggridDomain "github.com/adiatma85/new-go-template/src/business/domain/ag_grid"
	"github.com/adiatma85/new-go-template/src/business/domain/parameter"
	"github.com/adiatma85/new-go-template/src/business/domain/parameter_column_index"
	"github.com/adiatma85/new-go-template/src/business/domain/pond"
	"github.com/adiatma85/new-go-template/src/business/entity"
	"github.com/adiatma85/own-go-sdk/jwtAuth"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Interface interface {
	GetListMetric(ctx context.Context, metricParam entity.AgGridMetricParam) ([]entity.AgGridMetric, error)
	SchedulerToPopulateAgGridData(ctx context.Context) error
	InsertingMetric(ctx context.Context, inputMetrics []entity.AggridMetricInput) error
}

type InitParam struct {
	AggridDomain         aggridDomain.Interface // this is mongo db
	JwtAuth              jwtAuth.Interface
	Parameter            parameter.Interface
	ParameterColumnIndex parameter_column_index.Interface
	Pond                 pond.Interface
}

type aggrid struct {
	aggridDomain         aggridDomain.Interface
	jwtAuth              jwtAuth.Interface
	parameter            parameter.Interface
	parameterColumnIndex parameter_column_index.Interface
	pond                 pond.Interface
}

func Init(param InitParam) Interface {
	pm := &aggrid{
		aggridDomain:         param.AggridDomain,
		jwtAuth:              param.JwtAuth,
		parameter:            param.Parameter,
		parameterColumnIndex: param.ParameterColumnIndex,
		pond:                 param.Pond,
	}

	return pm
}

func (a *aggrid) GetListMetric(ctx context.Context, metricParam entity.AgGridMetricParam) ([]entity.AgGridMetric, error) {
	result, _, err := a.aggridDomain.GetList(ctx, metricParam)
	if err != nil {
		return []entity.AgGridMetric{}, err
	}

	return result, nil
}

// Kind of clumsy, but it's okay
func (a *aggrid) InsertingMetric(ctx context.Context, inputMetrics []entity.AggridMetricInput) error {
	for _, inputeMetric := range inputMetrics {
		fmt.Println(inputeMetric)
		inserItem := a.convertMetricToBsonD(inputeMetric.AttributeName, inputeMetric.Value)

		objectID, err := primitive.ObjectIDFromHex(inputeMetric.ID)
		if err != nil {
			return err
		}

		filter := bson.M{
			"_id": objectID,
		}

		err = a.aggridDomain.UpdateV2(ctx, filter, inserItem)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *aggrid) convertMetricToBsonD(key string, value interface{}) bson.M {
	insertItem := bson.M{
		"$set": bson.M{
			key: value,
		},
	}

	return insertItem
}

func (a *aggrid) SchedulerToPopulateAgGridData(ctx context.Context) error {
	return nil
}
