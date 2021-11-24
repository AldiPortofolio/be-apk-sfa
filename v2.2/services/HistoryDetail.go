package services

import (
	"fmt"
	"ottosfa-api-apk/constants"
	"ottosfa-api-apk/hosts/rose"
	"ottosfa-api-apk/models"
	"ottosfa-api-apk/redis"
	"ottosfa-api-apk/utils"
)

// HistoryDetail ..
func (svc *Service) HistoryDetail(bearer string, req models.ReportHistoryDetailReq, res *models.Response) {
	fmt.Println(">>> HistoryDetail - Service <<<")

	token := utils.DecodeBearer(bearer)
	_, validateToken := redis.GetRedisKey(utils.RedisKeyAuth + token)
	if validateToken != nil {
		res.Meta = utils.GetMetaResponse("invalid-token")
		return
	}

	dataRose, err := rose.PencapaianSalesAll(req.Phone)
	if err != nil {
		fmt.Println("HistoryDetailvv")
		res.Meta = utils.GetMetaResponse(constants.KeyResponseDefault)
		return
	}

	details := []models.ReportHistoryDetailRes2{}
	for _, val := range dataRose.AcquisitionData {
		//fmt.Println("val.JoinAt: ", val.JoinAt)
		//var createdAt string //= jodaTime.Format("yyyy-MM-dd HH:mm:ss", val.JoinAt)
		//if val.JoinAt != "" { //2021-08-30T20:14:17.454985Z
		//	//layout := "2006-01-02T15:04:05.000Z"
		//	//t,_ := time.Parse(layout, val.JoinAt)
		//	//fmt.Println("t: ", t)
		//	//createdAt = jodaTime.Format("yyyy-MM-dd HH:mm:ss", t)
		//	createdAt = fmt.Sprintf("%s-%s-%s %s", val.JoinAt[0:4], val.JoinAt[5:7], val.JoinAt[8:10], val.JoinAt[12:10])
		//	//createdAt = jodaTime.Format("yyyy-MM-dd HH:mm:ss", t)
		//}

		acq := models.ReportHistoryDetailRes2{
			Created:       val.JoinAt,
			Address:       val.Address,
			MerchantPhoto: val.ImageMerchant,
			StoreName:     val.Name,
			Status:        val.MerchantStatus,
			MerchantID:    val.MerchantID,
			Phone:         val.PhoneNumber,
		}
		details = append(details, acq)
	}

	res.Data = models.ReportHistoryDetailRes{
		AcquisitionData: details,
	}
	res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)

	return
}


