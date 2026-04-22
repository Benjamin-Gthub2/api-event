/*
 * File: people_func_handler.go
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
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	restCore "github.com/smart0n3/api-shared/api-core/interfaces/rest"
	httpResponse "github.com/smart0n3/api-shared/custom-http/interfaces/rest"
	paramsDomain "github.com/smart0n3/api-shared/params/domain"

	"github.com/Benjamin-Gthub2/api-event/people/domain"
)

// GetPeople is a method to get people
// @Summary get people
// @Description get people
// @Tags People
// @Accept json
// @Produce json
// @Param search_name query string false "the document number of the people"
// @Param document query string false "the document number of the people"
// @Param document_type_id query string false "the document type id of the people"
// @Success 200 {object} peopleResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/people/ [get]
// @Security BearerAuth
func (h peopleHandler) GetPeople(c *gin.Context) {
	ctx := c.Request.Context()

	searchParams := domain.GetPeopleParams{}
	searchParams.QueryParamsToStruct(c.Request, &searchParams)
	pagination := paramsDomain.NewPaginationParams(c.Request)

	people, paginationRes, err := h.peopleUseCase.GetPeople(ctx, searchParams, pagination)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := peopleResult{
		Data:       people,
		Pagination: *paginationRes,
		Status:     http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// CreatePerson is a method to create a person
// @Summary Create a person
// @Description Create a person
// @Tags People
// @Accept json
// @Produce json
// @Param createPersonBody body domain.CreatePersonBody true "Create body of person"
// @Success 201 {object} httpResponse.IdResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/people [post]
// @Security BearerAuth
func (h peopleHandler) CreatePerson(c *gin.Context) {
	ctx := c.Request.Context()

	var personValidate createPersonValidate
	if err := c.ShouldBindJSON(&personValidate); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("CreatePerson").SetRaw(errors.New(
				"casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}
		messagesErr := make([]string, 0)

		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("CreatePerson").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	createPersonBody := domain.CreatePersonBody{
		UserId:         personValidate.UserId,
		TypeDocumentId: personValidate.TypeDocumentId,
		Document:       personValidate.Document,
		Names:          personValidate.Names,
		Surname:        personValidate.Surname,
		LastName:       personValidate.LastName,
		Phone:          personValidate.Phone,
		Email:          personValidate.Email,
		Gender:         personValidate.Gender,
		Enable:         personValidate.Enable,
	}
	id, err := h.peopleUseCase.CreatePerson(ctx, createPersonBody)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := httpResponse.IdResult{
		Data:   *id,
		Status: http.StatusCreated,
	}
	restCore.Json(c, http.StatusCreated, res)
}

// UpdatePerson is a method to update a person
// @Summary Update a person
// @Description Update a person
// @Tags People
// @Accept json
// @Produce json
// @Param personId path string true "person id"
// @Param updatePersonBody body domain.UpdatePersonBody true "Update person body"
// @Success 200 {object} httpResponse.StatusResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/people/{personId} [put]
// @Security BearerAuth
func (h peopleHandler) UpdatePerson(c *gin.Context) {
	ctx := c.Request.Context()
	personId := c.Param("personId")

	var personValidate UpdatePersonValidate
	if err := c.ShouldBindJSON(&personValidate); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("UpdatePerson").SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}

		messagesErr := make([]string, 0)
		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("UpdatePerson").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	updatePersonBody := domain.UpdatePersonBody{
		UserId:         personValidate.UserId,
		TypeDocumentId: personValidate.TypeDocumentId,
		Document:       personValidate.Document,
		Names:          personValidate.Names,
		Surname:        personValidate.Surname,
		LastName:       personValidate.LastName,
		Phone:          personValidate.Phone,
		Email:          personValidate.Email,
		Gender:         personValidate.Gender,
		Enable:         personValidate.Enable,
	}
	err := h.peopleUseCase.UpdatePerson(ctx, personId, updatePersonBody)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := httpResponse.StatusResult{
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// DeletePerson is a method to delete a person
// @Summary Delete a person
// @Description Delete a person
// @Tags People
// @Accept json
// @Produce json
// @Param personId path string true "person id"
// @Success 200 {object} deletePersonResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/people/{personId} [delete]
// @Security BearerAuth
func (h peopleHandler) DeletePerson(c *gin.Context) {
	ctx := c.Request.Context()
	personId := c.Param("personId")
	result, err := h.peopleUseCase.DeletePerson(ctx, personId)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}
	res := deletePersonResult{
		Data:   result,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}
