package dbmodels

import (
	"time"
)

// Attendance ..
type Attendance struct {
	ID                 uint      `json:"id"`
	SalesID            int       `json:"sales_id"`
	SalesPhone         string    `json:"sales_phone"`
	SalesName          string    `json:"sales_name"`
	Selfie             string    `json:"selfie"`
	ClocktimeServer    time.Time `json:"date"`
	ClocktimeLocal     time.Time `json:"clocktime_local"`
	Location           string    `json:"location"`
	Latitude           string    `json:"latitude"`
	Longitude          string    `json:"longitude"`
	AttendCategory     string    `json:"attend_category"`
	AttendCategoryType string    `json:"attend_category_type"`
	TypeAttendance     string    `json:"type_attendance"`
	Notes              string    `json:"notes"`
	PhotoAccuration    string    `json:"photo_accuration"`
	PhotoProfile       string    `json:"photo_profile"`
	MinAccPercentage   string    `json:"min_accuration_percentage"`
}

// AttendanceList ..
type AttendanceList struct {
	ID                 uint      `json:"id"`
	SalesID            int       `json:"sales_id"`
	SalesPhone         string    `json:"sales_phone"`
	SalesName          string    `json:"sales_name"`
	ClocktimeServer    time.Time `json:"date"`
	AttendCategory     string    `json:"attend_category"`
	AttendCategoryType string    `json:"attend_category_type"`
	Notes              string    `json:"notes"`
}

// AttendCategoryList ..
type AttendCategoryList struct {
	ID           uint   `json:"category_id"`
	CategoryName string `json:"category_name"`
	CategoryType string `json:"type"`
	TimeIn       string `json:"time_in"`
	TimeOut      string `json:"time_out"`
}

// AttendCategory ..
type AttendCategory struct {
	ID           uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	CategoryName string    `json:"category_name"`
	CategoryType string    `json:"type"`
	TimeIn       time.Time `json:"time_in"`
	TimeOut      time.Time `json:"time_out"`
}

// ReqAttendCategory ..
type ReqAttendCategory struct {
	ID           string `json:"id"`
	CategoryName string `json:"category_name"`
	CategoryType string `json:"type"`
	TimeIn       string `json:"time_in"`
	TimeOut      string `json:"time_out"`
}

// TableName ..
func (t *AttendCategory) TableName() string {
	return "attendance_categories"
}
