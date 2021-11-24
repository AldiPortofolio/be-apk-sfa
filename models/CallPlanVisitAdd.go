package models

// CallPlanVisitAddReq ..
type CallPlanVisitAddReq struct {

	//Call Plan Merchant
	CallPlanId      int64  `gorm:"column:call_plan_id" json:"call_plan_id"`
	IdMerchant      int64  `gorm:"column:merchant_id" json:"id_merchant"` //id_merchant
	MerchantPhone   string `gorm:"column:merchant_phone" json:"merchant_phone"`
	MerchantAddress string `gorm:"column:merchant_address" json:"merchant_address"`
	MerchantTypeId  int64  `gorm:"column:merchant_type_id" json:"merchant_type_id"`
	MerchantStatus  string `gorm:"column:merchant_status" json:"merchant_status"` //Found - Open
	MerchantId      string `gorm:"column:mid" json:"merchant_id"`                 //merchant_id
	MerchantName    string `gorm:"column:merchant_name" json:"merchant_name"`
	ClockIn         string `gorm:"column:action_date" json:"clock_in"` //clock_in
	Longitude       string `gorm:"column:longitude" json:"longitude"`
	Latitude        string `gorm:"column:latitude" json:"latitude"`
	Status          string `gorm:"column:status" json:"status"` //Visited

	//Call Plan Action
	CallPlanActionName string  `gorm:"column:name" json:"call_plan_action_name"`
	ActionId           int8    `gorm:"column:action" json:"action_id"`
	ActionName         string  `gorm:"column:action" json:"action_name"`
	ProductId          int8    `gorm:"column:product" json:"product_id"`
	ProductName        string  `gorm:"column:product" json:"product_name"`
	Description        string  `gorm:"column:description" json:"description"`
	Result             bool    `gorm:"column:result" json:"result"`
	MerchantAction     string  `gorm:"column:merchant_action" json:"merchant_action"`
	Amount             float32 `gorm:"column:amount" json:"amount"`
	Reason             string  `gorm:"column:reason" json:"reason"`
	Note               string  `gorm:"column:note" json:"note"`
	ActionStatus       string  `gorm:"column:status" json:"action_status"`    //Completed
	ActionType         string  `gorm:"column:action_type" json:"action_type"` //Visited
}
