package rest

import (
	paginationDomain "github.com/smart0n3/api-shared/params/domain"

	attendancesDomain "github.com/Benjamin-Gthub2/api-event/attendances/domain"
)

type attendanceByIdResult struct {
	Data   *attendancesDomain.Attendance `json:"data"`
	Status int                           `json:"status" binding:"required"`
}

type attendancesResult struct {
	Data       []attendancesDomain.Attendance     `json:"data"`
	Pagination paginationDomain.PaginationResults `json:"pagination" binding:"required"`
	Status     int                                `json:"status" binding:"required"`
}
