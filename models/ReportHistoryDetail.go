package models

// ReportHistoryDetailReq ..
type ReportHistoryDetailReq struct {
	Phone string `json:"phone" form:"phone"`
}

// ReportHistoryDetailRes ..
type ReportHistoryDetailRes struct {
	AcquisitionData []ReportHistoryDetailRes2 `json:"acquisition_data,omitempty"`
}

// ReportHistoryDetailRes2 ..
type ReportHistoryDetailRes2 struct {
	Created       string `json:"created"`
	Address       string `json:"address"`
	MerchantPhoto string `json:"merchant_photo"`
	StoreName     string `json:"store_name"`
	Status        string `json:"status"`
	MerchantID    string `json:"merchant_id"`
	Phone         string `json:"phone"`
}
