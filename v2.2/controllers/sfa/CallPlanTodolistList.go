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

// CallPlanTodolistList ..
// CallPlan Todolist List godoc
// @Summary CallPlan Todolist List
// @Description CallPlan Todolist List
// @ID CallPlan Todolist List v2.2
// @Tags OTTO SFA
// @Router /v2.2/sfa/call_plan/todolist/list [post]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Param Body body models.CallPlanTodolistListReq true "Body"
// @Success 200 {object} models.Response{data=models.CallPlanTodolistListRes} "CallPlan Todolist List Response EXAMPLE"
func CallPlanTodolistList(ctx *gin.Context) {
	fmt.Println(">>> CallPlanTodolistList - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	//langID := ctx.GetHeader("Language-Id")

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := models.CallPlanTodolistListReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	header := models.RequestHeader{
		Authorization: ctx.GetHeader("Authorization"),
	}

	reqBytes, _ := json.Marshal(req)
	log.Info("CallPlanTodolistList Controller",
		log.AddField("RequestHeader:", header),
		log.AddField("RequestBody:", string(reqBytes)))

	services.InitiateService(log).CallPlanTodolistList(header.Authorization, req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("CallPlanTodolistList Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
