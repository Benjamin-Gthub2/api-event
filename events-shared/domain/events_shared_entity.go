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
	//Description: registrations total of workshop
	TotalReg *int `json:"total_reg" example:"6"`
	//Description: payments total of workshop
	TotalPay *int `json:"total_pay" example:"2"`
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
	//Description: registrations total of event
	TotalReg *int `json:"total_reg" example:"6"`
	//Description: payments total of event
	TotalPay *int `json:"total_pay" example:"2"`
	//Description: presence total of event
	TotalPres *int `json:"total_pres" example:"3"`
}

type WorkshopTotals struct {
	//Description: registrations total of event
	TotalReg *int `json:"total_reg" example:"6"`
	//Description: payments total of event
	TotalPay *int `json:"total_pay" example:"2"`
	//Description: presence total of event
	TotalPres *int `json:"total_pres" example:"3"`
}

type SessionTotals struct {
	//Description: registrations total of event
	TotalReg *int `json:"total_reg" example:"6"`
	//Description: payments total of event
	TotalPay *int `json:"total_pay" example:"2"`
	//Description: presence total of event
	TotalPres *int `json:"total_pres" example:"3"`
}
