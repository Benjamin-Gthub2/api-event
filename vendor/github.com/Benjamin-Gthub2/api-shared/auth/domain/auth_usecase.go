/*
 * File: auth_usecase.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file defines the AuthUseCase interface
 *
 * Last Modified: 2023-11-26
 */

package domain

import (
	"context"
)

type AuthUseCase interface {
	GenerateToken(userId string) (*string, error)
	DecodeToken(ctx context.Context, myToken string) (userId *string, err error)
	GetMerchantStoresByUser(ctx context.Context, userId string) (modules map[string]ModuleMid, err error)
}
