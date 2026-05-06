package domain

import (
	"context"
)

type MerchantMid struct {
	Id       string   `json:"id"`
	StoreIds []string `json:"store_ids"`
}

type ModuleMid struct {
	Id          string        `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Code        string        `json:"code"`
	Merchants   []MerchantMid `json:"merchants"`
}

func GetStrStoreIds(ctx context.Context, moduleCode string) ([]string, error) {
	merchantApiKey, hasMerchantApiKey := ctx.Value("merchantApiKey").(string)
	storeApiKey, hasStoreApiKey := ctx.Value("storeApiKey").(string)

	modulesMap, ok := ctx.Value("modules").(map[string]ModuleMid)
	if !ok {
		return nil, ErrGetStoresMiddleNotFound
	}
	module, hasModule := modulesMap[moduleCode]
	if !hasModule {
		return nil, ErrUserInsufficientPermissions
	}
	storeIds := make([]string, 0)
	for _, merchant := range module.Merchants {
		storeIds = append(storeIds, merchant.StoreIds...)
	}
	if hasStoreApiKey && storeApiKey != "" {
		found := false
		for _, storeId := range storeIds {
			if storeId == storeApiKey {
				found = true
				break
			}
		}
		if !found {
			return nil, ErrUserInsufficientPermissions
		}
		return []string{storeApiKey}, nil
	} else if hasMerchantApiKey && merchantApiKey != "" {
		var merchantMid *MerchantMid
		for _, merchant := range module.Merchants {
			if merchant.Id == merchantApiKey {
				merchantMid = &merchant
				break
			}
		}
		if merchantMid == nil {
			return nil, ErrUserInsufficientPermissions
		}
		return merchantMid.StoreIds, nil
	} else {
		return storeIds, nil
	}
}

func GetStrMerchantIds(ctx context.Context, moduleCode string) ([]string, error) {
	merchantApiKey, hasMerchantApiKey := ctx.Value("merchantApiKey").(string)

	modulesMap, ok := ctx.Value("modules").(map[string]ModuleMid)
	if !ok {
		return nil, ErrGetStoresMiddleNotFound
	}
	module, hasModule := modulesMap[moduleCode]
	if !hasModule {
		return nil, ErrUserInsufficientPermissions
	}
	merchantIds := make([]string, 0)
	for _, merchant := range module.Merchants {
		merchantIds = append(merchantIds, merchant.Id)
	}
	if hasMerchantApiKey && merchantApiKey != "" {
		found := false
		for _, merchant := range module.Merchants {
			if merchant.Id == merchantApiKey {
				found = true
				break
			}
		}
		if !found {
			return nil, ErrUserInsufficientPermissions
		}
		return []string{merchantApiKey}, nil
	} else {
		return merchantIds, nil
	}
}
