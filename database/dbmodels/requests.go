package dbmodels

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

// Requests ..
type Requests struct {
	ID             int64       `gorm:"column:id;primary_key" json:"id"`
	RequestType    int         `gorm:"column:request_type" json:"request_type"`
	MakerId        int         `gorm:"column:maker_id" json:"maker_id"`
	Status         int         `gorm:"column:status" json:"status"`
	ApprovableId   int         `gorm:"column:approvable_id" json:"approvable_id"`
	ApprovableType string      `gorm:"column:approvable_type" json:"approvable_type"`
	ToBeChanged    ToBeChanged `gorm:"column:to_be_changed" json:"to_be_changed"`
	Note           string      `gorm:"column:note" json:"note"`
	Module         int         `gorm:"column:module" json:"module"`
	DeletedAt      time.Time   `gorm:"column:deleted_at" json:"deleted_at"`
	CreatedAt      time.Time   `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time   `gorm:"column:updated_at" json:"updated_at"`
}

// ToBeChanged ..
type ToBeChanged map[string]interface{}

// Value ..
func (a ToBeChanged) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan ..
func (a *ToBeChanged) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

// TableName ..
func (t *Requests) TableName() string {
	return "public.requests"
}
