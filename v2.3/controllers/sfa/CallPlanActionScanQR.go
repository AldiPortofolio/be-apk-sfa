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

// CallPlanActionScanQR ..
// CallPlan Action Scan QR godoc
// @Summary CallPlan Action Scan QR
// @Description CallPlan Action Scan QR
// @ID CallPlan Action Scan QR
// @Tags OTTO SFA
// @Router /v2.3/sfa/call_plan/action/scan_qr [post]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Param Body body models.CallPlanActionScanQRReq true "Body"
// @Success 200 {object} models.Response{data=models.CheckQRISRes} "CallPlan Action Scan QR Response EXAMPLE"
func CallPlanActionScanQR(ctx *gin.Context) {
	fmt.Println(">>> CallPlanActionScanQR - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	//langID := ctx.GetHeader("Language-Id")

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := models.CallPlanActionScanQRReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	header := models.RequestHeader{
		Authorization: ctx.GetHeader("Authorization"),
	}

	reqBytes, _ := json.Marshal(req)
	log.Info("CallPlanActionDetail Controller",
		log.AddField("RequestHeader:", header),
		log.AddField("RequestBody:", string(reqBytes)))

	services.InitiateService(log).CallPlanActionScanQR(header.Authorization, req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("CallPlanActionDetail Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
