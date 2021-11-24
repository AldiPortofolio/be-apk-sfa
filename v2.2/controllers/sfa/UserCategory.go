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

// UserCategory ..
// User Category godoc
// @Summary User Category
// @Description User Category
// @ID User Category v2.2
// @Tags OTTO SFA
// @Router /v2.2/sfa/user_category [get]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Success 200 {object} models.Response{data=[]models.UserCategoryRes} "User Category Response EXAMPLE"
func UserCategory(ctx *gin.Context) {
	fmt.Println(">>> UserCategory - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	header := models.RequestHeader{
		Authorization: ctx.GetHeader("Authorization"),
	}

	log.Info("UserCategory Controller",
		log.AddField("RequestHeader:", header))

	services.InitiateService(log).UserCategory(header.Authorization, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("UserCategory Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
