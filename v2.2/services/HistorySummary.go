package services

import (
	"ottosfa-api-apk/hosts/rose"
	"ottosfa-api-apk/models"
	"ottosfa-api-apk/utils"
	redis "ottosfa-api-apk/redis"
	"fmt"
	"ottosfa-api-apk/constants"
)

// HistorySummary ..
func (svc *Service) HistorySummary(bearer string, req models.ReportBySalesReq, res *models.Response) {
	fmt.Println(">>> HistorySummary - Service <<<")

	token := utils.DecodeBearer(bearer)
	_, validateToken := redis.GetRedisKey(utils.RedisKeyAuth + token)
	if validateToken != nil {
		res.Meta = utils.GetMetaResponse("invalid-token")
		return
	}

	data, err := rose.SalesHistoryReport(req.Phone, req.DateFrom, req.DateTo)
	if err != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseDefault)
		return
	}

	acqData := []models.AcquisitionSummary{}
	for _, val := range data.AcquisitionData {
		acq := models.AcquisitionSummary{
			Created:   val.Created,
			Address:   val.Address,
			Status:    val.Status,
			StoreName: val.StoreName,
		}
		acqData = append(acqData, acq)
	}

	res.Data = models.ReportHistorySummaryRes{
		AchievementDay:  data.AchievementDay,
		AcquisitionData: acqData,
		Target:          data.Target,
	}
	res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)

	return
}


