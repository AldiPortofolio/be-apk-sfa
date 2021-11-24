package sales

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

// Profile ..
// Check Profile Sales godoc
// @Summary Profile Sales
// @Description Profile Sales
// @ID Profile Sales
// @Tags OTTO SFA
// @Router /v2.2/sales/profile [get]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Success 200 {object} models.Response{data=models.CheckProfilSalesRes} "Profile Sales Response EXAMPLE"
func Profile(ctx *gin.Context) {
	fmt.Println(">>> Profile - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	//langID := ctx.GetHeader("Language-Id")

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	header := models.RequestHeader{
		Authorization: ctx.GetHeader("Authorization"),
	}

	log.Info("Profile Controller",
		log.AddField("RequestHeader:", header))

	services.InitiateService(log).Profile(header.Authorization, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("Profile Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
