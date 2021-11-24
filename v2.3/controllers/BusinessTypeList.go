package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ottosfa-api-apk/constants"
	"ottosfa-api-apk/models"
	"ottosfa-api-apk/utils"
	"ottosfa-api-apk/v2.3/services"

	"github.com/gin-gonic/gin"
	ottologger "ottodigital.id/library/logger/v2"
)

// BusinessTypeList ..
// Business Type List godoc
// @Summary Business Type List
// @Description Business Type List
// @ID Business Type List v2.3
// @Tags OTTO SFA
// @Router /v2.3/merchants/business_type/list [get]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Success 200 {object} models.Response{data=models.BusinessTypeListRes} "Business Type List Response EXAMPLE"
func BusinessTypeList(ctx *gin.Context) {
	fmt.Println(">>> BusinessTypeList - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	header := models.RequestHeader{
		Authorization: ctx.GetHeader("Authorization"),
	}

	log.Info("BusinessTypeList Controller",
		log.AddField("RequestHeader:", header))

	services.InitiateService(log).BusinessTypeList(header.Authorization, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("BusinessTypeList Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
