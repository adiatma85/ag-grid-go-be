package entity

import (
	"time"

	"github.com/adiatma85/own-go-sdk/null"
	"github.com/adiatma85/own-go-sdk/query"
)

var (
	MapPondType = map[int64]string{
		1: "Kolam Produksi",
		2: "Kolam Tandon",
		3: "Kolam Treatment",
		4: "IPAL",
		5: "Sumber Air (Laut)",
	}
)

type Pond struct {
	ID                         int64        `db:"id" json:"id"`
	FarmID                     int64        `db:"fk_farm_id" json:"farmId"`
	ModuleID                   int64        `db:"fk_module_id" json:"moduleId"`
	CurrentDOC                 int64        `db:"-" json:"currentDoc"`
	CurrentDOP                 int64        `db:"-" json:"currentDop"`
	LatestMetricDate           null.Date    `db:"latest_metric_date" json:"latestMetricDate" swaggertype:"string" example:"2022-06-21T10:32:29Z"`
	LatestSamplingDate         null.Date    `db:"latest_sampling_date" json:"latestSamplingDate" swaggertype:"string" example:"2022-06-21T10:32:29Z"`
	CurrentWqi                 string       `db:"-" json:"currentWqi"`
	CurrentWqiScore            null.Float64 `db:"current_wqi_score" json:"currentWqiScore" swaggertype:"number"`
	CurrentCycleID             null.Int64   `db:"current_fk_cycle_id" json:"currentFkCycleId"  swaggertype:"integer"` // For easier query, add it when create module, make id 0 or null?? when cycle end
	CurrentCycleName           null.String  `db:"current_cycle_name" json:"currentCycleName" swaggertype:"string"`
	CurrentCycleStatus         null.String  `db:"current_cycle_status" json:"currentCycleStatus" swaggertype:"string"`
	IsFinalHarvest             null.Bool    `db:"is_final_harvest" json:"isFinalHarvest"`
	CurrentCycleSeverity       null.String  `db:"current_cycle_severity" json:"currentCycleSeverity" swaggertype:"string"`
	CurrentCycleSeverityRating null.Float64 `db:"current_cycle_severity_rating" json:"currentCycleSeverityRating"  swaggertype:"number"`
	CurrentCycleDopStartDate   null.Date    `db:"current_cycle_dop_start_date" json:"currentCycleDopStartDate" swaggertype:"string" example:"2022-06-21T10:32:29Z"`
	CurrentCycleDopEndDate     null.Date    `db:"current_cycle_dop_end_date" json:"currentCycleDopEndDate" swaggertype:"string" example:"2022-06-21T10:32:29Z"`
	CurrentCycleDocStartDate   null.Date    `db:"current_cycle_doc_start_date" json:"currentCycleDocStartDate" swaggertype:"string" example:"2022-06-21T10:32:29Z"`
	CurrentCycleDocEndDate     null.Date    `db:"current_cycle_doc_end_date" json:"currentCycleDocEndDate" swaggertype:"string" example:"2022-06-21T10:32:29Z"`
	CurrentCycleHatchery       null.String  `db:"current_cycle_hatchery" json:"currentCycleHatchery" swaggertype:"string"`
	CurrentCycleStockingDate   null.Date    `db:"current_cycle_stocking_date" json:"currentCycleStockingDate" swaggertype:"string" example:"2022-06-21T10:32:29Z"`
	CurrentCycleStockingNetto  null.Int64   `db:"current_cycle_stocking_netto" json:"currentCycleStockingNetto" swaggertype:"integer"`
	CurrentCycleStockingGross  null.Int64   `db:"current_cycle_stocking_gross" json:"currentCycleStockingGross" swaggertype:"integer"`
	CurrentCycleStockingActual null.Int64   `db:"current_cycle_stocking_actual" json:"currentCycleStockingActual" swaggertype:"integer"`
	CurrentCycleDensityNetto   null.Int64   `db:"current_cycle_density_netto" json:"currentCycleDensityNetto" swaggertype:"integer"`
	CurrentCycleDensityGross   null.Int64   `db:"current_cycle_density_gross" json:"currentCycleDensityGross" swaggertype:"integer"`
	CurrentCycleDensityActual  null.Int64   `db:"current_cycle_density_actual" json:"currentCycleDensityActual" swaggertype:"integer"`
	CurrentCycleAccFeed        null.Int64   `db:"current_cycle_acc_feed" json:"currentCycleAccFeed" swaggertype:"integer"`
	CurrentCycleSampling       null.Int64   `db:"current_cycle_sampling" json:"currentCycleSampling" swaggertype:"integer"`
	HarvestPlanMinShrimpSize   int64        `db:"harvest_plan_min_shrimp_size" json:"harvestPlanMinShrimpSize" swaggertype:"integer"`
	HarvestPlanMaxShrimpSize   int64        `db:"harvest_plan_max_shrimp_size" json:"harvestPlanMaxShrimpSize" swaggertype:"integer"`
	HarvestPlanDayInterval     int64        `db:"harvest_plan_day_interval" json:"harvestPlanDayInterval" swaggertype:"integer"`
	HarvestPlanDocLength       int64        `db:"harvest_plan_doc_length" json:"harvestPlanDocLength" swaggertype:"integer"`
	HarvestPlanBiomassTarget   int64        `db:"harvest_plan_biomass_target" json:"harvestPlanBiomassTarget" swaggertype:"integer"`
	Name                       string       `db:"name" json:"name"`
	Code                       string       `db:"code" json:"code"`
	PondType                   int64        `db:"pond_type" json:"pondType"`
	Category                   string       `db:"" json:"category"`
	AreaM2                     null.Float64 `db:"area_m2" json:"areaM2" swaggertype:"number"`
	DepthM                     null.Float64 `db:"depth_m" json:"depthM" swaggertype:"number"`
	LengthM                    null.Float64 `db:"length_m" json:"lengthM" swaggertype:"number"`
	WidthM                     null.Float64 `db:"width_m" json:"widthM" swaggertype:"number"`
	WaterHeightM               null.Float64 `db:"-" json:"waterHeightM" swaggertype:"number"`
	WaterVolumeL               null.Float64 `db:"-" json:"waterVolumeL" swaggertype:"number"`
	DiseaseDescription         null.String  `db:"disease_description" json:"diseaseDescription" example:"disease"`
	Status                     int64        `db:"status" json:"status"`
	Flag                       int64        `db:"flag" json:"flag"`
	Meta                       null.String  `db:"meta" json:"meta" swaggertype:"string"`
	CreatedAt                  null.Time    `db:"created_at" json:"createdAt" swaggertype:"string" example:"2022-06-21T10:32:29Z"`
	CreatedBy                  null.String  `db:"created_by" json:"createdBy" swaggertype:"string"`
	UpdatedAt                  null.Time    `db:"updated_at" json:"updatedAt" swaggertype:"string" example:"2022-06-21T10:32:29Z"`
	UpdatedBy                  null.String  `db:"updated_by" json:"updatedBy" swaggertype:"string"`
	DeletedAt                  null.Time    `db:"deleted_at" json:"deletedAt" swaggertype:"string" example:"2022-06-21T10:32:29Z"`
	DeletedBy                  null.String  `db:"deleted_by" json:"deletedBy" swaggertype:"string"`
	ModuleName                 string       `db:"module_name" json:"moduleName"`
	CycleStatus                string       `db:"cycle_status" json:"cycleStatus"`
	PondProgress               int64        `db:"-" json:"pondProgress"`
	PondTotalSample            int64        `db:"-" json:"pondTotalSample"`
	HasRecommendation          bool         `db:"-" json:"hasRecommendation"`
	PondStatus                 null.String  `db:"-" json:"pondStatus" example:"normal/waspada/nonaktif"`
	LabelColor                 null.String  `db:"-" json:"labelColor" example:"green/yellow/grey"`
	LabelTextColor             null.String  `db:"-" json:"labelTextColor" example:"green/yellow/grey"`
	LabelBackgroundColorCode   null.String  `db:"-" json:"labelBackgroundColorCode" example:"green/yellow/grey"`
	LabelTextColorCode         null.String  `db:"-" json:"labelTextColorCode" example:"green/yellow/grey"`
	LabelBorderColorCode       null.String  `db:"-" json:"labelborderColorCode" example:"green/yellow/grey"`
}

type PondParam struct {
	FarmID               null.Int64 `uri:"farm_id" param:"fk_farm_id" db:"fk_farm_id" form:"farm_id" swaggertype:"integer"`
	FarmIDs              []int64    `param:"farm_ids" db:"fk_farm_id" form:"farm_ids"`
	CycleID              null.Int64 `uri:"cycle_id" param:"current_fk_cycle_id" db:"current_fk_cycle_id" form:"cycle_id" swaggertype:"integer"`
	CycleIDs             []int64    `param:"current_fk_cycle_ids" db:"current_fk_cycle_id"`
	PondID               null.Int64 `uri:"pond_id" param:"id" db:"id" swaggertype:"integer" form:"pond_id"`
	PondIDs              []int64    `param:"pond_ids" db:"id"`
	Code                 string     `param:"code" db:"code" json:"code" form:"code"`
	PondTypeIN           []int64    `param:"pond_types" db:"pond_type"`
	PondTypeNIN          []int64    `param:"pond_type__nin" db:"pond_type"`
	CurrentCycleStatus   string     `json:"current_cycle_status" db:"current_cycle_status" form:"current_cycle_status"`
	CurrentCycleSampling null.Int64 `param:"current_cycle_sampling" db:"current_cycle_sampling" form:"currentCycleSampling" swaggertype:"integer"`
	ModuleID             null.Int64 `param:"fk_module_id" db:"fk_module_id" form:"module_id"`
	ModuleIDs            []int64    `param:"module_ids" db:"fk_module_id" form:"module_ids"`
	CycleStatus          string     `json:"-" db:"-" form:"cycle_status"`
	WQIScoreLowerRange   null.Int64 `param:"current_wqi_score__gt" db:"current_wqi_score" swaggertype:"integer"`
	WQIScoreUpperRange   null.Int64 `param:"current_wqi_score__lte" db:"current_wqi_score" swaggertype:"integer"`
	ProgressDate         time.Time  `param:"-" db:"-" form:"progress_date" swaggertype:"string" example:"2022-06-21T10:32:29Z"`
	BypassCache          bool       `json:"-"`
	LatestSamplingDate   null.Date  `param:"latest_sampling_date" db:"latest_sampling_date"`
	LatestMetricDate     null.Date  `param:"latest_metric_date" db:"latest_metric_date"`
	BypassUserCheck      bool       `form:"bypass_user_check"`
	FetchPondMetric      bool       `form:"fetch_pond_metric"`
	Codes                []string   `param:"codes" db:"code" json:"codes" form:"codes"`
	PaginationParam
	QueryOption query.Option
}

func (p *Pond) GetPondType() string {
	if value, ok := MapPondType[p.PondType]; ok {
		return value
	}

	return ""
}
