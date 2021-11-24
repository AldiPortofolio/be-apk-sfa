package models

// CheckIdCardReq ..
type CheckIdCardReq struct {
	IdCard	string `json:"id_card"`
}

// CheckIdCardRes ..
type CheckIdCardRes struct {
	NumIdCard	string `json:"num_id_card"`
}