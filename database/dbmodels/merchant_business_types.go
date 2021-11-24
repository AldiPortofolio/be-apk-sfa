package dbmodels

// ActionMerchants ..
type MerchantBusinessTypes struct {
	Code string `gorm:"column:code" json:"code"`
	Name string `gorm:"column:name" json:"name"`
}

// TableName ..
func (t *MerchantBusinessTypes) TableName() string {
	return "public.merchant_business_types"
}
