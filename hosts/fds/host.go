package fds

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
	ottologger "ottodigital.id/library/logger"
	"ottosfa-api-apk/models"
	"ottosfa-api-apk/models/fdsmodels"
	https "ottosfa-api-apk/utils/http"
)

// Env ..
type Env struct {
	Host                         string `envconfig:"FDS_HOST" default:"https://fintech-dev.pactindo.com:8443/merchant/rest/sfa"`
	EndpointCheckMerchant        string `envconfig:"FDS_ENDPOINT_CHECK_MERCHANT" default:"/checkMerchantData"`
	EndpointCheckMerchantByID    string `envconfig:"FDS_ENDPOINT_CHECK_MERCHANT_BY_ID" default:"/checkMerchantId"`
	EndpointReportHistoryDetail  string `envconfig:"FDS_ENDPOINT_REPORT_HISTORY_DETAIL" default:"/historyDetail"`
	EndpointReportHistorySummary string `envconfig:"FDS_ENDPOINT_REPORT_HISTORY_SUMMARY" default:"/historySummary"`
	EndpointReportBySales        string `envconfig:"FDS_ENDPOINT_REPORT_BY_SALES" default:"/sfaDashboardBySales"`
	EndpointChangeMerchantPhone  string `envconfig:"FDS_ENDPOINT_CHANGE_MERCHANT_PHONE" default:"/updatePhone"`
	EndpointCheckProfilSales     string `envconfig:"FDS_ENDPOINT_CHECK_PROFIL_SALES" default:"/historyDetail"`
	EndpointGetLongLatMerchant   string `envconfig:"FDS_ENDPOINT_GET_LONGITUDE_LATITUDE" default:"/getGeotag"`
	HealthCheckKey               string `envconfig:"FDS_HEALTH_CHECK_KEY" default:"OTTOSFA-API-APK_HEALTH_CHECK:FDS1"`
}

var (
	fdsEnv Env
)

func init() {
	err := envconfig.Process("FDS", &fdsEnv)
	if err != nil {
		fmt.Println("Failed to get FDS env:", err)
	}
}

// Send ..
func Send(msgReq interface{}, typeTrans string, spanID string) ([]byte, error) {
	sugarLogger := ottologger.GetLogger()

	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	urlSvr := ""
	jsonData := map[string]string{}
	x, _ := json.Marshal(msgReq)
	switch typeTrans {
	case "CHECKMERCHANT":
		fmt.Println("CheckMerchantByPhonenumber")
		urlSvr = fdsEnv.Host + fdsEnv.EndpointCheckMerchant
		req := models.CheckMerchantOttopayReq{}
		json.Unmarshal(x, &req)
		jsonData = map[string]string{"phone": req.Phone}
		break
	case "CHECKMERCHANTBYID":
		fmt.Println("CheckMerchantByID")
		urlSvr = fdsEnv.Host + fdsEnv.EndpointCheckMerchantByID
		req := models.CheckMerchantByIDReq{}
		json.Unmarshal(x, &req)
		jsonData = map[string]string{"merchantId": req.MerchantID}
		break
	case "REPORTHISTORYDETAIL":
		fmt.Println("ReportHistoryDetail")
		urlSvr = fdsEnv.Host + fdsEnv.EndpointReportHistoryDetail
		req := models.ReportHistoryDetailReq{}
		json.Unmarshal(x, &req)
		jsonData = map[string]string{"phoneNumber": req.Phone}
		break
	case "REPORTHISTORYSUMMARY":
		fmt.Println("ReportHistorySummary")
		urlSvr = fdsEnv.Host + fdsEnv.EndpointReportHistorySummary
		req := models.ReportHistorySummaryReq{}
		json.Unmarshal(x, &req)
		jsonData = map[string]string{"phoneNumber": req.Phone, "dateFrom": req.DateFrom, "dateTo": req.DateTo}
		break
	case "REPORTBYSALES":
		fmt.Println("ReportBySales")
		urlSvr = fdsEnv.Host + fdsEnv.EndpointReportBySales
		req := models.ReportBySalesReq{}
		json.Unmarshal(x, &req)
		jsonData = map[string]string{"salesPhone": req.Phone, "dateFrom": req.DateFrom, "dateTo": req.DateTo}
		break
	case "CHANGEMERCHANTPHONE":
		fmt.Println("ChangeMerchantPhone")
		urlSvr = fdsEnv.Host + fdsEnv.EndpointReportHistoryDetail
		req := models.ChangeMerchantPhoneReq{}
		json.Unmarshal(x, &req)
		jsonData = map[string]string{"merchantId": req.MerchantID, "newPhone": req.NewPhone}
		break
	case "CHECKPROFILSALES":
		fmt.Println("CheckProfilSales")
		urlSvr = fdsEnv.Host + fdsEnv.EndpointCheckProfilSales
		req := models.CheckProfilSalesFDSReq{}
		json.Unmarshal(x, &req)
		jsonData = map[string]string{"phoneNumber": req.Phone}
		break
	case "GETLONGLATMERCHANT":
		fmt.Println("GetLongLatMerchant")
		urlSvr = fdsEnv.Host + fdsEnv.EndpointGetLongLatMerchant
		req := fdsmodels.LongLatMerchantReq{}
		json.Unmarshal(x, &req)
		jsonData = map[string]string{"phone": req.Phone}
		break
	}

	dataReq, _ := json.Marshal(msgReq)

	sugarLogger.Info("Level: Info",
		zap.String("SpanID: ", spanID),
		zap.String("Request: ", string(dataReq)),
		zap.String("TypeTrans: ", typeTrans))

	data, err := https.HTTPPostXForm(urlSvr, jsonData, fdsEnv.HealthCheckKey)

	sugarLogger.Info("Level: Info",
		zap.String("SpanID: ", spanID),
		zap.String("Response: ", string(data)),
		zap.String("TypeTrans: ", typeTrans))

	return data, err
}
