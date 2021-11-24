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

// ScanAndUpdateQRIS ..
// Scan And Update QRIS godoc
// @Summary Scan And Update QRIS
// @Description Scan And Update QRIS
// @ID Scan And
// @Tags OTTO SFA
// @Router /v2.2/qris/scan [post]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Param Body body models.ScanAndUpdateQRISReq true "Body"
// @Success 200 {object} models.Response{data=models.CheckQRISRes} "Scan And Update QRIS Response EXAMPLE"
func ScanAndUpdateQRIS(ctx *gin.Context) {
	fmt.Println(">>> ScanAndUpdateQRIS - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	//langID := ctx.GetHeader("Language-Id")

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := models.ScanAndUpdateQRISReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	header := models.RequestHeader{
		Authorization: ctx.GetHeader("Authorization"),
	}

	reqBytes, _ := json.Marshal(req)
	log.Info("ScanAndUpdateQRIS Controller",
		log.AddField("RequestHeader:", header),
		log.AddField("RequestBody:", string(reqBytes)))

	services.InitiateService(log).ScanAndUpdateQRIS(header.Authorization, req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("ScanAndUpdateQRIS Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
