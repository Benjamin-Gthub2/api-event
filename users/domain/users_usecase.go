/*
 * File: users_usecase.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file defines the UserUseCase interface, which declares methods for interacting with users entities.
 * It includes methods for retrieving, creating, updating, and deleting users data.
 *
 * Last Modified: 2023-11-23
 */

package domain

import (
	"context"

	paramsDomain "github.com/smart0n3/api-shared/params/domain"
)

type UserUseCase interface {
	GetUser(ctx context.Context, userId string) (*User, error)
	GetUsers(ctx context.Context, searchParams GetUsersParams, pagination paramsDomain.PaginationParams) (
		[]UserMultiple, *paramsDomain.PaginationResults, error)
	GetMenuByUser(ctx context.Context, userId string) ([]*MenuModule, error)
	GetMeByUser(ctx context.Context, userId string) (*UserMe, error)
	CreateUser(ctx context.Context, body CreateUserBody) (*string, error)
	UpdateUser(ctx context.Context, userId string, body UpdateUserBody) error
	DeleteUser(ctx context.Context, userId string) (bool, error)
	ResetPasswordUser(ctx context.Context, userId string, body ResetUserPasswordBody) (bool, error)
	LoginUser(ctx context.Context, body LoginUserBody) (*string, *string, error)
	VerifyPermissionsByUser(ctx context.Context, userId string, storeId string, codePermission string) (bool, error)
	GetModulePermissions(ctx context.Context, userId string, codeModule string) ([]Permissions, error)
	GetRolesByUser(ctx context.Context, userId string) (*GeneralRolesByUser, error)
	GetUserPermissionsByRole(ctx context.Context, userId string, roleId string) ([]GetUserPermissionsByRole, error)
	UpdatePasswordUser(ctx context.Context, body UpdatePasswordBody, userId string) error
	UpdateUserTheme(ctx context.Context, body UpdateUserThemeBody, userId string) error
	GenerateQRCode(ctx context.Context, content string) ([]byte, error)
}
