package models

// PushNotifReq ..
type PushNotifReq struct {
	PhoneNumber string `json:"phone_number"`
	Title       string `json:"title"`
	Body        string `json:"body"`
	Target      string `json:"target"`
}
