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

// CallPlanActionMerchantUnknown ..
// CallPlan Action Merchant Unknown godoc
// @Summary CallPlan Action Merchant Unknown
// @Description CallPlan Action Merchant Unknown
// @ID CallPlan Action Merchant Unknown v2.3
// @Tags OTTO SFA
// @Router /v2.3/sfa/call_plan/action/merchant_unknown [post]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Param Body body models.CallPlanActionMerchantUnknownReq true "Body"
// @Success 200 {object} models.Response "CallPlan Action Merchant Unknown Response EXAMPLE"
func CallPlanActionMerchantUnknown(ctx *gin.Context) {
	fmt.Println(">>> CallPlanActionMerchantUnknown - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	//langID := ctx.GetHeader("Language-Id")

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := models.CallPlanActionMerchantUnknownReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	header := models.RequestHeader{
		Authorization: ctx.GetHeader("Authorization"),
	}

	reformatReq := req
	reformatReq.PhotoLocation = utils.ReformatReq(req.PhotoLocation)
	log.Info("CallPlanActionMerchantUnknown Controller",
		log.AddField("RequestHeader:", header),
		log.AddField("RequestBody:", reformatReq))

	services.InitiateService(log).CallPlanActionMerchantUnknown(header.Authorization, req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("CallPlanActionMerchantUnknown Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}