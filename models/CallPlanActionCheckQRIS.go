package models

// CallPlanActionCheckQRISReq ..
type CallPlanActionCheckQRISReq struct {
	QRContent          string `json:"qr_content"`
	CallPlanMerchantId int64  `json:"call_plan_merchant_id"`
	Longitude          string `json:"longitude"`
	Latitude           string `json:"latitude"`
}
