package models

// CheckMerchantByIDReq ..
type CheckMerchantByIDReq struct {
	MerchantID string `json:"merchant_id" form:"merchant_id"`
}

// CheckMerchantByIDRes ..
type CheckMerchantByIDRes struct {
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
