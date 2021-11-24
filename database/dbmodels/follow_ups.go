package dbmodels

import (
	"time"
)

// FollowUps ..
type FollowUps struct {
	ID          int       `gorm:"column:id;primary_key" json:"id"`
	Label       string    `gorm:"column:label" json:"label"`
	ContentType string    `gorm:"column:content_type" json:"content_type"`
	Body        int       `gorm:"column:body" json:"body"`
	TaskID      string    `gorm:"column:task_id" json:"task_id"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName ..
func (t *FollowUps) TableName() string {
	return "public.follow_ups"
}
