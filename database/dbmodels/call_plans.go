package dbmodels

import "time"

// CallPlans ..
type CallPlans struct {
	Id                  int64     `gorm:"column:id;primary_key" json:"id"`
	SalesId             int64     `gorm:"column:sales_id" json:"sales_id"`
	SalesName           string    `gorm:"column:sales_name" json:"sales_name"`
	SalesPhone          string    `gorm:"column:sales_phone" json:"sales_phone"`
	Cycle               int64     `gorm:"column:cycle" json:"cycle"`
	SubArea             string    `gorm:"column:sub_area" json:"sub_area"`
	Clusters            string    `gorm:"column:clusters" json:"clusters"`
	SuccessCall         string    `gorm:"column:success_call" json:"success_call"`
	EffectiveCall       string    `gorm:"column:effective_call" json:"effective_call"`
	ClusterCoverageArea string    `gorm:"column:cluster_coverage_area" json:"cluster_coverage_area"`
	Status              int64     `gorm:"column:status" json:"status"`
	CallPlanDate        time.Time `gorm:"column:call_plan_date" json:"call_plan_date"`
	Period              int64     `gorm:"column:period" json:"period"`
	CreatedAt           time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt           time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName ..
func (t *CallPlans) TableName() string {
	return "public.call_plans"
}
