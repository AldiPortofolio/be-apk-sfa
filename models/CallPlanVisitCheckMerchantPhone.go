package models

// CallPlanVisitCheckMerchantPhoneReq ..
type CallPlanVisitCheckMerchantPhoneReq struct {
	MerchantPhone string `json:"merchant_phone"`
}

// CallPlanVisitMerchantRes ..
type CallPlanVisitMerchantRes struct {
	MerchantName     string `gorm:"merchant_name" json:"merchant_name"`
	MerchantId       string `gorm:"merchant_id" json:"merchant_id"`
	Mpan             string `gorm:"mpan" json:"mpan"`
	IdMerchant       int64  `gorm:"id_merchant" json:"id_merchant"`
	MerchantAddress  string `gorm:"merchant_address" json:"merchant_address"`
	MerchantPhone    string `gorm:"merchant_phone" json:"merchant_phone"`
	MerchantTypeId   int    `gorm:"merchant_type_id" json:"merchant_type_id"`
	MerchantTypeName string `gorm:"merchant_type_name" json:"merchant_type_name"`
	ClockIn          string `json:"clock_in"`
}
