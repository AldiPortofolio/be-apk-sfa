package rose

import (
	"encoding/json"
	"fmt"
	https "ottosfa-api-apk/utils/http"
	"strconv"

	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	ottologger "ottodigital.id/library/logger"
	// ottologgerv2 "ottodigital.id/library/logger/v2"
)

// Env ..
type Env struct {
	Host   string `envconfig:"ROSE_API_SERVICE_HOST" default:"http://13.228.25.85:8899/rose-api-service/v0.0.1"`
	HostV2 string `envconfig:"ROSE_API_SERVICE_HOST_V2" default:"http://13.228.25.85:8899/rose-api-service/v0.0.2"`

	EndpointInquiryMerchant string `envconfig:"ROSE_API_SERVICE_ENDPOINT_INQUIRY_MERCHANT" default:"/inquiry-merchant/find"`
	EndpointLookUpGroup     string `envconfig:"ROSE_API_SERVICE_ENDPOINT_LOOK_UP_GROUP" default:"/lookup/lookupgroup"`
	EndpointUserCategory    string `envconfig:"ROSE_API_SERVICE_ENDPOINT_USER_CATEGORY" default:"/user-category/find"`
	EndpointFindByMid       string `envconfig:"ROSE_API_SERVICE_ENDPOINT_FIND_BY_MID" default:"/inquiry-merchant/find-by-mid"`
	EndpointReplaceMIDMpan  string `envconfig:"ROSE_API_SERVICE_ENDPOINT_FIND_BY_MID" default:"/merchant/replace-mid-mpan"`
	EndpointCheckIdCard     string `envconfig:"ROSE_API_SERVICE_ENDPOINT_FIND_BY_MID" default:"/owner/get-merchant"`
	EndpointSalesByReport   string `envconfig:"ROSE_API_SERVICE_ENDPOINT_SALES_BY_REPORT" default:"/pencapaian-sales"`
	EndpointSalesHistoryReport   string `envconfig:"ROSE_API_SERVICE_ENDPOINT_SALES_BY_REPORT" default:"/pencapaian-sales/history"`
	EndpointPencapaianSales string `envconfig:"ROSE_API_SERVICE_ENDPOINT_PENCAPAIAN_SALES_DETAIL" default:"/pencapaian-sales/detail"`
	KeyAppId                string `envconfig:"ROSE_API_SERVICE_ENDPOINT_KEY_APP_ID" default:"3"`
	HealthCheckKey          string `envconfig:"ROSE_API_SERVICE_HEALTH_CHECK_KEY" default:"OTTOSFA-API-APK_HEALTH_CHECK:ROSE_API_SERVICE"`

	HostRoseOPService      string `envconfig:"ROSE_OP_SERVICE_HOST" default:"http://13.228.25.85:8914/rose-op-service/v0.0.1"`
	EndpointMerchantProfil string `envconfig:"ROSE_OP_SERVICE_ENDPOINT_MERCHANT_PROFIL" default:"/merchant/info/profile"`

	EndpointUpdateMerchant string `envconfig:"ROSE_API_SERVICE_ENDPOINT_UPDATE_MERCHANT" default:"/updated-data-merchant"`
}

var (
	roseEnv Env
)

// init ..
func init() {
	err := envconfig.Process("ROSE API SERVICE", &roseEnv)
	if err != nil {
		fmt.Println("Failed to get ROSE API SERVICE env:", err)
	}
}

// InquiryMerchant ..
func InquiryMerchant(TID string, MPAN string) (InquiryMerchantResponse, error) {
	sugarLogger := ottologger.GetLogger()
	var res Response
	var inquiryMerchantRes InquiryMerchantResponse

	urlSvr := roseEnv.Host + roseEnv.EndpointInquiryMerchant

	jsonData := map[string]string{"tid": TID, "mpan": MPAN}
	data, err := https.HTTPPost(urlSvr, jsonData)
	if err != nil {
		sugarLogger.Error("Level: Error", zap.String("Error: ", err.Error()))
		return inquiryMerchantRes, err
	}

	if err = json.Unmarshal(data, &res); err != nil {
		sugarLogger.Error("Level: Error", zap.String("Failed to unmarshaling response: ", err.Error()))
		return inquiryMerchantRes, err
	}

	if res.Rc == "04" {
		sugarLogger.Error("Level: Error", zap.String("Failed to get data: ", errors.New(res.Msg).Error()))
		return inquiryMerchantRes, errors.New(res.Msg)
	}

	bdata, _ := json.Marshal(res.Data)
	if err = json.Unmarshal(bdata, &inquiryMerchantRes); err != nil {
		sugarLogger.Error("Level: Error", zap.String("Failed to unmarshaling response data: ", err.Error()))
		return inquiryMerchantRes, err
	}

	return inquiryMerchantRes, nil
}

// LookUpGroup ..
func LookUpGroup(lookupGroup string) ([]LookUpGroupResponse, error) {
	sugarLogger := ottologger.GetLogger()
	var res Response
	var lookUpGroupResponse []LookUpGroupResponse

	urlSvr := roseEnv.Host + roseEnv.EndpointLookUpGroup

	jsonData := map[string]string{"lookupGroup": lookupGroup}
	data, err := https.HTTPPost(urlSvr, jsonData)
	if err != nil {
		sugarLogger.Error("Level: Error", zap.String("Error: ", err.Error()))
		return lookUpGroupResponse, err
	}

	if err = json.Unmarshal(data, &res); err != nil {
		sugarLogger.Error("Level: Error", zap.String("Failed to unmarshaling response: ", err.Error()))
		return lookUpGroupResponse, err
	}

	if res.Rc == "04" {
		sugarLogger.Error("Level: Error", zap.String("Failed to get data: ", errors.New(res.Msg).Error()))
		return lookUpGroupResponse, errors.New(res.Msg)
	}

	bdata, _ := json.Marshal(res.Data)
	if err = json.Unmarshal(bdata, &lookUpGroupResponse); err != nil {
		sugarLogger.Error("Level: Error", zap.String("Failed to unmarshaling response data: ", err.Error()))
		return lookUpGroupResponse, err
	}

	return lookUpGroupResponse, nil
}

// UserCategory ..
func UserCategory() ([]UserCategoryResponse, error) {
	sugarLogger := ottologger.GetLogger()
	var res Response
	var userCategoryResponse []UserCategoryResponse

	urlSvr := roseEnv.Host + roseEnv.EndpointUserCategory

	jsonData := map[string]string{"appId": roseEnv.KeyAppId}
	data, err := https.HTTPPost(urlSvr, jsonData)
	if err != nil {
		sugarLogger.Error("Level: Error", zap.String("Error: ", err.Error()))
		return userCategoryResponse, err
	}

	if err = json.Unmarshal(data, &res); err != nil {
		sugarLogger.Error("Level: Error", zap.String("Failed to unmarshaling response: ", err.Error()))
		return userCategoryResponse, err
	}

	if res.Rc == "04" {
		sugarLogger.Error("Level: Error", zap.String("Failed to get data: ", errors.New(res.Msg).Error()))
		return userCategoryResponse, errors.New(res.Msg)
	}

	bdata, _ := json.Marshal(res.Data)
	if err = json.Unmarshal(bdata, &userCategoryResponse); err != nil {
		sugarLogger.Error("Level: Error", zap.String("Failed to unmarshaling response data: ", err.Error()))
		return userCategoryResponse, err
	}

	return userCategoryResponse, nil
}

// FindByMid ..
func FindByMid(mid string) (FindByMIDResponse, error) {
	sugarLogger := ottologger.GetLogger()
	var res Response
	var findByMIDRes FindByMIDResponse

	urlSvr := roseEnv.Host + roseEnv.EndpointFindByMid

	jsonData := map[string]string{"mid": mid}
	data, err := https.HTTPPost(urlSvr, jsonData)
	if err != nil {
		sugarLogger.Error("Level: Error", zap.String("Error: ", err.Error()))
		return findByMIDRes, err
	}

	if err = json.Unmarshal(data, &res); err != nil {
		sugarLogger.Error("Level: Error", zap.String("Failed to unmarshaling response: ", err.Error()))
		return findByMIDRes, err
	}

	if res.Rc == "04" {
		sugarLogger.Error("Level: Error", zap.String("Failed to get data: ", errors.New(res.Msg).Error()))
		return findByMIDRes, errors.New(res.Msg)
	}

	bdata, _ := json.Marshal(res.Data)
	if err = json.Unmarshal(bdata, &findByMIDRes); err != nil {
		sugarLogger.Error("Level: Error", zap.String("Failed to unmarshaling response data: ", err.Error()))
		return findByMIDRes, err
	}

	return findByMIDRes, nil
}

// ReplaceMidMpan ..
func ReplaceMidMpan(jsonData ReplaceMidMpanRequest) error {
	sugarLogger := ottologger.GetLogger()
	var res Response

	urlSvr := roseEnv.Host + roseEnv.EndpointReplaceMIDMpan

	//jsonData := map[string]string{ "mid": mid}
	data, err := https.HTTPPost(urlSvr, jsonData)
	if err != nil {
		sugarLogger.Error("Level: Error", zap.String("Error: ", err.Error()))
		return err
	}

	if err = json.Unmarshal(data, &res); err != nil {
		sugarLogger.Error("Level: Error", zap.String("Failed to unmarshaling response: ", err.Error()))
		return err
	}

	if res.Rc == "04" {
		sugarLogger.Error("Level: Error", zap.String("Failed to get data: ", errors.New(res.Msg).Error()))
		return errors.New(res.Msg)
	}

	return nil
}

// FindByPhoneNumber ..
func FindByPhoneNumber(phoneNumber string) (FindByPhoneNumberResponse, error) {
	sugarLogger := ottologger.GetLogger()
	var res Response
	var findByPhoneNumberRes FindByPhoneNumberResponse

	urlSvr := roseEnv.HostRoseOPService + roseEnv.EndpointMerchantProfil

	jsonData := map[string]string{"noHp": phoneNumber}
	data, err := https.HTTPPost(urlSvr, jsonData)
	if err != nil {
		sugarLogger.Error("Level: Error", zap.String("Error: ", err.Error()))
		return findByPhoneNumberRes, err
	}

	if err = json.Unmarshal(data, &res); err != nil {
		sugarLogger.Error("Level: Error", zap.String("Failed to unmarshaling response: ", err.Error()))
		return findByPhoneNumberRes, err
	}

	if res.Rc == "04" {
		sugarLogger.Error("Level: Error", zap.String("Failed to get data: ", errors.New(res.Msg).Error()))
		return findByPhoneNumberRes, errors.New(res.Msg)
	}

	bdata, _ := json.Marshal(res.Data)
	if err = json.Unmarshal(bdata, &findByPhoneNumberRes); err != nil {
		sugarLogger.Error("Level: Error", zap.String("Failed to unmarshaling response data: ", err.Error()))
		return findByPhoneNumberRes, err
	}

	return findByPhoneNumberRes, nil
}

// CheckIdCard ..
func CheckIdCard(idCard string) (Response, error) {
	sugarLogger := ottologger.GetLogger()
	var res Response

	urlSvr := roseEnv.Host + roseEnv.EndpointCheckIdCard

	jsonData := map[string]string{"noKtp": idCard}
	data, err := https.HTTPPost(urlSvr, jsonData)
	if err != nil {
		sugarLogger.Error("Level: Error", zap.String("Error: ", err.Error()))
		return res, err
	}

	if err = json.Unmarshal(data, &res); err != nil {
		sugarLogger.Error("Level: Error", zap.String("Failed to unmarshaling response: ", err.Error()))
		return res, err
	}

	//if res.Rc == "04" {
	//	sugarLogger.Error("Level: Error", zap.String("Failed to get data: ", errors.New(res.Msg).Error()))
	//	return res, errors.New(res.Msg)
	//}

	return res, nil
}

//UpdateDataMerchant
func UpdateDataMerchant(req UpdateDataMerchantReq) (Response, error) {
	sugarLogger := ottologger.GetLogger()
	var res Response

	urlSvr := roseEnv.Host + roseEnv.EndpointUpdateMerchant

	jsonData := map[string]string{
		"MerchantGroupId":   req.MerchantGroupID,
		"PartnerCustomerId": req.PartnerCustomerID,
		"UserCategoryCode":  req.UserCategoryCode,
		"address":           req.Address,
		"city":              req.City,
		"district":          req.District,
		"expireDate":        req.ExpireDate,
		"loanBankAccount":   req.LoanBankAccount,
		"mid":               req.MID,
		"partnerCode":       req.PartnerCode,
		"postalCode":        req.PostalCode,
		"province":          req.Province,
		"rt":                req.Rt,
		"rw":                req.Rw,
		"village":           req.Village,
	}

	data, err := https.HTTPPost(urlSvr, jsonData)
	if err != nil {
		sugarLogger.Error("Level: Error", zap.String("Error: ", err.Error()))
		return res, err
	}

	if err = json.Unmarshal(data, &res); err != nil {
		sugarLogger.Error("Level: Error", zap.String("Failed to unmarshaling response: ", err.Error()))
		return res, err
	}

	if res.Rc != "00" {
		sugarLogger.Error("Level: Error", zap.String("Failed to post data: ", errors.New(res.Msg).Error()))
		return res, errors.New(res.Msg)
	}

	return res, nil
}

//UpdateDataMerchantV2
func UpdateDataMerchantV2(req UpdateDataMerchantV2Req) (Response, error) {
	sugarLogger := ottologger.GetLogger()
	var res Response

	urlSvr := roseEnv.HostV2 + roseEnv.EndpointUpdateMerchant
	data, err := https.HTTPPost(urlSvr, req)
	if err != nil {
		sugarLogger.Error("Level: Error", zap.String("Error: ", err.Error()))
		return res, err
	}

	if err = json.Unmarshal(data, &res); err != nil {
		sugarLogger.Error("Level: Error", zap.String("Failed to unmarshaling response: ", err.Error()))
		return res, err
	}

	if res.Rc != "00" {
		sugarLogger.Error("Level: Error", zap.String("Failed to post data: ", errors.New(res.Msg).Error()))
		return res, errors.New(res.Msg)
	}

	return res, nil
}

// SalesByReport ..
func SalesByReport(phone string, dateFrom string, dateTo string) (SalesByReportResponse, error) {
	sugarLogger := ottologger.GetLogger()
	var res Response
	var salesByReportRes SalesByReportResponse
	urlSvr := roseEnv.Host + roseEnv.EndpointSalesByReport
	
	jsonData := map[string]string{"salesPhone": phone , "dateFrom": dateFrom, "dateTo": dateTo}
	fmt.Println(jsonData)
	data, err := https.HTTPPost(urlSvr, jsonData)
	if err != nil {
		sugarLogger.Error("Level: Error", zap.String("Error: ", err.Error()))
		return salesByReportRes, err
	}

	if err = json.Unmarshal(data, &res); err != nil {
		sugarLogger.Error("Level: Error", zap.String("Failed to unmarshaling response: ", err.Error()))
		return salesByReportRes, err
	}

	bdata, _ := json.Marshal(res.Data)
	if err = json.Unmarshal(bdata, &salesByReportRes); err != nil {
		sugarLogger.Error("Level: Error", zap.String("Failed to unmarshaling response data: ", err.Error()))
		return salesByReportRes, err
	}

	return salesByReportRes, nil
}

// SalesHistoryReport ..
func SalesHistoryReport(phone string, dateFrom string, dateTo string) (ResHistoryPencapaianSalesDto, error) {
	sugarLogger := ottologger.GetLogger()
	var res Response
	var salesHistoryReportRes ResHistoryPencapaianSalesDto
	urlSvr := roseEnv.Host + roseEnv.EndpointSalesHistoryReport
	
	jsonData := map[string]string{"salesPhone": phone , "dateFrom": dateFrom, "dateTo": dateTo}
	fmt.Println(jsonData)
	data, err := https.HTTPPost(urlSvr, jsonData)
	if err != nil {
		sugarLogger.Error("Level: Error", zap.String("Error: ", err.Error()))
		return salesHistoryReportRes, err
	}

	if err = json.Unmarshal(data, &res); err != nil {
		sugarLogger.Error("Level: Error", zap.String("Failed to unmarshaling response: ", err.Error()))
		return salesHistoryReportRes, err
	}

	bdata, _ := json.Marshal(res.Data)
	if err = json.Unmarshal(bdata, &salesHistoryReportRes); err != nil {
		sugarLogger.Error("Level: Error", zap.String("Failed to unmarshaling response data: ", err.Error()))
		return salesHistoryReportRes, err
	}

	return salesHistoryReportRes, nil
}

// PencapaianSales ..
func PencapaianSales(salesPhone string, srId int, villageId []string) (PencapaianSalesResponse, error) {
	sugarLogger := ottologger.GetLogger()
	var res Response
	var pencapaianSalesRes PencapaianSalesResponse

	urlSvr := roseEnv.Host + roseEnv.EndpointPencapaianSales

	jsonData := PencapaianSalesRequest{
		SalesPhone: salesPhone,
		SrId:       strconv.Itoa(srId),
		VillageId:  villageId,
	}

	data, err := https.HTTPPost(urlSvr, jsonData)
	if err != nil {
		sugarLogger.Error("Level: Error", zap.String("Error: ", err.Error()))
		return pencapaianSalesRes, err
	}

	if err = json.Unmarshal(data, &res); err != nil {
		sugarLogger.Error("Level: Error", zap.String("Failed to unmarshaling response: ", err.Error()))
		return pencapaianSalesRes, err
	}

	if res.Rc == "04" {
		sugarLogger.Error("Level: Error", zap.String("Failed to get data: ", errors.New(res.Msg).Error()))
		return pencapaianSalesRes, errors.New(res.Msg)
	}

	bdata, _ := json.Marshal(res.Data)
	if err = json.Unmarshal(bdata, &pencapaianSalesRes); err != nil {
		sugarLogger.Error("Level: Error", zap.String("Failed to unmarshaling response data: ", err.Error()))
		return pencapaianSalesRes, err
	}

	return pencapaianSalesRes, nil
}

// PencapaianSalesAll ..
func PencapaianSalesAll(salesPhone string) (PencapaianSalesResponse, error) {
	sugarLogger := ottologger.GetLogger()
	var res Response
	var pencapaianSalesRes PencapaianSalesResponse

	urlSvr := roseEnv.Host + roseEnv.EndpointPencapaianSales

	jsonData := PencapaianSalesRequest{
		SalesPhone: salesPhone,
	}

	data, err := https.HTTPPost(urlSvr, jsonData)
	if err != nil {
		sugarLogger.Error("Level: Error", zap.String("Error: ", err.Error()))
		return pencapaianSalesRes, err
	}

	if err = json.Unmarshal(data, &res); err != nil {
		sugarLogger.Error("Level: Error", zap.String("Failed to unmarshaling response: ", err.Error()))
		return pencapaianSalesRes, err
	}

	if res.Rc == "04" {
		sugarLogger.Error("Level: Error", zap.String("Failed to get data: ", errors.New(res.Msg).Error()))
		return pencapaianSalesRes, errors.New(res.Msg)
	}

	bdata, _ := json.Marshal(res.Data)
	if err = json.Unmarshal(bdata, &pencapaianSalesRes); err != nil {
		sugarLogger.Error("Level: Error", zap.String("Failed to unmarshaling response data: ", err.Error()))
		return pencapaianSalesRes, err
	}

	return pencapaianSalesRes, nil
}
