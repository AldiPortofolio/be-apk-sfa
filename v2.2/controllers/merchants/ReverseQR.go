package merchants

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

// ReverseQR ..
// Reverse QR godoc
// @Summary Reverse QR
// @Description Reverse QR
// @ID Reverse QR
// @Tags OTTO SFA
// @Router /v2.2/merchants/reverse_qr [post]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Param Body body models.CheckQRISReq true "Body"
// @Success 200 {object} models.Response{data=models.ReverseQRRes} "Reverse QR Response EXAMPLE"
func ReverseQR(ctx *gin.Context) {
	fmt.Println(">>> ReverseQR - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := models.CheckQRISReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	header := models.RequestHeader{
		Authorization: ctx.GetHeader("Authorization"),
	}

	reqBytes, _ := json.Marshal(req)
	log.Info("ReverseQR Controller",
		log.AddField("RequestHeader:", header),
		log.AddField("RequestBody:", string(reqBytes)))

	services.InitiateService(log).ReverseQR(header.Authorization, req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("ReverseQR Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
