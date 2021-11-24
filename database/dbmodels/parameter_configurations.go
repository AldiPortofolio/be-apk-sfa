package dbmodels

import (
	"time"
)

// ParameterConfigurations ..
type ParameterConfigurations struct {
	ID          int       `gorm:"column:id;primary_key" json:"id"`
	Name        string    `gorm:"column:name" json:"name"`
	ParamValue  string    `gorm:"column:param_value" json:"param_value"`
	ParamType   string    `gorm:"column:param_type" json:"param_type"`
	Module      string    `gorm:"column:module" json:"module"`
	Description string    `gorm:"column:description" json:"description"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName ..
func (t *ParameterConfigurations) TableName() string {
	return "public.parameter_configurations"
}
