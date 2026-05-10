/*
 * File: registrations_handler_helper_entity.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file content the helper entity in the handler.
 *
 * Last Modified: 2026-04-21
 */

package rest

type createRegistrationValidated struct {
	EventId       string `json:"event_id" binding:"required" example:"200bbbc9-7e93-11ee-89fd-0242ac110016"`
	BeneficiaryId string `json:"beneficiary_id" binding:"required" example:"200bbbc9-7e93-11ee-89fd-0242ac110016"`
}
