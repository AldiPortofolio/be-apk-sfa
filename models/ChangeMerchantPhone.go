package models

// ChangeMerchantPhoneReq ..
type ChangeMerchantPhoneReq struct {
	MerchantID string `json:"merchant_id" form:"merchant_id"`
	NewPhone   string `json:"new_phone" form:"new_phone"`
}

// ChangeMerchantPhoneRes ..
type ChangeMerchantPhoneRes struct {
	DescriptionCode string `json:"description_code"`
	ResponseCode    string `json:"respon_code"`
}
