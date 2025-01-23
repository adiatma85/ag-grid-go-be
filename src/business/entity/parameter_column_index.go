package entity

import (
	"github.com/adiatma85/own-go-sdk/null"
	"github.com/adiatma85/own-go-sdk/query"
)

type ParameterColumnIndex struct {
	ID              int64                    `db:"id" json:"id"`
	ParameterName   string                   `db:"parameter_name" json:"parameterName"`
	NumberFormat    null.String              `db:"number_format" json:"numberFormat"`
	IsGeneralColumn bool                     `db:"is_general_column" json:"isGeneralColumn"`
	IsConstant      bool                     `db:"is_constant" json:"isConstant"`
	ColIndex        int64                    `db:"col_index" json:"colIndex"`
	SheetName       string                   `db:"sheet_name" json:"sheetName"`
	Formula         null.String              `db:"formula" json:"formula"`
	Dropdown        null.String              `db:"dropdown" json:"dropdown"`
	IsZeroAllowed   bool                     `db:"is_zero_allowed" json:"isZeroAllowed"`
	Status          int64                    `db:"status" json:"status"`
	Flag            int64                    `db:"flag" json:"flag"`
	Meta            null.String              `db:"meta" json:"meta" swaggertype:"string"`
	MetaObj         MetaParameterColumnIndex `db:"-" json:"metaObj"`
	CreatedAt       null.Time                `db:"created_at" json:"createdAt" swaggertype:"string" example:"2022-06-21T10:32:29Z"`
	CreatedBy       null.String              `db:"created_by" json:"createdBy" swaggertype:"string"`
	UpdatedAt       null.Time                `db:"updated_at" json:"updatedAt" swaggertype:"string" example:"2022-06-21T10:32:29Z"`
	UpdatedBy       null.String              `db:"updated_by" json:"updatedBy" swaggertype:"string"`
	DeletedAt       null.Time                `db:"deleted_at" json:"deletedAt" swaggertype:"string" example:"2022-06-21T10:32:29Z"`
	DeletedBy       null.String              `db:"deleted_by" json:"deletedBy" swaggertype:"string"`
}

type ParameterColumnIndexParam struct {
	ID                    int64     `db:"id" param:"id" uri:"parameter_column_index_id"`
	ParameterName         string    `db:"parameter_name" param:"parameter_name"`
	ParameterNames        []string  `db:"parameter_name" param:"parameter_name"`
	IsGeneralColumn       null.Bool `db:"is_general_column" param:"is_general_column"`
	IsConstant            null.Bool `db:"is_constant" json:"isConstant"`
	ColIndex              int64     `db:"col_index" param:"col_index" uri:"col_index"`
	Formulas              []string  `db:"formula" param:"formula" json:"formulas" form:"formulas"`
	NumberFormatIsNotNull bool      `db:"number_format" param:"number_format__neq_null"`
	NumberFormatIsNull    bool      `db:"number_format" param:"number_format__eq_null"`
	FormulaIsNotNull      bool      `db:"formula" param:"formula__neq_null"`
	FormulaIsNull         bool      `db:"formula" param:"formula__eq_null"`
	DropdownIsNotNull     bool      `db:"dropdown" param:"dropdown__neq_null"`
	SheetName             string    `db:"sheet_name" param:"sheet_name" form:"sheetName"`
	DropdownIsNull        bool      `db:"dropdown" param:"dropdown__eq_null"`
	IsZeroAllowed         null.Bool `db:"is_zero_allowed" param:"is_zero_allowed"`
	QueryOption           query.Option
	PaginationParam
	BypassCache bool
}

type ParameterColumnIndexInputParam struct {
	ParameterName   string      `db:"parameter_name" json:"parameterName"`
	NumberFormat    null.String `db:"number_format" json:"numberFormat"`
	IsGeneralColumn null.Bool   `db:"is_general_column" json:"isGeneralColumn"`
	IsConstant      null.Bool   `db:"is_constant" json:"isConstant"`
	ColIndex        int64       `db:"col_index" json:"colIndex"`
	SheetName       string      `db:"sheet_name" json:"sheetName"`
	Formula         null.String `db:"formula" json:"formula"`
	Dropdown        null.String `db:"dropdown" json:"dropdown"`

	CreatedAt null.Time   `db:"created_at" json:"-"`
	CreatedBy null.String `db:"created_by" json:"-"`
}

type ParameterColumnIndexUpdateParam struct {
	ParameterName   string      `db:"parameter_name" json:"parameterName"`
	NumberFormat    null.String `db:"number_format" json:"numberFormat"`
	IsGeneralColumn null.Bool   `db:"is_general_column" json:"isGeneralColumn"`
	IsConstant      null.Bool   `db:"is_constant" json:"isConstant"`
	ColIndex        int64       `db:"col_index" json:"colIndex"`
	SheetName       string      `db:"sheet_name" json:"sheetName"`
	Formula         null.String `db:"formula" json:"formula"`
	Dropdown        null.String `db:"dropdown" json:"dropdown"`
	Status          null.Int64  `db:"status" json:"-"`

	UpdatedAt null.Time   `db:"updated_at" json:"-"`
	UpdatedBy null.String `db:"updated_by" json:"-"`
	DeletedAt null.Time   `db:"deleted_at" json:"-"`
	DeletedBy null.String `db:"deleted_by" json:"-"`
}

type MetaParameterColumnIndex struct {
	IsSkip       bool   `json:"isSkip"`
	DropdownType string `json:"dropdownType"`
}
