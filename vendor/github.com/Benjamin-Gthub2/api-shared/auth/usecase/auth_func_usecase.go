/*
 * File: auth_func_usecase.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Implementation of use cases to auth.
 *
 * Last Modified: 2023-11-26
 */

package usecase

import (
	"context"

	"github.com/Benjamin-Gthub2/api-shared/auth/domain"
)

func (u authUseCase) GenerateToken(userId string) (*string, error) {
	return u.authRepository.GenerateToken(userId)
}

func (u authUseCase) DecodeToken(ctx context.Context, tokenStr string) (userId *string, err error) {
	return u.authRepository.DecodeToken(ctx, tokenStr)
}

func (u authUseCase) GetMerchantStoresByUser(ctx context.Context, userId string) (modules map[string]domain.ModuleMid, err error) {
	modulesMiddle, err := u.authRepository.GetMerchantStoresByUser(ctx, userId)
	if err != nil {
		return nil, err
	}
	modules = make(map[string]domain.ModuleMid)
	for _, module := range modulesMiddle {
		modules[module.Code] = module
	}
	return modules, nil
}
