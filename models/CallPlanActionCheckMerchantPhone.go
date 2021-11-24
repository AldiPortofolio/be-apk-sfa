package models

// CallPlanActionCheckMerchantPhoneReq ..
type CallPlanActionCheckMerchantPhoneReq struct {
	CallPlanMerchantId int64  `json:"call_plan_merchant_id"`
	MerchantPhone      string `json:"merchant_phone"`
	Longitude          string `json:"longitude"`
	Latitude           string `json:"latitude"`
}
