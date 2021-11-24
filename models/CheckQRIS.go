package models

// CheckQRISReq ..
type CheckQRISReq struct {
	QRContent string `json:"qr_content"`
	//CallPlanMerchantId  int64  `json:"call_plan_merchant_id"`
}

// CheckQRISRes ..
type CheckQRISRes struct {
	MID                 string `json:"mid"`
	MPAN                string `json:"mpan"`
	NMID                string `json:"nmid"`
	StoreNamePreprinted string `json:"store_name_preprinted"`
}
