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

// CallPlanActionMerchantList ..
// CallPlan Action Merchant List godoc
// @Summary CallPlan Action Merchant List
// @Description CallPlan Action Merchant List
// @ID CallPlan Action Merchant List
// @Tags OTTO SFA
// @Router /v2.2/call_plan/action_merchant/list [get]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Success 200 {object} models.Response{data=[]dbmodels.ActionMerchants} "CallPlan Action Merchant List Response EXAMPLE"
func CallPlanActionMerchantList(ctx *gin.Context) {
	fmt.Println(">>> CallPlanActionMerchantList - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	//langID := ctx.GetHeader("Language-Id")

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	header := models.RequestHeader{
		Authorization: ctx.GetHeader("Authorization"),
	}

	log.Info("CallPlanActionMerchantList Controller",
		log.AddField("RequestHeader:", header))

	services.InitiateService(log).CallPlanActionMerchantList(header.Authorization, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("CallPlanActionMerchantList Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
