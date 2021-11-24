package models

// IndomarcoLoginReq ..
type IndomarcoLoginReq struct {
	IdNumber      string `form:"idNumber" json:"idNumber"`
	Pin           string `form:"pin" json:"pin"`
	DeviceID      string `form:"deviceId" json:"deviceId"`
	DeviceToken   string `form:"deviceToken" json:"deviceToken"`
	SalesID       string `form:"salesId" json:"salesId"`
	VersionCode   string `form:"version_code" json:"version_code"`
	Role          string `form:"role" json:"role"`
	FirebaseToken string `form:"firebase_token" json:"firebase_token"`
}

// IndomarcoLoginRes ..
type IndomarcoLoginRes struct {
	ResponseCode string `json:"response_code"`
	SalesName    string `json:"sales_name"`
	Email        string `json:"email"`
	Description  string `json:"description_code"`
	Phone        string `json:"phone"`
	Photo        string `json:"photo"`
	SessionToken string `json:"session_token"`
	Status       string `json:"status"`
	SalesId      string `json:"sales_id"`
	SFAID        string `json:"sfa_id"`
	ForceUpdate  bool   `json:"force_update"`
	FunctionalId string `json:"functional_id"`
}
