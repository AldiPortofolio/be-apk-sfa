package models

// ReportBySalesReq ..
type ReportBySalesReq struct {
	Phone    string `json:"phone" form:"phone"`
	DateFrom string `json:"date_from" form:"date_from"`
	DateTo   string `json:"date_to" form:"date_to"`
}

// ReportBySalesRes ..
type ReportBySalesRes struct {
	OttoPay    int    `json:"ottopay"`
	SfaOnly    int    `json:"sfa_only"`
	SalesPhone string `json:"sales_phone"`
}