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

// TodolistTaskBySubCategory ..
// Todolist TaskBySubCategory godoc
// @Summary Todolist TaskBySubCategory
// @Description Todolist TaskBySubCategory
// @ID Todolist TaskBySubCategory
// @Tags OTTO SFA
// @Router /v2.2/sfa/todolist/taskbysubcategory [post]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Param Body body models.TodolistTaskBySubCategoryReq true "Body"
// @Success 200 {object} models.Response{[]models.TodolistTaskBySubCategoryRes} "Todolist TaskBySubCategory Response EXAMPLE"
func TodolistTaskBySubCategory(ctx *gin.Context) {
	fmt.Println(">>> TodolistTaskBySubCategory - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	//langID := ctx.GetHeader("Language-Id")

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := models.TodolistTaskBySubCategoryReq{}
	if err := ctx.ShouldBind(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	header := models.RequestHeader{
		Authorization: ctx.GetHeader("Authorization"),
	}

	log.Info("TodolistTaskBySubCategory Controller",
		log.AddField("RequestHeader:", header),
		log.AddField("RequestBody:", req))

	services.InitiateService(log).TodolistTaskBySubCategory(header.Authorization, req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("TodolistTaskBySubCategory Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
