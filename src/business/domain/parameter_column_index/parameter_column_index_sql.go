package parameter_column_index

import (
	"context"
	"fmt"
	"strings"

	"github.com/adiatma85/new-go-template/src/business/entity"

	"github.com/adiatma85/own-go-sdk/codes"
	"github.com/adiatma85/own-go-sdk/errors"
	"github.com/adiatma85/own-go-sdk/query"
	"github.com/adiatma85/own-go-sdk/sql"
)

func (p *parameterColumnIndex) getSQL(ctx context.Context, param entity.ParameterColumnIndexParam) (entity.ParameterColumnIndex, error) {
	parameterColumnIndex := entity.ParameterColumnIndex{}

	p.log.Debug(ctx, fmt.Sprintf("get parameter column index with body: %v", param))

	param.QueryOption.DisableLimit = true
	qb := query.NewSQLQueryBuilder(p.db, "param", "db", &param.QueryOption)
	queryExt, queryArgs, _, _, err := qb.Build(&param)
	if err != nil {
		return parameterColumnIndex, errors.NewWithCode(codes.CodeSQLBuilder, err.Error())
	}

	row, err := p.db.Follower().QueryRow(ctx, "rParameterColumnIndex", readParameterColumnIndex+queryExt, queryArgs...)
	if err != nil && !errors.Is(err, sql.ErrNotFound) {
		return parameterColumnIndex, errors.NewWithCode(codes.CodeSQLRead, err.Error())
	}

	if err := row.StructScan(&parameterColumnIndex); err != nil && errors.Is(err, sql.ErrNotFound) {
		return parameterColumnIndex, errors.NewWithCode(codes.CodeSQLRecordDoesNotExist, err.Error())
	} else if err != nil {
		return parameterColumnIndex, errors.NewWithCode(codes.CodeSQLRowScan, err.Error())
	}

	p.log.Debug(ctx, fmt.Sprintf("success get parameter column index with body: %v", param))

	return parameterColumnIndex, nil
}

func (p *parameterColumnIndex) getListSQL(ctx context.Context, param entity.ParameterColumnIndexParam) ([]entity.ParameterColumnIndex, *entity.Pagination, error) {
	parameterColumnIndexes := []entity.ParameterColumnIndex{}
	pg := entity.Pagination{}

	p.log.Debug(ctx, fmt.Sprintf("get parameter column index list with body: %v", param))

	qb := query.NewSQLQueryBuilder(p.db, "param", "db", &param.QueryOption)
	queryExt, queryArgs, countExt, countArgs, err := qb.Build(&param)
	if err != nil {
		return parameterColumnIndexes, &pg, errors.NewWithCode(codes.CodeSQLBuilder, err.Error())
	}

	rows, err := p.db.Follower().Query(ctx, "rParameterColumnIndexList", readParameterColumnIndex+queryExt, queryArgs...)
	if err != nil && !errors.Is(err, sql.ErrNotFound) {
		return parameterColumnIndexes, &pg, errors.NewWithCode(codes.CodeSQLRead, err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		parameterColumnIndex := entity.ParameterColumnIndex{}
		err := rows.StructScan(&parameterColumnIndex)
		if err != nil {
			p.log.Error(ctx, errors.NewWithCode(codes.CodeSQLRowScan, err.Error()))
			continue
		}

		// If the meta is valid
		if parameterColumnIndex.Meta.Valid {
			metaParameterColumnIndex := entity.MetaParameterColumnIndex{}
			err := p.json.Unmarshal([]byte(parameterColumnIndex.Meta.String), &metaParameterColumnIndex)
			if err != nil {
				p.log.Error(ctx, errors.NewWithCode(codes.CodeUnmarshal, fmt.Sprintf("failed to unmarshall meta for parameter colum index id: %d", parameterColumnIndex.ID)))
				continue
			}

			parameterColumnIndex.MetaObj = metaParameterColumnIndex
		}

		parameterColumnIndexes = append(parameterColumnIndexes, parameterColumnIndex)
	}

	pg = entity.Pagination{
		CurrentPage:     param.PaginationParam.Page,
		CurrentElements: int64(len(parameterColumnIndexes)),
		SortBy:          param.SortBy,
	}

	if !param.QueryOption.DisableLimit && len(parameterColumnIndexes) > 0 {
		err := p.db.Follower().Get(ctx, "cParameterColumnIndexList", countParameterColumnIndex+countExt, &pg.TotalElements, countArgs...)
		if err != nil {
			return parameterColumnIndexes, &pg, errors.NewWithCode(codes.CodeSQLRead, err.Error())
		}
	}

	pg.ProcessPagination(param.Limit)

	p.log.Debug(ctx, fmt.Sprintf("success get parameter column index list with body: %v", param))

	return parameterColumnIndexes, &pg, nil
}

func (p *parameterColumnIndex) createSQL(ctx context.Context, inputParam entity.ParameterColumnIndexInputParam) (entity.ParameterColumnIndex, error) {
	feedback := entity.ParameterColumnIndex{}

	p.log.Debug(ctx, fmt.Sprintf("create parameter column index with body: %v", inputParam))

	tx, err := p.db.Leader().BeginTx(ctx, "txParameterColumnIndex", sql.TxOptions{})
	if err != nil {
		return feedback, errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	res, err := tx.NamedExec("iNewParameterColumnIndex", insertParameterColumnIndex, inputParam)
	if err != nil && strings.Contains(err.Error(), entity.DuplicateEntryErrMessage) {
		return feedback, errors.NewWithCode(codes.CodeSQLUniqueConstraint, err.Error())
	} else if err != nil {
		return feedback, errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		return feedback, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	} else if rowCount < 1 {
		return feedback, errors.NewWithCode(codes.CodeSQLNoRowsAffected, "no parameter column index created")
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return feedback, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	}
	if err := tx.Commit(); err != nil {
		return feedback, errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}

	p.log.Debug(ctx, fmt.Sprintf("success create parameter column index with body: %v", inputParam))

	feedback = entity.ParameterColumnIndex{
		ID:              lastID,
		ParameterName:   inputParam.ParameterName,
		NumberFormat:    inputParam.NumberFormat,
		IsGeneralColumn: inputParam.IsGeneralColumn.Bool,
		IsConstant:      inputParam.IsConstant.Bool,
		ColIndex:        inputParam.ColIndex,
		SheetName:       inputParam.SheetName,
		Formula:         inputParam.Formula,
		Dropdown:        inputParam.Dropdown,
		Status:          1,
		CreatedAt:       inputParam.CreatedAt,
		CreatedBy:       inputParam.CreatedBy,
	}

	return feedback, nil
}

func (p *parameterColumnIndex) updateSQL(ctx context.Context, updateParam entity.ParameterColumnIndexUpdateParam, selectParam entity.ParameterColumnIndexParam) error {
	p.log.Debug(ctx, fmt.Sprintf("update parameter column index %v with body: %v", selectParam.ID, updateParam))

	qb := query.NewSQLQueryBuilder(p.db, "param", "db", &selectParam.QueryOption)
	queryUpdate, args, err := qb.BuildUpdate(&updateParam, &selectParam)
	if err != nil {
		return errors.NewWithCode(codes.CodeSQLBuilder, err.Error())
	}

	tx, err := p.db.Leader().BeginTx(ctx, "txParameterColumnIndex", sql.TxOptions{})
	if err != nil {
		return errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	res, err := tx.Exec("uParameterColumnIndex", updateParameterColumnIndex+queryUpdate, args...)
	if err != nil && strings.Contains(err.Error(), entity.DuplicateEntryErrMessage) {
		return errors.NewWithCode(codes.CodeSQLUniqueConstraint, err.Error())
	} else if err != nil {
		return errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		return errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	} else if rowCount < 1 {
		return errors.NewWithCode(codes.CodeSQLNoRowsAffected, "no parameter column index updated")
	}

	if err := tx.Commit(); err != nil {
		return errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}

	p.log.Debug(ctx, fmt.Sprintf("success update parameter column index %v with body: %v", selectParam.ID, updateParam))

	return nil
}
