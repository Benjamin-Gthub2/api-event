/*
 * File: events_shared_entity.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file content the entity in the domain.
 *
 * Last Modified: 2026-04-24
 */

package domain

type UpdateEventTotals struct {
	//Description: registrations total of event
	TotalReg *int `json:"total_reg" example:"6"`
	//Description: payments total of event
	TotalPay *int `json:"total_pay" example:"2"`
	//Description: presence total of event
	TotalPres *int `json:"total_pres" example:"3"`
}

type UpdateWorkshopTotals struct {
	//Description: presence total of workshop
	TotalPres *int `json:"total_pres" example:"3"`
}

type UpdateSessionTotals struct {
	//Description: registrations total of session
	TotalReg *int `json:"total_reg" example:"6"`
	//Description: payments total of session
	TotalPay *int `json:"total_pay" example:"2"`
	//Description: presence total of session
	TotalPres *int `json:"total_pres" example:"3"`
}

type EventTotals struct {
	//DescriptionL the id of event
	Id string `json:"status_id" binding:"required" example:"219bbbc9-7e93-11ee-89fd-0242ac110016"`
	//Description: registrations total of event
	TotalReg int `json:"total_reg" binding:"required" example:"6"`
	//Description: payments total of event
	TotalPay int `json:"total_pay" binding:"required" example:"2"`
	//Description: presence total of event
	TotalPres int `json:"total_pres" binding:"required" example:"3"`
}

type WorkshopTotals struct {
	//DescriptionL the id of workshop
	Id string `json:"id" binding:"required" example:"219bbbc9-7e93-11ee-89fd-0242ac110016"`
	//Description: presence total of event
	TotalPres int `json:"total_pres" binding:"required" example:"3"`
}

type SessionTotals struct {
	//DescriptionL the id of session
	Id string `json:"id" binding:"required" example:"219bbbc9-7e93-11ee-89fd-0242ac110016"`
	//Description: registrations total of event
	TotalReg int `json:"total_reg" binding:"required" example:"6"`
	//Description: payments total of event
	TotalPay int `json:"total_pay" binding:"required" example:"2"`
	//Description: presence total of event
	TotalPres int `json:"total_pres" binding:"required" example:"3"`
}

type EventWorkshopSession struct {
	//Description: the id of workshop
	WorkshopId string `json:"workshop_id" binding:"required" example:"1"`
	//Description: the id of event
	EventId string `json:"event_id" binding:"required" example:"1"`
}
