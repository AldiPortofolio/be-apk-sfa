package dbmodels

import (
	"time"
)

// LabelTasks ..
type LabelTasks struct {
	ID            int       `gorm:"column:id;primary_key" json:"id"`
	Name          string    `gorm:"column:name" json:"name"`
	LabelType     string    `gorm:"column:label_type" json:"label_type"`
	SubCategoryID int       `gorm:"column:sub_category_id" json:"sub_category_id"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`
	Condition     string    `gorm:"column:condition" json:"condition"`
	Step          int       `gorm:"column:step" json:"step"`
	SupplierName  string    `gorm:"column:supplier_name" json:"supplier_name"`
}

// TableName ..
func (t *LabelTasks) TableName() string {
	return "public.label_tasks"
}
