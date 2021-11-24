package models

// CallPlanTodolistMerchantNotFoundReq ..
type CallPlanTodolistMerchantNotFoundReq struct {
	MerchantId    string   `json:"merchant_id"`
	MerchantPhone string   `json:"merchant_phone"`
	MerchantImage string   `json:"merchant_image"`
	Issue         []string `json:"issue"`
	Latitude      string   `json:"latitude"`
	Longitude     string   `json:"longitude"`
}
