package dbmodels

import (
	"time"
)

// Merchant ..
type Merchant struct {
	ID                    int       `gorm:"column:id;primary_key" json:"id"`
	MerchantId            string    `gorm:"column:merchant_id" json:"merchant_id"`
	Name                  string    `gorm:"column:name" json:"name"`             //
	PhoneArea             string    `gorm:"column:phone_area" json:"phone_area"` //?
	PhoneNumber           string    `gorm:"column:phone_number" json:"phone_number"`
	SalesmanId            int       `gorm:"column:salesman_id" json:"salesman_id"`
	IdCard                string    `gorm:"column:id_card" json:"id_card"`                                 //
	ImageIdCard           string    `gorm:"column:image_id_card" json:"image_id_card"`                     //my_id_6464648454543464.jpg
	ImageMerchant         string    `gorm:"column:image_merchant" json:"image_merchant"`                   // merchant.jpg
	ImageMerchantLocation string    `gorm:"column:image_merchant_location" json:"image_merchant_location"` //location.jpg
	Longitude             string    `gorm:"column:longitude" json:"longitude"`
	Latitude              string    `gorm:"column:latitude" json:"latitude"`
	Address               string    `gorm:"column:address" json:"address"`
	BusinessLocation      string    `gorm:"column:business_location" json:"business_location"`
	BusinessType          string    `gorm:"column:business_type" json:"business_type"`
	CategoryType          string    `gorm:"column:category_type" json:"category_type"`
	OwnerName             string    `gorm:"column:owner_name" json:"owner_name"`
	Signature             string    `gorm:"column:signature" json:"signature"` //signature.jpeg
	OperationHour         string    `gorm:"column:operation_hour" json:"operation_hour"`
	CreatedAt             time.Time `gorm:"column:created_at" json:"created_at"` //with nanosecond
	UpdatedAt             time.Time `gorm:"column:updated_at" json:"updated_at"` //with nanosecond
	BestVisit             string    `gorm:"column:best_visit" json:"best_visit"`
	MerchantLocation      int       `gorm:"column:merchant_location" json:"merchant_location"`
	//SecurityQuestionId               int
	//PasswordDigest                   string
	//Answer                           string
	ProvinceId            int    `gorm:"column:province_id" json:"province_id"`
	CityId                int    `gorm:"column:city_id" json:"city_id"`
	DistrictId            int    `gorm:"column:district_id" json:"district_id"`
	VillageId             int    `gorm:"column:village_id" json:"village_id"`
	Dob                   string `gorm:"column:dob" json:"dob"` //YYYY-mm-dd
	Note                  string `gorm:"column:note" json:"note"`
	PhotoMerchantLocation string `gorm:"column:photo_merchant_location" json:"photo_merchant_location"` //merchant_location.jpg
	IndomarcoStatus       bool   `gorm:"column:indomarco_status" json:"indomarco_status"`               //false
	//ActivationIdmStatus              string
	ImageMerchantLocationAdditional1 string `gorm:"column:image_merchant_location_additional_1" json:"image_merchant_location_additional_1"`
	ImageMerchantLocationAdditional2 string `gorm:"column:image_merchant_location_additional_2" json:"image_merchant_location_additional_2"`
	CustomerCode                     string `gorm:"column:customer_code" json:"customer_code"`
	//MerchantLogo                     string
	//PostalCode                       string
	//OwnerTitle                       string
	//OwnerFirstName                   string
	//OwnerLastName                    string
	//OwnerAddress                     string
	//OwnerRt                          int
	//OwnerRw                          int
	//OwnerPostalCode                  string
	//IdType                           int
	//Gender                           int
	//OwnerPhone                       string
	//OwnerOtherPhone                  string
	//Job                              string
	//Email                            string
	//BirthPlace                       string
	//MotherMaidenName                 string
	//ReferralCode                     string
	InstitutionId string `gorm:"column:institution_id" json:"institution_id"`
	//MerchantOutletSign               string
	//DeviceType                       string
	//DeviceGroup                      string //non_youtap
	//DeviceBrand                      string
	//TerminalLabel                    string
	//TerminalProvider                 string
	//RegistrationLongitude            string
	//RegistrationLatitude             string
	//ImageMerchantSelfie              string
	//GroupName                        string
	//PhotoSelfie                      string
	//AgentId                          string
	//AgentName                        string
	//OwnerKabupatenKota               int
	//OwnerKecamatan                   int
	//OwnerKelurahan                   int
	//OwnerProvinsi                    int
	//PaymentMethod                    string
	//TerminalPhoneNumber              string
	//RoseStatus                       bool
	//RoseDesc                         string
	//VersionId                        int
	WithDevice       bool   `gorm:"column:with_device" json:"with_device"`
	ClusterID        string `gorm:"column:cluster_id" json:"cluster_id"`
	Status           string `gorm:"column:status" json:"status"`
	MerchantTypeIdd  int    `gorm:"column:merchant_type_id" json:"merchant_type_id"`
	MPAN             string `gorm:"column:mpan"`
	Salesman         bool   `gorm:"column:salesman"`
	AddressBenchmark string `gorm:"address_benchmark"`
	SalesTypeId      int    `gorm:"column:sales_type_id"`
}

// TableName ..
func (t *Merchant) TableName() string {
	return "public.merchants"
}

// MerchantStatus ..
type MerchantStatus struct {
	ID             uint   `json:"id" gorm:"primary_key:true"`
	MerchantID     string `json:"merchant_id"`
	PhoneNumber    string `json:"merchant_phone"`
	Name           string `json:"merchant_name"`
	Status         string `json:"status"`
	MerchantTypeID uint   `json:"merchant_type_id"`
}

// TableName ..
func (t *MerchantStatus) TableName() string {
	return "public.merchants"
}

// MerchantType ..
type MerchantType struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Priority int    `json:"priority"`
}

// MerchantLinkingIndomarco ..
type MerchantLinkingIndomarco struct {
	ID              uint   `json:"id" gorm:"primary_key:true"`
	MerchantID      string `json:"merchant_id"`
	BusinessType    string `json:"business_type"`
	SalesTypeID     string `json:"sales_type_id"`
	CustomerCode    string `json:"customer_code"`
	IndomarcoStatus bool   `json:"indomarco_status"`
}

// TableName ..
func (t *MerchantLinkingIndomarco) TableName() string {
	return "public.merchants"
}
