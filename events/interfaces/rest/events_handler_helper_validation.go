/*
 * File: events_handler_helper_validation.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file defines validation structures for events related data.
 *
 * Last Modified: 2026-04-15
 */

package rest

type createMerchantsValidate struct {
	Name        string  `json:"name" binding:"required" example:"Odin Corp"`
	Description string  `json:"description" binding:"required" example:"Proveedor de servicios de mantenimiento"`
	Code        string  `json:"code" binding:"required" example:"CODE1"`
	Phone       *string `json:"phone" example:"+1234567890"`
	Document    string  `json:"document" binding:"required" example:"123456789"`
	Address     string  `json:"address" binding:"required" example:"123 Main Street"`
	Industry    string  `json:"industry" binding:"required" example:"Mantenimiento"`
	Enable      bool    `json:"enable" example:"true"`
}

type enableDisableMerchantValidate struct {
	Enable *bool `json:"enable" example:"true"`
}
