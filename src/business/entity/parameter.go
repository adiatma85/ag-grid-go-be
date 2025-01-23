package entity

import (
	"github.com/adiatma85/own-go-sdk/null"
	"github.com/adiatma85/own-go-sdk/query"
)

const (
	NoRowsAffected           = "no rows affected"
	DuplicateEntryErrMessage = "Duplicate entry"
)

type Parameter struct {
	ID                    int64               `db:"id" json:"id"`
	ParameterType         string              `db:"parameter_type" json:"parameterType"`
	Group                 string              `db:"group" json:"group"`
	Name                  string              `db:"name" json:"name"`
	HasPrediction         bool                `db:"has_prediction" json:"-"`
	InputType             null.String         `db:"input_type" json:"inputType"`
	FieldType             null.String         `db:"field_type" json:"fieldType"`
	InputOptionsString    null.String         `db:"input_options" json:"-"`
	InputOptions          []string            `db:"-" json:"inputOptions"`
	CascadingInputOptions map[string][]string `db:"-" json:"cascadingInputOptions,omitempty"`
	CascadingParameterId  int64               `db:"-" json:"cascadingParameterId,omitempty"`
	DisplayHintID         null.String         `db:"hint_id" json:"hintId"`
	DisplayHintEn         null.String         `db:"hint_en" json:"hintEn"`
	DisplayHint           null.String         `db:"-" json:"hint"`
	SubGroupString        null.String         `db:"sub_group" json:"-"`
	SubGroup              []int64             `db:"-" json:"subGroup"`
	TimeClassification    string              `db:"time_classification" json:"timeClassification"`
	TimeClassificationIds []int64             `db:"-" json:"timeClassificationIds"`
	Notes                 null.String         `db:"notes" json:"notes"`
	Priority              int                 `db:"priority" json:"priority"`
	Source                null.String         `db:"source" json:"source"`
	Type                  null.String         `db:"type" json:"type"`
	Unit                  string              `db:"unit" json:"unit"`
	OptionalUnitString    null.String         `db:"optional_units" json:"-"`
	OptionalUnits         []string            `db:"-" json:"optionalUnits"`
	DisplayNameID         string              `db:"display_name_id" json:"displayNameId"`
	DisplayNameEn         string              `db:"display_name_en" json:"displayNameEn"`
	ExportHeaderId        string              `db:"export_header_id" json:"exportHeaderId"`
	DisplayName           string              `db:"-" json:"displayName"`
	YRangeMin             null.Float64        `db:"y_range_min" json:"yRangeMin"`
	YRangeMax             null.Float64        `db:"y_range_max" json:"yRangeMax"`
	XRangeMin             null.Float64        `db:"x_range_min" json:"xRangeMin"`
	XRangeMax             null.Float64        `db:"x_range_max" json:"xRangeMax"`
	XLegend               null.String         `db:"x_legend" json:"xLegend"`
	ThresholdType         string              `db:"threshold_type" json:"thresholdType"` // minimal, maximal, between
	AlertTrackingType     string              `db:"alert_tracking_type" json:"alertTrackingType"`
	ThresholdMin          float64             `db:"threshold_min" json:"thresholdMin"`
	ThresholdMax          float64             `db:"threshold_max" json:"thresholdMax"`
	ExtremeChangesLower   null.Int64          `db:"extreme_changes_lower" json:"extremeChangesLower"`
	ExtremeChangesUpper   null.Int64          `db:"extreme_changes_upper" json:"extremeChangesUpper"`
	AlertThresholdMin     float64             `db:"alert_threshold_min" json:"alertThresholdMin"`
	AlertThresholdMax     float64             `db:"alert_threshold_max" json:"alertThresholdMax"`
	AlertDifference       float64             `db:"alert_difference" json:"alertDifference"`
	GraphType             string              `db:"graph_type" json:"graphType"`
	MetricsType           string              `db:"metrics_type" json:"metricsType"`
	DataLevel             string              `db:"data_level" json:"dataLevel"`
	JSONKey               string              `db:"json_key" json:"jsonKey"`
	TableColumnKey        string              `db:"table_column_key" json:"tableColumnKey"`
	TableSource           string              `db:"table_source" json:"tableSource"`
	AdditionalConfig      string              `db:"additional_config" json:"additionalConfig"`
	APIParams             string              `db:"api_params" json:"apiParams"`
	APIURL                string              `db:"api_url" json:"apiUrl"`
	Tooltip               string              `db:"-" json:"tooltip"`
	Status                int16               `db:"status" json:"status"`
	Flag                  int                 `db:"flag" json:"flag"`

	Autofill  bool        `db:"-" json:"autofill"`
	Meta      null.String `db:"meta" json:"meta" swaggertype:"string"`
	CreatedAt null.Time   `db:"created_at" json:"createdAt" swaggertype:"string" example:"2022-06-21T10:32:29Z"`
	CreatedBy null.String `db:"created_by" json:"createdBy" swaggertype:"string"`
	UpdatedAt null.Time   `db:"updated_at" json:"updatedAt" swaggertype:"string" example:"2022-06-21T10:32:29Z"`
	UpdatedBy null.String `db:"updated_by" json:"updatedBy" swaggertype:"string"`
	DeletedAt null.Time   `db:"deleted_at" json:"deletedAt" swaggertype:"string" example:"2022-06-21T10:32:29Z"`
	DeletedBy null.String `db:"deleted_by" json:"deletedBy" swaggertype:"string"`

	// Not Needed for now
	// Formula               ParameterFormula    `db:"-" json:"formula,omitempty"`
	Validation ValidationParameter `db:"-" json:"validation"`
}

type ParameterParam struct {
	ParameterType               string     `form:"parameterType" param:"parameter_type" db:"parameter_type"`
	ParameterTimeClassification string     `form:"parameterTimeClassification" param:"time_classification" db:"time_classification"`
	ParameterGroup              string     `form:"parameterGroup" param:"group" db:"group"`
	HasPrediction               bool       `form:"hasPrediction" param:"has_prediction" db:"has_prediction"`
	ParameterID                 null.Int64 `uri:"parameter_id" form:"parameter_id" param:"id" db:"id"`
	Status                      null.Int64 `uri:"status" form:"status" param:"status" db:"status"`
	ParameterIDs                []int64    `param:"parameter_ids" db:"id"`
	Name                        string     `param:"name" db:"name"`
	Names                       []string   `param:"name" db:"name"`
	Priority                    null.Int64 `param:"priority" db:"priority"`
	AlertTrackingType           string     `param:"alert_tracking_type" db:"alert_tracking_type"`
	AlertTrackingTypes          []string   `param:"alert_tracking_types" db:"alert_tracking_type"`
	ExtremeChangesLowerGt       null.Int64 `db:"extreme_changes_lower" param:"extreme_changes_lower__gt"`
	ExtremeChangesUpperGt       null.Int64 `db:"extreme_changes_upper" param:"extreme_changes_upper__gt"`
	AppLang                     string
	BypassCache                 bool `json:"-"`
	PaginationParam
	QueryOption query.Option
}

type ParameterUpdateParam struct {
	ParameterType      string       `db:"parameter_type" json:"parameterType"`
	Group              null.String  "db:\"`group`\" json:\"group\""
	Name               string       `db:"name" json:"name"`
	HasPrediction      bool         `db:"has_prediction" json:"-"`
	InputType          null.String  `db:"input_type" json:"inputType"`
	FieldType          null.String  `db:"field_type" json:"fieldType"`
	InputOptionsString null.String  `db:"input_options" json:"-"`
	InputOptions       []string     `db:"-" json:"inputOptions"`
	DisplayHintID      null.String  `db:"hint_id" json:"-"`
	DisplayHintEn      null.String  `db:"hint_en" json:"-"`
	DisplayHint        null.String  `db:"-" json:"hint"`
	SubGroupString     null.String  `db:"sub_group" json:"-"`
	SubGroup           []int64      `db:"-" json:"subGroup"`
	TimeClassification string       `db:"time_classification" json:"timeClassification"`
	Notes              null.String  `db:"notes" json:"notes"`
	Priority           int          `db:"priority" json:"priority"`
	Source             null.String  `db:"source" json:"source"`
	Type               null.String  `db:"type" json:"type"`
	Unit               string       `db:"unit" json:"unit"`
	DisplayNameID      string       `db:"display_name_id" json:"displayNameId"`
	DisplayNameEn      string       `db:"display_name_en" json:"displayNameEn"`
	ExportHeaderId     string       `db:"export_header_id" json:"exportHeaderId"`
	DisplayName        string       `db:"-" json:"displayName"`
	YRangeMin          null.Float64 `db:"y_range_min" json:"yRangeMin"`
	YRangeMax          null.Float64 `db:"y_range_max" json:"yRangeMax"`
	XRangeMin          null.Float64 `db:"x_range_min" json:"xRangeMin"`
	XRangeMax          null.Float64 `db:"x_range_max" json:"xRangeMax"`
	XLegend            null.String  `db:"x_legend" json:"xLegend"`
	ThresholdType      string       `db:"threshold_type" json:"thresholdType"` // minimal, maximal, between
	ThresholdMin       float64      `db:"threshold_min" json:"thresholdMin"`
	ThresholdMax       float64      `db:"threshold_max" json:"thresholdMax"`
	GraphType          string       `db:"graph_type" json:"graphType"`
	MetricsType        string       `db:"metrics_type" json:"metricsType"`
	DataLevel          string       `db:"data_level" json:"dataLevel"`
	JSONKey            string       `db:"json_key" json:"jsonKey"`
	TableColumnKey     string       `db:"table_column_key" json:"tableColumnKey"`
	TableSource        string       `db:"table_source" json:"tableSource"`
	AdditionalConfig   string       `db:"additional_config" json:"additionalConfig"`
	APIParams          string       `db:"api_params" json:"apiParams"`
	APIURL             string       `db:"api_url" json:"apiUrl"`
	AlertThresholdMin  null.Int64   `db:"alert_threshold_min" json:"alertThresholdMin"`
	AlertThresholdMax  null.Int64   `db:"alert_threshold_max" json:"alertThresholdMax"`
	Status             null.Int64   `db:"status" json:"status"`
	UpdatedAt          null.Time    `db:"updated_at" json:"updatedAt"`
	UpdatedBy          null.String  `db:"updated_by" json:"updatedBy"`
	DeletedAt          null.Time    `db:"deleted_at" json:"deletedAt"`
	DeletedBy          null.String  `db:"deleted_by" json:"deletedBy"`
}

type ValidationParameter struct {
	IsRequired       bool   `db:"-" json:"isRequired"`
	ErrorMessage     string `db:"-" json:"errorMessage"`
	ErrorDescription string `db:"-" json:"errorDescription"`
}

type MetaParameter struct {
	IsRequired       bool   `db:"-" json:"isRequired"`
	ErrorMessage     string `db:"-" json:"errorMessage"`
	ErrorDescription string `db:"-" json:"errorDescription"`
	IsAutofill       bool   `db:"-" json:"isAutofill"`
	Tooltip          string `db:"-" json:"tooltip"`
}
