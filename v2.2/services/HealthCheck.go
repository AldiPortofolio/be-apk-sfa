package services

import (
	"fmt"
	"ottodigital.id/library/healthcheck"
	"ottodigital.id/library/utils"
	"ottosfa-api-apk/models"
)

// HealthCheck ..
func (svc *Service) HealthCheck(res *models.Response) {
	fmt.Println(">>> HealthCheck - Service <<<")

	// === database ===
	databaseReq := healthcheck.HealthCheckDBReq{
		Host:     utils.GetEnv("DB_POSTGRES_ADDRESS", ""),
		Port:     utils.GetEnv("DB_POSTGRES_PORT", "5432"),
		User:     utils.GetEnv("DB_POSTGRES_USER", "ubuntu"),
		Password: utils.GetEnv("DB_POSTGRES_PASS", "Ubuntu!23"),
		DBName:   utils.GetEnv("DB_POSTGRES_NAME", "otto-sfa-admin-api_copy_production"),
	}
	database := healthcheck.GenerateHealthCheckPostgres(databaseReq)

	// === redis ===
	redisReq := healthcheck.HealthCheckRedisReq{
		HostCluster: []string{
			utils.GetEnv("REDIS_HOST_CLUSTER1", ""),
			utils.GetEnv("REDIS_HOST_CLUSTER2", ""),
			utils.GetEnv("REDIS_HOST_CLUSTER3", ""),
		},
	}
	clusterRedis := healthcheck.GenerateHealthCheckRedisCluster(redisReq)

	// === service ===
	stringBuilderServiceReq := healthcheck.CheckServiceReq{
		ServiceName: utils.GetEnv("STRINGBUILDER_NAME", "QRIS.STRINGBUILDER"),
		Host:        utils.GetEnv("STRINGBUILDER_HOST", ""),
		Endpoint:    utils.GetEnv("STRINGBUILDER_ENDPOINT_REVERSE_QR", "/merchant/reverseqr"),
		Method:      "",
	}
	stringBuilderService := healthcheck.GenerateHealthCheckService(stringBuilderServiceReq)

	acquitisionServiceReq := healthcheck.CheckServiceReq{
		ServiceName: "otto-worker-akuisisi",
		Host:        utils.GetEnv("ACQUITISIONS_HOST", ""),
		Endpoint:    utils.GetEnv("ACQUITISIONS_ENDPOINT", ""),
		Method:      "",
	}
	acquitisionService := healthcheck.GenerateHealthCheckService(acquitisionServiceReq)

	healthCheckData := make([]healthcheck.DataHealthCheck, 0)
	healthCheckData = append(healthCheckData,
		<-database,
		<-clusterRedis,
		<-stringBuilderService,
		<-acquitisionService,
	)

	resHealthCheck := healthcheck.GenerateResponseHealthCheck(healthCheckData...)

	res.Meta = models.MetaData{
		Status:  true,
		Code:    resHealthCheck.ResponseCode,
		Message: resHealthCheck.Message,
	}
	res.Data = resHealthCheck.Data
	return
}
