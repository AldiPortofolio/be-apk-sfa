package models

import "time"

// CallPlanTodolistListReq ..
type CallPlanTodolistListReq struct {
	MerchantPhone string   `json:"merchant_phone"`
	Status        []string `json:"status,omitempty"`
	Page          int64    `json:"page"`
	Limit         int64    `json:"limit"`
}

// TodoList ..
type TodoList struct {
	TodoList interface{} `json:"todo_list"`
}

// CallPlanTodolistListRes ..
type CallPlanTodolistListRes struct {
	MerchantName          string    `gorm:"merchant_name" json:"merchant_name"`
	TaskDate              time.Time `gorm:"task_date"`
	TaskDateString        string    `json:"task_date"`
	MerchantAddress       string    `gorm:"merchant_address" json:"merchant_address"`
	MerchantID            string    `gorm:"merchant_id" json:"merchant_id"`
	PhoneNumber           string    `gorm:"phone_number" json:"phone_number"`
	NameCategory          string    `gorm:"name_category" json:"name_category"`
	Status                string    `gorm:"status" json:"status"`
	ID                    int64     `gorm:"id" json:"id"`
	Reason                string    `gorm:"reason" json:"reason"`
	Longitude             string    `gorm:"longitude" json:"longitude"`
	Latitude              string    `gorm:"latitude" json:"latitude"`
	PendingTaskDate       time.Time `gorm:"pending_task_date"`
	PendingTaskDateString string    `json:"pending_task_date"`
	VillageID             int       `json:"village_id"`
}

// CallPlanTodolistListResv23 ..
type CallPlanTodolistListResv23 struct {
	MerchantName          string    `gorm:"merchant_name" json:"merchant_name"`
	TaskDate              time.Time `gorm:"task_date"`
	TaskDateString        string    `json:"task_date"`
	MerchantAddress       string    `gorm:"merchant_address" json:"merchant_address"`
	MerchantID            string    `gorm:"merchant_id" json:"merchant_id"`
	PhoneNumber           string    `gorm:"phone_number" json:"phone_number"`
	NameCategory          string    `gorm:"name_category" json:"name_category"`
	Status                string    `gorm:"status" json:"status"`
	ID                    int64     `gorm:"id" json:"id"`
	Reason                string    `gorm:"reason" json:"reason"`
	Longitude             string    `gorm:"longitude" json:"longitude"`
	Latitude              string    `gorm:"latitude" json:"latitude"`
	PendingTaskDate       time.Time `gorm:"pending_task_date"`
	PendingTaskDateString string    `json:"pending_task_date"`
	VillageID             int       `json:"village_id"`
	AddressBenchmark      string    `gorm:"address_benchmark" json:"address_benchmark"`
}
