package fdsmodels

// CheckProfilSalesRes ..
type CheckProfilSalesRes struct {
	DescriptionCode string                 `json:"DescriptionCode"`
	AcquisitionData []CheckProfilSalesRes1 `json:"acquisitionData"`
	ResponCode      string                 `json:"responCode"`
}

// CheckProfilSalesRes1 ..
type CheckProfilSalesRes1 struct {
	MerchantPhoto string `json:"MerchantPhoto"`
	Phone         string `json:"Phone"`
	StoreName     string `json:"StoreName"`
	Address       string `json:"address"`
	Created       string `json:"created"`
	MerchantID    string `json:"merchantId"`
	Status        string `json:"status"`
}
