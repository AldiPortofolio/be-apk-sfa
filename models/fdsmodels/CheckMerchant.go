package fdsmodels

// CheckMerchantRes ..
type CheckMerchantRes struct {
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
