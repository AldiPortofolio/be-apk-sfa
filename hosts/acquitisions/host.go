package acquitisions

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
	Host                 string `envconfig:"ACQUITISIONS_HOST" default:"http://13.228.25.85:8871/v3"`
	EndpointAcquitisions string `envconfig:"ACQUITISIONS_ENDPOINT" default:"/merchants/new"`
	HealthCheckKey       string `envconfig:"ACQUITISIONS_HEALTH_CHECK_KEY" default:"OTTOSFA-API-APK_HEALTH_CHECK:ACQUITISIONS"`
	Name                 string `envconfig:"ACQUITISIONS_NAME" default:"ACQUITISIONS"`
}

var (
	acquitisionsEnv Env
)

func init() {
	err := envconfig.Process("ACQUITISIONS", &acquitisionsEnv)
	if err != nil {
		fmt.Println("Failed to get ACQUITISIONS env:", err)
	}
}

// Send ..
func Send(msgReq interface{}, spanID string) (res AcquitisionsRes, err error) {
	fmt.Println("<<< Acquitisions >>>")
	sugarLogger := ottologger.GetLogger()

	urlSvr := acquitisionsEnv.Host + acquitisionsEnv.EndpointAcquitisions

	data, err := https.HTTPUploadMinio(urlSvr, msgReq, acquitisionsEnv.HealthCheckKey)
	if err != nil {
		sugarLogger.Error("Level: Error", zap.String("Error: ", err.Error()))
		return res, err
	}

	err = json.Unmarshal(data, &res)
	if err != nil {
		sugarLogger.Error("Level: Error", zap.String("Error: ", "Failed to unmarshaling response"))
		return res, err
	}

	sugarLogger.Info("Level: Info",
		zap.String("SpanID: ", spanID),
		zap.String("Response: ", string(data)))

	return res, err
}
