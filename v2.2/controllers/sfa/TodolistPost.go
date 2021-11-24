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

// TodolistPost ..
// Todolist Post godoc
// @Summary Todolist Post
// @Description Todolist Post
// @ID Todolist Post v2.2
// @Tags OTTO SFA
// @Router /v2.2/sfa/todolist/post [post]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Param Body body models.TodolistPostReq true "Body"
// @Success 200 {object} models.Response{} "Todolist Post Response EXAMPLE"
func TodolistPost(ctx *gin.Context) {
	fmt.Println(">>> TodolistPost - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	//langID := ctx.GetHeader("Language-Id")

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := models.TodolistPostReq{}
	if err := ctx.ShouldBind(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	header := models.RequestHeader{
		Authorization: ctx.GetHeader("Authorization"),
	}

	reformatReq := req
	reformatReq.PhotoMerchant1 = utils.ReformatReq(req.PhotoMerchant1)
	reformatReq.PhotoMerchant2 = utils.ReformatReq(req.PhotoMerchant2)
	reformatReq.PhotoMerchant3 = utils.ReformatReq(req.PhotoMerchant3)
	log.Info("TodolistPost Controller",
		log.AddField("RequestHeader:", header),
		log.AddField("RequestBody:", reformatReq))

	services.InitiateService(log).TodolistPost(header.Authorization, req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("TodolistPost Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
