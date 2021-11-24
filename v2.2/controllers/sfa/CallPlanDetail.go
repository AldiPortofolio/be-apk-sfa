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

// CallPlanDetail ..
// CallPlan Action Check Merchant Phone godoc
// @Summary CallPlan Action Check Merchant Phone
// @Description CallPlan Action Check Merchant Phone
// @ID CallPlan Action Check Merchant Phone
// @Tags OTTO SFA
// @Router /v2.2/sfa/call_plan/action/check_merchant_phone [post]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Param Body body models.CallPlanActionCheckMerchantPhoneReq true "Body"
// @Success 200 {object} models.Response{data=[]dbmodels.CallPlanActions} "CallPlan Action Check Merchant Phone Response EXAMPLE"
func CallPlanDetail(ctx *gin.Context) {
	fmt.Println(">>> CallPlanDetail - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	//langID := ctx.GetHeader("Language-Id")

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := models.CallPlanDetailReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	header := models.RequestHeader{
		Authorization: ctx.GetHeader("Authorization"),
	}

	reqBytes, _ := json.Marshal(req)
	log.Info("CallPlanDetail Controller",
		log.AddField("RequestHeader:", header),
		log.AddField("RequestBody:", string(reqBytes)))

	services.InitiateService(log).CallPlanDetail(header.Authorization, req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("CallPlanDetail Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
