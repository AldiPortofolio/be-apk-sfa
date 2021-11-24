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
	"ottosfa-api-apk/v2.4/services"
)

// TodolistMerchantNotFoundList ..
// Todolist Merchant Not Found List godoc
// @Summary Todolist Merchant Not Found List
// @Description Todolist Merchant Not Found List
// @ID Todolist Merchant Not Found List v2.4
// @Tags OTTO SFA
// @Router /v2.4/sfa/todolist/merchant_not_found/list [post]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Param Body body models.TodolistMerchantNotFoundListReq true "Body"
// @Success 200 {object} models.Response{data=[]models.TodolistMerchantNotFoundListRes} "Todolist Merchant Not Found List Response EXAMPLE"
func TodolistMerchantNotFoundList(ctx *gin.Context) {
	fmt.Println(">>> TodolistMerchantNotFoundList - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	//langID := ctx.GetHeader("Language-Id")

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := models.TodolistMerchantNotFoundListReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	header := models.RequestHeader{
		Authorization: ctx.GetHeader("Authorization"),
	}

	reqBytes, _ := json.Marshal(req)
	log.Info("TodolistMerchantNotFoundList Controller",
		log.AddField("RequestHeader:", header),
		log.AddField("RequestBody:", string(reqBytes)))

	services.InitiateService(log).TodolistMerchantNotFoundList(header.Authorization, req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("TodolistMerchantNotFoundList Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
