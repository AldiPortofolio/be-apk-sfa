package dbmodels

// BusinessCategory ..
type BusinessCategory struct {
	Code string `gorm:"column:code" json:"code"`
	Name string `gorm:"column:name" json:"name"`
}

// TableName ..
func (t *BusinessCategory) TableName() string {
	return "public.merchant_business_categories"
}
