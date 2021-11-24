package models

// ChangeStatusMerchantReq ..
type ChangeStatusMerchantReq struct {
	Name         string `json:"nama_merchant"`
	MID          string `json:"mid"`
	PhoneNumber  string `json:"no_hp_merchant"`
	Status       string `json:"status"`
	MerchantType string `json:"merchant_type"`
}

// DataMerchantErrorByRow ..
type DataMerchantErrorByRow struct {
	NoRow         int    `csv:"Nomor Row"`
	Name          string `csv:"Nama Merchant"`
	MID           string `csv:"MID"`
	PhoneNumber   string `csv:"No Hp Merchant"`
	Status        string `csv:"Status"`
	MerchantType  string `csv:"Merchant Type"`
	ErrorMessages string `csv:"Error Messages"`
}
