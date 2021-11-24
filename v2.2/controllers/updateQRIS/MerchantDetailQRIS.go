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

// MerchantDetailQRIS ..
// Merchant Detail QRIS godoc
// @Summary Merchant Detail QRIS
// @Description Merchant Detail QRIS
// @ID Merchant List QRIS v2.2
// @Tags OTTO SFA
// @Router /v2.2/qris/merchant/detail [post]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Param Body body models.MerchantDetailQRISReq true "Body"
// @Success 200 {object} models.Response{data=models.MerchantDetailQRISRes} "Merchant Detail QRIS Response EXAMPLE"
func MerchantDetailQRIS(ctx *gin.Context) {
	fmt.Println(">>> MerchantDetailQRIS - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	//langID := ctx.GetHeader("Language-Id")

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := models.MerchantDetailQRISReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	header := models.RequestHeader{
		Authorization: ctx.GetHeader("Authorization"),
	}

	reqBytes, _ := json.Marshal(req)
	log.Info("MerchantDetailUpdateQRIS Controller",
		log.AddField("RequestHeader:", header),
		log.AddField("RequestBody:", string(reqBytes)))

	services.InitiateService(log).MerchantDetailQRIS(header.Authorization, req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("MerchantDetailUpdateQRIS Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
