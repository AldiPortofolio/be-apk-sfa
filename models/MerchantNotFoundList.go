package models

import "time"

// TodolistMerchantNotFoundListReq ..
type TodolistMerchantNotFoundListReq struct {
	TodolistCategoryId int64 `json:"todolist_category_id"`
}

// TodolistMerchantNotFoundListRes ..
type TodolistMerchantNotFoundListRes struct {
	ID        int64     `gorm:"column:id;primary_key" json:"ID"`
	Label     string    `gorm:"column:label" json:"name"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}
