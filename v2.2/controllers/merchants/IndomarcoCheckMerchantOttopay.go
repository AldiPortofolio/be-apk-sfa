package merchants

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

// IndomarcoCheckMerchantOttopay ..
// Indomarco Check Merchant Ottopay godoc
// @Summary Indomarco Check Merchant Ottopay
// @Description Indomarco Check Merchant Ottopay
// @ID Indomarco Check Merchant Ottopay
// @Tags OTTO SFA
// @Router /v2.2/merchants/indomarco/check_merchant_ottopay [post]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Param Body body models.CheckMerchantOttopayReq true "Body"
// @Success 200 {object} models.Response{} "Indomarco Check Merchant Ottopay Response EXAMPLE"
func IndomarcoCheckMerchantOttopay(ctx *gin.Context) {
	fmt.Println(">>> IndomarcoCheckMerchantOttopay - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	//langID := ctx.GetHeader("Language-Id")

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := models.CheckMerchantOttopayReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	reqBytes, _ := json.Marshal(req)
	log.Info("IndomarcoCheckMerchantOttopay Controller",
		log.AddField("RequestBody:", string(reqBytes)))

	services.InitiateService(log).IndomarcoCheckMerchantOttopay(req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("IndomarcoCheckMerchantOttopay Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
