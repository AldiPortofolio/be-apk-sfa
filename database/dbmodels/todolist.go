package dbmodels

import "time"

// TodoLists ..
type TodoLists struct {
	ID                    int64     `gorm:"column:id;primary_key" json:"id"`
	TodoListCategoryId    int       `gorm:"column:category" json:"category"`
	TaskDate              time.Time `gorm:"task_date"`
	TaskDateString        string    `json:"task_date"`
	ActionDate            time.Time `gorm:"action_date"`
	MerchantID            string    `gorm:"column:mid" json:"merchant_id"`
	MerchantName          string    `gorm:"column:merchant_name" json:"merchant_name"`
	SalesPhone            string    `gorm:"column:sales_phone" json:"sales_phone"`
	VillageID             int       `json:"village_id"`
	Status                string    `gorm:"column:status" json:"status"`
	Notes                 string    `gorm:"column:notes" json:"notes"`
	Longitude             string    `gorm:"longitude" json:"longitude"`
	Latitude              string    `gorm:"latitude" json:"latitude"`
	PendingTaskDate       time.Time `gorm:"pending_task_date"`
	PendingTaskDateString string    `json:"pending_task_date"`
	CreatedAt             time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt             time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName ..
func (t *TodoLists) TableName() string {
	return "public.todolists"
}
