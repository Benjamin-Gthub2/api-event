/*
 * File: people_usecase.go
 * Author: lady
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is the entry point for the application.
 *
 * Last Modified: 2023-12-12
 */

package domain

import (
	"context"

	paramsDomain "github.com/smart0n3/api-shared/params/domain"
)

type PeopleUseCase interface {
	GetPeople(ctx context.Context, searchParams GetPeopleParams, pagination paramsDomain.PaginationParams) ([]People,
		*paramsDomain.PaginationResults, error)
	CreatePerson(ctx context.Context, body CreatePersonBody) (*string, error)
	UpdatePerson(ctx context.Context, peopleId string, body UpdatePersonBody) error
	DeletePerson(ctx context.Context, peopleId string) (bool, error)
}
