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

// CallPlanDescriptionMerchant ..
// CallPlan Description Merchant godoc
// @Summary CallPlan Description Merchant
// @Description CallPlan Description Merchant
// @ID CallPlan Description Merchant
// @Tags OTTO SFA
// @Router /v2.2/call_plan/description_merchant [post]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Param Body body models.CallPlanDescriptionMerchantReq true "Body"
// @Success 200 {object} models.Response{data=dbmodels.CallPlanDescriptions} "CallPlan Description Merchant Response EXAMPLE"
func CallPlanDescriptionMerchant(ctx *gin.Context) {
	fmt.Println(">>> CallPlanDescriptionMerchant - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	//langID := ctx.GetHeader("Language-Id")

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := models.CallPlanDescriptionMerchantReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	header := models.RequestHeader{
		Authorization: ctx.GetHeader("Authorization"),
	}

	reqBytes, _ := json.Marshal(req)
	log.Info("CallPlanDescriptionMerchant Controller",
		log.AddField("RequestHeader:", header),
		log.AddField("RequestBody:", string(reqBytes)))

	services.InitiateService(log).CallPlanDescriptionMerchant(header.Authorization, req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("CallPlanDescriptionMerchant Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
