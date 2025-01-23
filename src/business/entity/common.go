package entity

import (
	"github.com/adiatma85/own-go-sdk/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PaginationParamMongo struct {
	GroupBy           []string `param:"-" db:"-"`
	SortBy            []string `param:"sort_by" db:"sort_by"`
	Limit             int64    `form:"limit" param:"limit" db:"limit"`
	Page              int64    `form:"page" param:"page" db:"page"`
	IncludePagination bool
}

func (pp *PaginationParamMongo) GetLimitAndSkip() (int64, int64) {
	// Set the default page to 1
	if pp.Page < 1 {
		pp.Page = 1
	}

	// Set the default limit to 10
	if pp.Limit < 1 {
		pp.Limit = 10
	}

	return pp.Limit, pp.Page*pp.Limit - pp.Limit
}

// TODO: Add Logic for Group By and Sort By with this reference
// 1. https://www.mongodb.com/docs/drivers/go/master/fundamentals/crud/read-operations/sort/
// See the `Sort` and `Aggregation`
// Return options
func (pp *PaginationParamMongo) GetFindOptions(option mongo.Option) *options.FindOptions {
	options := options.Find()

	if option.DisableLimit {
		return options
	}

	limit, skip := pp.GetLimitAndSkip()
	options.SetLimit(limit)
	options.SetSkip(skip)

	if len(pp.SortBy) > 0 {
		sorts := bson.D{}
		for _, item := range pp.SortBy {
			if item[0] == '-' {
				sorts = append(sorts, bson.E{Key: item[1:], Value: -1})
			} else {
				sorts = append(sorts, bson.E{Key: item, Value: 1})
			}
		}

		options.SetSort(sorts)
	}

	return options
}
