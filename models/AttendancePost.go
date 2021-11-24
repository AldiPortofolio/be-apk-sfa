package models

// AttendancePostReq ..
type AttendancePostReq struct {
	SalesID                int64  `json:"sales_id,omitempty"`
	SalesPhone             string `json:"sales_phone,omitempty"`
	SalesName              string `json:"sales_name,omitempty"`
	AttendanceCategory     string `json:"attendance_category"`
	AttendanceCategoryType string `json:"attendance_category_type"`
	TypeAttendance         string `json:"type_attendance"`
	Notes                  string `json:"notes"`
	PhotoSelfie            string `json:"photo_selfie"`
	Location               string `json:"location,omitempty"`
	Longitude              string `json:"longitude"`
	Latitude               string `json:"latitude"`
	Time                   string `json:"time"`
	TypeTimezone           string `json:"type_timezone"`
	AccurationPhoto        string `json:"photo_accuration"`
}
