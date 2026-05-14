/*
 * File: users_func_handler.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Implementation to handlers to users.
 *
 * Last Modified: 2023-11-23
 */

package rest

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	restCore "github.com/Benjamin-Gthub2/api-shared/api-core/interfaces/rest"
	httpResponse "github.com/Benjamin-Gthub2/api-shared/custom-http/interfaces/rest"
	_ "github.com/Benjamin-Gthub2/api-shared/error-core/domain"
	paramsDomain "github.com/Benjamin-Gthub2/api-shared/params/domain"

	usersDomain "github.com/Benjamin-Gthub2/api-event/users/domain"
)

// GetUser is a method to get user by id
// @Summary get user
// @Description get user by id
// @Tags Users
// @Accept json
// @Produce json
// @Param userId path string true "user id"
// @Success 200 {object} userResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/users/{userId} [get]
// @Security BearerAuth
func (h usersHandler) GetUser(c *gin.Context) {
	ctx := c.Request.Context()
	userId := c.Param("userId")

	user, err := h.usersUseCase.GetUser(ctx, userId)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}
	res := userResult{
		Data:   *user,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// GetUsers is a method to get users
// @Summary get users
// @Description get users
// @Tags Users
// @Accept json
// @Produce json
// @Param type_id query string false "the user type id"
// @Param username query string false "the username of the user"
// @Param role_id query string false "the role id of the user"
// @Success 200 {object} multipleUsersResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/users/ [get]
// @Security BearerAuth
func (h usersHandler) GetUsers(c *gin.Context) {
	ctx := c.Request.Context()

	searchParams := usersDomain.GetUsersParams{}
	searchParams.QueryParamsToStruct(c.Request, &searchParams)
	pagination := paramsDomain.NewPaginationParams(c.Request)

	users, paginationRes, err := h.usersUseCase.GetUsers(ctx, searchParams, pagination)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := multipleUsersResult{
		Data:       users,
		Pagination: *paginationRes,
		Status:     http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// GetMenuByUser is a method to get menu by user
// @Summary get menu
// @Description get menu by user
// @Tags Users
// @Accept json
// @Produce json
// @Param userId path string true "user id"
// @Success 200 {object} menuByUserResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/users/{userId}/menu [get]
// @Security BearerAuth
func (h usersHandler) GetMenuByUser(c *gin.Context) {
	ctx := c.Request.Context()
	userId := c.Param("userId")

	menu, err := h.usersUseCase.GetMenuByUser(ctx, userId)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := menuByUserResult{
		Data:   menu,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// GetMenuByUserToken is a method to get menu by user
// @Summary Get menu by user using their token
// @Description Get menu by user using their token
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} menuByUserResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/users/menu [get]
// @Security BearerAuth
func (h usersHandler) GetMenuByUserToken(c *gin.Context) {
	ctx := c.Request.Context()
	userId := c.GetString("userId")

	menu, err := h.usersUseCase.GetMenuByUser(ctx, userId)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := menuByUserResult{
		Data:   menu,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// GetMeByUser is a method to get user me
// @Summary Get user me using their token
// @Description Get user me using their token
// @Tags Users
// @Accept json
// @Produce json
// @Param userId path string true "user id"
// @Success 200 {object} GetMeByUser "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/users/me [get]
// @Security BearerAuth
func (h usersHandler) GetMeByUser(c *gin.Context) {
	ctx := c.Request.Context()
	userId := c.GetString("userId")

	userMe, err := h.usersUseCase.GetMeByUser(ctx, userId)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := GetMeByUser{
		Data:   *userMe,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// CreateUser is a method to create a user
// @Summary Create a user
// @Description Create a user
// @Tags Users
// @Accept json
// @Produce json
// @Param createUserBody body usersDomain.CreateUserBody true "Create user body"
// @Success 201 {object} httpResponse.IdResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/users [post]
// @Security BearerAuth
func (h usersHandler) CreateUser(c *gin.Context) {
	ctx := c.Request.Context()
	var usersValidate createUsersValidate
	if err := c.ShouldBindJSON(&usersValidate); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("CreateUser").SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}

		messagesErr := make([]string, 0)
		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("CreateUser").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	createUserBody := usersDomain.CreateUserBody{
		UserName:   usersValidate.UserName,
		Password:   usersValidate.Password,
		UserTypeId: usersValidate.UserTypeId,
		PersonId:   usersValidate.PersonId,
	}

	if usersValidate.Person != nil {
		createUserBody.Person = &usersDomain.Person{
			TypeDocumentId: usersValidate.Person.TypeDocumentId,
			Document:       usersValidate.Person.Document,
			Names:          usersValidate.Person.Names,
			Surname:        usersValidate.Person.Surname,
			LastName:       usersValidate.Person.LastName,
			Phone:          usersValidate.Person.Phone,
			Email:          usersValidate.Person.Email,
			Gender:         usersValidate.Person.Gender,
			Enable:         usersValidate.Person.Enable,
		}
	}

	id, err := h.usersUseCase.CreateUser(ctx, createUserBody)
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

// UpdateUser is a method to update a user
// @Summary Update a user
// @Description Update a user
// @Tags Users
// @Accept json
// @Produce json
// @Param userId path string true "user id"
// @Param updateUserBody body usersDomain.UpdateUserBody true "Update user body"
// @Success 200 {object} httpResponse.StatusResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/users/{userId} [put]
// @Security BearerAuth
func (h usersHandler) UpdateUser(c *gin.Context) {
	ctx := c.Request.Context()
	userId := c.Param("userId")

	var usersValidate UpdateUserValidate
	if err := c.ShouldBindJSON(&usersValidate); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("UpdateUser").SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}

		messagesErr := make([]string, 0)
		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("UpdateUser").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	userBody := usersDomain.UpdateUserBody{
		UserName:   usersValidate.UserName,
		UserTypeId: usersValidate.UserTypeId,
		PersonId:   usersValidate.PersonId,
	}

	if usersValidate.Person != nil {
		userBody.Person = &usersDomain.Person{
			TypeDocumentId: usersValidate.Person.TypeDocumentId,
			Document:       usersValidate.Person.Document,
			Names:          usersValidate.Person.Names,
			Surname:        usersValidate.Person.Surname,
			LastName:       usersValidate.Person.LastName,
			Phone:          usersValidate.Person.Phone,
			Email:          usersValidate.Person.Email,
			Gender:         usersValidate.Person.Gender,
			Enable:         usersValidate.Person.Enable,
		}
	}

	err := h.usersUseCase.UpdateUser(ctx, userId, userBody)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := httpResponse.StatusResult{
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// UpdatePasswordUser is a method to update a password of user
// @Summary Update a password of user
// @Description Update a password of user
// @Tags Users
// @Accept json
// @Produce json
// @Param UpdatePasswordBody body usersDomain.UpdatePasswordBody true "Update password data"
// @Success 200 {object} httpResponse.StatusResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/users/update/password [PUT]
// @Security BearerAuth
func (h usersHandler) UpdatePasswordUser(c *gin.Context) {
	ctx := c.Request.Context()
	userId := c.GetString("userId")

	var usersValidate UpdatePasswordDataValidate
	if err := c.ShouldBindJSON(&usersValidate); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("UpdateUser").SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}

		messagesErr := make([]string, 0)
		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("UpdateUser").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	data := usersDomain.UpdatePasswordBody{
		CurrentPassword: usersValidate.CurrentPassword,
		NewPassword:     usersValidate.NewPassword,
	}

	err := h.usersUseCase.UpdatePasswordUser(ctx, data, userId)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := httpResponse.StatusResult{
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// DeleteUser is a method to delete a user
// @Summary Delete a user
// @Description Delete a user
// @Tags Users
// @Accept json
// @Produce json
// @Param userId path string true "user id"
// @Success 200 {object} deleteUsersResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/users/{userId} [delete]
// @Security BearerAuth
func (h usersHandler) DeleteUser(c *gin.Context) {
	ctx := c.Request.Context()
	userId := c.Param("userId")
	result, err := h.usersUseCase.DeleteUser(ctx, userId)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}
	res := deleteUsersResult{
		Data:   result,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// ResetPasswordUser is a method to reset password of user
// @Summary Reset password
// @Description Reset password
// @Tags Users
// @Accept json
// @Produce json
// @Param userId path string true "user id"
// @Success 200 {object} ResetPasswordUserResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/users/{userId}/password [put]
// @Security BearerAuth
func (h usersHandler) ResetPasswordUser(c *gin.Context) {
	ctx := c.Request.Context()
	userId := c.Param("userId")

	var userPasswordValidate resetUserPasswordValidate
	if err := c.ShouldBindJSON(&userPasswordValidate); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("ResetPasswordUser").SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}
		messagesErr := make([]string, 0)
		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("ResetPasswordUser").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}
	resetUserPasswordBody := usersDomain.ResetUserPasswordBody{
		NewPassword: userPasswordValidate.NewPassword,
	}

	success, err := h.usersUseCase.ResetPasswordUser(ctx, userId, resetUserPasswordBody)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	if !success {
		err = h.err.Clone().SetFunction("ResetPasswordUser").SetRaw(errors.New("casting ValidationErrors"))
		restCore.ErrJson(c, err) // todo
		return
	}

	res := ResetPasswordUserResult{
		Data:   success,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

func GetHostWithoutPort(req *http.Request) string {
	host := req.Host
	if index := strings.Index(host, ":"); index != -1 {
		host = host[:index]
	}
	return host
}

// LoginUser is a method to logs in a user
// @Summary Login
// @Description Login a user
// @Tags Users
// @Accept json
// @Produce json
// @Param loginBody body usersDomain.LoginUserBody true "Login Body"
// @Success 201 {object} LoginUserResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/auth/login [post]
func (h usersHandler) LoginUser(c *gin.Context) {
	ctx := c.Request.Context()

	host := GetHostWithoutPort(c.Request)
	ctx = context.WithValue(ctx, "xTenantId", host)
	c.Header("X-Tenant-Host", host)

	var loginValidate loginUserValidate
	if err := c.ShouldBindJSON(&loginValidate); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("LoginUser").SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}
		messagesErr := make([]string, 0)
		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("LoginUser").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}
	loginUserBody := usersDomain.LoginUserBody{
		UserName: loginValidate.UserName,
		Password: loginValidate.Password,
	}

	tkn, xTenantId, err := h.usersUseCase.LoginUser(ctx, loginUserBody)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}
	res := LoginUserResult{
		Data:   *tkn,
		Status: http.StatusOK,
	}
	if xTenantId != nil {
		c.Header("X-Tenant-Id", *xTenantId)
	}
	restCore.Json(c, http.StatusOK, res)
}

// VerifyPermissionsByUser is a method to verify permissions of a user
// @Summary is a method to verify permissions of a user
// @Description is a method to verify permissions of a user
// @Tags Users
// @Accept json
// @Produce json
// @Param store_id query string false "store id"
// @Param codePermission path string true "code permission"
// @Success 200 {object} httpResponse.BoolResponse "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/users/me/permissions/{codePermission} [get]
// @Security BearerAuth
func (h usersHandler) VerifyPermissionsByUser(c *gin.Context) {
	ctx := c.Request.Context()
	userId := c.GetString("userId")
	storeId := c.Query("store_id")
	codePermission := c.Param("codePermission")

	result, err := h.usersUseCase.VerifyPermissionsByUser(ctx, userId, storeId, codePermission)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}
	res := httpResponse.BoolResponse{
		Data: result,
	}
	restCore.Json(c, http.StatusOK, res)
}

// GetModulePermissions is a method to list permissions of a user in a module
// @Summary is a method to list permissions of a user in a module
// @Description is a method to list permissions of a user in a module
// @Tags Users
// @Accept json
// @Produce json
// @Param codeModule path string true "code module"
// @Success 200 {object} PermissionsResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/users/me/modules/{codeModule}/permissions [get]
// @Security BearerAuth
func (h usersHandler) GetModulePermissions(c *gin.Context) {
	ctx := c.Request.Context()
	userId := c.GetString("userId")
	codeModule := c.Param("codeModule")

	result, err := h.usersUseCase.GetModulePermissions(ctx, userId, codeModule)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}
	res := PermissionsResult{
		Data:   result,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// GetRolesByUser is a method to list roles of a user
// @Summary is a method to list roles of a user
// @Description is a method to list roles of a user
// @Tags Users
// @Accept json
// @Produce json
// @Param userId path string true "user id"
// @Success 200 {object} rolesByUserResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/users/{userId}/merchants/roles [get]
// @Security BearerAuth
func (h usersHandler) GetRolesByUser(c *gin.Context) {
	ctx := c.Request.Context()
	userId := c.Param("userId")

	roles, err := h.usersUseCase.GetRolesByUser(ctx, userId)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := rolesByUserResult{
		Data:   roles,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// GetUserPermissionsByRole is a method to list permissions of a user by role
// @Summary is a method to list permissions of a user by role
// @Description is a method to list permissions of a user by role
// @Tags Users
// @Accept json
// @Produce json
// @Param userId path string true "user id"
// @Param userRoleId path string true "user role id"
// @Success 200 {object} PermissionsResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/users/{userId}/roles/{userRoleId}/permissions [get]
// @Security BearerAuth
func (h usersHandler) GetUserPermissionsByRole(c *gin.Context) {
	ctx := c.Request.Context()
	userId := c.Param("userId")
	userRoleId := c.Param("userRoleId")

	result, err := h.usersUseCase.GetUserPermissionsByRole(ctx, userId, userRoleId)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}
	res := permissionsByRoleResult{
		Data:   result,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// UpdateUserTheme is a method to update the theme of a user
// @Summary Update the theme of a user
// @Description Update the theme of a user
// @Tags Users
// @Accept json
// @Produce json
// @Param UpdateUserThemeBody body usersDomain.UpdateUserThemeBody true "Update user theme body"
// @Param userId path string true "user id"
// @Success 200 {object} httpResponse.StatusResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/users/{userId}/theme [put]
// @Security BearerAuth
func (h usersHandler) UpdateUserTheme(c *gin.Context) {
	ctx := c.Request.Context()
	userId := c.Param("userId")

	var usersValidate UpdateUserThemeValidate
	if err := c.ShouldBindJSON(&usersValidate); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("UpdateUserTheme").SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}

		messagesErr := make([]string, 0)
		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("UpdateUserTheme").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	body := usersDomain.UpdateUserThemeBody{
		Theme: usersValidate.Theme,
	}

	err := h.usersUseCase.UpdateUserTheme(ctx, body, userId)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := httpResponse.StatusResult{
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// GenerateQRCode is a method to generate a QR code
// @Summary Generate a QR code
// @Description Generate a QR code in PNG format
// @Tags Users
// @Accept json
// @Produce image/png
// @Success 200 {string} string "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/users/qr [get]
// @Security BearerAuth
func (h usersHandler) GenerateQRCode(c *gin.Context) {
	ctx := c.Request.Context()
	userId := c.GetString("userId")

	png, err := h.usersUseCase.GenerateQRCode(ctx, userId)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	c.Data(http.StatusOK, "image/png", png)
}

// GetViewsByUserToken is a method to get  views by user
// @Summary Get views by user using their token
// @Description Get views by user using their token
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} viewsByUserResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/users/views [get]
// @Security BearerAuth
func (h usersHandler) GetViewsByUserToken(c *gin.Context) {
	ctx := c.Request.Context()
	userId := c.GetString("userId")

	menu, err := h.usersUseCase.GetViewsByUser(ctx, userId)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := viewsByUserResult{
		Data:   menu,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}
