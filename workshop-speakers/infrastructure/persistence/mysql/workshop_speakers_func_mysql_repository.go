package mysql

import (
	"context"
	"database/sql"
	_ "embed"

	"github.com/jackskj/carta"
	"github.com/Benjamin-Gthub2/api-shared/db"
	logErrorCoreDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"
	paramsDomain "github.com/Benjamin-Gthub2/api-shared/params/domain"
	"github.com/stroiman/go-automapper"

	workshopSpeakersDomain "github.com/Benjamin-Gthub2/api-event/workshop-speakers/domain"
)

//go:embed sql/get_workshop_speaker_by_id.sql
var QueryGetWorkshopSpeakerById string

//go:embed sql/get_workshop_speakers.sql
var QueryGetWorkshopSpeakers string

//go:embed sql/get_total_workshop_speakers.sql
var QueryGetTotalWorkshopSpeakers string

//go:embed sql/create_workshop_speaker.sql
var QueryCreateWorkshopSpeaker string

//go:embed sql/delete_workshop_speaker.sql
var QueryDeleteWorkshopSpeaker string

func (r workshopSpeakersMySQLRepo) GetWorkshopSpeakerById(
	ctx context.Context,
	workshopSpeakerId string,
) (
	workshopSpeakerById *workshopSpeakersDomain.WorkshopSpeaker,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetWorkshopSpeakerById").SetRaw(err)
	}
	results, err := client.QueryContext(ctx, QueryGetWorkshopSpeakerById, workshopSpeakerId)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetWorkshopSpeakerById").SetRaw(err)
	}
	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &errClose)
		}
	}(results)

	workshopSpeakersTmp := make([]WorkshopSpeaker, 0)
	err = carta.Map(results, &workshopSpeakersTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetWorkshopSpeakerById").SetRaw(err)
	}

	workshopSpeakersAux := make([]workshopSpeakersDomain.WorkshopSpeaker, 0)
	automapper.Map(workshopSpeakersTmp, &workshopSpeakersAux)
	if len(workshopSpeakersAux) == 0 {
		return nil, workshopSpeakersDomain.ErrWorkshopSpeakerNotFound
	}
	return &workshopSpeakersAux[0], nil
}

func (r workshopSpeakersMySQLRepo) GetWorkshopSpeakers(
	ctx context.Context,
	pagination paramsDomain.PaginationParams,
	searchParams workshopSpeakersDomain.GetWorkshopSpeakersParams,
) (
	workshopSpeakersRows []workshopSpeakersDomain.WorkshopSpeaker,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	sizePage := pagination.GetSizePage()
	offset := pagination.GetOffset()

	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetWorkshopSpeakers").SetRaw(err)
	}

	results, err := client.QueryContext(
		ctx,
		QueryGetWorkshopSpeakers,
		searchParams.WorkshopId,
		searchParams.WorkshopId,
		searchParams.SpeakerId,
		searchParams.SpeakerId,
		sizePage,
		offset,
	)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetWorkshopSpeakers").SetRaw(err)
	}
	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &errClose)
		}
	}(results)

	workshopSpeakersTmp := make([]WorkshopSpeaker, 0)
	err = carta.Map(results, &workshopSpeakersTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetWorkshopSpeakers").SetRaw(err)
	}

	var workshopSpeakers = make([]workshopSpeakersDomain.WorkshopSpeaker, 0)
	automapper.Map(workshopSpeakersTmp, &workshopSpeakers)
	return workshopSpeakers, nil
}

func (r workshopSpeakersMySQLRepo) GetTotalWorkshopSpeakers(
	ctx context.Context,
	searchParams workshopSpeakersDomain.GetWorkshopSpeakersParams,
) (
	total *int,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var totalTmp int
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetTotalWorkshopSpeakers").SetRaw(err)
	}
	err = client.QueryRowContext(
		ctx,
		QueryGetTotalWorkshopSpeakers,
		searchParams.WorkshopId,
		searchParams.WorkshopId,
		searchParams.SpeakerId,
		searchParams.SpeakerId,
	).Scan(&totalTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetTotalWorkshopSpeakers").SetRaw(err)
	}
	total = &totalTmp
	return total, nil
}

func (r workshopSpeakersMySQLRepo) CreateWorkshopSpeaker(
	ctx context.Context,
	body workshopSpeakersDomain.CreateWorkshopSpeaker,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	now := r.clock.Now().Format("2006-01-02 15:04:05")
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("CreateWorkshopSpeaker").SetRaw(err)
	}
	_, err = client.ExecContext(ctx,
		QueryCreateWorkshopSpeaker,
		body.Id,
		body.WorkshopId,
		body.DegreeAbbreviation,
		body.SpeakerId,
		body.CreatedBy,
		now,
	)
	if err != nil {
		return r.err.Clone().SetFunction("CreateWorkshopSpeaker").SetRaw(err)
	}
	return
}

func (r workshopSpeakersMySQLRepo) DeleteWorkshopSpeaker(
	ctx context.Context,
	body workshopSpeakersDomain.DeleteWorkshopSpeaker,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	now := r.clock.Now().Format("2006-01-02 15:04:05")
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("DeleteWorkshopSpeaker").SetRaw(err)
	}
	_, err = client.ExecContext(ctx,
		QueryDeleteWorkshopSpeaker,
		now,
		body.DeletedBy,
		body.Id,
	)
	if err != nil {
		return r.err.Clone().SetFunction("DeleteWorkshopSpeaker").SetRaw(err)
	}
	return
}
