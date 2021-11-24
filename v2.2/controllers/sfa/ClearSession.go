package sfa

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

// ClearSession ..
// Clear Session godoc
// @Summary Clear Session
// @Description Clear Session
// @ID Clear Session
// @Tags OTTO SFA
// @Router /v2.2/sfa/clear_session [get]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Success 200 {object} models.Response{} "Clear Session Response EXAMPLE"
func ClearSession(ctx *gin.Context) {
	fmt.Println(">>> ClearSession - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	log.Info("ClearSession Controller")

	services.InitiateService(log).ClearSession(&res)

	resBytes, _ := json.Marshal(res)
	log.Info("ClearSession Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
