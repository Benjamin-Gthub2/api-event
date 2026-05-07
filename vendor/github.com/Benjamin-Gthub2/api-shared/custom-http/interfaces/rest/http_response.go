/*
 * File: http_handler.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the http generic responses handler for the application.
 *
 * Last Modified: 2023-12-19
 */

package httpResponse

type StatusResult struct {
	Status int `json:"status" binding:"required" example:"200"`
}

type IdResult struct {
	Data   string `json:"data" binding:"required" example:"201"`
	Status int    `json:"status" binding:"required"`
}

type IdsResult struct {
	Data   []string `json:"data" binding:"required" example:"201"`
	Status int      `json:"status" binding:"required"`
}

type IntIdResponse struct {
	Id int `json:"id" binding:"required"`
}

type BoolResponse struct {
	Data bool `json:"data" binding:"required"`
}

type StringIdResponse struct {
	Id string `json:"id" binding:"required"`
}

type DataResultStringID struct {
	Data   StringIdResponse `json:"data" binding:"required"`
	Status int              `json:"status" binding:"required"`
}

type DataResultIntId struct {
	Data   IntIdResponse `json:"data" binding:"required"`
	Status int           `json:"status" binding:"required"`
}
