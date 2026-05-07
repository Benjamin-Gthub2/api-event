package jwt

type MerchantMid struct {
	Id       string   `db:"merchant_id_middle"`
	StoreIds []string `db:"store_id_middle"`
}

type ModuleMid struct {
	Id          string `db:"module_id"`
	Name        string `db:"module_name"`
	Description string `db:"module_description"`
	Code        string `db:"module_code_middle"`
	Merchants   []MerchantMid
}
