package indomarco

import (
	"fmt"
	"ottosfa-api-apk/models"
	https "ottosfa-api-apk/utils/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	ottologger "ottodigital.id/library/logger"
)

// Env ..
type Env struct {
	Host                     string `envconfig:"IDM_HOST" default:"http://ottopayqavm.southeastasia.cloudapp.azure.com:8885/indomarco/v0.1.0"`
	EndpointCheckAccount     string `envconfig:"IDM_ENDPOINT_CHECK_ACCOUNT" default:"/users/check_account"`
	EndpointUpdateMerchant   string `envconfig:"IDM_ENDPOINT_UPDATE_MERCHANT" default:"/users/update_merchant"`
	HealthCheckKey           string `envconfig:"IDM_HEALTH_CHECK_KEY" default:"OTTOSFA-API-APK_HEALTH_CHECK:INDOMARCO"`
}

var (
	// IDMEnv ..
	IDMEnv Env
)

func init() {
	err := envconfig.Process("INDOMARCO", &IDMEnv)
	if err != nil {
		fmt.Println("Failed to get INDOMARCO env:", err)
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
	case "CHECKACCOUNT":
		fmt.Println("CheckAccountIndomarco")
		urlSvr = IDMEnv.Host + IDMEnv.EndpointCheckAccount
		req := models.CheckMerchantOttopayReq{}
		json.Unmarshal(x, &req)
		jsonData = map[string]string{"customer_code": req.CustomerCode}
		break
	case "UPDATEMERCHANT":
		fmt.Println("UpdateMerchant")
		urlSvr = IDMEnv.Host + IDMEnv.EndpointUpdateMerchant
		req := models.UpdateMerchantIndomarcoReq{}
		json.Unmarshal(x, &req)
		jsonData = map[string]string{"otto_phone": req.Phone, "customer_code": req.CustomerCode, "merchant_id": req.MerchantID}
		break
	}

	dataReq, _ := json.Marshal(msgReq)

	sugarLogger.Info("Level: Info",
		zap.String("SpanID: ", spanID),
		zap.String("Response: ", string(dataReq)),
		zap.String("TypeTrans: ", typeTrans))

	data, err := https.HTTPPost(urlSvr, jsonData)

	sugarLogger.Info("Level: Info",
		zap.String("SpanID: ", spanID),
		zap.String("Response: ", string(data)),
		zap.String("TypeTrans: ", typeTrans))

	return data, err
}

func SendV2(msgReq interface{}, typeTrans string, spanID string) (UpdateMerchantResponse, error) {
	sugarLogger := ottologger.GetLogger()
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var res Response
	var updateMerchantResponse UpdateMerchantResponse

	urlSvr := ""
	jsonData := map[string]string{}
	x, _ := json.Marshal(msgReq)

	switch typeTrans {
	case "UPDATEMERCHANT":
		fmt.Println("UpdateMerchant")
		urlSvr = IDMEnv.Host + IDMEnv.EndpointUpdateMerchant
		req := models.UpdateMerchantIndomarcoReq{}
		json.Unmarshal(x, &req)
		jsonData = map[string]string{"otto_phone": req.Phone, "customer_code": req.CustomerCode, "merchant_id": req.MerchantID}
		break
	}

	dataReq, _ := json.Marshal(msgReq)

	sugarLogger.Info("Level: Info",
		zap.String("SpanID: ", spanID),
		zap.String("Response: ", string(dataReq)),
		zap.String("TypeTrans: ", typeTrans))
	data, err := https.HTTPPost(urlSvr, jsonData)

	sugarLogger.Info("Level: Info",
		zap.String("SpanID: ", spanID),
		zap.String("Response: ", string(data)),
		zap.String("TypeTrans: ", typeTrans))

	if err != nil {
		sugarLogger.Error("Level: Error", zap.String("Error: ", err.Error()))
		return updateMerchantResponse, err
	}

	if err = json.Unmarshal(data, &res); err != nil {
		sugarLogger.Error("Level: Error", zap.String("Failed to unmarshaling response: ", err.Error()))
		return updateMerchantResponse, err
	}

	fmt.Println("===========  res  =============", res)

	if !res.Meta["status"].(bool) {
		err := res.Data.Msg
		fmt.Println("===========  err  =============", err)
		sugarLogger.Error("Level: Error", zap.String("Failed to get data: ", errors.New(err).Error()))
		return updateMerchantResponse, errors.New(err)
	}

	fmt.Println("===========  res.Data.Customer  =============", res.Data.Customer)
	bdata, _ := json.Marshal(res.Data.Customer)
	if err = json.Unmarshal(bdata, &updateMerchantResponse); err != nil {
		sugarLogger.Error("Level: Error", zap.String("Failed to unmarshaling response data: ", err.Error()))
		return updateMerchantResponse, err
	}

	fmt.Println("===========  updateMerchantResponse  =============", updateMerchantResponse)

	return updateMerchantResponse, nil
}
