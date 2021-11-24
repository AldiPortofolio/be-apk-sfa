package dbmodels

// SalesType ..
type SalesType struct {
	ID        int    `gorm:"column:id;primary_key" json:"id"`
	FirstName string `gorm:"column:name" json:"name"`
}

// TableName ..
func (t *SalesType) TableName() string {
	return "public.sales_types"
}
