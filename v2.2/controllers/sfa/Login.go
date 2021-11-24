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

// Login ..
// Login godoc
// @Summary Login
// @Description Login
// @ID Login v2.2
// @Tags OTTO SFA
// @Router /v2.2/sfa/auth/login [post]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Param Body body models.LoginReq true "Body"
// @Success 200 {object} models.Response{data=models.LoginRes} "Login Response EXAMPLE"
func Login(ctx *gin.Context) {
	fmt.Println(">>> Login - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	//langID := ctx.GetHeader("Language-Id")

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := models.LoginReq{}
	if err := ctx.ShouldBind(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	reqBytes, _ := json.Marshal(req)
	log.Info("Login Controller",
		log.AddField("RequestBody:", string(reqBytes)))

	services.InitiateService(log).Login(req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("Login Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
