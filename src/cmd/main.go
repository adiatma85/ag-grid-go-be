package main

import (
	"github.com/adiatma85/new-go-template/src/business/domain"
	"github.com/adiatma85/new-go-template/src/business/usecase"
	"github.com/adiatma85/new-go-template/src/handler/rest"
	"github.com/adiatma85/new-go-template/src/handler/scheduler"
	"github.com/adiatma85/new-go-template/utils/config"
	"github.com/adiatma85/own-go-sdk/configreader"
	"github.com/adiatma85/own-go-sdk/instrument"
	"github.com/adiatma85/own-go-sdk/jwtAuth"
	"github.com/adiatma85/own-go-sdk/log"
	"github.com/adiatma85/own-go-sdk/mongo"
	"github.com/adiatma85/own-go-sdk/parser"
	"github.com/adiatma85/own-go-sdk/redis"
	"github.com/adiatma85/own-go-sdk/sql"
	"github.com/adiatma85/own-go-sdk/timelib"
)

// @contact.name   Rahmadhani Lucky Adiatma
// @contact.email  adiatma85@gmail.com

// @securitydefinitions.apikey BearerAuth
// @in header
// @name Authorization

const (
	configfile   string = "./etc/cfg/conf.json"
	templatefile string = "./etc/tpl/conf.template.json"
)

func main() {
	// Build the config
	// Assume the config is exist in the first place

	// Read the Config first
	cfg := config.Init()
	configreader := configreader.Init(configreader.Options{
		ConfigFile: configfile,
	})
	configreader.ReadConfig(&cfg)

	// init logger
	log := log.Init(cfg.Log)

	// init the instrument
	instr := instrument.Init(cfg.Instrument)

	// Init the DB
	db := sql.Init(cfg.SQL, log, instr)

	// init the parser
	parsers := parser.InitParser(log, cfg.Parser)

	// init the redis
	cache := redis.Init(cfg.Redis, log)

	// Init the jwt
	jwt := jwtAuth.Init(cfg.JwtAuth)

	// mongo db
	mongo := mongo.Init(cfg.DocumentDB, log, instr)

	// timelib
	timelib := timelib.Init()

	// Init the domain
	d := domain.Init(domain.InitParam{Log: log, Db: db, Json: parsers.JSONParser(), Redis: cache, Mongo: mongo, Timelib: timelib})

	// Init the usecase
	uc := usecase.Init(usecase.InitParam{Log: log, Dom: d, JwtAuth: jwt})

	// Scheduler
	scheduler := scheduler.Init(scheduler.InitParam{Conf: cfg.Scheduler, MetaConf: cfg.Meta, Log: log, JwtAuth: jwt, Uc: uc, Instr: instr})

	// Init the GIN
	rest := rest.Init(rest.InitParam{Conf: cfg.Gin, Json: parsers.JSONParser(), Log: log, Uc: uc, Instrument: instr, JwtAuth: jwt, Scheduler: scheduler})

	// run scheduler
	scheduler.Run()

	// Run the REST
	rest.Run()
}
