package dbmodels

import (
	"time"
)

// Salesmen is dbmodels for table salesmen
type Salesmen struct {
	ID                   int       `gorm:"column:id;primary_key" json:"id"`
	FirstName            string    `gorm:"column:first_name" json:"first_name"`
	LastName             string    `gorm:"column:last_name" json:"last_name"`
	PasswordDigest       string    `gorm:"column:password_digest" json:"password_digest"`
	IDNumber             string    `gorm:"column:id_number" json:"id_number"`
	Dob                  time.Time `gorm:"column:dob" json:"dob"`
	PhoneArea            string    `gorm:"column:phone_area" json:"phone_area"`
	PhoneNumber          string    `gorm:"column:phone_number" json:"phone_number"`
	Gender               int       `gorm:"column:gender" json:"gender"`
	CompanyCode          string    `gorm:"column:company_code" json:"company_code"`
	ProvinceID           int       `gorm:"column:province_id" json:"province_id"`
	CreatedAt            time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt            time.Time `gorm:"column:updated_at" json:"updated_at"`
	FailedLoginAttempts  int       `gorm:"column:failed_login_attempts" json:"failed_login_attempts"`
	LastLoginAt          time.Time `gorm:"last_login_at" json:"last_login_at"`
	DeviceID             string    `gorm:"column:device_id" json:"device_id"`
	Photo                string    `gorm:"column:photo" json:"photo"`
	IDCard               string    `gorm:"column:id_card" json:"id_card"`
	Email                string    `gorm:"column:email" json:"email"`
	Status               int       `gorm:"column:status" json:"status"`
	DeviceToken          string    `gorm:"column:device_token" json:"device_token"`
	SessionToken         string    `gorm:"column:session_token" json:"session_token"`
	SalesId              string    `gorm:"column:sales_id" json:"sales_id"`
	SessionExpiredAt     time.Time `gorm:"column:session_expired_at" json:"session_expired_at"`
	Address              string    `gorm:"column:address" json:"address"`
	CityID               int       `gorm:"column:city_id" json:"city_id"`
	DistrictID           int       `gorm:"column:district_id" json:"district_id"`
	VillageID            int       `gorm:"column:village_id" json:"village_id"`
	PostCode             string    `gorm:"column:postcode" json:"postcode"`
	BirthPlace           string    `gorm:"column:birth_place" json:"birth_place"`
	SalesID              string    `gorm:"column:sales_id" json:"sales_id"`
	WorkDate             time.Time `gorm:"column:work_date" json:"work_date"`
	Occupation           string    `gorm:"column:occupation" json:"occupation"`
	SfaID                string    `gorm:"column:sfa_id" json:"sfa_id"`
	FunctionalPositionId int       `gorm:"column:functional_position_id" json:"functional_position_id"`
	FunctionalPosition   string    `gorm:"column:functional_position" json:"functional_position"`
	SalesType            string    `gorm:"column:sales_type" json:"sales_type"`
	FirebaseToken        string    `gorm:"column:firebase_token" json:"firebase_token"`
	SalesTypeId          int  	   `gorm:"column:sales_type_id" json:"sales_type_id"`
}

// TableName ..
func (t *Salesmen) TableName() string {
	return "public.salesmen"
}
