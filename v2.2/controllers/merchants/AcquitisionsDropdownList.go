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

// AcquitisionDropdownList ..
// Acquitisions Dropdown List QR godoc
// @Summary Acquitisions Dropdown List
// @Description Acquitisions Dropdown List
// @ID Acquitisions Dropdown List
// @Tags OTTO SFA
// @Router /v2.2/merchants/acquitision/dropdown/list [post]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Success 200 {object} models.Response{data=models.ReverseQRRes} "Acquitisions Dropdown List Response EXAMPLE"
func AcquitisionDropdownList(ctx *gin.Context) {
	fmt.Println(">>> AcquitisionsDropdownList - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	header := models.RequestHeader{
		Authorization: ctx.GetHeader("Authorization"),
	}

	log.Info("AcquitisionsDropdownList Controller",
		log.AddField("RequestHeader:", header))

	services.InitiateService(log).AcquitisionsDropdownList(header.Authorization, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("AcquitisionsDropdownList Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
