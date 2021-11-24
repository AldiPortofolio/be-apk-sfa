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
	"ottosfa-api-apk/v2.3/services"
)

// MerchantDetailTokoOttopay ..
// Merchant Detail Toko Ottopay godoc
// @Summary Merchant Detail Toko Ottopay
// @Description Merchant Detail Toko Ottopay
// @ID Merchant List Toko Ottopay v2.3
// @Tags OTTO SFA
// @Router /v2.3/toko-ottopay/merchant/detail [post]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Param Body body models.MerchantDetailQRISReq true "Body"
// @Success 200 {object} models.Response{data=models.MerchantDetailTokoOttopayV23Res} "Merchant Detail QRIS Response EXAMPLE"
func MerchantDetailTokoOttopay(ctx *gin.Context) {
	fmt.Println(">>> MerchantDetailTokoOttopay - Controller <<<")

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
	log.Info("MerchantDetailTokoOttopay Controller",
		log.AddField("RequestHeader:", header),
		log.AddField("RequestBody:", string(reqBytes)))

	services.InitiateService(log).MerchantDetailTokoOttopay(header.Authorization, req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("MerchantDetailTokoOttopay Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
