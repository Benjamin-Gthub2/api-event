/*
 * File: users_mysql_repository.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Repository for users.
 *
 * Last Modified: 2023-11-23
 */

package mysql

import (
	"time"

	smartClock "github.com/smart0n3/api-shared/clock"

	userDomain "github.com/Benjamin-Gthub2/api-event/users/domain"
	errDomain "github.com/smart0n3/api-shared/error-core/domain"
)

type usersMySQLRepo struct {
	clock   smartClock.Clock
	timeout time.Duration
	err     *errDomain.SmartError
}

func NewUsersRepository(
	clock smartClock.Clock,
	mongoTimeout int,
) userDomain.UserRepository {
	rep := &usersMySQLRepo{
		clock:   clock,
		timeout: time.Duration(mongoTimeout) * time.Second,
		err:     errDomain.NewErr().SetLayer(errDomain.Infra),
	}
	return rep
}
