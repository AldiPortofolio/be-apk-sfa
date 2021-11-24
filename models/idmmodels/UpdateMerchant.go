package idmmodels

// UpdateMerchantRes ..
type UpdateMerchantRes struct {
	Data GetData `json:"data"`
	Meta GetMeta `json:"meta"`
}

// GetData ..
type GetData struct {
	Customer GetCustomer `json:"customer"`
	Message  string      `json:"message"`
}

// GetCustomer ..
type GetCustomer struct {
	AreaCode       string      `json:"area_code"`
	AreaID         interface{} `json:"area_id"`
	CreatedAt      string      `json:"created_at"`
	CustomerCode   string      `json:"customer_code"`
	CustomerGroup  string      `json:"customer_group"`
	CustomerType   string      `json:"customer_type"`
	DeliveryCode   interface{} `json:"delivery_code"`
	ID             int         `json:"id"`
	Inactive       string      `json:"inactive"`
	LocationCode   string      `json:"location_code"`
	MarsToken      interface{} `json:"mars_token"`
	MerchantID     string      `json:"merchant_id"`
	Name           string      `json:"name"`
	OttoPhone      string      `json:"otto_phone"`
	OttoStatus     bool        `json:"otto_status"`
	OutletCode     string      `json:"outlet_code"`
	PasswordDigest interface{} `json:"password_digest"`
	Phone          string      `json:"phone"`
	SeqID          int         `json:"seq_id"`
	StockpointID   int         `json:"stockpoint_id"`
	UpdatedAt      string      `json:"updated_at"`
}

// GetMeta ..
type GetMeta struct {
	Status  bool   `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}
