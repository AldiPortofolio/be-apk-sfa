package dbmodels

// Provinces ..
type Provinces struct {
	ID   int64  `gorm:"column:id;primary_key" json:"id"`
	Name string `gorm:"column:name" json:"name"`
	//CountryId			int8	`gorm:"column:country_id" json:"country_id"`
	//ProvinceCode		string	`gorm:"column:province_code" json:"province_code"`
	//CreatedAt           time.Time `gorm:"column:created_at" json:"created_at"`
	//UpdatedAt           time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName ..
func (t *Provinces) TableName() string {
	return "public.provinces"
}
