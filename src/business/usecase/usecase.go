package usecase

import (
	"github.com/adiatma85/new-go-template/src/business/domain"
	aggrid "github.com/adiatma85/new-go-template/src/business/usecase/ag_grid"
	"github.com/adiatma85/new-go-template/src/business/usecase/user"
	"github.com/adiatma85/own-go-sdk/jwtAuth"
	"github.com/adiatma85/own-go-sdk/log"
)

type Usecase struct {
	User   user.Interface
	Aggrid aggrid.Interface
}

type InitParam struct {
	Log     log.Interface
	Dom     *domain.Domain
	JwtAuth jwtAuth.Interface
}

func Init(param InitParam) *Usecase {
	usecase := &Usecase{
		User:   user.Init(user.InitParam{Log: param.Log, User: param.Dom.User, JwtAuth: param.JwtAuth}),
		Aggrid: aggrid.Init(aggrid.InitParam{AggridDomain: param.Dom.AgGrid, JwtAuth: param.JwtAuth, Parameter: param.Dom.Parameter, ParameterColumnIndex: param.Dom.ParameterColumnIndex, Pond: param.Dom.Pond}),
	}

	return usecase
}
