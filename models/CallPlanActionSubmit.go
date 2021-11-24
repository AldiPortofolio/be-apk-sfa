package models

// CallPlanActionSubmitReq ..
type CallPlanActionSubmitReq struct {
	CallPlanMerchantId int64  `json:"call_plan_merchant_id"`
	MerchantStatus     string `json:"merchant_status"`
	Status             string `json:"status"`
	Notes              string `json:"notes"`
}
