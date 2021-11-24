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

// PushNotif ..
// Clear Session godoc
// @Summary Clear Session
// @Description Clear Session
// @ID Clear Session v2.2
// @Tags OTTO SFA
// @Router /v2.2/sfa/clear_session [get]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Success 200 {object} models.Response{} "Clear Session Response EXAMPLE"
func PushNotif(ctx *gin.Context) {
	fmt.Println(">>> PushNotif - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := models.PushNotifReq{}
	if err := ctx.ShouldBind(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	reqBytes, _ := json.Marshal(req)
	log.Info("PushNotif Controller",
		log.AddField("RequestBody:", string(reqBytes)))

	services.InitiateService(log).PushNotif(req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("PushNotif Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
