package models

type MerchantIndomarcoReq struct {
	Phone string `json:"phone"`
	CustomerCode string `json:"customer_code"`
	CustomerID string `json:"customer_id"`
	MerchantID string `json:"merchant_id"`
}

type UpdateMerchantIndomarcoV2Req struct {
	Phone        string `json:"phone" form:"otto_phone"`
	CustomerCode string `json:"customer_code" form:"customer_code"`
	MerchantID   string `json:"merchant_id" form:"merchant_id"`
}
