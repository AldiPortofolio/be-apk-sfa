package models

import "time"

// TodolistDetailReq ..
type TodolistDetailReq struct {
	MerchantPhone      string `json:"merchant_phone" example:"089898988897"`
	CustomerCode       string `json:"customer_code" example:"543543422556.0"`
	TodolistId         string `json:"todolist_id" example:"322"`
	TodolistCategoryId int8   `json:"todolist_category_id" example:"5"`
}

// TodolistDetailRes ..
type TodolistDetailRes struct {
	TodolistID         int64             `json:"todolist_id"`
	IdMerchant         int64             `json:"id_merchant"`
	MerchantName       string            `json:"merchant_name"`
	TaskDateString     string            `json:"task_date"`
	CreatedAtString    string            `json:"created_at"`
	MerchantAddress    string            `json:"merchant_address"`
	MerchantID         string            `json:"merchant_id"`
	MerchantPhone      string            `json:"merchant_phone"`
	CustomerCode       string            `json:"customer_code"`
	NameCategory       string            `json:"name_category"`
	TodolistCategoryId string            `json:"todolist_category_id"`
	Status             string            `json:"status"`
	Reason             string            `json:"reason"`
	IdCard             string            `json:"id_card"`
	Notes              string            `json:"notes"`
	Task               []TaskTodolistRes `json:"task"`
}

// TodolistDetailDBRes ..
type TodolistDetailDBRes struct {
	TodolistID            int64     `gorm:"todolist_id" json:"todolist_id"`
	IdMerchant            int64     `gorm:"id_merchant" json:"id_merchant"`
	MerchantName          string    `gorm:"merchant_name" json:"merchant_name"`
	TaskDate              time.Time `gorm:"task_date"`
	TaskID                int64     `gorm:"task_id" json:"task_id"`
	MerchantAddress       string    `gorm:"merchant_address" json:"merchant_address"`
	MerchantID            string    `gorm:"merchant_id" json:"merchant_id"`
	MerchantPhone         string    `gorm:"merchant_phone"`
	CustomerCode          string    `gorm:"customer_code"`
	NameCategory          string    `gorm:"name_category" json:"name_category"`
	TodolistCategoryId    string    `gorm:"todolist_category_id"`
	Status                string    `gorm:"status" json:"status"`
	CreatedAt             time.Time `gorm:"created_at"`
	UpdatedAt             time.Time `gorm:"updated_at"`
	TodolistSubCategoryID string    `gorm:"todolist_sub_category_id" json:"todolist_sub_category_id"`
	Code                  string    `gorm:"code" json:"code"`
	Name                  string    `gorm:"name" json:"name"`
	Reason                string    `gorm:"reason" json:"reason"`
	Notes                 string    `gorm:"notes" json:"notes"`
	IdCard                string    `gorm:"id_card" json:"id_card"`
	LabelType             string    `gorm:"label_type" json:"type"`
	Link                  string    `gorm:"link" json:"link"`
}

// TaskTodolistRes ..
type TaskTodolistRes struct {
	TodolistSubCategoryID string `json:"todolist_sub_category_id"`
	Code                  string `json:"code"`
	Name                  string `json:"name"`
	TaskID                int64  `json:"task_id"`
	Status                bool   `json:"status"`
	LabelType             string `json:"type"`
	Link                  string `json:"link"`
}

// CheckStatusTaskTodolistRes ..
type CheckStatusTaskTodolistRes struct {
	TaskID int64 `json:"task_id"`
}