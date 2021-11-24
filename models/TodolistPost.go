package models

// TodolistPostReq ..
type TodolistPostReq struct {
	TaskID      []int64  `json:"task_id"`
	Label       []string `json:"label"`
	ContentType []string `json:"content_type"`
	Body        []string `json:"body"`

	LabelPhotoMerchant1 string `json:"label_photo_merchant1,omitempty"`
	LabelPhotoMerchant2 string `json:"label_photo_merchant2,omitempty"`
	LabelPhotoMerchant3 string `json:"label_photo_merchant3,omitempty"`
	PhotoMerchant1      string `json:"photo_merchant1,omitempty"`
	PhotoMerchant2      string `json:"photo_merchant2,omitempty"`
	PhotoMerchant3      string `json:"photo_merchant3,omitempty"`

	NewTaskDate string `json:"new_task_date,omitempty"`
	Reason      string `json:"reason,omitempty"`
	TodolistID  int64  `json:"todolist_id,omitempty"`
	OldTaskDate string `json:"old_task_date,omitempty"`

	Village   string `json:"village,omitempty"`
	District  string `json:"district,omitempty"`
	City      string `json:"city,omitempty"`
	Province  string `json:"province,omitempty"`
	Longitude string `json:"longitude,omitempty"`
	Latitude  string `json:"latitude,omitempty"`
	Alamat    string `json:"alamat,omitempty"`
	Patokan   string `json:"patokan,omitempty"`

	Status string `json:"status"`
	Long   string `json:"longitude_merchant"`
	Lat    string `json:"latitude_merchant"`

	VersionID int64 `json:"versionId"`
}

// PostTodolistToDBReq ..
type PostTodolistToDBReq struct {
	TaskID      int64  `json:"task_id"`
	Label       string `json:"label"`
	ContentType string `json:"content_type"`
	Body        string `json:"body"`
}

// PostHistoryTodolistReq ..
type PostHistoryTodolistReq struct {
	TodolistID   int64  `json:"todolist_id"`
	OldTaskDate  string `json:"old_task_date"`
	NewTaskDate  string `json:"new_task_date"`
	Reason       string `json:"reason"`
	FotoLocation string `json:"foto_location,omitempty"`
	Latitude     string `json:"latitude,omitempty"`
	Longitude    string `json:"longitude,omitempty"`
}

// TodolistPostV23Req ..
type TodolistPostV23Req struct {
	TodolistCategoryId string `json:"todolist_category_id"`
	IdMerchant         int64  `json:"id_merchant"`

	TaskID      []int64  `json:"task_id"`
	Label       []string `json:"label"`
	ContentType []string `json:"content_type"`
	Body        []string `json:"body"`

	LabelPhotoMerchant1 string `json:"label_photo_merchant1,omitempty"`
	LabelPhotoMerchant2 string `json:"label_photo_merchant2,omitempty"`
	LabelPhotoMerchant3 string `json:"label_photo_merchant3,omitempty"`
	PhotoMerchant1      string `json:"photo_merchant1,omitempty"`
	PhotoMerchant2      string `json:"photo_merchant2,omitempty"`
	PhotoMerchant3      string `json:"photo_merchant3,omitempty"`

	NewTaskDate string `json:"new_task_date,omitempty"`
	Reason      string `json:"reason,omitempty"`
	TodolistID  int64  `json:"todolist_id,omitempty"`
	OldTaskDate string `json:"old_task_date,omitempty"`

	Village   string `json:"village,omitempty"`
	District  string `json:"district,omitempty"`
	City      string `json:"city,omitempty"`
	Province  string `json:"province,omitempty"`
	Longitude string `json:"longitude,omitempty"`
	Latitude  string `json:"latitude,omitempty"`
	Alamat    string `json:"alamat,omitempty"`
	Patokan   string `json:"patokan,omitempty"`

	Status string `json:"status"`
	Long   string `json:"longitude_merchant"`
	Lat    string `json:"latitude_merchant"`

	VersionID int64 `json:"versionId"`

	Acquitisions Acquitisions `json:"acquitisions"`
}

// Acquitisions ..
type Acquitisions struct {
	MerchantId  string `form:"merchant_id" json:"merchant_id"`
	Name        string `form:"name" json:"name"`
	PhoneNumber string `form:"phone_number" json:"phone_number"`
	SalesPhone  string `form:"sales_phone" json:"sales_phone"`
	IdCard      string `form:"id_card" json:"id_card"`
	ImageIdCard string `form:"image_id_card" json:"image_id_card"` //images

	ImageMerchant         string `form:"image_merchant" json:"image_merchant"`                   //images
	ImageMerchantLocation string `form:"image_merchant_location" json:"image_merchant_location"` //images
	Longitude             string `form:"longitude" json:"longitude"`
	Latitude              string `form:"latitude" json:"latitude"`
	Address               string `form:"address" json:"address"`

	LokasiUsaha  string `form:"lokasi_usaha" json:"lokasi_usaha"`
	BusinessType string `form:"business_type" json:"business_type"`
	CategoryType string `form:"category_type" json:"category_type"`

	OwnerName        string `form:"owner_name" json:"owner_name"`
	Signature        string `form:"signature" json:"signature"` //images
	OperationHour    string `form:"operation_hour" json:"operation_hour"`
	BusinessLocation string `form:"business_location" json:"business_location"`
	BestVisit        string `form:"best_visit" json:"best_visit"`

	ProvinceId                      string `form:"province_id" json:"province_id"`
	City                            string `form:"city_id" json:"city_id"`
	District                        string `form:"district_id" json:"district_id"`
	Village                         string `form:"village_id" json:"village_id"`
	BirthDate                       string `form:"dob" json:"dob"`
	ImageMerchantLocationAdditional string `form:"image_merchant_location_additional_1" json:"image_merchant_location_additional_1"`
	Note                            string `form:"note" json:"note"`

	PhotoMerchantLocation            string `form:"photo_merchant_location" json:"photo_merchant_location"` //images
	Institution                      string `form:"institution" json:"institution"`
	CustomerCode                     string `form:"customer_code" json:"customer_code" `
	ImageMerchantLocationAdditional2 string `form:"image_merchant_location_additional_2" json:"image_merchant_location_additional_2"`
	Device                           int16  `form:"device" json:"device"`

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

// TodolistPostV24Req ..
type TodolistPostV24Req struct {
	TodolistCategoryId string `json:"todolist_category_id"`
	IdMerchant         int64  `json:"id_merchant"`

	TaskID      		[]int64  `json:"task_id"`
	Label       		[]string `json:"label"`
	ContentType 		[]string `json:"content_type"`
	Body        		[]string `json:"body"`

	LabelPhotoMerchant1 string `json:"label_photo_merchant1,omitempty"`
	LabelPhotoMerchant2 string `json:"label_photo_merchant2,omitempty"`
	LabelPhotoMerchant3 string `json:"label_photo_merchant3,omitempty"`
	PhotoMerchant1      string `json:"photo_merchant1,omitempty"`
	PhotoMerchant2      string `json:"photo_merchant2,omitempty"`
	PhotoMerchant3      string `json:"photo_merchant3,omitempty"`

	NewTaskDate 		string `json:"new_task_date,omitempty"`
	Reason      		string `json:"reason,omitempty"`
	TodolistID  		int64  `json:"todolist_id,omitempty"`
	OldTaskDate 		string `json:"old_task_date,omitempty"`

	Village   			string `json:"village,omitempty"`
	District  			string `json:"district,omitempty"`
	City      			string `json:"city,omitempty"`
	Province  			string `json:"province,omitempty"`
	Longitude 			string `json:"longitude,omitempty"`
	Latitude  			string `json:"latitude,omitempty"`
	Alamat    			string `json:"alamat,omitempty"`
	Patokan   			string `json:"patokan,omitempty"`

	Status 				string `json:"status"`
	Long   				string `json:"longitude_merchant"`
	Lat    				string `json:"latitude_merchant"`

	VersionID 			int64 `json:"versionId"`

	Acquitisions 		AcquitisionsV24 `json:"acquitisions"`
}

// AcquitisionsV24 ..
type AcquitisionsV24 struct {
	MerchantId  						string `form:"merchant_id" json:"merchant_id"`
	Name        						string `form:"name" json:"name"`
	PhoneNumber 						string `form:"phone_number" json:"phone_number" `
	NMID			 					string 	`form:"nmid" json:"nmid"`
	MPAN			 					string 	`form:"mpan" json:"mpan"`
	StoreNamePrePrinted					string 	`form:"store_name_preprinted" json:"store_name_preprinted"`
	EducationMerchant		 			bool 	`form:"education" json:"education"`

	ImageIdCard 						string `form:"image_id_card" json:"image_id_card"` //images
	ImageMerchant         				string `form:"image_merchant" json:"image_merchant"` //images
	ImageMerchantLocation 				string `form:"image_merchant_location" json:"image_merchant_location" ` //images
	ImageMerchantLocationAdditional 	string `form:"image_merchant_location_additional_1" json:"image_merchant_location_additional_1"`
	ImageMerchantLocationAdditional2    string `form:"image_merchant_location_additional_2" json:"image_merchant_location_additional_2"`
	PhotoMerchantLocation 				string `form:"photo_merchant_location" json:"photo_merchant_location"`  //images
	Signature        					string `form:"signature" json:"signature" ` //images
	PhotoLocationLeft					string 	`form:"photo_location_left" json:"photo_location_left"`
	PhotoLocationRight					string 	`form:"photo_location_right" json:"photo_location_right"`
	PhotoQRPreprinted					string 	`form:"photo_qr_preprinted" json:"photo_qr_preprinted"`

	Longitude             				string `form:"longitude" json:"longitude"`
	Latitude              				string `form:"latitude" json:"latitude"`
	Address               				string `form:"address" json:"address"`
	AddressBenchmark					string 	`form:"address_benchmark" json:"address_benchmark"`
	ProvinceId              			string `form:"province_id" json:"province_id"`
	City                    			string `form:"city_id" json:"city_id"`
	District                			string `form:"district_id" json:"district_id"`
	Village                 			string `form:"village_id" json:"village_id"`

	LokasiUsaha  						string `form:"lokasi_usaha" json:"lokasi_usaha"`
	LokasiUsahaValue  					string `form:"lokasi_usaha_value" json:"lokasi_usaha_value"`
	BusinessType 						string `form:"business_type" json:"business_type"`
	CategoryType 						string `form:"category_type" json:"category_type"`
	OperationHour    					string `form:"operation_hour" json:"operation_hour"`
	BusinessLocation 					string `form:"business_location" json:"business_location"`
	BusinessLocationCode 				string `form:"business_location_code" json:"business_location_code"`
	BestVisit        					string `form:"best_visit" json:"best_visit"`

	Note                            	string `form:"note" json:"note"`
	Institution           				string `form:"institution" json:"institution"`
	Customer_code         				string `form:"customer_code" json:"customer_code" `
	Device                				int16  `form:"device" json:"device"`
	VersionId 							int    `form:"version_id" json:"version_id"`

	SalesPhone 							string 	`form:"sales_phone" json:"sales_phone"`
	Salesman 							bool 	`form:"sales" json:"sales;omitempty"`
	SalesID 							int 	`form:"sales_id" json:"sales_id"`
	SalesName 							string 	`form:"sales_name" json:"sales_name"`
	SalesTypeId 						int 	`form:"sales_type_id" json:"sales_type_id"`
	SalesCompanyCode 					string 	`form:"company_code" json:"company_code"`

	IdCard      						string `form:"id_card" json:"id_card"`
	OwnerName        					string `form:"owner_name" json:"owner_name"`
	BirthDate                       	string `form:"dob" json:"dob"`

	OwnerBirthPlace                     string `form:"owner_birth_place" json:"owner_birth_place"`
	OwnerGender							string `form:"owner_gender" json:"owner_gender"`
	OwnerAddress               			string `form:"owner_address" json:"owner_address"`
	OwnerProvinceId              		string `form:"owner_province_id" json:"owner_province_id"`
	OwnerCityId                    		string `form:"owner_city_id" json:"owner_city_id"`
	OwnerDistrictId                		string `form:"owner_district_id" json:"owner_district_id"`
	OwnerVillageId                 		string `form:"owner_village_id" json:"owner_village_id"`
	OwnerRT								string `form:"owner_rt" json:"owner_rt"`
	OwnerRW								string `form:"owner_rw" json:"owner_rw"`
	OwnerJob							string `form:"owner_job" json:"owner_job"`
	IdCardExpiredDate					string `form:"id_card_expired_date" json:"id_card_expired_date"`
	OcrPercentage						float32 `form:"ocr_percentage" json:"ocr_percentage"`
	MerchantGroupName					string `form:"merchant_group_name" json:"merchant_group_name"`
}
