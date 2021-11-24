package opmodels

// UpdateMerchantRes ..
type UpdateMerchantRes struct {
	Data GetData `json:"data"`
	Meta GetMeta `json:"meta"`
}

// GetData ..
type GetData struct {
	Phone string `json:"phone"`
}

// GetMeta ..
type GetMeta struct {
	Status  bool   `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}
