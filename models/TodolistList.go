package models

import "time"

// TodolistListReq ..
type TodolistListReq struct {
	Keyword       string   `json:"keyword,omitempty" example:"089898988897"`
	TaskDateStart string   `json:"task_date_start,omitempty" example:"089898988897"`
	TaskDateEnd   string   `json:"task_date_end,omitempty" example:"089898988897"`
	Status        []string `json:"status,omitempty" example:"['Late','Open']"`
	CategoryID    []string `json:"id_category,omitempty" example:"['2','1']"`
	VillageID     []string `json:"id_village,omitempty" example:"['111111212142','122214214132']"`
	ClusterID     []string `json:"id_cluster,omitempty" example:"['111111212142','122214214132']"`
	Page          int64    `json:"page"`
	Limit         int64    `json:"limit"`
}

// TodolistListRes ..
type TodolistListRes struct {
	TodoList interface{} `json:"todo_list"`
}

// TodolistListDBRes ..
type TodolistListDBRes struct {
	MerchantName    string    `gorm:"merchant_name" json:"merchant_name"`
	TaskDate        time.Time `gorm:"task_date"`
	TaskDateString  string    `json:"task_date"`
	MerchantAddress string    `gorm:"merchant_address" json:"merchant_address"`
	MerchantID      string    `gorm:"merchant_id" json:"merchant_id"`
	//MerchantPhone     		string 		`gorm:"merchant_phone" json:"merchant_phone"`
	CustomerCode          string    `gorm:"customer_code" json:"customer_code"`
	PhoneNumber           string    `gorm:"phone_number" json:"phone_number"`
	TodolistCategoryId    int8      `gorm:"todolist_category_id" json:"todolist_category_id"`
	MerchantNewRecId      int       `gorm:"merchant_new_rec_id" json:"merchant_new_recruitment_id"`
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

// MerchantTodolistListDBRes ..
type MerchantTodolistListDBRes struct {
	Longitude    string `gorm:"longitude" json:"longitude"`
	Latitude     string `gorm:"latitude" json:"latitude"`
	Address      string `gorm:"address" json:"address"`
	CustomerCode string `gorm:"customer_code" json:"customer_code"`
	PhoneNumber  string `gorm:"phone_number" json:"phone_number"`
}

// TodolistVillageID ..
type TodolistVillageID struct {
	VillageId string `gorm:"column:village_id"`
}

// MerchantTodolistListDBResV24 ..
type MerchantTodolistListDBResV24 struct {
	Longitude     		string `gorm:"longitude" json:"longitude"`
	Latitude      		string `gorm:"latitude" json:"latitude"`
	Address       		string `gorm:"address" json:"address"`
	AddressBenchmark	string `gorm:"address_benchmark" json:"address_benchmark"`
	CustomerCode  		string `gorm:"customer_code" json:"customer_code"`
	PhoneNumber   		string `gorm:"phone_number" json:"phone_number"`
	SalesTypeId   		int    `json:"sales_type_id"`
	SalesTypeName 		string `json:"sales_type_name" gorm:"column:sales_type_name"`
}

// TodolistListDBResV24 ..
type TodolistListDBResV24 struct {
	MerchantName    string    `gorm:"merchant_name" json:"merchant_name"`
	TaskDate        time.Time `gorm:"task_date"`
	TaskDateString  string    `json:"task_date"`
	MerchantAddress string    `gorm:"merchant_address" json:"merchant_address"`
	MerchantID      string    `gorm:"merchant_id" json:"merchant_id"`
	//MerchantPhone     		string 		`gorm:"merchant_phone" json:"merchant_phone"`
	CustomerCode          string    `gorm:"customer_code" json:"customer_code"`
	PhoneNumber           string    `gorm:"phone_number" json:"phone_number"`
	TodolistCategoryId    int8      `gorm:"todolist_category_id" json:"todolist_category_id"`
	MerchantNewRecId      int       `gorm:"merchant_new_rec_id" json:"merchant_new_recruitment_id"`
	NameCategory          string    `gorm:"name_category" json:"name_category"`
	Status                string    `gorm:"status" json:"status"`
	ID                    int64     `gorm:"id" json:"id"`
	Reason                string    `gorm:"reason" json:"reason"`
	Longitude             string    `gorm:"longitude" json:"longitude"`
	Latitude              string    `gorm:"latitude" json:"latitude"`
	PendingTaskDate       time.Time `gorm:"pending_task_date"`
	PendingTaskDateString string    `json:"pending_task_date"`
	VillageID             int       `json:"village_id"`
	SalesTypeId           int       `json:"sales_type_id" gorm:"column:sales_type_id"`
	SalesTypeName         string    `json:"sales_type_name" gorm:"column:sales_type_name"`
	AddressBenchmark      string 	`json:"address_benchmark" gorm:"column:address_benchmark"`
	SalesPhone		      string 	`json:"sales_phone" gorm:"column:sales_phone"`
}