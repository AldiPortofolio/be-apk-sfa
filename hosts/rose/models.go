package rose

// Response ..
type Response struct {
	Rc   string      `json:"rc"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// InquiryMerchantResponse ..
type InquiryMerchantResponse struct {
	PhoneNumber            string `json:"phoneNumber"`
	Mid                    string `json:"mid"`
	MerchantName           string `json:"merchantName"`
	OutletId               string `json:"outletId"`
	HostType               string `json:"hostType"`
	MerchantCriteria       string `json:"merchantCriteria"`
	MerchantNamePreprinted string `json:"merchantNamePreprinted"`
}

// LookUpGroupResponse ..
type LookUpGroupResponse struct {
	ID          int    `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	LookupGroup string `json:"lookupGroup"`
}

// UserCategoryResponse ..
type UserCategoryResponse struct {
	ID     int    `json:"id"`
	Code   string `json:"code"`
	Name   string `json:"name"`
	Logo   string `json:"logo"`
	AppId  string `json:"appId"`
	Notes  string `json:"notes"`
	Seq    int    `json:"seq"`
	Status string `json:"status"`
}

// FindByMIDResponse ..
type FindByMIDResponse struct {
	MerchantCriteria  string `json:"merchantCriteria"`
	MerchantGroup     string `json:"merchantGroup"`
	MerchantGroupId   string `json:"merchantGroupId"`
	MerchantName      string `json:"merchantName"`
	HostType          string `json:"mpan"`
	OtpOption         string `json:"otpOption"`
	PhoneNumber       string `json:"phoneNumber"`
	ProfilePictureUrl string `json:"profilePictureUrl"`
}

// ReplaceMidMpanRequest ..
type ReplaceMidMpanRequest struct {
	Mid                 string `json:"mid"`
	Mpan                string `json:"mpan"`
	Nmid                string `json:"nmid"`
	StoreNamePreprinted string `json:"storeNamePreprinted"`
	StorePhoneNumber    string `json:"storePhoneNumber"`
}

// FindByPhoneNumberResponse ..
type FindByPhoneNumberResponse struct {
	MID               string `json:"mid"`
	OwnerName         string `json:"ownerName"`
	StoreName         string `json:"storeName"`
	Address           string `json:"completeAddress"`
	AddressProvince   string `json:"addressProvince"`
	AddressCity       string `json:"addressCity"`
	AddressDistrict   string `json:"addressDistrict"`
	AddressVillage    string `json:"addressVillage"`
	AddressProvinceId int64  `json:"addressProvinceId"`
	AddressCityId     int64  `json:"addressCityId"`
	AddressDistrictId int64  `json:"addressDistrictId"`
	AddressVillageId  int64  `json:"addressVillageId"`
	MemberType        string `json:"memberType"`
	Category          string `json:"category"`
	NMID              string `json:"nmid"`
	Tid               string `json:"tid"`
	Mpan              string `json:"mpan"`
	Email             string `json:"email"`

	//"merchantGroupName": "Ottopay",
	//"merchantStatus": false,
	//"partnerCustomerId": "",
	//"profilePictureUrl": "",
	//"otpOption": "",
	//"tagList": null
}

type UpdateDataMerchantReq struct {
	MerchantGroupID   string `json:"MerchantGroupId"`
	PartnerCustomerID string `json:"PartnerCustomerId"`
	UserCategoryCode  string `json:"UserCategoryCode"`
	Address           string `json:"address"`
	City              string `json:"city"`
	District          string `json:"district"`
	ExpireDate        string `json:"expireDate"`
	LoanBankAccount   string `json:"loanBankAccount"`
	MID               string `json:"mid"`
	PartnerCode       string `json:"partnerCode"`
	PostalCode        string `json:"postalCode"`
	Province          string `json:"province"`
	Rt                string `json:"rt"`
	Rw                string `json:"rw"`
	Village           string `json:"village"`
}

type UpdateDataMerchantV2Req struct {
	MerchantGroupID   int    `json:"MerchantGroupId"`
	PartnerCustomerID string `json:"PartnerCustomerId"`
	UserCategoryCode  string `json:"UserCategoryCode"`
	Address           string `json:"address"`
	BestVisit         string `json:"bestVisit"`
	CategoryBisnis    string `json:"categoryBisnis"`
	City              string `json:"city"`
	District          string `json:"district"`
	ExpireDate        string `json:"expireDate"`
	JenisLokasiBisnis string `json:"jenisLokasiBisnis"`
	Latitude          string `json:"latitude"`
	LoanBankAccount   string `json:"loanBankAccount"`
	LokasiBisnis      string `json:"lokasiBisnis"`
	Longitude         string `json:"longitude"`
	MID               string `json:"mid"`
	OperationHour     string `json:"operationHour"`
	OwnerAddress      string `json:"ownerAddress"`
	OwnerCity         string `json:"ownerCity"`
	OwnerKecamatan    string `json:"OwnerKecamatan"`
	OwnerKelurahan    string `json:"ownerKelurahan"`
	OwnerName         string `json:"ownerName"`
	OwnerNoId         string `json:"ownerNoId"`
	OwnerPhoneNumber  string `json:"ownerPhoneNumber"`
	OwnerPostalCode   string `json:"ownerPostalCode"`
	OwnerProvinsi     string `json:"ownerProvinsi"`
	OwnerRt           string `json:"ownerRt"`
	OwnerRw           string `json:"ownerRw"`
	OwnerTanggalLahir string `json:"ownerTanggalLahir"`
	PartnerCode       string `json:"partnerCode"`
	Patokan           string `json:"patokan"`
	PostalCode        string `json:"postalCode"`
	Province          string `json:"province"`
	SrID              string `json:"srId"`
	StoreName         string `json:"storeName"`
	StorePhoneNumber  string `json:"storePhoneNumber"`
	TipeBisnis        string `json:"tipeBisnis"`
	UserCategory      string `json:"userCategory"`
	Village           string `json:"village"`
}

// SalesByReportResponse ..
type SalesByReportResponse struct {
	OttoPay    int    `json:"ottopay"`
	SfaOnly    int    `json:"sfaOnly"`
	SalesPhone string `json:"salesPhone"`
}

// ResHistoryPencapaianSalesDto ...
type ResHistoryPencapaianSalesDto struct {
	AchievementDay   int                         `json:"achievementDay"`
	AcquisitionData []HistoryPencapaianSalesDto `json:"acquisitionData"`
	Target          int                         `json:"target"`
}

// HistoryPencapaianSalesDto ...
type HistoryPencapaianSalesDto struct {
	Created   string `json:"created"`
	Address   string `json:"Address"`
	Status    string `json:"status"`
	StoreName string `json:"storeName"`
}

// PencapaianSalesRequest ..
type PencapaianSalesRequest struct {
	SalesPhone             string 	`json:"salesPhone"`
	SrId                   string 	`json:"srId"`
	VillageId	           []string `json:"villageId"`
}

// PencapaianSalesResponse ..
type PencapaianSalesResponse struct {
	AcquisitionData            []AcquisitionDataPencapaianSales `json:"acquisitionData"`
}

type AcquisitionDataPencapaianSales struct {
	Address          string 	`json:"address"`
	ID               int64  	`json:"id"`
	ImageMerchant    string 	`json:"image_merchant"`
	JoinAt           string	    `json:"join_at"`
	MerchantCategory string 	`json:"merchant_category"`
	MerchantID       string 	`json:"merchant_id"`
	MerchantLevel    string 	`json:"merchant_level"`
	MerchantStatus   string 	`json:"merchant_status"`
	MerchantType     string 	`json:"merchant_type"`
	Name             string 	`json:"name"`
	PhoneNumber      string 	`json:"phone_number"`
}