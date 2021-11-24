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

// TodolistDetail ..
// Todolist Detail godoc
// @Summary Todolist Detail
// @Description Todolist Detail
// @ID Todolist Detail v2.3
// @Tags OTTO SFA
// @Router /v2.3/sfa/todolist/detail [post]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Param Body body models.TodolistDetailReq true "Body"
// @Success 200 {object} models.Response{data=models.ChangePhotoPinKTPRes} "Todolist Detail Response EXAMPLE"
func TodolistDetail(ctx *gin.Context) {
	fmt.Println(">>> TodolistDetail - Controller <<<")

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

	log.Info("TodolistDetail Controller",
		log.AddField("RequestHeader:", header),
		log.AddField("RequestBody:", req))

	services.InitiateService(log).TodolistDetail(header.Authorization, req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("TodolistDetail Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
