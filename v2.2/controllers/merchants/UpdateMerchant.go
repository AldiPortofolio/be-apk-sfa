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

// UpdateMerchant ..
// Update Merchant godoc
// @Summary Update Merchant
// @Description Update Merchant
// @ID Update Merchant
// @Tags OTTO SFA
// @Router /v2.2/merchants/update [post]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Param Body body models.UpdateMerchantReq true "Body"
// @Success 200 {object} models.Response{} "Update Merchant Response EXAMPLE"
func UpdateMerchant(ctx *gin.Context) {
	fmt.Println(">>> UpdateMerchant - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	//langID := ctx.GetHeader("Language-Id")

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := models.UpdateMerchantReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	reqBytes, _ := json.Marshal(req)
	log.Info("UpdateMerchant Controller",
		log.AddField("RequestBody:", string(reqBytes)))

	services.InitiateService(log).UpdateMerchant(req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("UpdateMerchant Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
