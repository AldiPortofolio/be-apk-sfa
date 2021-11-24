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

// CallPlanMerchantList ..
// CallPlan Merchant List godoc
// @Summary CallPlan Merchant List
// @Description CallPlan Merchant List
// @ID CallPlan Merchant List v2.2
// @Tags OTTO SFA
// @Router /v2.2/sfa/call_plan/merchant/list [post]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Param Body body models.CallPlanMerchantListReq true "Body"
// @Success 200 {object} models.Response{data=models.CallPlanMerchantListRes} "CallPlan Merchant List Response EXAMPLE"
func CallPlanMerchantList(ctx *gin.Context) {
	fmt.Println(">>> CallPlanMerchantList - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	//langID := ctx.GetHeader("Language-Id")

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := models.CallPlanMerchantListReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	header := models.RequestHeader{
		Authorization: ctx.GetHeader("Authorization"),
	}

	reqBytes, _ := json.Marshal(req)
	log.Info("CallPlanMerchantList Controller",
		log.AddField("RequestHeader:", header),
		log.AddField("RequestBody:", string(reqBytes)))

	services.InitiateService(log).CallPlanMerchantList(header.Authorization, req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("CallPlanMerchantList Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
