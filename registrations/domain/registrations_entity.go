package domain

import "time"

type Registration struct {
	//Description: The id of registration.
	Id string `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	//Description: The date of the creation of the registration.
	CreatedAt   *time.Time  `json:"created_at" example:"2023-11-10 08:10:00"`
	Session     Session     `json:"session" binding:"required"`
	Beneficiary Beneficiary `json:"beneficiary" binding:"required"`
}

type Session struct {
	//Description: the id of session.
	Id string `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	//Description: the start date
	StartDate *time.Time `json:"start_date" example:"2023-11-10 08:10:00"`
	//Description: the end date
	EndDate *time.Time `json:"end_date" example:"2023-11-10 08:10:00"`
	//Description: the date of creation
	CreatedAt *time.Time `json:"created_at" example:"2023-11-10 08:10:00"`
	WorkShop  WorkShop   `json:"work_shop" binding:"required"`
}

type WorkShop struct {
	//Description: the id of workshop.
	Id string `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	//Description: the name of workshop.
	Name string `json:"name" binding:"required" example:"FIRST WORKSHOP"`
}

type Beneficiary struct {
	//Description: the id of beneficiary
	Id
}
