package dbmodels

import "time"

// ProductMerchants ..
type ProductMerchants struct {
	Id        int64     `gorm:"column:id;primary_key" json:"product_merchant_id"`
	Name      string    `gorm:"column:name" json:"product_merchant_name"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName ..
func (t *ProductMerchants) TableName() string {
	return "public.product_merchants"
}
