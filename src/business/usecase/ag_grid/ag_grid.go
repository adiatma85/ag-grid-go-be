package aggrid

import (
	"context"
	"fmt"
	"strconv"
	"time"

	aggridDomain "github.com/adiatma85/new-go-template/src/business/domain/ag_grid"
	"github.com/adiatma85/new-go-template/src/business/domain/parameter"
	"github.com/adiatma85/new-go-template/src/business/domain/parameter_column_index"
	"github.com/adiatma85/new-go-template/src/business/domain/pond"
	"github.com/adiatma85/new-go-template/src/business/entity"
	"github.com/adiatma85/own-go-sdk/jwtAuth"
	"github.com/adiatma85/own-go-sdk/log"
	"github.com/adiatma85/own-go-sdk/query"
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
	Log                  log.Interface
}

type aggrid struct {
	aggridDomain         aggridDomain.Interface
	jwtAuth              jwtAuth.Interface
	parameter            parameter.Interface
	parameterColumnIndex parameter_column_index.Interface
	pond                 pond.Interface
	log                  log.Interface
}

func Init(param InitParam) Interface {
	pm := &aggrid{
		aggridDomain:         param.AggridDomain,
		jwtAuth:              param.JwtAuth,
		parameter:            param.Parameter,
		parameterColumnIndex: param.ParameterColumnIndex,
		pond:                 param.Pond,
		log:                  param.Log,
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

// Kind of clumsy, but it's okay for prototype
func (a *aggrid) InsertingMetric(ctx context.Context, inputMetrics []entity.AggridMetricInput) error {
	for _, inputeMetric := range inputMetrics {
		inserItem, err := a.convertMetricToBsonD(ctx, inputeMetric.AttributeName, inputeMetric.Value)
		if err != nil {
			a.log.Error(ctx, fmt.Sprintf("Error converting metric to bson with key: %v", inputeMetric.AttributeName))
			continue
		}

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

func (a *aggrid) convertMetricToBsonD(ctx context.Context, key string, value interface{}) (bson.M, error) {

	// must check the type of the formula
	parameterParam := entity.ParameterParam{
		Name: key,
		QueryOption: query.Option{
			IsActive: true,
		},
	}

	parameter, err := a.parameter.Get(ctx, parameterParam)
	if err != nil {
		return bson.M{}, err
	}

	if parameter.FieldType.String == "float" || parameter.FieldType.String == "integer" {
		switch v := value.(type) {
		case float64:
			value = v
		case string:
			validFloat, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return bson.M{}, err
			}

			value = validFloat
		}

	}

	insertItem := bson.M{
		"$set": bson.M{
			key: value,
		},
	}

	return insertItem, nil
}

func (a *aggrid) SchedulerToPopulateAgGridData(ctx context.Context) error {
	// Fetch all the active pond
	pondParam := entity.PondParam{
		QueryOption: query.Option{
			DisableLimit: true,
			IsActive:     true,
		},
	}

	ponds, _, _, err := a.pond.GetList(ctx, pondParam)
	if err != nil {
		return err
	}

	todayDate := time.Now()

	for _, pond := range ponds {
		insertionParam := bson.M{
			"farm_id":     pond.FarmID,
			"pond_id":     pond.ID,
			"metric_date": todayDate.Format("2006-01-02"),
		}
		err = a.aggridDomain.Create(ctx, insertionParam)
		if err != nil {
			return err
		}
	}

	return nil
}
