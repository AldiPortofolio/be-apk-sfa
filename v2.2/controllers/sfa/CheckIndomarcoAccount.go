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

// CheckIndomarcoAccount ..
// Check Indomarco Account godoc
// @Summary Check Indomarco Account
// @Description Check Indomarco Account
// @ID Check Indomarco Account
// @Tags OTTO SFA
// @Router /v2.2/sfa/check_indomarco_account [post]
// @Accept json
// @Produce json
// @Param Body body models.CheckIndomarcoAccountReq true "Body"
// @Param Authorization header string true "Bearer <token>"
// @Success 200 {object} models.Response{data=[]models.CheckAccountIndomarcoRes} "Check Indomarco Account Response EXAMPLE"
func CheckIndomarcoAccount(ctx *gin.Context) {
	fmt.Println(">>> CheckIndomarcoAccount - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	//langID := ctx.GetHeader("Language-Id")

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := models.CheckIndomarcoAccountReq{}
	if err := ctx.ShouldBind(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}


	header := models.RequestHeader{
		Authorization: ctx.GetHeader("Authorization"),
	}

	log.Info("CheckIndomarcoAccount Controller",
		log.AddField("RequestHeader:", header))

	services.InitiateService(log).CheckIndomarcoAccount(header.Authorization, req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("CheckIndomarcoAccount Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
