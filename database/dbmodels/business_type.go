package dbmodels

// BusinessType ..
type BusinessType struct {
	Code string `gorm:"column:code" json:"code"`
	Name string `gorm:"column:name" json:"name"`
}

// TableName ..
func (t *BusinessType) TableName() string {
	return "public.merchant_business_types"
}
