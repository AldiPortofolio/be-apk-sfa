package indomarco

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ottosfa-api-apk/constants"
	"ottosfa-api-apk/models"
	"ottosfa-api-apk/utils"
	"ottosfa-api-apk/v2.5/services"

	"github.com/gin-gonic/gin"
	ottologger "ottodigital.id/library/logger/v2"
)

// UpdateMerchantIndomarco ..
// UpdateMerchantIndomarco godoc
// @Summary UpdateMerchantIndomarco Update
// @Description UpdateMerchantIndomarco Update
// @ID UpdateMerchantIndomarco Update
// @Tags OTTO SFA-IDM
// @Router /v2.5/indomarco/merchant/update [post]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Param Body body models.UpdateMerchantIndomarcoReq true "Body"
// @Success 200 {object} models.Response{data=models.Response} "UpdateMerchantIndomarcoReq Response EXAMPLE"
func UpdateMerchantIndomarco(ctx *gin.Context) {
	fmt.Println(">>> UpdateMerchantIndomarco - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	//langID := ctx.GetHeader("Language-Id")

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := models.UpdateMerchantIndomarcoReq{}
	if err := ctx.ShouldBind(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	header := models.RequestHeader{
		Authorization: ctx.GetHeader("Authorization"),
	}

	log.Info("UpdateMerchantIndomarco Controller",
		log.AddField("RequestHeader:", header),
		log.AddField("RequestBody:", req))

	services.InitiateService(log).UpdateMerchantIndomarco(req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("UpdateMerchantIndomarco Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
