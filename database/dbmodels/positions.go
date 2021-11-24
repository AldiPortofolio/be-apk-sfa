package dbmodels

import "time"

// Positions ..
type Positions struct {
	Id             int       `gorm:"column:id;primary_key" json:"id"`
	RoleNAme       string    `gorm:"column:role_name" json:"role_name"`
	PostCode       string    `gorm:"column:post_code" json:"post_code"`
	SalesRoleId    int       `gorm:"column:sales_role_id" json:"sales_role_id"`
	SalesmenId     int       `gorm:"column:salesman_id" json:"salesman_id"`
	RegionableId   int       `gorm:"column:regionable_id" json:"regionable_id"`
	RegionableType string    `gorm:"column:regionable_type" json:"regionable_type"`
	CreatedAt      time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName ..
func (t *Positions) TableName() string {
	return "public.positions"
}
