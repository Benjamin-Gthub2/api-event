/*
 * File: validations_mysql_repository.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the functions of the validations.
 *
 * Last Modified: 2023-11-10
 */

package mysql

import (
	"context"
	"fmt"
	"time"

	"github.com/Benjamin-Gthub2/api-shared/db"
	errDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"

	"github.com/Benjamin-Gthub2/api-shared/validations/domain"
)

type validationsMySQLRepo struct {
	timeout time.Duration
}

func NewValidationsRepository(mongoTimeout int) domain.ValidationRepository {
	rep := &validationsMySQLRepo{
		timeout: time.Duration(mongoTimeout) * time.Second,
	}
	return rep
}

func (r validationsMySQLRepo) RecordExists(
	ctx context.Context,
	params domain.RecordExistsParams,
) (bool, error) {
	var exists int
	var query string
	var args []interface{}

	query = fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE %s = ?", params.Table, params.IdColumnName)
	args = append(args, params.IdValue)

	if params.StatusColumnName != nil && params.StatusValue != nil {
		query += fmt.Sprintf(" AND %s = ?", *params.StatusColumnName)
		args = append(args, *params.StatusValue)
	}

	if params.StatusColumnName != nil && params.StatusValue == nil {
		query += fmt.Sprintf(" AND %s is null", *params.StatusColumnName)
	}

	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return false, errDomain.NewErr().
			SetRaw(err).
			SetLayer(errDomain.Infra).
			SetFunction("ValidateExistence")
	}

	err = client.QueryRowContext(ctx, query, args...).Scan(&exists)
	if err != nil {
		return false, errDomain.NewErr().
			SetRaw(err).
			SetLayer(errDomain.Infra).
			SetFunction("RecordExists")
	}
	if exists == 0 {
		return false, nil
	}

	return true, nil
}

func (r validationsMySQLRepo) ValidateExistence(
	ctx context.Context,
	params domain.RecordExistsParams,
) (bool, error) {
	var exists int
	var query string
	var args []interface{}

	query = fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE %s = ?", params.Table, params.IdColumnName)
	args = append(args, params.IdValue)

	if params.StatusColumnName != nil && params.StatusValue != nil {
		query += fmt.Sprintf(" AND %s = ?", *params.StatusColumnName)
		args = append(args, *params.StatusValue)
	}

	if params.StatusColumnName != nil && params.StatusValue == nil {
		query += fmt.Sprintf(" AND %s is null", *params.StatusColumnName)
	}

	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return false, errDomain.NewErr().
			SetRaw(err).
			SetLayer(errDomain.Infra).
			SetFunction("ValidateExistence")
	}

	err = client.QueryRowContext(ctx, query, args...).Scan(&exists)
	if err != nil {
		return false, errDomain.NewErr().
			SetRaw(err).
			SetLayer(errDomain.Infra).
			SetFunction("ValidateExistence")
	}
	if exists != 0 {
		return true, nil
	}
	return false, nil
}

func (r validationsMySQLRepo) ValidateUniqueField(
	ctx context.Context,
	params domain.ValidateUniqueFieldParams,
) (
	bool, error,
) {
	var count int
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE %s = ? AND %s != ?", params.Table, params.ColumnName, params.RecordIdName)
	args := []interface{}{params.Value, params.RecordIdValue}

	if params.StatusColumnName != nil && params.StatusValue == nil {
		query += fmt.Sprintf(" AND %s is null", *params.StatusColumnName)
	}

	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return false, errDomain.NewErr().
			SetRaw(err).
			SetLayer(errDomain.Infra).
			SetFunction("ValidateUniqueField")
	}
	err = client.QueryRowContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return false, errDomain.NewErr().
			SetRaw(err).
			SetLayer(errDomain.Infra).
			SetFunction("ValidateUniqueField")
	}

	if count > 0 {
		return false, nil
	}
	return true, nil
}
