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

// Province godoc
// @Summary Province
// @Description Province
// @ID Province
// @Tags OTTO SFA
// @Router /v2.2/province/:country_id [get]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Success 200 {object} models.Response{data=models.Province} "Province Response EXAMPLE"
func Province(ctx *gin.Context) {
	fmt.Println(">>> Province - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	countryId := ctx.Params.ByName("country_id")

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	log.Info("Province Controller",
		log.AddField("RequestBody-CountryId:", countryId))

	services.InitiateService(log).Province(countryId, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("Province Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
