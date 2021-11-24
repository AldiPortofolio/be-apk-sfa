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

// ReportBySales ...
// ReportBySales godoc
// @Summary ReportBySales
// @Description ReportBySales
// @ID ReportBySales V2.2
// @Tags OTTO SFA
// @Router /v2.2/merchants/report_by_sales [post]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Param Body body models.ReportBySalesReq true "Body"
// @Success 200 {object} models.Response{data=models.Response}"
func ReportBySales(ctx *gin.Context) {
	fmt.Println(">>> ReportBySales - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)
	
	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := models.ReportBySalesReq{}
	if err := ctx.ShouldBind(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	header := models.RequestHeader{
		Authorization: ctx.GetHeader("Authorization"),
	}

	log.Info("ReportBySales Controller",
		log.AddField("RequestHeader:", header),
		log.AddField("RequestBody:", req))

	services.InitiateService(log).ReportBySales(header.Authorization, req, &res)
	logs.Info("[SERVICESEND][ReportSales]Req:%s, Res:%v", req.Phone, res.Meta)
	ctx.JSON(http.StatusOK, res)
}

