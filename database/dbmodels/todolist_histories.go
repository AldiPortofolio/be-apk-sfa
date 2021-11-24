package dbmodels

import "time"

// TodoListHistories ..
type TodoListHistories struct {
	ID           int64     `gorm:"column:id;primary_key" json:"id"`
	TodoListId   int64     `gorm:"column:todolist_id" json:"todolist_id"`
	Description  string    `gorm:"column:description" json:"description"`
	Status       string    `gorm:"column:status" json:"status"`
	NewTaskDate  time.Time `gorm:"new_task_date"`
	OldTaskDate  time.Time `gorm:"old_task_date"`
	FotoLocation string    `gorm:"foto_location"`
	Longitude    string    `gorm:"longitude" json:"longitude"`
	Latitude     string    `gorm:"latitude" json:"latitude"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName ..
func (t *TodoListHistories) TableName() string {
	return "public.todolist_histories"
}
