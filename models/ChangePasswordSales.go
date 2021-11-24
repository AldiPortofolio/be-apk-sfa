package models

// ChangePasswordSalesReq ..
type ChangePasswordSalesReq struct {
	PhoneNumber string `form:"phone_number" json:"phone_number"`
	Pin         string `form:"old_pin" json:"old_pin"`
	NewPin      string `form:"new_pin" json:"new_pin"`
}

// ChangePasswordSalesRes ..
type ChangePasswordSalesRes struct {
	ResponseCode string `json:"response_code"`
	Description  string `json:"description_code"`
}
