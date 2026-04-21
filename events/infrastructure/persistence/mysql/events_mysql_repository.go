/*
 * File: events_mysql_repository.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file initializes the repository layer to manage event related data.
 *
 * Last Modified: 2026-04-15
 */

package mysql

import (
	"time"

	smartClock "github.com/smart0n3/api-shared/clock"
	errDomain "github.com/smart0n3/api-shared/error-core/domain"

	eventDomain "github.com/Benjamin-Gthub2/api-event/events/domain"
	policiesDomain "github.com/Benjamin-Gthub2/api-event/policies/domain"
	policyPermissionsDomain "github.com/Benjamin-Gthub2/api-event/policy-permissions/domain"
	rolePoliciesDomain "github.com/Benjamin-Gthub2/api-event/role-policies/domain"
	rolesDomain "github.com/Benjamin-Gthub2/api-event/roles/domain"
)

type eventsMySQLRepo struct {
	clock                       smartClock.Clock
	timeout                     time.Duration
	err                         *errDomain.SmartError
	roleRepository              rolesDomain.RoleRepository
	policyRepository            policiesDomain.PolicyRepository
	rolePoliciesRepository      rolePoliciesDomain.RolePolicyRepository
	policyPermissionsRepository policyPermissionsDomain.PolicyPermissionRepository
}

func NewEventsRepository(
	clock smartClock.Clock,
	mongoTimeout int,
	roleRepository rolesDomain.RoleRepository,
	policyRepository policiesDomain.PolicyRepository,
	rolePoliciesRepository rolePoliciesDomain.RolePolicyRepository,
	policyPermissionsRepository policyPermissionsDomain.PolicyPermissionRepository,
) eventDomain.EventRepository {
	rep := &eventsMySQLRepo{
		clock:                       clock,
		timeout:                     time.Duration(mongoTimeout) * time.Second,
		err:                         errDomain.NewErr().SetLayer(errDomain.Infra),
		roleRepository:              roleRepository,
		policyRepository:            policyRepository,
		rolePoliciesRepository:      rolePoliciesRepository,
		policyPermissionsRepository: policyPermissionsRepository,
	}
	return rep
}
