package notif

// NotifRequest ..
type NotifRequest struct {
	CollapseKey  string           `json:"collapse_key"`
	Notification NotificationData `json:"notification"`
	Data         DataNotif        `json:"data"`
}

// NotificationData ..
type NotificationData struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Sound string `json:"sound"`
}

// DataNotif ..
type DataNotif struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	Target string `json:"target"`
}
