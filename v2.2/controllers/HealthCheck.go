package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	ottologger "ottodigital.id/library/logger/v2"
	"ottosfa-api-apk/constants"
	"ottosfa-api-apk/models"
	"ottosfa-api-apk/utils"
	"ottosfa-api-apk/v2.2/services"
)

// HealthCheck ..
// Health Check godoc
// @Summary Health Check
// @Description Health Check
// @ID HealthCheck v2.2
// @Tags OTTO SFA
// @Router /v2.2/ottosfa/health_check [get]
// @Accept json
// @Produce json
// @Success 200 {object} models.Response{} "Health Check Response EXAMPLE"
func HealthCheck(ctx *gin.Context) {
	fmt.Println(">>> HealthCheck - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}
	log.Info("HealthCheck Controller")

	services.InitiateService(log).HealthCheck(&res)

	resBytes, _ := json.Marshal(res)
	log.Info("HealthCheck Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
