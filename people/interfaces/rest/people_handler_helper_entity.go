/*
 * File: people_handler_helper_entity.go
 * Author: lady
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is the entry point for the application.
 *
 * Last Modified: 2023-12-12
 */

package rest

import (
	paginationDomain "github.com/Benjamin-Gthub2/api-shared/params/domain"

	"github.com/Benjamin-Gthub2/api-event/people/domain"
)

type peopleResult struct {
	Data       []domain.People                    `json:"data" binding:"required"`
	Pagination paginationDomain.PaginationResults `json:"pagination" binding:"required"`
	Status     int                                `json:"status" binding:"required"`
}

type deletePersonResult struct {
	Data   bool `json:"data" binding:"required"`
	Status int  `json:"status" binding:"required"`
}
