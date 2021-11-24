package stringbuilder

import (
	"encoding/json"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
	ottologger "ottodigital.id/library/logger"
	https "ottosfa-api-apk/utils/http"
)

// Env ..
type Env struct {
	Host              string `envconfig:"STRINGBUILDER_HOST" default:"http://13.228.25.85:8995"`
	EndpointReverseQR string `envconfig:"STRINGBUILDER_ENDPOINT_REVERSE_QR" default:"/merchant/reverseqr"`
	HealthCheckKey    string `envconfig:"STRINGBUILDER_HEALTH_CHECK_KEY" default:"OTTOSFA-API-APK_HEALTH_CHECK:STRINGBUILDER"`
	Name              string `envconfig:"STRINGBUILDER_NAME" default:"QRIS.STRINGBUILDER"`
}

var (
	stringbuilderEnv Env
)

func init() {
	err := envconfig.Process("STRINGBUILDER", &stringbuilderEnv)
	if err != nil {
		fmt.Println("Failed to get STRINGBUILDER env:", err)
	}
}

// ReverseQR ...
func ReverseQR(qrData string) (*TagResponse, error) {
	fmt.Println("<<< Get Reverse QR from String Builder >>>")
	sugarLogger := ottologger.GetLogger()

	var res TagResponse

	urlSvr := stringbuilderEnv.Host + stringbuilderEnv.EndpointReverseQR
	jsonData := map[string]string{"qrData": qrData}

	data, err := https.HTTPPostJson(urlSvr, jsonData, stringbuilderEnv.HealthCheckKey)
	if err != nil {
		sugarLogger.Error("Level: Error",
			zap.String("Error: ", err.Error()))
		return &res, err
	}

	err = json.Unmarshal(data, &res)
	if err != nil {
		sugarLogger.Error("Level: Error",
			zap.String("Error: ", "Failed to unmarshaling response from stringbuilder"))
		return &res, err
	}

	sugarLogger.Info("Level: Info",
		zap.String("Response: ", string(data)))

	return &res, nil
}
