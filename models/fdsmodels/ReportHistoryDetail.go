package fdsmodels

// ReportHistoryDetailRes ..
type ReportHistoryDetailRes struct {
	ResponCode      string                    `json:"responCode"`
	DescriptionCode string                    `json:"DescriptionCode"`
	AcquisitionData []ReportHistoryDetailRes1 `json:"acquisitionData"`
}

// ReportHistoryDetailRes1 ..
type ReportHistoryDetailRes1 struct {
	Created       string `json:"created"`
	Address       string `json:"address"`
	MerchantPhoto string `json:"MerchantPhoto"`
	StoreName     string `json:"StoreName"`
	Status        string `json:"status"`
	MerchantID    string `json:"merchantId"`
	Phone         string `json:"Phone"`
}
