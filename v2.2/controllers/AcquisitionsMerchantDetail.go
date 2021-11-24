package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	ottologger "ottodigital.id/library/logger/v2"
	"ottosfa-api-apk/constants"
	"ottosfa-api-apk/v2.2/services"
	"ottosfa-api-apk/models"
	"ottosfa-api-apk/utils"
)

// AcquisitionsMerchantDetail ..
// Acquisitions Merchant Detail godoc
// @Summary Acquisitions Merchant Detail
// @Description Acquisitions Merchant Detail
// @ID AcquisitionsMerchantDetail
// @Tags OTTO SFA
// @Router /v2.2/acquitisions/merchant/detail [post]
// @Accept json
// @Produce json
// @Param Body body models.AcquisitionsMerchantDetailReq true "Body"
// @Success 200 {object} models.Response{data=models.AcquisitionsMerchantDetailRes} "Acquisitions Merchant Detail Response EXAMPLE"
func AcquisitionsMerchantDetail(ctx *gin.Context) {
	fmt.Println(">>> AcquisitionsMerchantDetail - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := models.AcquisitionsMerchantDetailReq{}
	if err := ctx.ShouldBind(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	reqBytes, _ := json.Marshal(req)
	log.Info("AcquisitionsMerchantDetail Controller",
		log.AddField("RequestBody", string(reqBytes)))

	services.InitiateService(log).AcquisitionsMerchantDetail(req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("AcquisitionsMerchantDetail Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)

}
