package models

// Count ..
type Count struct {
	Count string `gorm:"column:count" json:"count"`
}
