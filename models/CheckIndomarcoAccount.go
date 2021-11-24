package models

// CheckIndomarcoAccountReq ..
type CheckIndomarcoAccountReq struct {
	CustomerCode string `json:"customer_code" form:"customer_code"`
}

// CheckIndomarcoAccountRes ..
type CheckAccountIndomarcoRes struct {
	Data		 interface{} 	`json:"data"`
	Meta		 MetaData 		`json:"meta"`
}

//// Customers ..
//type Customers struct {
//	Area          interface{}   `json:"area"`
//	AuthToken     string        `json:"auth_token"`
//	CustomerCode  string        `json:"customer_code"`
//	CustomerGroup string        `json:"customer_group"`
//	ID            int           `json:"id"`
//	MerchantID    string        `json:"merchant_id"`
//	Name          string        `json:"name"`
//	OttoPhone     string        `json:"otto_phone"`
//	OttoStatus    bool          `json:"otto_status"`
//	Phone         string        `json:"phone"`
//	Stockpoint    GetStockpoint `json:"stockpoint"`
//}
//
//// GetStockpoint ..
//type GetStockpoint struct {
//	CompanyCode    string      `json:"company_code"`
//	CreatedAt      string      `json:"created_at"`
//	Email          interface{} `json:"email"`
//	ID             int         `json:"id"`
//	LocationCode   string      `json:"location_code"`
//	Name           string      `json:"name"`
//	PlantCode      string      `json:"plant_code"`
//	PlantName      string      `json:"plant_name"`
//	ProxyVoucher   interface{} `json:"proxy_voucher"`
//	SalesPointCode string      `json:"sales_point_code"`
//	SeqID          int         `json:"seq_id"`
//	UpdatedAt      string      `json:"updated_at"`
//}
//
