package services

import (
	"ottosfa-api-apk/hosts/rose"
	"ottosfa-api-apk/models"
	"ottosfa-api-apk/utils"
	redis "ottosfa-api-apk/redis"
	"fmt"
	"ottosfa-api-apk/constants"
)

// ReportBySales ..
func (svc *Service) ReportBySales(bearer string, req models.ReportBySalesReq,res *models.Response) {
	fmt.Println(">>> ReportBySales - Service <<<")

	token := utils.DecodeBearer(bearer)
	_, validateToken := redis.GetRedisKey(utils.RedisKeyAuth + token)
	if validateToken != nil {
		res.Meta = utils.GetMetaResponse("invalid-token")
		return
	}

	fmt.Println("request: ", req)

	data, err := rose.SalesByReport(req.Phone, req.DateFrom, req.DateTo)
	if err != nil {
		res.Meta = utils.GetMetaResponse("default")
		return
	}

	res.Data = models.ReportBySalesRes{
		OttoPay:    data.OttoPay,
		SfaOnly:    data.SfaOnly,
		SalesPhone: data.SalesPhone,
	}
	res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)

	return
}
