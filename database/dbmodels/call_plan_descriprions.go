package dbmodels

import "time"

// CallPlanDescriptions ..
type CallPlanDescriptions struct {
	Id                int64     `gorm:"column:id;primary_key" json:"description_id"`
	ActionMerchantId  int64     `gorm:"column:action_merchant_id" json:"action_merchant_id"`
	ProductMerchantId int64     `gorm:"column:product_merchant_id" json:"product_merchant_id"`
	Description       string    `gorm:"column:description" json:"description"`
	CreatedAt         time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt         time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName ..
func (t *CallPlanDescriptions) TableName() string {
	return "public.call_plan_descriptions"
}
