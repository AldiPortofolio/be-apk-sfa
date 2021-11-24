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

// CallPlanTodolistMerchantNotFound ..
// CallPlan Todolist Merchant Not Found godoc
// @Summary CallPlan Todolist Merchant Not Found List
// @Description CallPlan Todolist Merchant Not Found
// @ID CallPlan Todolist Merchant Not Found
// @Tags OTTO SFA
// @Router /v2.2/sfa/call_plan/todolist/merchant_not_found [post]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Param Body body models.CallPlanTodolistMerchantNotFoundReq true "Body"
// @Success 200 {object} models.Response "CallPlan Todolist Merchant Not Found Response EXAMPLE"
func CallPlanTodolistMerchantNotFound(ctx *gin.Context) {
	fmt.Println(">>> CallPlanTodolistMerchantNotFound - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	//langID := ctx.GetHeader("Language-Id")

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := models.CallPlanTodolistMerchantNotFoundReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	header := models.RequestHeader{
		Authorization: ctx.GetHeader("Authorization"),
	}

	reformatReq := req
	reformatReq.MerchantImage = utils.ReformatReq(req.MerchantImage)
	log.Info("CallPlanTodolistMerchantNotFound Controller",
		log.AddField("RequestHeader:", header),
		log.AddField("RequestBody:", reformatReq))

	services.InitiateService(log).CallPlanTodolistMerchantNotFound(header.Authorization, req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("CallPlanTodolistMerchantNotFound Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
