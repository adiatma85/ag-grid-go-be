package parameter

import (
	"context"
	"fmt"

	"github.com/adiatma85/new-go-template/src/business/entity"
	"github.com/adiatma85/own-go-sdk/appcontext"
	"github.com/adiatma85/own-go-sdk/codes"
	"github.com/adiatma85/own-go-sdk/errors"
	"github.com/adiatma85/own-go-sdk/language"
	"github.com/adiatma85/own-go-sdk/query"
	"github.com/adiatma85/own-go-sdk/sql"
)

func (f *parameter) getSQL(ctx context.Context, params entity.ParameterParam) (entity.Parameter, error) {
	f.log.Debug(ctx, fmt.Sprintf("getting parameter ID: %v", params.ParameterID))

	parameter := entity.Parameter{}
	qb := query.NewSQLQueryBuilder(f.db, "param", "db", &params.QueryOption).AddAliasPrefix("p", &params)
	queryExt, queryArgs, _, _, err := qb.Build(&params)
	if err != nil {
		return parameter, errors.NewWithCode(codes.CodeSQLBuilder, err.Error())
	}

	row, err := f.db.Follower().QueryRow(ctx, "rParameterByID", readParameter+queryExt, queryArgs...)
	if err != nil && !errors.Is(err, sql.ErrNotFound) {
		return parameter, errors.NewWithCode(codes.CodeSQLRead, err.Error())
	} else if errors.Is(err, sql.ErrNotFound) {
		return parameter, errors.NewWithCode(codes.CodeSQLRecordDoesNotExist, err.Error())
	}

	if err := row.StructScan(&parameter); err != nil {
		if errors.Is(err, sql.ErrNotFound) {
			return parameter, errors.NewWithCode(codes.CodeSQLRecordDoesNotExist, err.Error())
		}
		return parameter, errors.NewWithCode(codes.CodeSQLRowScan, err.Error())
	}

	if parameter.InputOptionsString.Valid {
		var err = f.json.Unmarshal([]byte(parameter.InputOptionsString.String), &parameter.InputOptions)
		if err != nil {
			f.log.Error(ctx, errors.NewWithCode(codes.CodeUnmarshal, err.Error()))
		}
	}

	// Optional unit
	if parameter.OptionalUnitString.Valid {
		err := f.json.Unmarshal([]byte(parameter.OptionalUnitString.String), &parameter.OptionalUnits)
		if err != nil {
			return parameter, err
		}
	}

	if parameter.SubGroupString.Valid {
		var err = f.json.Unmarshal([]byte(parameter.SubGroupString.String), &parameter.SubGroup)
		if err != nil {
			f.log.Error(ctx, errors.NewWithCode(codes.CodeUnmarshal, err.Error()))
		}
	}

	return parameter, nil
}

func (f *parameter) getListSQL(ctx context.Context, params entity.ParameterParam) ([]entity.Parameter, map[int64]entity.Parameter, map[string]entity.Parameter, *entity.Pagination, error) {
	f.log.Debug(ctx, "getting parameter list")

	parameters := []entity.Parameter{}
	mapParamater := make(map[int64]entity.Parameter)
	mapParameterByName := make(map[string]entity.Parameter)

	qb := query.NewSQLQueryBuilder(f.db, "param", "db", &params.QueryOption).AddAliasPrefix("p", &params)

	queryExt, queryArgs, countExt, countArgs, err := qb.Build(&params)
	if err != nil {
		return parameters, mapParamater, mapParameterByName, nil, errors.NewWithCode(codes.CodeSQLBuilder, err.Error())
	}

	rows, err := f.db.Follower().Query(ctx, "rParameterList", readParameter+queryExt, queryArgs...)
	if err != nil && !errors.Is(err, sql.ErrNotFound) {
		return parameters, mapParamater, mapParameterByName, nil, errors.NewWithCode(codes.CodeSQLRead, err.Error())
	}

	selectedLang := appcontext.GetAcceptLanguage(ctx)

	defer rows.Close()

	for rows.Next() {
		parameter := entity.Parameter{}
		if err = rows.StructScan(&parameter); err != nil {
			f.log.Error(ctx, errors.NewWithCode(codes.CodeSQLRowScan, err.Error()))
			continue
		}

		parameter, err := f.languageAndParameterInput(parameter, selectedLang)
		if err != nil {
			f.log.Error(ctx, err)
			continue
		}

		parameters = append(parameters, parameter)
		mapParamater[parameter.ID] = parameter
		mapParameterByName[parameter.Name] = parameter
	}

	pg := entity.Pagination{
		CurrentPage:     params.Page,
		CurrentElements: int64(len(parameters)),
	}
	if len(parameters) > 0 && !params.QueryOption.DisableLimit && params.IncludePagination {
		if err := f.db.Follower().Get(ctx, "cParameterList", readParameterCount+countExt, &pg.TotalElements, countArgs...); err != nil {
			return parameters, mapParamater, mapParameterByName, nil, errors.NewWithCode(codes.CodeSQLRead, err.Error())
		}
	}

	pg.ProcessPagination(params.Limit)

	return parameters, mapParamater, mapParameterByName, &pg, nil
}

// Need to marshal data parameter for inputOptions and subGroup on level usecase
func (f *parameter) createSQL(ctx context.Context, parameter entity.Parameter) (entity.Parameter, error) {
	f.log.Debug(ctx, "creating new parameter")

	var err error

	tx, err := f.db.Leader().BeginTx(ctx, "txParameter", sql.TxOptions{})
	if err != nil {
		return parameter, errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	res, err := tx.NamedExec("iNewParameter", createParameter, parameter)

	if err != nil {
		return parameter, errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		return parameter, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	} else if rowCount < 1 {
		return parameter, errors.NewWithCode(codes.CodeSQLNoRowsAffected, entity.NoRowsAffected)
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return parameter, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	}

	parameter.ID = lastID

	if err := tx.Commit(); err != nil {
		return parameter, errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}

	f.log.Debug(ctx, fmt.Sprintf("successfully created new parameter: %v", parameter))

	if err = f.deleteParameterCache(ctx); err != nil {
		f.log.Error(ctx, err)
	}

	return parameter, nil
}

func (f *parameter) updateSQL(ctx context.Context, updateParam entity.ParameterUpdateParam, selectParam entity.ParameterParam) error {
	f.log.Debug(ctx, fmt.Sprintf("updating parameter ID: %v", selectParam.ParameterID))

	var err error

	qb := query.NewSQLQueryBuilder(f.db, "param", "db", &selectParam.QueryOption)
	queryUpdate, args, err := qb.BuildUpdate(&updateParam, &selectParam)
	if err != nil {
		return errors.NewWithCode(codes.CodeSQLBuilder, err.Error())
	}

	query := updateParameter + queryUpdate

	res, err := f.db.Leader().Exec(ctx, "uParameter", query, args...)
	if err != nil {
		return errors.NewWithCode(codes.CodeSQLBuilder, err.Error())
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		return errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	} else if rowCount < 1 {
		return errors.NewWithCode(codes.CodeSQLNoRowsAffected, entity.NoRowsAffected)
	}

	f.log.Debug(ctx, fmt.Sprintf("successfully updated parameter: %v", selectParam.ParameterID.Int64))

	if err = f.deleteParameterCache(ctx); err != nil {
		f.log.Error(ctx, err)
	}

	return nil
}

func (f *parameter) languageAndParameterInput(parameter entity.Parameter, selectedLang string) (entity.Parameter, error) {
	switch selectedLang {
	case language.Indonesian:
		parameter.DisplayName = parameter.DisplayNameID
		parameter.DisplayHint = parameter.DisplayHintID
	default:
		parameter.DisplayName = parameter.DisplayNameEn
		parameter.DisplayHint = parameter.DisplayHintEn
	}

	// Input options
	if parameter.InputOptionsString.Valid {
		err := f.json.Unmarshal([]byte(parameter.InputOptionsString.String), &parameter.InputOptions)
		if err != nil {
			return parameter, err
		}
	}

	// Optional unit
	if parameter.OptionalUnitString.Valid {
		err := f.json.Unmarshal([]byte(parameter.OptionalUnitString.String), &parameter.OptionalUnits)
		if err != nil {
			return parameter, err
		}
	}

	// Sub group
	if parameter.SubGroupString.Valid {
		err := f.json.Unmarshal([]byte(parameter.SubGroupString.String), &parameter.SubGroup)
		if err != nil {
			return parameter, err
		}
	}

	// Time classification ids
	if parameter.TimeClassification != "" && string(parameter.TimeClassification[0]) == "[" {
		err := f.json.Unmarshal([]byte(parameter.TimeClassification), &parameter.TimeClassificationIds)
		if err != nil {
			return parameter, err
		}
	}

	return parameter, nil
}
