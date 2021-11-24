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

// CallPlanActionAddOrEdit ..
// CallPlan Action Save godoc
// @Summary CallPlan Action Save
// @Description CallPlan Action Save
// @ID CallPlan Action Save v2.3
// @Tags OTTO SFA
// @Router /v2.3/sfa/call_plan/action/save [post]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Param Body body models.CallPlanActionScanQRReq true "Body"
// @Success 200 {object} models.Response{data=[]dbmodels.CallPlanActions} "CallPlan Action Save Response EXAMPLE"
func CallPlanActionAddOrEdit(ctx *gin.Context) {
	fmt.Println(">>> CallPlanActionAddOrEdit - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	//langID := ctx.GetHeader("Language-Id")

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := models.CallPlanActionAddOrEditReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	header := models.RequestHeader{
		Authorization: ctx.GetHeader("Authorization"),
	}

	reqBytes, _ := json.Marshal(req)
	log.Info("CallPlanActionAddOrEdit Controller",
		log.AddField("RequestHeader:", header),
		log.AddField("RequestBody:", string(reqBytes)))

	services.InitiateService(log).CallPlanActionAddOrEdit(header.Authorization, req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("CallPlanActionAddOrEdit Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
