package models

// MerchantListReq ..
type MerchantListReq struct {
	MerchantCategory string `json:"merchant_category"`
	Keyword          string `json:"keyword"`
	Page             int    `json:"page"`
	Limit            int    `json:"limit"`

	//with filter
	ProvinceId []string `json:"province_id"`
	CityId     []string `json:"city_id"`
	DistrictId []string `json:"district_id"`
	VillageId  []string `json:"village_id"`
}

// MerchantListRes ..
type MerchantListRes struct {
	ID               int    `json:"id"`
	MerchantId       string `json:"merchant_id"`
	Name             string `json:"name"`
	PhoneNumber      string `json:"phone_number"`
	ImageMerchant    string `json:"image_merchant"`
	Address          string `json:"address"`
	JoinAt           string `json:"join_at"`
	MerchantCategory string `json:"merchant_category"`
	MerchantStatus   string `json:"merchant_status"`
}

// MerchantListv23Req ..
type MerchantListv23Req struct {
	//with filter
	ProvinceId []string `json:"province_id"`
	CityId     []string `json:"city_id"`
	DistrictId []string `json:"district_id"`
	VillageId  []string `json:"village_id"`
}

// MerchantListv23Res ..
type MerchantListv23Res struct {
	ID               int64  `json:"id"`
	MerchantId       string `json:"merchant_id"`
	Name             string `json:"name"`
	PhoneNumber      string `json:"phone_number"`
	ImageMerchant    string `json:"image_merchant"`
	Address          string `json:"address"`
	JoinAt           string `json:"join_at"`
	MerchantCategory string `json:"merchant_category"`
	MerchantStatus   string `json:"merchant_status"`
}