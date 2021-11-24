package models

// MerchantDetailTokoOttopayRes ..
type MerchantDetailTokoOttopayRes struct {
	MerchantId    string `json:"merchant_id"`
	MerchantPhone string `json:"merchant_phone"`
	MerchantName  string `json:"merchant_name"`
	OwnerName     string `json:"owner_name"`
	Address       string `json:"address"`
}

// MerchantDetailTokoOttopayV23Res ..
type MerchantDetailTokoOttopayV23Res struct {
	MerchantId    string `json:"merchant_id"`
	MerchantPhone string `json:"merchant_phone"`
	MerchantName  string `json:"merchant_name"`
	OwnerName     string `json:"owner_name"`
	Address       string `json:"address"`
	AddressProvinceId	int64  `json:"province_id"`
	AddressCityId		int64  `json:"city_id"`
	AddressDistrictId	int64  `json:"district_id"`
	AddressVillageId	int64  `json:"villager_id"`
}
