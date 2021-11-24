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

// BusinessTypeList ..
// Business Type List QR godoc
// @Summary Business Type List
// @Description Business Type List
// @ID Business Type List v2.2
// @Tags OTTO SFA
// @Router /v2.2/business_type/list [get]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Success 200 {object} models.Response{data=models.BusinessTypeListRes} "Business Type List Response EXAMPLE"
func BusinessTypeList(ctx *gin.Context) {
	fmt.Println(">>> BusinessTypeList - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	header := models.RequestHeader{
		Authorization: ctx.GetHeader("Authorization"),
	}

	log.Info("BusinessTypeList Controller",
		log.AddField("RequestHeader:", header))

	services.InitiateService(log).BusinessTypeList(header.Authorization, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("BusinessTypeList Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
