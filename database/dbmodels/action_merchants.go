package dbmodels

import "time"

// ActionMerchants ..
type ActionMerchants struct {
	Id        int64     `gorm:"column:id;primary_key" json:"action_merchant_id"`
	Name      string    `gorm:"column:name" json:"action_merchant_name"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName ..
func (t *ActionMerchants) TableName() string {
	return "public.action_merchants"
}
