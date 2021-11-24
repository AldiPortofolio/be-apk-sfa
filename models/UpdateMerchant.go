package models

// UpdateMerchantReq ..
type UpdateMerchantReq struct {
	Name                  string    `gorm:"column:name" json:"name" example:""`
	OwnerName             string    `gorm:"column:owner_name" json:"owner_name"`
	//SR Merchant
	PhoneNumber           string    `gorm:"column:phone_number" json:"phone_number"`
	IdCard                string    `gorm:"column:id_card" json:"id_card"`
	ImageIdCard           string    `gorm:"column:image_id_card" json:"image_id_card"`
	ImageMerchant         string    `gorm:"column:image_merchant" json:"image_merchant"` //Foto Selfie dengan KTP
	ImageMerchantLocation string    `gorm:"column:image_merchant_location" json:"image_merchant_location"` //Foto Pemilik dan Lokasi
	PhotoMerchantLocation string    `gorm:"column:photo_merchant_location" json:"photo_merchant_location"` //Foto Lokasi Bisnis
	Signature             string    `gorm:"column:signature" json:"signature"`
	Longitude             string    `gorm:"column:longitude" json:"longitude"`
	Latitude              string    `gorm:"column:latitude" json:"latitude"`
	Address               string    `gorm:"column:address" json:"address"`
	BusinessLocation      string    `gorm:"column:business_location" json:"business_location"` //Jenis Lokasi Bisnis
	MerchantLocation      int       `gorm:"column:merchant_location" json:"merchant_location"`  //Lokasi Bisnis
	BusinessType          string    `gorm:"column:business_type" json:"business_type"` //Tipe Bisnis
	CategoryType          string    `gorm:"column:category_type" json:"category_type"` //Category Bisnis
	OperationHour         string    `gorm:"column:operation_hour" json:"operation_hour"`
	BestVisit             string    `gorm:"column:best_visit" json:"best_visit"`
	Dob                   string	`gorm:"column:dob" json:"dob"` //YYYY-mm-dd
	ProvinceId            int       `gorm:"column:province_id" json:"province_id"`
	CityId                int       `gorm:"column:city_id" json:"city_id"`
	DistrictId            int       `gorm:"column:district_id" json:"district_id"`
	VillageId             int       `gorm:"column:village_id" json:"village_id"`
	MerchantId            string    `gorm:"column:merchant_id" json:"merchant_id"`
	AddressBenchmark	  string	`gorm:"address_benchmark" json:"address_benchmark"`
}