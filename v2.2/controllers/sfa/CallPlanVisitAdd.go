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

// CallPlanVisitAdd ..
// Call Plan Visit Add godoc
// @Summary Call Plan Visit Add
// @Description Call Plan Visit Add
// @ID Call Plan Visit Add
// @Tags OTTO SFA
// @Router /v2.2/sfa/call_plan/visit/add [post]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Param Body body models.CallPlanVisitAddReq true "Body"
// @Success 200 {object} models.Response "Call Plan Visit Add Response EXAMPLE"
func CallPlanVisitAdd(ctx *gin.Context) {
	fmt.Println(">>> CallPlanVisitAdd - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	//langID := ctx.GetHeader("Language-Id")

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := models.CallPlanVisitAddReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	header := models.RequestHeader{
		Authorization: ctx.GetHeader("Authorization"),
	}

	reqBytes, _ := json.Marshal(req)
	log.Info("CallPlanVisitAdd Controller",
		log.AddField("RequestHeader:", header),
		log.AddField("RequestBody:", string(reqBytes)))

	services.InitiateService(log).CallPlanVisitAdd(header.Authorization, req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("CallPlanVisitAdd Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
