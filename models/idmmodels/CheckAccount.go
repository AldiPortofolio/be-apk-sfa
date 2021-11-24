package idmmodels

// CheckAccountRes ..
type CheckAccountRes struct {
	Data Datas `json:"data"`
	Meta Metas `json:"meta"`
}

// Datas ..
type Datas struct {
	Customer Customers `json:"customer"`
}

// Customers ..
type Customers struct {
	Area          struct{}       `json:"area"`
	AuthToken     string         `json:"auth_token"`
	CustomerCode  string         `json:"customer_code"`
	CustomerGroup string         `json:"customer_group"`
	ID            int            `json:"id"`
	MerchantID    string         `json:"merchant_id"`
	Name          string         `json:"name"`
	OttoPhone     string         `json:"otto_phone"`
	OttoStatus    bool           `json:"otto_status"`
	Phone         string         `json:"phone"`
	Stockpoint    DataStockpoint `json:"stockpoint"`
}

// DataStockpoint ..
type DataStockpoint struct {
	CompanyCode    string      `json:"company_code"`
	CreatedAt      string      `json:"created_at"`
	Email          interface{} `json:"email"`
	ID             int         `json:"id"`
	LocationCode   string      `json:"location_code"`
	Name           string      `json:"name"`
	PlantCode      string      `json:"plant_code"`
	PlantName      string      `json:"plant_name"`
	ProxyVoucher   interface{} `json:"proxy_voucher"`
	SalesPointCode string      `json:"sales_point_code"`
	SeqID          int         `json:"seq_id"`
	UpdatedAt      string      `json:"updated_at"`
}

// Metas ..
type Metas struct {
	Status  bool   `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}
