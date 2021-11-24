package models

// UpdatePhotoProfilSalesReq ..
type UpdatePhotoProfilSalesReq struct {
	Phone string `json:"phone" form:"phone"`
	Photo string `json:"photo" form:"photo"`
}

// UpdatePhotoProfilSalesRes ..
type UpdatePhotoProfilSalesRes struct {
	ResponseCode    string                     `json:"response_code" form:"response_code"`
	SalesName       string                     `json:"sales_name" form:"sales_name"`
	Email           string                     `json:"email" form:"email"`
	DescriptionCode string                     `json:"description_code" form:"description_code"`
	Phone           string                     `json:"phone" form:"phone"`
	Photo           string                     `json:"photo" form:"photo"`
	SessionToken    string                     `json:"session_token" form:"session_token"`
	Status          string                     `json:"status" form:"status"`
	SalesID         string                     `json:"sales_id" form:"sales_id"`
	SfaID           string                     `json:"sfa_id" form:"sfa_id"`
	AreaAquisitions UpdatePhotoProfilSalesRes1 `json:"area_aquisitions" form:"area_aquisitions"`
}

// UpdatePhotoProfilSalesRes1 ..
type UpdatePhotoProfilSalesRes1 struct {
	Province  string                       `json:"province" form:"province"`
	City      string                       `json:"city" form:"city"`
	Locations []UpdatePhotoProfilSalesRes2 `json:"locations" form:"locations"`
}

// UpdatePhotoProfilSalesRes2 ..
type UpdatePhotoProfilSalesRes2 struct {
	Village  string `json:"village" form:"village"`
	District string `json:"district" form:"district"`
}

// GetDataSalesLocationDB table salesmen and location_areas
type GetDataSalesLocationDB struct {
	ID           int         `gorm:"column:id;primary_key" json:"id"`
	FirstName    string      `gorm:"column:first_name" json:"first_name"`
	LastName     string      `gorm:"column:last_name" json:"last_name"`
	PhoneNumber  string      `gorm:"column:phone_number" json:"phone_number"`
	Email        string      `gorm:"column:email" json:"email"`
	Status       int         `gorm:"column:status" json:"status"`
	SessionToken string      `gorm:"column:session_token" json:"session_token"`
	SfaID        string      `gorm:"column:sfa_id" json:"sfa_id"`
	ProvinceID   int         `gorm:"column:province_id" json:"province_id"`
	CityID       int         `gorm:"column:city_id" json:"city_id"`
	SalesId      string      `gorm:"column:sales_id" json:"sales_id"`
	Locations    []Locations `gorm:"column:locations" json:"locations"`
}

// Locations ..
type Locations struct {
	Village  []Villages `json:"village"`
	District Districts  `json:"district"`
}

// Villages ..
type Villages struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

// Districts ..
type Districts struct {
	Label string `json:"label"`
	Value string `json:"value"`
}
