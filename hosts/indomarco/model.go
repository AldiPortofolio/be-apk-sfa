package indomarco

//Response
type Response struct {
	Meta map[string]interface{} `json:"meta"`
	Data Data                   `json:"data"`
}

//Data
type Data struct {
	Msg      string      `json:"message"`
	Customer interface{} `json:"customer"`
}

// UpdateMerchantResponse ..
type UpdateMerchantResponse struct {
	ID            int64      `json:"id"`
	CustomerCode  string     `json:"customer_code"`
	CustomerGroup string     `json:"customer_group"`
	CustomerType  string     `json:"customer_type"`
	DeliveryCode  string     `json:"delivery_code"`
	OutletCode    string     `json:"outlet_code"`
	Name          string     `json:"name"`
	Phone         string     `json:"phone"`
	AreaID        int64      `json:"area_id"`
	StockPoint    StockPoint `json:"stockpoint"`
	Role          string     `json:"role"`
	CreditLimits  []string   `json:"credit_limits"`
}

type StockPoint struct {
	ID             int64  `json:"id"`
	BaCode         string `json:"ba_code"`
	CompanyCode    string `json:"company_code"`
	CompanyName    string `json:"company_name"`
	Email          string `json:"email"`
	LocationCode   string `json:"location_code"`
	Name           string `json:"name"`
	PlantCode      string `json:"plant_code"`
	PlantName      string `json:"plant_name"`
	ProxyVoucher   string `json:"proxy_voucher"`
	SalesPointCode string `json:"sales_point_code"`
	SeqID          int    `json:"seq_id"`
}
