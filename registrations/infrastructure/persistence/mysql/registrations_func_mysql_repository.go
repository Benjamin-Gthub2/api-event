/*
 * File: registrations_func_mysql_repository.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the repository functions of the registrations.
 *
 * Last Modified: 2026-04-21
 */

package mysql

import (
	"context"
	"database/sql"
	_ "embed"
	"time"

	"github.com/Benjamin-Gthub2/api-shared/db"
	"github.com/jackskj/carta"
	"github.com/stroiman/go-automapper"

	logErrorCoreDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"
	paramsDomain "github.com/Benjamin-Gthub2/api-shared/params/domain"

	registrationsDomain "github.com/Benjamin-Gthub2/api-event/registrations/domain"
)

var limaLoc = time.FixedZone("America/Lima", -5*60*60)

//go:embed sql/get_registration_by_id.sql
var QueryGetRegistrationById string

//go:embed sql/get_registrations.sql
var QueryGetRegistrations string

//go:embed sql/get_total_registrations.sql
var QueryGetTotalRegistrations string

//go:embed sql/create_registration.sql
var QueryCreateRegistration string

//go:embed sql/update_registration_status.sql
var QueryUpdateRegistrationStatus string

//go:embed sql/update_registration_send_qr.sql
var QueryUpdateRegistrationSendQr string

func intToPtr(value int) *int {
	return &value
}

func (r registrationsMySQLRepo) GetRegistrationById(
	ctx context.Context,
	registrationId string,
) (
	registrationById *registrationsDomain.Registration,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetRegistrationById").SetRaw(err)
	}
	results, err := client.
		QueryContext(
			ctx,
			QueryGetRegistrationById,
			registrationId,
		)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetRegistrationById").SetRaw(err)
	}
	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &errClose)
		}
	}(results)

	RegistrationByIdTmp := make([]Registration, 0)
	err = carta.Map(results, &RegistrationByIdTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetRegistrationById").SetRaw(err)
	}

	registrationAux := make([]registrationsDomain.Registration, 0)
	automapper.Map(RegistrationByIdTmp, &registrationAux)
	if len(registrationAux) == 0 {
		return nil, registrationsDomain.ErrRegistrationsNotFound
	}

	return &registrationAux[0], nil
}

func (r registrationsMySQLRepo) GetRegistrations(
	ctx context.Context,
	pagination paramsDomain.PaginationParams,
	searchParams registrationsDomain.GetRegistrationsParams,
) (
	registrationsRows []registrationsDomain.Registration,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	sizePage := pagination.GetSizePage()
	offset := pagination.GetOffset()

	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetRegistrations").SetRaw(err)
	}

	results, err := client.QueryContext(
		ctx,
		QueryGetRegistrations,
		searchParams.StartDate,
		searchParams.StartDate,
		searchParams.EndDate,
		searchParams.CreatedBy,
		searchParams.CreatedBy,
		searchParams.SearchValue, searchParams.SearchValue, searchParams.SearchValue, searchParams.SearchValue, searchParams.SearchValue, searchParams.SearchValue,
		sizePage,
		offset,
	)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetRegistrations").SetRaw(err)
	}

	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &err)
		}
	}(results)

	registrationsTmp := make([]Registration, 0)
	err = carta.Map(results, &registrationsTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetRegistrations").SetRaw(err)
	}
	var registrations = make([]registrationsDomain.Registration, 0)
	automapper.Map(registrationsTmp, &registrations)

	return registrations, nil
}

func (r registrationsMySQLRepo) GetTotalRegistrations(
	ctx context.Context,
	searchParams registrationsDomain.GetRegistrationsParams,
) (
	total *int,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var totalTmp int
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetTotalRegistrations").SetRaw(err)
	}
	err = client.
		QueryRowContext(
			ctx,
			QueryGetTotalRegistrations,
			searchParams.StartDate,
			searchParams.StartDate,
			searchParams.EndDate,
			searchParams.CreatedBy,
			searchParams.CreatedBy,
			searchParams.SearchValue, searchParams.SearchValue, searchParams.SearchValue, searchParams.SearchValue, searchParams.SearchValue, searchParams.SearchValue,
		).Scan(&totalTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetTotalRegistrations").SetRaw(err)
	}
	total = &totalTmp
	return total, nil
}

func (r registrationsMySQLRepo) MainCreateRegistration(
	ctx context.Context,
	body registrationsDomain.CreateRegistration,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var tx *sql.Tx

	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("MainCreateRegistration").SetRaw(err)
	}
	tx, err = client.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	//aqui no ira esto,s era en attendances
	//var sessionWorkshopEventById *eventSharedDomain.EventWorkshopSession
	//sessionWorkshopEventById, err = r.eventsSharedRepository.GetSessionWorkshopEventById(ctx, tx, body.SessionId)
	//if err != nil {
	//	return err
	//}
	//
	//eventId := sessionWorkshopEventById.EventId
	//workshopId := sessionWorkshopEventById.WorkshopId
	//sessionId := sessionWorkshopEventById.SessionId
	//
	//var eventTotals *eventSharedDomain.EventTotals
	//var workshopTotals *eventSharedDomain.WorkshopTotals
	//var sessionTotals *eventSharedDomain.SessionTotals
	//var updateEventTotals eventSharedDomain.UpdateEventTotals
	//var updateWorkshopTotals eventSharedDomain.UpdateWorkshopTotals
	//var updateSessionTotals eventSharedDomain.UpdateSessionTotals
	//
	//eventTotals, err = r.eventsSharedRepository.GetEventTotals(ctx, tx, eventId)
	//if err != nil {
	//	return err
	//}
	//updateEventTotals = eventSharedDomain.UpdateEventTotals{
	//	TotalReg:  intToPtr(eventTotals.TotalReg + 1),
	//	TotalPay:  intToPtr(eventTotals.TotalPay + 1),
	//	TotalPres: intToPtr(eventTotals.TotalPres + 1),
	//}
	//err = r.eventsSharedRepository.UpdateEventTotals(ctx, tx, eventId, updateEventTotals)
	//if err != nil {
	//	return err
	//}
	//
	//workshopTotals, err = r.eventsSharedRepository.GetWorkshopTotals(ctx, tx, workshopId)
	//if err != nil {
	//	return err
	//}
	//updateWorkshopTotals = eventSharedDomain.UpdateWorkshopTotals{
	//	TotalReg:  intToPtr(workshopTotals.TotalReg + 1),
	//	TotalPay:  intToPtr(workshopTotals.TotalPay + 1),
	//	TotalPres: intToPtr(workshopTotals.TotalPres + 1),
	//}
	//err = r.eventsSharedRepository.UpdateWorkshopTotals(ctx, tx, workshopId, updateWorkshopTotals)
	//if err != nil {
	//	return err
	//}
	//
	//sessionTotals, err = r.eventsSharedRepository.GetSessionTotals(ctx, tx, sessionId)
	//if err != nil {
	//	return err
	//}
	//updateSessionTotals = eventSharedDomain.UpdateSessionTotals{
	//	TotalReg:  intToPtr(sessionTotals.TotalReg + 1),
	//	TotalPay:  intToPtr(sessionTotals.TotalPay + 1),
	//	TotalPres: intToPtr(sessionTotals.TotalPres + 1),
	//}
	//err = r.eventsSharedRepository.UpdateSessionTotals(ctx, tx, sessionId, updateSessionTotals)
	//if err != nil {
	//	return err
	//}

	err = r.CreateRegistration(ctx, tx, body)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}
	return err
}

func (r registrationsMySQLRepo) CreateRegistration(
	ctx context.Context,
	tx *sql.Tx,
	body registrationsDomain.CreateRegistration,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	now := r.clock.Now().In(limaLoc).Format("2006-01-02 15:04:05")
	_, err = tx.ExecContext(ctx,
		QueryCreateRegistration,
		body.Id,
		body.StatusId,
		body.EventId,
		body.BeneficiaryId,
		body.CreatedBy,
		now,
	)
	if err != nil {
		return r.err.Clone().SetFunction("CreateRegistration").SetRaw(err)
	}
	return
}

func (r registrationsMySQLRepo) UpdateRegistrationStatus(
	ctx context.Context,
	registrationId string,
	statusCode string,
) (err error) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("UpdateRegistrationStatus").SetRaw(err)
	}
	_, err = client.ExecContext(ctx, QueryUpdateRegistrationStatus, statusCode, registrationId)
	if err != nil {
		return r.err.Clone().SetFunction("UpdateRegistrationStatus").SetRaw(err)
	}
	return
}

func (r registrationsMySQLRepo) UpdateRegistrationSendQr(
	ctx context.Context,
	registrationId string,
) (err error) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("UpdateRegistrationSendQr").SetRaw(err)
	}
	_, err = client.ExecContext(ctx, QueryUpdateRegistrationSendQr, registrationId)
	if err != nil {
		return r.err.Clone().SetFunction("UpdateRegistrationSendQr").SetRaw(err)
	}
	return
}
