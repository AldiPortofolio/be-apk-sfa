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

// CheckQRIS ..
// Check QR godoc
// @Summary Check QR
// @Description Check QR
// @ID Check QR
// @Tags OTTO SFA
// @Router /v2.2/merchants/check_qris [post]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Param Body body models.CheckQRISReq true "Body"
// @Success 200 {object} models.Response{data=models.CheckQRISRes} "Check QR Response EXAMPLE"
func CheckQRIS(ctx *gin.Context) {
	fmt.Println(">>> CheckQRIS - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	//langID := ctx.GetHeader("Language-Id")

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
	log.Info("CheckQRIS Controller",
		log.AddField("RequestHeader:", header),
		log.AddField("RequestBody:", string(reqBytes)))

	services.InitiateService(log).CheckQRIS(header.Authorization, req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("CheckQRIS Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
