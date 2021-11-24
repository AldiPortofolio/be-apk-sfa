package dbmodels

import "time"

// Villages ..
type Villages struct {
	ID   int64  `gorm:"column:id;primary_key" json:"id"`
	Name string `gorm:"column:name" json:"name"`
	//CityId			int8	`gorm:"column:city_id" json:"city_id"`
	//DistrictCode	string	`gorm:"column:district_code" json:"district_code"`
	VillageCode string `json:"village_code"`
	DistrictId  uint   `json:"district_id"`
	//SubAreas            []SubArea `gorm:"many2many:sub_areas_villages;"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName ..
func (t *Villages) TableName() string {
	return "public.villages"
}
