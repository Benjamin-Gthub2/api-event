/*
 * File: people_repository.go
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

	paramsDomain "github.com/Benjamin-Gthub2/api-shared/params/domain"
)

type PeopleRepository interface {
	GetPeople(ctx context.Context, searchParams GetPeopleParams, pagination paramsDomain.PaginationParams) (
		[]People, error)
	GetTotalPeople(ctx context.Context, searchParams GetPeopleParams,
		pagination paramsDomain.PaginationParams) (*int, error)
	CreatePerson(ctx context.Context, peopleId string, body CreatePersonBody) (*string, error)
	UpdatePerson(ctx context.Context, peopleId string, body UpdatePersonBody) error
	DeletePerson(ctx context.Context, peopleId string) (bool, error)
	GetPeopleByDocument(ctx context.Context, document string) ([]People, error)
}
