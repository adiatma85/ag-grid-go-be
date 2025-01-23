package rest

import (
	"github.com/adiatma85/new-go-template/src/business/entity"
	"github.com/adiatma85/own-go-sdk/codes"
	"github.com/gin-gonic/gin"
)

// @Summary Get List Metric AgGrid
// @Description Get list all metric Get List Metric AgGrid
// @Tags AgGrid
// @Param limit query integer false "limit"
// @Param page query integer false "page"
// @Param disableLimit query boolean false "disable limit" Enums(true, false)
// @Produce json
// @Success 200 {object} entity.HTTPResp{data=[]entity.AgGridMetric{}}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /public/v1/ag-grid/metrics [GET]
func (r *rest) GetListMetricAgGrid(ctx *gin.Context) {
	params := entity.AgGridMetricParam{}

	if err := r.BindParams(ctx, &params); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	metric, err := r.uc.Aggrid.GetListMetric(ctx.Request.Context(), params)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, metric, nil)
}

// @Summary Get List Metric AgGrid
// @Description Get list all metric Get List Metric AgGrid
// @Tags AgGrid
// @Param user body []entity.AggridMetricInput true "insertionMetricData"
// @Produce json
// @Success 200 {object} entity.HTTPResp{}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /public/v1/ag-grid/metrics [POST]
func (r *rest) InsertAgMetric(ctx *gin.Context) {
	insertionMetric := []entity.AggridMetricInput{}

	if err := r.Bind(ctx, &insertionMetric); err != nil {
		r.httpRespError(ctx, err)
		return
	}
	if err := r.uc.Aggrid.InsertingMetric(ctx.Request.Context(), insertionMetric); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, nil, nil)
}
