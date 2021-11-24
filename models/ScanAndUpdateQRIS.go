package models

// ScanAndUpdateQRISReq ..
type ScanAndUpdateQRISReq struct {
	QRContent     string `json:"qr_content"`
	MerchantId    string `json:"merchant_id"`
	MerchantPhone string `json:"merchant_phone"`
	Mpan          string `json:"mpan"`
}

// ScanAndUpdateQRISRes ..
type ScanAndUpdateQRISRes struct {
}
