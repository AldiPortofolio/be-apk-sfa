package models

// CheckMerchantOttopayReq ..
type CheckMerchantOttopayReq struct {
	Phone        string `json:"phone" form:"phone"`
	CustomerCode string `json:"customer_code" form:"customer_code"`
	IdCard       string `json:"id_card" form:"id_card"`
}

// CheckMerchantOttopayRes ..
type CheckMerchantOttopayRes struct {
	AccontNumber    string   `json:"accont_number"`
	Address         string   `json:"address"`
	BankFavorite    []string `json:"bank_favorite"`
	BirthDate       string   `json:"birth_date"`
	CustomerID      int      `json:"customer_id"`
	DescriptionCode string   `json:"description_code"`
	Email           string   `json:"email"`
	Gender          string   `json:"gender"`
	IDCardNumber    string   `json:"id_card_number"`
	MerchantID      string   `json:"merchant_id"`
	Nama            string   `json:"nama"`
	Nickname        string   `json:"nickname"`
	OwnerName       string   `json:"owner_name"`
	ResponCode      string   `json:"respon_code"`
	Status          string   `json:"status"`
	VerifyStatus    string   `json:"verify_status"`
	VirtualAccount  string   `json:"virtual_account"`
}

// CheckMerchantOttopayNotExistRes ..
type CheckMerchantOttopayNotExistRes struct {
	ResponCode      string `json:"respon_code"`
	DescriptionCode string `json:"description_code"`
	MerchantID      string `json:"merchant_id,omitempty"`
}

// CheckMerchantOttopayFDSRes ..
type CheckMerchantOttopayFDSRes struct {
	DescriptionCode   string   `json:"DescriptionCode"`
	AccontNumber      string   `json:"accontNumber"`
	Address           string   `json:"address"`
	BankAccountName   string   `json:"bankAccountName"`
	BankAccountNumber string   `json:"bankAccountNumber"`
	BankFavorite      []string `json:"bankFavorite"`
	BankID            string   `json:"bankId"`
	BirthDate         string   `json:"birthDate"`
	CustomerID        int      `json:"customerId"`
	Email             string   `json:"email"`
	Gender            string   `json:"gender"`
	IDCardNumber      string   `json:"idCardNumber"`
	MerchantID        string   `json:"merchantId"`
	Nama              string   `json:"nama"`
	Nickname          string   `json:"nickname"`
	OwnerName         string   `json:"ownerName"`
	ResponCode        string   `json:"responCode"`
	Status            string   `json:"status"`
	VerifyStatus      string   `json:"verifyStatus"`
	VirtualAccount    string   `json:"virtualAccount"`
}

// CheckMerchantOttopayAlreadyExistRes ..
type CheckMerchantOttopayAlreadyExistRes struct {
	DescriptionCode string `json:"description_code"`
	MerchantID 		string `json:"merchant_id"`
	ResponCode 		string `json:"respon_code"`
}