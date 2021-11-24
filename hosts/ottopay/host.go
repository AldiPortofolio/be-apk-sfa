package ottopay

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
	ottologger "ottodigital.id/library/logger"
	"ottosfa-api-apk/models"
	https "ottosfa-api-apk/utils/http"
)

// Env ..
type Env struct {
	Host                   string `envconfig:"OTTOPAY_HOST" default:"http://ottopay2-api.clappingape.com/v1"`
	EndpointUpdateMerchant string `envconfig:"OTTOPAY_ENDPOINT_UPDATE_MERCHANT" default:"/indomarco/auth/approve-access"`
	HealthCheckKey         string `envconfig:"OTTOPAY_HEALTH_CHECK_KEY" default:"OTTOSFA-API-APK_HEALTH_CHECK:OTTOPAY"`
}

var (
	// OPEnv ..
	OPEnv Env
)

func init() {
	err := envconfig.Process("OTTOPAY", &OPEnv)
	if err != nil {
		fmt.Println("Failed to get OTTOPAY env:", err)
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
	case "UPDATEMERCHANT":
		fmt.Println("UpdateMerchant")
		urlSvr = OPEnv.Host + OPEnv.EndpointUpdateMerchant
		req := models.UpdateMerchantIndomarcoOttopayReq{}
		json.Unmarshal(x, &req)
		jsonData = map[string]string{"otto_phone": req.Phone, "accont_number": req.AccountNumber, "owner_name": req.OwnerName, "merchant_id": req.MerchantID, "customer_id": req.CustomerID, "name": req.Name}
		break
	}

	dataReq, _ := json.Marshal(msgReq)

	sugarLogger.Info("Level: Info",
		zap.String("SpanID: ", spanID),
		zap.String("Response: ", string(dataReq)),
		zap.String("TypeTrans: ", typeTrans))
	data, err := https.HTTPPostJson(urlSvr, jsonData, OPEnv.HealthCheckKey)

	sugarLogger.Info("Level: Info",
		zap.String("SpanID: ", spanID),
		zap.String("Response: ", string(data)),
		zap.String("TypeTrans: ", typeTrans))

	return data, err
}
