package sfa

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

// AcquisitionList ..
// AcquisitionList List godoc
// @Summary AcquisitionList List
// @Description AcquisitionList List
// @ID AcquisitionList List v2.5
// @Tags OTTO SFA
// @Router /v2.5/sfa/acquisition/list [post]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Param Body body models.GetAcquisitionBySalesTypeID true "Body"
// @Success 200 {object} models.Response{data=dbmodels.Acquisitions} "AcquisitionList List Response EXAMPLE"
func AcquisitionList(ctx *gin.Context) {
	fmt.Println(">>> AcquisitionList - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	//langID := ctx.GetHeader("Language-Id")

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := models.GetAcquisitionBySalesTypeID{}
	if err := ctx.ShouldBind(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	header := models.RequestHeader{
		Authorization: ctx.GetHeader("Authorization"),
	}

	log.Info("AcquisitionList Controller",
		log.AddField("RequestHeader:", header),
		log.AddField("RequestBody:", req))

	services.InitiateService(log).AcquisitionListBySR(header.Authorization, req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("AcquisitionList Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
