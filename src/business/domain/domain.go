package domain

import (
	aggrid "github.com/adiatma85/new-go-template/src/business/domain/ag_grid"
	"github.com/adiatma85/new-go-template/src/business/domain/user"
	"github.com/adiatma85/own-go-sdk/log"
	"github.com/adiatma85/own-go-sdk/mongo"
	"github.com/adiatma85/own-go-sdk/parser"
	"github.com/adiatma85/own-go-sdk/redis"
	"github.com/adiatma85/own-go-sdk/sql"
)

// This comment used to test the action

type Domain struct {
	User   user.Interface
	AgGrid aggrid.Interface
}

type InitParam struct {
	Log   log.Interface
	Db    sql.Interface
	Json  parser.JSONInterface
	Redis redis.Interface
	Mongo mongo.Interface
}

func Init(param InitParam) *Domain {
	domain := &Domain{
		User:   user.Init(user.InitParam{Log: param.Log, Db: param.Db, Json: param.Json, Redis: param.Redis}),
		AgGrid: aggrid.Init(aggrid.InitParam{Log: param.Log, Redis: param.Redis, Json: param.Json, Mongo: param.Mongo}),
	}

	return domain
}
