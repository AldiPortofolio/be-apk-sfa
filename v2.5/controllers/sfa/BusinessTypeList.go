package sfa

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ottosfa-api-apk/constants"
	"ottosfa-api-apk/models"
	"ottosfa-api-apk/utils"
	"ottosfa-api-apk/v2.5/services"

	"github.com/gin-gonic/gin"
	ottologger "ottodigital.id/library/logger/v2"
)

// BusinessTypeList ..
// BusinessType List godoc
// @Summary BusinessTypeList List
// @Description BusinessTypeList List
// @ID BusinessTypeList List
// @Tags OTTO SFA
// @Router /v2.5/sfa/business_type/list [post]
// @Accept json
// @Produce json
// @Param Body body models.GetAcquisitionByName true "Body"
// @Success 200 {object} models.Response{data=dbmodels.MerchantBusinessTypes} "BusinessTypeList List Response EXAMPLE"
func BusinessTypeList(ctx *gin.Context) {
	fmt.Println(">>> BusinessTypeList - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	//langID := ctx.GetHeader("Language-Id")

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := models.GetAcquisitionByName{}
	if err := ctx.ShouldBind(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	log.Info("BusinessTypeList Controller",
		log.AddField("RequestBody:", req))

	services.InitiateService(log).BusinessTypeListByName(req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("BusinessTypeList Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}

