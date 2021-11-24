package models

import "time"

// TodolistTaskBySubCategoryReq ..
type TodolistTaskBySubCategoryReq struct {
	MerchantPhone string `json:"merchant_phone"`
	TodolistId    string `json:"todolist_id"`
	SubCategoryId string `json:"sub_category_id"`
}

// TodolistTaskBySubCategoryRes ..
type TodolistTaskBySubCategoryRes struct {
	ID            int       `json:"id"`
	Name          []string  `json:"name"`
	LabelType     string    `json:"label_type"`
	SubCategoryID int       `json:"sub_category_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Condition     string    `json:"condition"`
	Step          int       `json:"step"`
}
