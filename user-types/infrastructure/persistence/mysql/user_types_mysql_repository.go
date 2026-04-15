/*
 * File: user_types_mysql_repository.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Repository for NewUserTypesRepository.
 *
 * Last Modified: 2023-11-23
 */

package mysql

import (
	"time"

	smartClock "github.com/smart0n3/api-shared/clock"

	userTypeDomain "github.com/Benjamin-Gthub2/api-event/user-types/domain"
	errDomain "github.com/smart0n3/api-shared/error-core/domain"
)

type userTypesMySQLRepo struct {
	clock   smartClock.Clock
	timeout time.Duration
	err     *errDomain.SmartError
}

func NewUserTypesRepository(
	clock smartClock.Clock,
	mongoTimeout int,
) userTypeDomain.UserTypeRepository {
	rep := &userTypesMySQLRepo{
		clock:   clock,
		timeout: time.Duration(mongoTimeout) * time.Second,
		err:     errDomain.NewErr().SetLayer(errDomain.Infra),
	}
	return rep
}
