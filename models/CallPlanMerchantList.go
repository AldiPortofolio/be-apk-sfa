package models

// CallPlanMerchantListReq ..
type CallPlanMerchantListReq struct {
	CallPlanId int64  `json:"call_plan_id"`
	Status     string `json:"status"` //Incompleted, Completed, Visited
	Limit      int    `json:"limit"`  //25
	Page       int    `json:"page"`
}

// CallPlanMerchantListRes ..
type CallPlanMerchantListRes struct {
	CallPlanMerchantId    int64  `json:"call_plan_merchant_id"`
	MerchantName          string `json:"merchant_name"`
	MerchantPhone         string `json:"merchant_phone"`
	Longitude             string `json:"longitude"`
	Latitude              string `json:"latitude"`
	MerchantId            string `json:"merchant_id"`
	MerchantTypeName      string `json:"merchant_type_name"`
	MerchantAddress       string `json:"merchant_address"`
	Priority              int    `json:"priority"`
	AmountAction          int    `json:"amount_action"`
	AmountActionCompleted int    `json:"amount_action_completed"`
	MerchantStatus        string `json:"merchant_status"` //Found - Open, Found - Closed, Not Found
	TodolistStatus        int    `json:"todolist_status"`
}

// CallPlanMerchantListResV23 ..
type CallPlanMerchantListResV23 struct {
	CallPlanMerchantId    int64  `json:"call_plan_merchant_id"`
	MerchantName          string `json:"merchant_name"`
	MerchantPhone         string `json:"merchant_phone"`
	Longitude             string `json:"longitude"`
	Latitude              string `json:"latitude"`
	MerchantId            string `json:"merchant_id"`
	MerchantTypeName      string `json:"merchant_type_name"`
	MerchantAddress       string `json:"merchant_address"`
	Priority              int    `json:"priority"`
	AmountAction          int    `json:"amount_action"`
	AmountActionCompleted int    `json:"amount_action_completed"`
	MerchantStatus        string `json:"merchant_status"` //Found - Open, Found - Closed, Not Found
	TodolistStatus        int    `json:"todolist_status"`
	SalesTypeId           int8   `json:"sales_type_id"`
	SalesTypeName         string `json:"sales_type_name" gorm:"column:sales_type_name"`
}

// CallPlanMerchantListResV24 ..
type CallPlanMerchantListResV24 struct {
	CallPlanMerchantId    int64  `json:"call_plan_merchant_id"`
	MerchantName          string `json:"merchant_name"`
	MerchantPhone         string `json:"merchant_phone"`
	Longitude             string `json:"longitude"`
	Latitude              string `json:"latitude"`
	MerchantId            string `json:"merchant_id"`
	MerchantTypeName      string `json:"merchant_type_name"`
	MerchantAddress       string `json:"merchant_address"`
	Priority              int    `json:"priority"`
	AmountAction          int    `json:"amount_action"`
	AmountActionCompleted int    `json:"amount_action_completed"`
	MerchantStatus        string `json:"merchant_status"` //Found - Open, Found - Closed, Not Found
	TodolistStatus        int    `json:"todolist_status"`
	SalesTypeId           int8   `json:"sales_type_id"`
	SalesTypeName         string `json:"sales_type_name" gorm:"column:sales_type_name"`
	AddressBenchmark      string `json:"address_benchmark" gorm:"column:address_benchmark"`
}
