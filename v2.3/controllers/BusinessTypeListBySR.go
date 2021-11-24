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

// BusinessTypeListBySR ..
// Business Type List By SR godoc
// @Summary Business Type List  By SR
// @Description Business Type List  By SR
// @ID Business Type List  By SR
// @Tags OTTO SFA
// @Router /v2.3/merchants/business_type/list [post]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Param Body body models.BusinessTypeListBySRReq true "Body"
// @Success 200 {object} models.Response{data=models.BusinessTypeListRes} "Business Type List By SR Response EXAMPLE"
func BusinessTypeListBySR(ctx *gin.Context) {
	fmt.Println(">>> BusinessTypeListBySR - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := models.BusinessTypeListBySRReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	header := models.RequestHeader{
		Authorization: ctx.GetHeader("Authorization"),
	}

	log.Info("BusinessTypeListBySR Controller",
		log.AddField("RequestHeader:", header))

	services.InitiateService(log).BusinessTypeListBySR(header.Authorization, req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("BusinessTypeListBySR Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
