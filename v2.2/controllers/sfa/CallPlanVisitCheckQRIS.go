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

// CallPlanVisitCheckQRIS ..
// CallPlan Visit Check QRIS godoc
// @Summary CallPlan Visit Check QRIS
// @Description CallPlan Visit Check QRIS
// @ID CallPlan Visit Check QRIS v2.2
// @Tags OTTO SFA
// @Router /v2.2/sfa/call_plan/visit/check_qris [post]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Param Body body models.CheckQRISReq true "Body"
// @Success 200 {object} models.Response{data=models.CallPlanVisitMerchantRes} "CallPlan Visit Check QRIS Response EXAMPLE"
func CallPlanVisitCheckQRIS(ctx *gin.Context) {
	fmt.Println(">>> CallPlanVisitCheckQRIS - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	//langID := ctx.GetHeader("Language-Id")

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := models.CheckQRISReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	header := models.RequestHeader{
		Authorization: ctx.GetHeader("Authorization"),
	}

	reqBytes, _ := json.Marshal(req)
	log.Info("CallPlanVisitCheckQRIS Controller",
		log.AddField("RequestHeader:", header),
		log.AddField("RequestBody:", string(reqBytes)))

	services.InitiateService(log).CallPlanVisitCheckQRIS(header.Authorization, req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("CallPlanVisitCheckQRIS Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
