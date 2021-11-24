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

// TodolistCount ..
// Todolist Count godoc
// @Summary Todolist Count
// @Description Todolist Count
// @ID Todolist Count
// @Tags OTTO SFA
// @Router /v2.2/sfa/todolist/count [get]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Success 200 {object} models.Response{models.Count} "Todolist Count Response EXAMPLE"
func TodolistCount(ctx *gin.Context) {
	fmt.Println(">>> TodolistCount - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	//langID := ctx.GetHeader("Language-Id")

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	header := models.RequestHeader{
		Authorization: ctx.GetHeader("Authorization"),
	}

	log.Info("TodolistCount Controller",
		log.AddField("RequestHeader:", header))

	services.InitiateService(log).TodolistCount(header.Authorization, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("TodolistCount Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
