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

// District godoc
// @Summary District
// @Description District
// @ID District
// @Tags OTTO SFA
// @Router /v2.2/district/:city_id [get]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Success 200 {object} models.Response{data=models.District} "District Response EXAMPLE"
func District(ctx *gin.Context) {
	fmt.Println(">>> District - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	cityId := ctx.Params.ByName("city_id")

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	log.Info("District Controller",
		log.AddField("RequestBody-CityId:", cityId))

	services.InitiateService(log).District(cityId, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("District Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
