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

// CallPlanVisitCheckMerchantPhone ..
// CallPlan Visit Check Merchant Phone godoc
// @Summary CallPlan Visit Check Merchant Phone
// @Description CallPlan Visit Check Merchant Phone
// @ID CallPlan Visit Check Merchant Phone
// @Tags OTTO SFA
// @Router /v2.3/sfa/call_plan/visit/check_merchant_phone [post]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Param Body body models.CallPlanVisitCheckMerchantPhoneReq true "Body"
// @Success 200 {object} models.Response{data=models.CallPlanVisitMerchantRes} "CallPlan Visit Check Merchant Phone Response EXAMPLE"
func CallPlanVisitCheckMerchantPhone(ctx *gin.Context) {
	fmt.Println(">>> CallPlanVisitCheckMerchantPhone - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	//langID := ctx.GetHeader("Language-Id")

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := models.CallPlanVisitCheckMerchantPhoneReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	header := models.RequestHeader{
		Authorization: ctx.GetHeader("Authorization"),
	}

	reqBytes, _ := json.Marshal(req)
	log.Info("CallPlanVisitCheckMerchantPhone Controller",
		log.AddField("RequestHeader:", header),
		log.AddField("RequestBody:", string(reqBytes)))

	services.InitiateService(log).CallPlanVisitCheckMerchantPhone(header.Authorization, req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("CallPlanVisitCheckMerchantPhone Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
