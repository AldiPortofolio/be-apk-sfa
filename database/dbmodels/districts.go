package dbmodels

// Districts ..
type Districts struct {
	ID   int64  `gorm:"column:id;primary_key" json:"id"`
	Name string `gorm:"column:name" json:"name"`
	//CityId			int8	`gorm:"column:city_id" json:"city_id"`
	//DistrictCode	string	`gorm:"column:district_code" json:"district_code"`
	//CreatedAt		time.Time `gorm:"column:created_at" json:"created_at"`
	//UpdatedAt		time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName ..
func (t *Districts) TableName() string {
	return "public.districts"
}
