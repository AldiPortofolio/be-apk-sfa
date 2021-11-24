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

// IDMTodolistDetail ..
// Todolist Detail godoc
// @Summary Todolist Detail
// @Description Todolist Detail
// @ID Todolist Detail IDM
// @Tags OTTO SFA-IDM
// @Router /v2.3/sfa-idm/todolist/detail [post]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Param Body body models.TodolistDetailReq true "Body"
// @Success 200 {object} models.Response{data=models.ChangePhotoPinKTPRes} "Todolist Detail Response EXAMPLE"
func IDMTodolistDetail(ctx *gin.Context) {
	fmt.Println(">>> IDMTodolistDetail - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	//langID := ctx.GetHeader("Language-Id")

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := models.TodolistDetailReq{}
	if err := ctx.ShouldBind(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	header := models.RequestHeader{
		Authorization: ctx.GetHeader("Authorization"),
	}

	log.Info("IDMTodolistDetail Controller",
		log.AddField("RequestHeader:", header),
		log.AddField("RequestBody:", req))

	services.InitiateService(log).IDMTodolistDetail(header.Authorization, req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("IDMTodolistDetail Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
