package dbmodels

import "time"

// CallPlanMerchants ..
type CallPlanMerchants struct {
	Id              int64     `gorm:"column:id;primary_key" json:"id"`
	CallPlanId      int64     `gorm:"column:call_plan_id" json:"call_plan_id"`
	MerchantId      int64     `gorm:"column:merchant_id" json:"merchant_id"` //id_merchant
	MerchantPhone   string    `gorm:"column:merchant_phone" json:"merchant_phone"`
	MerchantAddress string    `gorm:"column:merchant_address" json:"merchant_address"`
	MerchantTypeId  int64     `gorm:"column:merchant_type_id" json:"merchant_type_id"`
	Status          string    `gorm:"column:status" json:"status"`
	MerchantStatus  string    `gorm:"column:merchant_status" json:"merchant_status"`
	EffectiveCall   bool      `gorm:"column:effective_call" json:"effective_call"`
	Amount          float32   `gorm:"column:amount" json:"amount"`
	ClockOut        string    `gorm:"column:clock_time" json:"clock_time"` //clock_out
	CreatedAt       time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at" json:"updated_at"`
	MID             string    `gorm:"column:mid" json:"mid"` //merchant_id
	MerchantName    string    `gorm:"column:merchant_name" json:"merchant_name"`
	ClockIn         string    `gorm:"column:action_date" json:"action_date"` //clock_in
	Longitude       string    `gorm:"column:longitude" json:"longitude"`
	Latitude        string    `gorm:"column:latitude" json:"latitude"`
	PhotoLocation   string    `gorm:"column:photo_location" json:"photo_location"`
	Notes           string    `gorm:"column:notes" json:"notes"`
}

// TableName ..
func (t *CallPlanMerchants) TableName() string {
	return "public.call_plan_merchants"
}
