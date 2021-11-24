package models

// CallPlanActionAddOrEditReq ..
type CallPlanActionAddOrEditReq struct {
	Id                 int64   `gorm:"column:id;primary_key" json:"id"`
	CallPlanMerchantId int64   `gorm:"column:call_plan_merchant_id" json:"call_plan_merchant_id"`
	Name               string  `gorm:"column:name" json:"name"`
	Action             string  `gorm:"column:action" json:"action"`
	ActionType         string  `gorm:"column:action_type" json:"action_type"`
	Product            string  `gorm:"column:product" json:"product"`
	Description        string  `gorm:"column:description" json:"description"`
	MerchantAction     string  `gorm:"column:merchant_action" json:"merchant_action"`
	Result             bool    `gorm:"column:result" json:"result"`
	Amount             float32 `gorm:"column:amount" json:"amount"`
	Reason             string  `gorm:"column:reason" json:"reason"`
	Note               string  `gorm:"column:note" json:"note"`
	Status             string  `gorm:"column:status" json:"status"`
	CreatedAt          string  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt          string  `gorm:"column:updated_at" json:"updated_at"`

	MerchantId    string `json:"merchant_id"`
	MerchantPhone string `json:"merchant_phone"`

	NMID                string `json:"nmid"`
	StoreNamePreprinted string `json:"store_name_preprinted"`
}
