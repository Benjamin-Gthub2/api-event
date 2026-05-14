/*
 * File: users_func_usecase.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Implementation of use cases to users.
 *
 * Last Modified: 2023-11-23
 */

package usecase

import (
	"context"
	"sort"
	"strings"
	"sync"

	"github.com/google/uuid"
	"github.com/skip2/go-qrcode"

	errDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"
	logErrorCoreDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"
	paramsDomain "github.com/Benjamin-Gthub2/api-shared/params/domain"
	validationsDomain "github.com/Benjamin-Gthub2/api-shared/validations/domain"

	usersDomain "github.com/Benjamin-Gthub2/api-event/users/domain"
)

func (u usersUseCase) UpdatePasswordUser(
	ctx context.Context,
	body usersDomain.UpdatePasswordBody,
	userId string,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "core_users",
		IdColumnName:     "id",
		IdValue:          userId,
		StatusColumnName: nil,
		StatusValue:      nil,
	}
	exist, err := u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return err
	}
	if !exist {
		return usersDomain.ErrUserNotFound.SetFunction("UpdatePasswordUser").SetLayer(errDomain.UseCase)
	}
	err = u.usersRepository.UpdatePasswordUserMain(ctx, body, userId)
	if err != nil {
		return err
	}
	return nil
}

func (u usersUseCase) GetUser(
	ctx context.Context,
	userId string,
) (
	user *usersDomain.User,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	user, err = u.usersRepository.GetUser(ctx, userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u usersUseCase) GetUsers(
	ctx context.Context,
	searchParams usersDomain.GetUsersParams,
	pagination paramsDomain.PaginationParams,
) (
	users []usersDomain.UserMultiple,
	paginationResults *paramsDomain.PaginationResults,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var errGetUsers, errGetTotalUsers error
	var total *int
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer logErrorCoreDomain.PanicThreadRecovery(&ctx, &errGetUsers, &wg)
		users, errGetUsers = u.usersRepository.GetUsers(ctx, searchParams, pagination)
		wg.Done()
	}()
	go func() {
		defer logErrorCoreDomain.PanicThreadRecovery(&ctx, &errGetTotalUsers, &wg)
		total, errGetTotalUsers = u.usersRepository.GetTotalUsers(ctx, searchParams, pagination)
		wg.Done()
	}()
	wg.Wait()

	if errGetUsers != nil {
		return nil, nil, errGetUsers
	}
	if errGetTotalUsers != nil {
		return nil, nil, errGetTotalUsers
	}

	paginationRes := paramsDomain.PaginationResults{}
	paginationRes.FromParams(pagination, *total)

	return users, &paginationRes, nil
}

func (u usersUseCase) GetMenuByUser(
	ctx context.Context,
	userId string,
) (
	menuAux []*usersDomain.MenuModule,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	modules, err := u.usersRepository.GetModules(ctx)
	if err != nil {
		return nil, err
	}

	modulesByUser, err := u.usersRepository.GetMenuByUser(ctx, userId)
	if err != nil {
		return nil, err
	}

	var menu2 = make([]*usersDomain.MenuModule, 0)
	menu2Map := map[string]*usersDomain.MenuModule{}

	// Crear mapas para búsqueda rápida
	modulesMap := make(map[string]usersDomain.Module)
	for _, module := range modules {
		modulesMap[module.Code] = module
	}

	modulesByUserMap := make(map[string]usersDomain.ModuleMenuUser)
	for _, module := range modulesByUser {
		modulesByUserMap[module.Code] = module
	}

	for _, moduleByUser := range modulesByUser {
		moduleParts := strings.Split(moduleByUser.Code, ".")
		modulePath := ""
		var parentModule *usersDomain.MenuModule

		for i, part := range moduleParts {
			if i > 0 {
				modulePath += "."
			}
			modulePath += part

			if existingModule, exists := menu2Map[modulePath]; exists {
				parentModule = existingModule
				continue
			}

			var newModule *usersDomain.MenuModule

			// Verificar si el módulo tiene vistas en modulesByUser
			if userModule, exists := modulesByUserMap[modulePath]; exists {
				newModule = &usersDomain.MenuModule{
					ModuleMenuUser: usersDomain.ModuleMenuUser{
						Id:          userModule.Id,
						Name:        userModule.Name,
						Description: userModule.Description,
						Code:        modulePath,
						Icon:        userModule.Icon,
						Position:    userModule.Position,
						CreatedAt:   userModule.CreatedAt,
						Views:       userModule.Views, // Tomar vistas de modulesByUser si existen
					},
					Modules: []*usersDomain.MenuModule{},
				}
			} else if systemModule, exists := modulesMap[modulePath]; exists {
				// Si no tiene vistas, pero existe en modules, tomar su información
				newModule = &usersDomain.MenuModule{
					ModuleMenuUser: usersDomain.ModuleMenuUser{
						Id:          systemModule.Id,
						Name:        systemModule.Name,
						Description: systemModule.Description,
						Code:        modulePath,
						Icon:        systemModule.Icon,
						Position:    systemModule.Position,
						CreatedAt:   systemModule.CreatedAt,
						Views:       []usersDomain.ViewMenuUser{}, // Si no tiene vistas, asignar slice vacío
					},
					Modules: []*usersDomain.MenuModule{},
				}
			} else {
				continue // Si no se encuentra en ninguno de los dos, ignorarlo
			}

			menu2Map[modulePath] = newModule

			if parentModule != nil {
				parentModule.Modules = append(parentModule.Modules, newModule)
			} else {
				menu2 = append(menu2, newModule)
			}

			parentModule = newModule
		}
	}
	sortModules(menu2)
	return menu2, nil
}

func sortModules(modules []*usersDomain.MenuModule) {
	sort.Slice(modules, func(i, j int) bool {
		return modules[i].ModuleMenuUser.Position < modules[j].ModuleMenuUser.Position
	})
	for _, mod := range modules {
		sortModules(mod.Modules)
	}
}

func (u usersUseCase) GetMeByUser(
	ctx context.Context,
	userId string,
) (
	user *usersDomain.UserMe,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var errUserMe, errGetStoresByUser, errGetMerchantsByUser error
	var userMe *usersDomain.UserMeInfo
	var storesByUser []usersDomain.StoreByUser
	var merchantsByUser []usersDomain.MerchantByUser
	var wg sync.WaitGroup

	wg.Add(3)
	go func() {
		defer logErrorCoreDomain.PanicThreadRecovery(&ctx, &errUserMe, &wg)
		userMe, errUserMe = u.usersRepository.GetMeByUser(ctx, userId)
		wg.Done()
	}()
	go func() {
		defer logErrorCoreDomain.PanicThreadRecovery(&ctx, &errGetStoresByUser, &wg)
		storesByUser, errGetStoresByUser = u.usersRepository.GetStoresByUser(ctx, userId)
		wg.Done()
	}()
	go func() {
		defer logErrorCoreDomain.PanicThreadRecovery(&ctx, &errGetMerchantsByUser, &wg)
		merchantsByUser, errGetMerchantsByUser = u.usersRepository.GetMerchantsByUser(ctx, userId)
		wg.Done()
	}()
	wg.Wait()

	if errUserMe != nil {
		return nil, errUserMe
	}
	if errGetStoresByUser != nil {
		return nil, errGetStoresByUser
	}
	if errGetMerchantsByUser != nil {
		return nil, errGetMerchantsByUser
	}

	user = &usersDomain.UserMe{
		Id:        userMe.Id,
		UserName:  userMe.UserName,
		Theme:     userMe.Theme,
		CreatedAt: userMe.CreatedAt,
		Person:    userMe.Person,
		RoleUser:  userMe.RoleUser,
		Stores:    storesByUser,
		Merchants: merchantsByUser,
	}
	return user, nil
}

func (u usersUseCase) CreateUser(
	ctx context.Context,
	body usersDomain.CreateUserBody,
) (
	id *string, err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "core_users",
		IdColumnName:     "username",
		IdValue:          body.UserName,
		StatusColumnName: nil,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.ValidateExistence(ctx, recordExistsParams)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, usersDomain.ErrUserUsernameAlreadyExist
	}

	userId := uuid.New().String()
	// case basic
	if body.PersonId == nil && body.Person == nil {
		id, err = u.usersRepository.CreateUser(ctx, nil, userId, body)
	} else if body.PersonId != nil {
		body.Person = nil
		id, err = u.usersRepository.CreateUserMain(ctx, userId, *body.PersonId, body)
	} else {
		err = u.usersRepository.ValidateUniquePersonByDocument(ctx, body.Person.TypeDocumentId, body.Person.Document)
		if err != nil {
			return nil, err
		}
		personId := uuid.New().String()
		id, err = u.usersRepository.CreateUserMain(ctx, userId, personId, body)
	}
	return
}

func (u usersUseCase) UpdateUser(
	ctx context.Context,
	userId string,
	body usersDomain.UpdateUserBody,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "core_users",
		IdColumnName:     "id",
		IdValue:          userId,
		StatusColumnName: nil,
		StatusValue:      nil,
	}
	exist, err := u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return err
	}
	if !exist {
		return u.err.Clone().CopyCodeDescription(usersDomain.ErrUserNotFound).SetFunction("UpdateUser")
	}

	if body.PersonId == nil && body.Person == nil {
		err = u.usersRepository.VerifyIfUserExist(ctx, userId)
		if err != nil {
			return err
		}
		err = u.usersRepository.UpdateUser(ctx, userId, body)

	} else if body.PersonId != nil && body.Person != nil {
		err = u.usersRepository.VerifyIfPersonExist(ctx, *body.PersonId)
		if err != nil {
			return err
		}
		err = u.usersRepository.UpdateUserMain(ctx, userId, *body.PersonId, body)
	} else if body.PersonId != nil {
		body.Person = nil
		err = u.usersRepository.UpdateUserMain(ctx, userId, *body.PersonId, body)
	} else {
		personId := uuid.New().String()
		err = u.usersRepository.ValidateUniquePersonByDocument(ctx, body.Person.TypeDocumentId, body.Person.Document)
		if err != nil {
			return err
		}
		err = u.usersRepository.UpdateUserMain(ctx, userId, personId, body)
	}

	return
}

func (u usersUseCase) DeleteUser(
	ctx context.Context,
	userId string,
) (
	update bool,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	var deleted string
	deleted = "deleted_at"
	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "core_users",
		IdColumnName:     "id",
		IdValue:          userId,
		StatusColumnName: &deleted,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return false, err
	}
	if !exist {
		return false, usersDomain.ErrUserIdHasBeenDeleted
	}

	res, err := u.usersRepository.DeleteUser(ctx, userId)
	return res, err
}

func (u usersUseCase) ResetPasswordUser(
	ctx context.Context,
	userId string,
	body usersDomain.ResetUserPasswordBody,
) (
	updated bool,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	hashPassword, err := HashPasswordUser(ctx, body.NewPassword)
	if err != nil {
		return false, err
	}

	updated, err = u.usersRepository.ResetPasswordUser(ctx, userId, hashPassword)
	return
}

func (u usersUseCase) LoginUser(
	ctx context.Context,
	body usersDomain.LoginUserBody,
) (
	tokenString *string,
	xTenantId *string,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	hashPassword, err := HashPasswordUser(ctx, body.Password)
	if err != nil {
		return nil, nil, err
	}

	user, xTenantId, err := u.usersRepository.GetUserByUserNameAndPassword(ctx, body.UserName, hashPassword)
	if err != nil {
		return nil, xTenantId, err
	}

	tokenString, err = u.authRepository.GenerateToken(user.Id)
	return tokenString, xTenantId, nil
}

func HashPasswordUser(
	ctx context.Context,
	password string,
) (
	passwordHash string,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	// REVIEW generate hash password in base of password
	return password, nil
}

func (u usersUseCase) VerifyPermissionsByUser(
	ctx context.Context, userId string, storeId string, codePermission string,
) (
	res bool, err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	if storeId == "" {
		return res, usersDomain.ErrStoreIdEmpty
	}

	res, err = u.usersRepository.VerifyPermissionsByUser(ctx, userId, storeId, codePermission)
	if err != nil {
		return res, err
	}
	return
}

func (u usersUseCase) GetModulePermissions(
	ctx context.Context, userId string, codeModule string,
) (
	permissions []usersDomain.Permissions, err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "core_modules",
		IdColumnName:     "code",
		IdValue:          codeModule,
		StatusColumnName: nil,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return permissions, err
	}
	if !exist {
		return permissions, u.err.Clone().CopyCodeDescription(usersDomain.ErrInvalidCodeModule).SetFunction("GetModulePermissions")
	}

	permissions, err = u.usersRepository.GetModulePermissions(ctx, userId, codeModule)
	if err != nil {
		return permissions, err
	}
	return permissions, nil
}

func (u usersUseCase) GetRolesByUser(
	ctx context.Context,
	userId string,
) (
	generalRoles *usersDomain.GeneralRolesByUser,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var errRoles, errGlobalRoles error
	var roles []usersDomain.MerchantRoleByUser
	var globalRoles []usersDomain.RolesByUser
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer logErrorCoreDomain.PanicThreadRecovery(&ctx, &errRoles, &wg)
		roles, errRoles = u.usersRepository.GetRolesByUser(ctx, userId)
		wg.Done()
	}()
	go func() {
		defer logErrorCoreDomain.PanicThreadRecovery(&ctx, &errGlobalRoles, &wg)
		globalRoles, errGlobalRoles = u.usersRepository.GetGlobalRolesByUser(ctx, userId)
		wg.Done()
	}()
	wg.Wait()

	if errRoles != nil {
		err = errRoles
		return
	}
	if errGlobalRoles != nil {
		err = errGlobalRoles
		return
	}

	generalRoles = &usersDomain.GeneralRolesByUser{
		Merchants:   roles,
		GlobalRoles: globalRoles,
	}
	return generalRoles, nil
}

func (u usersUseCase) GetUserPermissionsByRole(
	ctx context.Context,
	userId string,
	roleId string,
) (
	permissions []usersDomain.GetUserPermissionsByRole,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	permissions, err = u.usersRepository.GetUserPermissionsByRole(ctx, userId, roleId)
	if err != nil {
		return nil, err
	}
	return permissions, nil
}

func (u usersUseCase) UpdateUserTheme(
	ctx context.Context,
	body usersDomain.UpdateUserThemeBody,
	userId string,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "core_users",
		IdColumnName:     "id",
		IdValue:          userId,
		StatusColumnName: nil,
		StatusValue:      nil,
	}
	exist, err := u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return err
	}
	if !exist {
		return usersDomain.ErrUserNotFound.SetFunction("UpdateUserTheme").SetLayer(errDomain.UseCase)
	}
	err = u.usersRepository.UpdateUserTheme(ctx, body, userId)
	if err != nil {
		return err
	}
	return nil
}

func (u usersUseCase) GenerateQRCode(
	ctx context.Context,
	userId string,
) (
	png []byte,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	png, err = qrcode.Encode(userId, qrcode.Medium, 256)
	if err != nil {
		return nil, err
	}
	return png, nil
}

func (u usersUseCase) GetViewsByUser(
	ctx context.Context,
	userId string,
) (
	result *usersDomain.ViewsByUserData,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var views []usersDomain.View
	var userMe *usersDomain.UserMeInfo
	var errViews, errMe error

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		views, errViews = u.usersRepository.GetViewsByUser(ctx, userId)
	}()
	go func() {
		defer wg.Done()
		userMe, errMe = u.usersRepository.GetMeByUser(ctx, userId)
	}()
	wg.Wait()

	if errViews != nil {
		return nil, errViews
	}

	result = &usersDomain.ViewsByUserData{
		Views: views,
	}
	if errMe == nil && userMe != nil {
		result.Person = userMe.Person
	}
	return result, nil
}
