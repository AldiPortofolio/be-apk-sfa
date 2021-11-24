package dbmodels

// Cities ..
type Cities struct {
	ID   int64  `gorm:"column:id;primary_key" json:"id"`
	Name string `gorm:"column:name" json:"name"`
	//ProvinceId		int8	`gorm:"column:province_id" json:"province_id"`
	//CityCode		string	`gorm:"column:city_code" json:"city_code"`
	//CreatedAt		time.Time `gorm:"column:created_at" json:"created_at"`
	//UpdatedAt		time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName ..
func (t *Cities) TableName() string {
	return "public.cities"
}
