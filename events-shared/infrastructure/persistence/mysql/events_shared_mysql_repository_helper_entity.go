/*
 * File: events_shared_mysql_repository_helper_entity.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file content the helper entity repository.
 *
 * Last Modified: 2026-04-24
 */

package mysql

type EventTotals struct {
	Id        string `db:"event_id"`
	TotalReg  int    `db:"event_total_reg"`
	TotalPay  int    `db:"event_total_pay"`
	TotalPres int    `db:"event_total_pres"`
}

type WorkshopTotals struct {
	Id        string `db:"workshop_id"`
	TotalReg  int    `db:"workshop_total_reg"`
	TotalPay  int    `db:"workshop_total_pay"`
	TotalPres int    `db:"workshop_total_pres"`
}

type SessionTotals struct {
	Id        string `db:"session_id"`
	TotalReg  int    `db:"session_total_reg"`
	TotalPay  int    `db:"session_total_pay"`
	TotalPres int    `db:"session_total_pres"`
}
