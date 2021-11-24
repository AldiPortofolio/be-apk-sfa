package models

// UpdateMerchantIndomarcoReq ..
type UpdateMerchantIndomarcoReq struct {
	Phone              string `json:"phone" form:"phone" example:"081122330042"`
	CustomerCode       string `json:"customer_code" form:"customer_code" example:"520052170070634329.0"`
	MerchantID         string `json:"merchant_id" form:"merchant_id" example:"OP1B00000041"`
	MerchantGroupID    int    `json:"rose_merchant_group" form:"rose_merchant_group" example:"74"`
	MerchantCategoryID string `json:"rose_merchant_category" form:"rose_merchant_category" example:"op"`
}

// UpdateMerchantIndomarcoOttopayReq ..
type UpdateMerchantIndomarcoOttopayReq struct {
	Phone         string `json:"phone" form:"phone"`
	AccountNumber string `json:"accont_number" form:"accont_number"`
	OwnerName     string `json:"owner_name" form:"owner_name"`
	MerchantID    string `json:"merchant_id" form:"merchant_id"`
	CustomerID    string `json:"customer_id" form:"customer_id"`
	Name          string `json:"name" form:"name"`
}

// UpdateMerchantIndomarcoRes ..
type UpdateMerchantIndomarcoRes struct {
	Customer UpdateMerchantIndomarcoRes1 `json:"customer"`
	Message  string                      `json:"message"`
}

// UpdateMerchantIndomarcoRes1 ..
type UpdateMerchantIndomarcoRes1 struct {
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
