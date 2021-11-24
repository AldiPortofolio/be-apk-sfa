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

// LoginIndomarco ..
// Login Indomarco godoc
// @Summary Login Indomarco
// @Description Login Indomarco
// @ID Login Indomarco
// @Tags OTTO SFA
// @Router /v2.2/sfa/indomarco/login [post]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Param Body body models.IndomarcoLoginReq true "Body"
// @Success 200 {object} models.Response{data=models.IndomarcoLoginRes} "Indomarco Login Response EXAMPLE"
func LoginIndomarco(ctx *gin.Context) {
	fmt.Println(">>> LoginIndomarco - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	//langID := ctx.GetHeader("Language-Id")

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := models.IndomarcoLoginReq{}
	if err := ctx.ShouldBind(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	log.Info("LoginIndomarco Controller")

	services.InitiateService(log).LoginIndomarco(req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("LoginIndomarco Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
