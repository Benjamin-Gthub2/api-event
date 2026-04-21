/*
 * File: users_route_handler.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Route handler to request for users.
 *
 * Last Modified: 2023-11-23
 */

package rest

import (
	"github.com/gin-gonic/gin"

	authMiddleware "github.com/smart0n3/api-shared/auth/interfaces/rest"
	errDomain "github.com/smart0n3/api-shared/error-core/domain"
	swaggerRest "github.com/smart0n3/api-shared/swagger/interfaces/rest"

	"github.com/Benjamin-Gthub2/api-event/users/docs"
	usersDomain "github.com/Benjamin-Gthub2/api-event/users/domain"
)

type usersHandler struct {
	usersUseCase   usersDomain.UserUseCase
	authMiddleware authMiddleware.AuthMiddleware
	err            *errDomain.SmartError
}

func NewUsersHandler(
	users usersDomain.UserUseCase,
	router *gin.Engine,
	authMiddleware authMiddleware.AuthMiddleware,
) {
	handler := &usersHandler{
		usersUseCase:   users,
		authMiddleware: authMiddleware,
		err:            errDomain.NewErr().SetLayer(errDomain.Interface),
	}
	swaggerRest.Handler(router, docs.SwaggerInfousers, docs.DocTemplateJson, "event", "users")

	apiAuth := router.Group("/api/v1/auth")
	apiAuth.Use(handler.authMiddleware.Cors)
	apiAuth.POST("/login", handler.LoginUser)

	api := router.Group("/api/v1/event")
	api.Use(handler.authMiddleware.Cors)
	api.Use(handler.authMiddleware.Auth)
	api.GET("/users/:userId", handler.GetUser)
	api.GET("/users", handler.GetUsers)
	api.GET("/users/:userId/menu", handler.GetMenuByUser)
	api.GET("/users/menu", handler.GetMenuByUserToken)
	api.GET("/users/me", handler.GetMeByUser)
	api.POST("/users", handler.CreateUser)
	api.GET("/users/qr", handler.GenerateQRCode)
	api.PUT("/users/:userId", handler.UpdateUser)
	api.DELETE("/users/:userId", handler.DeleteUser)
	api.PUT("/users/:userId/password", handler.ResetPasswordUser)
	api.PUT("/users/update/password", handler.UpdatePasswordUser)
	api.GET("/users/me/permissions/:codePermission", handler.VerifyPermissionsByUser)
	api.GET("/users/me/modules/:codeModule/permissions", handler.GetModulePermissions)
	api.GET("/users/:userId/merchants/roles", handler.GetRolesByUser)
	api.GET("/users/:userId/roles/:userRoleId/permissions", handler.GetUserPermissionsByRole)
	api.PUT("/users/:userId/theme", handler.UpdateUserTheme)
}
