package models

// CallPlanActionMerchantUnknownReq ..
type CallPlanActionMerchantUnknownReq struct {
	CallPlanMerchantId int64  `json:"call_plan_merchant_id"`
	Longitude          string `json:"longitude"`
	Latitude           string `json:"latitude"`
	PhotoLocation      string `json:"photo_location"`
	MerchantStatus     string `json:"merchant_status"` //Found - Close / Not Found
	Status             string `json:"status"`          //Completed
	Notes              string `json:"notes"`
	ClockOut           string `json:"clock_out"`
}

//kalau merchant_status itu isinya:
//Found - Open
//Found - Closed
//Not Found

// status
//- Incompleted
//- Completed
//- Visited
