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
	"ottosfa-api-apk/v2.3/services"
)

// AttendancePost ..
// Attendance Post godoc
// @Summary Attendance Post
// @Description Attendance Post
// @ID Attendance Post v2.3
// @Tags OTTO SFA
// @Router /v2.3/sfa/attendance/post [post]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Param Body body models.AttendancePostReq true "Body"
// @Success 200 {object} models.Response{data=models.ChangePhotoPinKTPRes} "Todolist Post Response EXAMPLE"
func AttendancePost(ctx *gin.Context) {
	fmt.Println(">>> AttendancePost - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	//langID := ctx.GetHeader("Language-Id")

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := models.AttendancePostReq{}
	if err := ctx.ShouldBind(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	header := models.RequestHeader{
		Authorization: ctx.GetHeader("Authorization"),
	}

	reformatReq := req
	reformatReq.PhotoSelfie = utils.ReformatReq(req.PhotoSelfie)
	log.Info("AttendancePost Controller",
		log.AddField("RequestHeader:", header),
		log.AddField("RequestBody:", reformatReq))

	services.InitiateService(log).AttendancePost(header.Authorization, req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("AttendancePost Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
