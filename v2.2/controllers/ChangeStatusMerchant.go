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

// ChangeStatusMerchant ..
// Change Status Merchant godoc
// @Summary Change Status Merchant
// @Description Change Status Merchant
// @ID Change Status Merchant v2.2
// @Tags OTTO SFA
// @Router /v2.2/ottosfa/merchants/change_status [post]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Success 200 {object} models.Response{data=models.District} "Change Status Merchant Response EXAMPLE"
func ChangeStatusMerchant(ctx *gin.Context) {
	fmt.Println(">>> ChangeStatusMerchant - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := models.ChangeStatusMerchantReq{}
	if err := ctx.ShouldBind(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	reqBytes, _ := json.Marshal(req)
	log.Info("ChangeStatusMerchant Controller",
		log.AddField("RequestBody", string(reqBytes)))

	services.InitiateService(log).ChangeStatusMerchant(req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("ChangeStatusMerchant Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
