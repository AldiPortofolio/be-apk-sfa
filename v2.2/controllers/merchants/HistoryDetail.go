package merchants

import (
	"fmt"
	"net/http"
	"ottosfa-api-apk/constants"
	"ottosfa-api-apk/models"
	"ottosfa-api-apk/utils"
	"ottosfa-api-apk/v2.2/services"

	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	ottologger "ottodigital.id/library/logger/v2"
)

// HistoryDetail ...
// HistoryDetail godoc
// @Summary HistoryDetail
// @Description HistoryDetail
// @ID HistoryDetail V2.2
// @Tags OTTO SFA
// @Router /v2.2/merchants/history_detail [post]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Param Body body models.ReportBySalesReq true "Body"
// @Success 200 {object} models.Response{data=models.Response}"
func HistoryDetail(ctx *gin.Context) {
	fmt.Println(">>> HistoryDetail - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := models.ReportHistoryDetailReq{}
	if err := ctx.ShouldBind(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	header := models.RequestHeader{
		Authorization: ctx.GetHeader("Authorization"),
	}

	log.Info("HistoryDetail Controller",
		log.AddField("RequestHeader:", header),
		log.AddField("RequestBody:", req))

	services.InitiateService(log).HistoryDetail(header.Authorization, req, &res)
	logs.Info("[SERVICESEND][HistoryDetail]Req:%s, Res:%v", req.Phone, res.Meta)
	ctx.JSON(http.StatusOK, res)
}
