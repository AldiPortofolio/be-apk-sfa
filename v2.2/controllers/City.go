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

// City godoc
// @Summary City
// @Description City
// @ID City v2.2
// @Tags OTTO SFA
// @Router /v2.2/city/:province_id [get]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Success 200 {object} models.Response{data=models.City} "City Response EXAMPLE"
func City(ctx *gin.Context) {
	fmt.Println(">>> City - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	provinceId := ctx.Params.ByName("province_id")

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	log.Info("City Controller",
		log.AddField("RequestBody-ProvinceId:", provinceId))

	services.InitiateService(log).City(provinceId, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("City Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
