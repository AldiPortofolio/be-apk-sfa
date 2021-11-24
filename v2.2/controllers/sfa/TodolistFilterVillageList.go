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

// TodolistFilterVillageList ..
// Todolist FilterVillageList godoc
// @Summary Todolist FilterVillageList
// @Description Todolist FilterVillageList
// @ID Todolist FilterVillageList
// @Tags OTTO SFA
// @Router /v2.2/sfa/todolist/filter/village [post]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Param Body body models.TodolistFilterVillageListReq true "Body"
// @Success 200 {object} models.Response{[]dbmodels.Villages} "Todolist FilterVillageList Response EXAMPLE"
func TodolistFilterVillageList(ctx *gin.Context) {
	fmt.Println(">>> TodolistFilterVillageList - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	//langID := ctx.GetHeader("Language-Id")

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := models.TodolistFilterVillageListReq{}
	if err := ctx.ShouldBind(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	header := models.RequestHeader{
		Authorization: ctx.GetHeader("Authorization"),
	}

	log.Info("TodolistFilterVillageList Controller",
		log.AddField("RequestHeader:", header),
		log.AddField("RequestBody:", req))

	services.InitiateService(log).TodolistFilterVillageList(header.Authorization, req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("TodolistFilterVillageList Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
