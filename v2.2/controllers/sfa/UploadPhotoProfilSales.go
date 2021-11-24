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

// UploadPhotoProfilSales ..
// Upload Photo Profil Sales godoc
// @Summary Upload Photo Profil Sales
// @Description Upload Photo Profil Sales
// @ID Upload Photo Profil Sales
// @Tags OTTO SFA
// @Router /v2.2/sfa/sales/upload/photo [post]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Param Body body models.UploadPhotoProfilSalesReq true "Body"
// @Success 200 {object} models.Response{} "Upload Photo Profil Sales Response EXAMPLE"
func UploadPhotoProfilSales(ctx *gin.Context) {
	fmt.Println(">>> UploadPhotoProfilSales - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	//langID := ctx.GetHeader("Language-Id")

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := models.UploadPhotoProfilSalesReq{}
	if err := ctx.ShouldBind(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	header := models.RequestHeader{
		Authorization: ctx.GetHeader("Authorization"),
	}

	reformatReq := req
	reformatReq.PhotoProfil = utils.ReformatReq(req.PhotoProfil)
	log.Info("UploadPhotoProfilSales Controller",
		log.AddField("RequestHeader:", header),
		log.AddField("RequestBody:", reformatReq))

	services.InitiateService(log).UploadPhotoProfilSales(header.Authorization, req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("UploadPhotoProfilSales Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
