package models

// MerchantDetailQRISReq ..
type MerchantDetailQRISReq struct {
	MerchantPhone string `json:"merchant_phone"`
}

// MerchantDetailQRISRes ..
type MerchantDetailQRISRes struct {
	MerchantId    string `json:"merchant_id"`
	Mpan          string `json:"mpan"`
	MerchantPhone string `json:"merchant_phone"`
	MerchantName  string `json:"merchant_name"`
}
