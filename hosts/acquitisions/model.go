package acquitisions

// AcquitisionsReq ..
type AcquitisionsReq struct {
	MerchantId  string `form:"merchant_id" json:"merchant_id"`
	Name        string `form:"name" json:"name" binding:"required"`
	PhoneNumber string `form:"phone_number" json:"phone_number" binding:"required"`
	SalesPhone  string `form:"sales_phone" json:"sales_phone" binding:"required"`
	IdCard      string `form:"id_card" json:"id_card" binding:"required"`
	ImageIdCard string `form:"image_id_card" json:"image_id_card" binding:"required"` //images

	ImageMerchant         string `form:"image_merchant" json:"image_merchant" binding:"required"`                   //images
	ImageMerchantLocation string `form:"image_merchant_location" json:"image_merchant_location" binding:"required"` //images
	Longitude             string `form:"longitude" json:"longitude" binding:"required"`
	Latitude              string `form:"latitude" json:"latitude" binding:"required"`
	Address               string `form:"address" json:"address" binding:"required"`

	LokasiUsaha  string `form:"lokasi_usaha" json:"lokasi_usaha" binding:"required"`
	BusinessType string `form:"business_type" json:"business_type" binding:"required"`
	CategoryType string `form:"category_type" json:"category_type" binding:"required"`

	OwnerName        string `form:"owner_name" json:"owner_name" binding:"required"`
	Signature        string `form:"signature" json:"signature" binding:"required"` //images
	OperationHour    string `form:"operation_hour" json:"operation_hour" binding:"required"`
	BusinessLocation string `form:"business_location" json:"business_location" binding:"required"`
	BestVisit        string `form:"best_visit" json:"best_visit"`

	ProvinceId                      string `form:"province_id" json:"province_id" binding:"required"`
	City                            string `form:"city_id" json:"city_id" binding:"required"`
	District                        string `form:"district_id" json:"district_id" binding:"required"`
	Village                         string `form:"village_id" json:"village_id" binding:"required"`
	BirthDate                       string `form:"dob" json:"dob" binding:"required"`
	ImageMerchantLocationAdditional string `form:"image_merchant_location_additional_1" json:"image_merchant_location_additional_1"`
	Note                            string `form:"note" json:"note"`

	PhotoMerchantLocation string `form:"photo_merchant_location" json:"photo_merchant_location"` //images
	Institution           string `form:"institution" json:"institution"`
	CustomerCode          string `form:"customer_code" json:"customer_code" `
	Businesslocation      string `form:"image_merchant_location_additional_2" json:"image_merchant_location_additional_2"`
	Device                int16  `form:"device" json:"device"`

	VersionId int  `form:"version_id" json:"version_id"`
	Salesman  bool `form:"sales" json:"sales;omitempty"`

	SalesID          int    `form:"sales_id" json:"sales_id"`
	SalesName        string `form:"sales_name" json:"sales_name"`
	SalesCompanyCode string `form:"company_code" json:"company_code"`

	EducationMerchant   bool   `form:"education" json:"education"`
	NMID                string `form:"nmid" json:"nmid"`
	MPAN                string `form:"mpan" json:"mpan"`
	StoreNamePrePrinted string `form:"store_name_preprinted" json:"store_name_preprinted"`
	PhotoLocationLeft   string `form:"photo_location_left" json:"photo_location_left"`
	PhotoLocationRight  string `form:"photo_location_right" json:"photo_location_right"`
	PhotoQRPreprinted   string `form:"photo_qr_preprinted" json:"photo_qr_preprinted"`
}

// AcquitisionsRes ..
type AcquitisionsRes struct {
	Data interface{} `json:"data"`
	Meta Meta        `json:"meta"`
}

// Meta ..
type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  bool   `json:"status"`
}
