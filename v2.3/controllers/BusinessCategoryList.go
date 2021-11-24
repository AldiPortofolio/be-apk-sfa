package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ottosfa-api-apk/constants"
	"ottosfa-api-apk/models"
	"ottosfa-api-apk/utils"
	"ottosfa-api-apk/v2.3/services"

	"github.com/gin-gonic/gin"
	ottologger "ottodigital.id/library/logger/v2"
)

// BusinessCategoryList ..
// BusinessCategory List godoc
// @Summary BusinessCategory List
// @Description BusinessCategory List
// @ID BusinessCategory List
// @Tags OTTO SFA
// @Router /v2.3/merchants/business_category/list [get]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Success 200 {object} models.Response{data=models.CallPlanListRes} "Profile Sales Response EXAMPLE"
func BusinessCategoryList(ctx *gin.Context) {
	fmt.Println(">>> BusinessCategory - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	//langID := ctx.GetHeader("Language-Id")

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	header := models.RequestHeader{
		Authorization: ctx.GetHeader("Authorization"),
	}

	log.Info("BussinessCategory Controller",
		log.AddField("RequestHeader:", header))

	services.InitiateService(log).BusinessCategoryList(header.Authorization, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("BussinessCategory Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
