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

// MerchantNotFoundList ..
// Merchant Not Found List godoc
// @Summary Merchant Not Found List
// @Description Merchant Not Found List
// @ID Merchant Not Found List
// @Tags OTTO SFA
// @Router /v2.2/sfa/merchant_not_found/list [post]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Success 200 {object} models.Response{data=[]models.TodolistMerchantNotFoundListRes} "Merchant Not Found List Response EXAMPLE"
func MerchantNotFoundList(ctx *gin.Context) {
	fmt.Println(">>> MerchantNotFoundList - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	//langID := ctx.GetHeader("Language-Id")

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	header := models.RequestHeader{
		Authorization: ctx.GetHeader("Authorization"),
	}

	log.Info("MerchantNotFoundList Controller",
		log.AddField("RequestHeader:", header))

	services.InitiateService(log).MerchantNotFoundList(header.Authorization, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("MerchantNotFoundList Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
