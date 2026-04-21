/*
 * File: setup_events.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is file content the setup of the events.
 *
 * Last Modified: 2026-04-15
 */

package setup

import (
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/smart0n3/api-shared/auth"
	authRepository "github.com/smart0n3/api-shared/auth/infrastructure/jwt"
	smartClock "github.com/smart0n3/api-shared/clock"
	validationsRepository "github.com/smart0n3/api-shared/validations/infrastructure/persistence/mysql"

	policiesRepository "github.com/Benjamin-Gthub2/api-event/policies/infrastructure/persistence/mysql"
	policyPermissionsRepository "github.com/Benjamin-Gthub2/api-event/policy-permissions/infrastructure/persistence/mysql"
	rolePoliciesRepository "github.com/Benjamin-Gthub2/api-event/role-policies/infrastructure/persistence/mysql"
	roleDefaultsRepository "github.com/Benjamin-Gthub2/api-event/roles-defaults/infrastructure/persistence/mysql"
	rolesRepository "github.com/Benjamin-Gthub2/api-event/roles/infrastructure/persistence/mysql"

	eventsRepository "github.com/Benjamin-Gthub2/api-event/events/infrastructure/persistence/mysql"
	eventsHttpDelivery "github.com/Benjamin-Gthub2/api-event/events/interfaces/rest"
	eventsUseCase "github.com/Benjamin-Gthub2/api-event/events/usecase"
)

func LoadEvents(router *gin.Engine) {
	timeoutContext := time.Duration(60) * time.Second
	clock := smartClock.NewClock()
	validationRepository := validationsRepository.NewValidationsRepository(60)
	roleRepository := rolesRepository.NewRolesRepository(clock, 60, nil, nil, nil)
	policyRepository := policiesRepository.NewPoliciesRepository(clock, 60)
	rolePolicyRepository := rolePoliciesRepository.NewRolePoliciesRepository(clock, 60)
	policyPermissionRepository := policyPermissionsRepository.NewPolicyPermissionsRepository(clock, 60)
	eventRepository := eventsRepository.NewEventsRepository(clock, 60, roleRepository,
		policyRepository, rolePolicyRepository, policyPermissionRepository)
	authJWTRepository := authRepository.NewAuthRepository()
	authMiddleware := auth.LoadAuthMiddleware()
	roleDefaultRepository := roleDefaultsRepository.NewRolesDefaultsRepository(clock, 60)
	eventUCase := eventsUseCase.NewEventsUseCase(
		eventRepository,
		validationRepository,
		authJWTRepository,
		roleDefaultRepository,
		timeoutContext,
	)
	eventsHttpDelivery.NewEventsHandler(eventUCase, router, authMiddleware)
}
