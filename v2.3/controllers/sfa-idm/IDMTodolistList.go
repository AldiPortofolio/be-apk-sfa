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

// IDMTodolistList ..
// Todolist List godoc
// @Summary Todolist List
// @Description Todolist List
// @ID Todolist List IDM
// @Tags SFA-IDM
// @Router /v2.3/sfa-idm/todolist/list [post]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Param Body body models.TodolistPostReq true "Body"
// @Success 200 {object} models.Response{data=models.TodolistListReq} "Todolist List Response EXAMPLE"
func IDMTodolistList(ctx *gin.Context) {
	fmt.Println(">>> IDMTodolistList - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	//langID := ctx.GetHeader("Language-Id")

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := models.TodolistListReq{}
	if err := ctx.ShouldBind(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	header := models.RequestHeader{
		Authorization: ctx.GetHeader("Authorization"),
	}

	log.Info("IDMTodolistList Controller",
		log.AddField("RequestHeader:", header),
		log.AddField("RequestBody:", req))

	services.InitiateService(log).IDMTodolistList(header.Authorization, req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("IDMTodolistList Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
