package models

// FeedingLongLatReq ..
type FeedingLongLatReq struct {
	SalesPhone    string `json:"sales_phone" example:"094343443243"`
	Longitude     string `json:"longitude" example:"654.333"`
	Latitude      string `json:"latitude" example:"-22.11"`
	Description   string `json:"description" example:"Sales Location"`
}
