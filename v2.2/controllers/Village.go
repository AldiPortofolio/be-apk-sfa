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

// Village godoc
// @Summary Village
// @Description Village
// @ID Village
// @Tags OTTO SFA
// @Router /v2.2/village/:district_id [get]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Success 200 {object} models.Response{data=models.Village} "Village Response EXAMPLE"
func Village(ctx *gin.Context) {
	fmt.Println(">>> Village - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	districtId := ctx.Params.ByName("district_id")

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	log.Info("District Controller",
		log.AddField("RequestBody-DistrictId:", districtId))

	services.InitiateService(log).Village(districtId, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("District Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
