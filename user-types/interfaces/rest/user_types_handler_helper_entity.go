/*
 * File: user_types_handler_helper_entity.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Entities helper to handler for user types.
 *
 * Last Modified: 2023-11-23
 */

package rest

import (
	paginationDomain "github.com/Benjamin-Gthub2/api-shared/params/domain"

	userTypesDomain "github.com/Benjamin-Gthub2/api-event/user-types/domain"
)

type userTypesResult struct {
	Data       []userTypesDomain.UserType         `json:"data" binding:"required"`
	Pagination paginationDomain.PaginationResults `json:"pagination" binding:"required"`
	Status     int                                `json:"status" binding:"required"`
}

type deleteUserTypesResult struct {
	Data   bool `json:"data" binding:"required"`
	Status int  `json:"status" binding:"required"`
}
