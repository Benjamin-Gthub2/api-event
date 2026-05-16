/*
 * File: people_func_mysql_repository.go
 * Author: lady
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is the entry point for the application.
 *
 * Last Modified: 2023-12-12
 */

package mysql

import (
	"context"
	"database/sql"
	_ "embed"
	"time"

	"github.com/jackskj/carta"
	"github.com/Benjamin-Gthub2/api-shared/db"
	"github.com/stroiman/go-automapper"

	logErrorCoreDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"
	paramsDomain "github.com/Benjamin-Gthub2/api-shared/params/domain"

	"github.com/Benjamin-Gthub2/api-event/people/domain"
)

var limaLoc = time.FixedZone("America/Lima", -5*60*60)

//go:embed sql/get_people.sql
var QueryGetPeople string

//go:embed sql/get_total_people.sql
var QueryGetTotalPeople string

//go:embed sql/create_person.sql
var QueryCreatePerson string

//go:embed sql/update_person.sql
var QueryUpdatePerson string

//go:embed sql/delete_person.sql
var QueryDeletePerson string

//go:embed sql/get_people_by_document.sql
var QueryGetPeopleByDocument string

func (r peopleMySQLRepo) GetPeople(
	ctx context.Context,
	searchParams domain.GetPeopleParams,
	pagination paramsDomain.PaginationParams,
) (
	peopleRows []domain.People,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	sizePage := pagination.GetSizePage()
	offset := pagination.GetOffset()
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetAccounts").SetRaw(err)
	}
	results, err := client.
		QueryContext(
			ctx,
			QueryGetPeople,
			searchParams.SearchName,
			searchParams.SearchName,
			searchParams.SearchName,
			searchParams.SearchName,
			searchParams.DocumentTypeId,
			searchParams.DocumentTypeId,
			searchParams.Document,
			searchParams.Document,
			searchParams.SearchValue, searchParams.SearchValue, searchParams.SearchValue, searchParams.SearchValue, searchParams.SearchValue,
			sizePage,
			offset,
		)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetPeople").SetRaw(err)
	}
	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &errClose)
		}
	}(results)
	peopleTmp := make([]People, 0)
	err = carta.Map(results, &peopleTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetPeople").SetRaw(err)
	}
	automapper.Map(peopleTmp, &peopleRows)
	for iPerson, person := range peopleRows {
		if person.User.Id == nil {
			peopleRows[iPerson].User = nil
		}
	}
	return peopleRows, nil
}

func (r peopleMySQLRepo) GetTotalPeople(
	ctx context.Context,
	searchParams domain.GetPeopleParams,
	pagination paramsDomain.PaginationParams,
) (
	total *int,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var totalTmp int
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetAccounts").SetRaw(err)
	}
	err = client.
		QueryRowContext(
			ctx,
			QueryGetTotalPeople,
			searchParams.SearchName,
			searchParams.SearchName,
			searchParams.SearchName,
			searchParams.SearchName,
			searchParams.DocumentTypeId,
			searchParams.DocumentTypeId,
			searchParams.Document,
			searchParams.Document,
			searchParams.SearchValue, searchParams.SearchValue, searchParams.SearchValue, searchParams.SearchValue, searchParams.SearchValue,
		).
		Scan(&totalTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetTotalPeople").SetRaw(err)
	}
	total = &totalTmp
	return total, nil
}

func (r peopleMySQLRepo) CreatePerson(
	ctx context.Context,
	peopleId string,
	body domain.CreatePersonBody,
) (
	lastId *string,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	now := r.clock.Now().In(limaLoc).Format("2006-01-02 15:04:05")
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetAccounts").SetRaw(err)
	}
	_, err = client.ExecContext(ctx,
		QueryCreatePerson,
		peopleId,
		body.UserId,
		body.TypeDocumentId,
		body.Document,
		body.Names,
		body.Surname,
		body.LastName,
		body.Phone,
		body.Email,
		body.Gender,
		body.Enable,
		now)
	if err != nil {
		return nil, r.err.Clone().SetFunction("CreatePerson").SetRaw(err)
	}
	lastId = &peopleId
	return
}

func (r peopleMySQLRepo) UpdatePerson(
	ctx context.Context,
	peopleId string,
	body domain.UpdatePersonBody,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("GetAccounts").SetRaw(err)
	}
	_, err = client.ExecContext(
		ctx,
		QueryUpdatePerson,
		body.UserId,
		body.TypeDocumentId,
		body.Document,
		body.Names,
		body.Surname,
		body.LastName,
		body.Phone,
		body.Email,
		body.Gender,
		body.Enable,
		peopleId,
	)
	if err != nil {
		return r.err.Clone().SetFunction("UpdatePerson").SetRaw(err)
	}
	return
}

func (r peopleMySQLRepo) DeletePerson(
	ctx context.Context,
	peopleId string,
) (
	updated bool,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	now := r.clock.Now().In(limaLoc).Format("2006-01-02 15:04:05")
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return false, r.err.Clone().SetFunction("GetAccounts").SetRaw(err)
	}
	_, err = client.ExecContext(
		ctx,
		QueryDeletePerson,
		now,
		peopleId)
	if err != nil {
		return false, r.err.Clone().SetFunction("DeletePerson").SetRaw(err)
	}
	return true, nil
}

func (r peopleMySQLRepo) GetPeopleByDocument(
	ctx context.Context,
	document string,
) (
	peopleRows []domain.People,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetPeopleByDocument").SetRaw(err)
	}
	results, err := client.
		QueryContext(
			ctx,
			QueryGetPeopleByDocument,
			document,
		)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetPeopleByDocument").SetRaw(err)
	}
	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &errClose)
		}
	}(results)
	peopleTmp := make([]People, 0)
	err = carta.Map(results, &peopleTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetPeopleByDocument").SetRaw(err)
	}
	automapper.Map(peopleTmp, &peopleRows)
	for iPerson, person := range peopleRows {
		if person.User.Id == nil {
			peopleRows[iPerson].User = nil
		}
	}
	return peopleRows, nil
}
