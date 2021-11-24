package models

import "ottosfa-api-apk/database/dbmodels"

// LoginReq ..
type LoginReq struct {
	PhoneNumber   string `form:"phoneNumber" json:"phoneNumber"`
	Pin           string `form:"pin" json:"pin"`
	DeviceID      string `form:"deviceId" json:"deviceId"`
	DeviceToken   string `form:"deviceToken" json:"deviceToken"`
	SalesID       string `form:"salesId" json:"salesId"`
	VersionCode   string `form:"version_code" json:"version_code"`
	Role          string `form:"role" json:"role"`
	FirebaseToken string `form:"firebase_token" json:"firebase_token"`
}

// LoginRes ..
type LoginRes struct {
	ResponseCode string               `json:"response_code"`
	SalesName    string               `json:"sales_name"`
	Email        string               `json:"email"`
	Description  string               `json:"description_code"`
	Phone        string               `json:"phone"`
	Photo        string               `json:"photo"`
	SessionToken string               `json:"session_token"`
	Status       string               `json:"status"`
	SalesId      string               `json:"sales_id"`
	SFAID        string               `json:"sfa_id"`
	ForceUpdate  bool                 `json:"force_update"`
	FunctionalId string               `json:"functional_id"`
	SalesType    []dbmodels.SalesType `json:"sales_type"`
}

// LoginV23Res ..
type LoginV23Res struct {
	ResponseCode string               `json:"response_code"`
	SalesName    string               `json:"sales_name"`
	Email        string               `json:"email"`
	Description  string               `json:"description_code"`
	Phone        string               `json:"phone"`
	Photo        string               `json:"photo"`
	SessionToken string               `json:"session_token"`
	Status       string               `json:"status"`
	SalesId      string               `json:"sales_id"`
	SFAID        string               `json:"sfa_id"`
	ForceUpdate  bool                 `json:"force_update"`
	FunctionalId string               `json:"functional_id"`
	SalesTypeId  int				  `json:"sales_type_id"`
}

// LoginV24Res ..
type LoginV24Res struct {
	ResponseCode string               `json:"response_code"`
	Email        string               `json:"email"`
	Description  string               `json:"description_code"`
	Photo        string               `json:"photo"`
	SessionToken string               `json:"session_token"`
	Status       string               `json:"status"`
	SFAID        string               `json:"sfa_id"`
	ForceUpdate  bool                 `json:"force_update"`
	FunctionalId string               `json:"functional_id"`
	SalesId      string               `json:"sales_id"`
	SalesName    string               `json:"sales_name"`
	Phone        string               `json:"phone"`
		SalesType	 string				  `json:"sales_type"`
	SalesTypeId  int				  `json:"sales_type_id"`
		SubArea		 string				  `json:"sub_area"`
		Region		 string				  `json:"region"`
		RegionCode	 string				  `json:"region_code"`
		RegionId	 int64				  `json:"region_id"`
}