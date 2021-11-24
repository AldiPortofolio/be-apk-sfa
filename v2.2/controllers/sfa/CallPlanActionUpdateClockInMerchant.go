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

// CallPlanActionUpdateClockInMerchant ..
// CallPlan Action Update ClockIn Merchant godoc
// @Summary CallPlan Action Update ClockIn Merchant
// @Description CallPlan Action Update ClockIn Merchant
// @ID CallPlan Action Update ClockIn Merchant
// @Tags OTTO SFA
// @Router /v2.2/sfa/call_plan/action/update_clock_in [post]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Param Body body models.CallPlanActionUpdateClockInMerchantReq true "Body"
// @Success 200 {object} models.Response "CallPlan Action Update ClockIn Merchant Response EXAMPLE"
func CallPlanActionUpdateClockInMerchant(ctx *gin.Context) {
	fmt.Println(">>> CallPlanActionUpdateClockInMerchant - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	//langID := ctx.GetHeader("Language-Id")

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := models.CallPlanActionUpdateClockInMerchantReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	header := models.RequestHeader{
		Authorization: ctx.GetHeader("Authorization"),
	}

	reqBytes, _ := json.Marshal(req)
	log.Info("CallPlanActionUpdateClockInMerchant Controller",
		log.AddField("RequestHeader:", header),
		log.AddField("RequestBody:", string(reqBytes)))

	services.InitiateService(log).CallPlanActionUpdateClockInMerchant(header.Authorization, req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("CallPlanActionUpdateClockInMerchant Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
