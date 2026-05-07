/*
 * File: auth_repository.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file defines the AuthRepository interface
 *
 * Last Modified: 2023-11-26
 */

package domain

import (
	"context"
)

type AuthRepository interface {
	GenerateToken(userId string) (*string, error)
	DecodeToken(ctx context.Context, tokenStr string) (userId *string, err error)
	GetMerchantStoresByUser(ctx context.Context, userId string) (modules []ModuleMid, err error)
}
