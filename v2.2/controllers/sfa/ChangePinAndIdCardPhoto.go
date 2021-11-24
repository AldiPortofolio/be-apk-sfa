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

// ChangePinAndIdCardPhoto ..
// Change Pin And Id Card Photo Profil godoc
// @Summary Change Pin And Id Card Photo Profil
// @Description Change Pin And Id Card Photo Profil
// @ID Change Pin And Id Card Photo Profil
// @Tags OTTO SFA
// @Router /v2.2/sfa/auth/change_pin_and_ktp [post]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Param Body body models.ChangePhotoPinKTPReq true "Body"
// @Success 200 {object} models.Response{data=models.ChangePhotoPinKTPRes} "Change Pin And Id Card Photo Profil Response EXAMPLE"
func ChangePinAndIdCardPhoto(ctx *gin.Context) {
	fmt.Println(">>> ChangePinAndIdCardPhoto - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	//langID := ctx.GetHeader("Language-Id")

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := models.ChangePhotoPinKTPReq{}
	if err := ctx.ShouldBind(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	header := models.RequestHeader{
		Authorization: ctx.GetHeader("Authorization"),
	}

	reformatReq := req
	reformatReq.Photo = utils.ReformatReq(req.Photo)
	reformatReq.IdCard = utils.ReformatReq(req.IdCard)
	log.Info("ChangePinAndIdCardPhoto Controller",
		log.AddField("RequestHeader:", header),
		log.AddField("RequestBody:", reformatReq))

	services.InitiateService(log).ChangePinAndIdCardPhoto(header.Authorization, req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("ChangePinAndIdCardPhoto Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
