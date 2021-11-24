package dbmodels

import (
	"time"
)

// MerchantTypes ..
type MerchantTypes struct {
	ID          int       `gorm:"column:id;primary_key" json:"id"`
	Name        string    `gorm:"column:name" json:"name"`
	Priority	int	      `gorm:"column:priority" json:"priority"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName ..
func (t *MerchantTypes) TableName() string {
	return "public.merchant_types"
}
