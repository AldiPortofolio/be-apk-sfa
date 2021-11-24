package models

// ChangePhotoPinKTPReq ..
type ChangePhotoPinKTPReq struct {
	NewPin string `json:"new_pin" form:"new_pin"`
	IdCard string `json:"id_card" form:"id_card"`
	Photo  string `json:"photo" form:"photo"`
}

// ChangePhotoPinKTPRes ..
type ChangePhotoPinKTPRes struct {
	ResponseCode    string `json:"response_code"`
	DescriptionCode string `json:"description_code"`
}
