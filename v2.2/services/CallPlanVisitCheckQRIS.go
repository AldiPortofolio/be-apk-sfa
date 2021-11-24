package services

import (
	"fmt"
	"github.com/vjeantet/jodaTime"
	"ottosfa-api-apk/constants"
	"ottosfa-api-apk/database/postgres"
	"ottosfa-api-apk/hosts/rose"
	"ottosfa-api-apk/hosts/stringbuilder"
	"ottosfa-api-apk/models"
	redis "ottosfa-api-apk/redis"
	"ottosfa-api-apk/utils"
	"time"
)

// CallPlanVisitCheckQRIS ..
func (svc *Service) CallPlanVisitCheckQRIS(bearer string, req models.CheckQRISReq, res *models.Response) {
	fmt.Println(">>> CallPlanVisitCheckQRIS - Service <<<")

	//log := svc.OttoLog

	// if qr content is empty, do not continue
	if req.QRContent == "" {
		res.Meta = utils.GetMetaResponse("default")
		return
	}

	token := utils.DecodeBearer(bearer)
	_, validateToken := redis.GetRedisKey(utils.RedisKeyAuth + token)
	if validateToken != nil {
		res.Meta = utils.GetMetaResponse("invalid-token")
		return
	}

	//connect to string builder
	dataStringBuilder, getDataStringBuilderErr := stringbuilder.ReverseQR(req.QRContent)
	if getDataStringBuilderErr != nil {
		res.Meta = utils.GetMetaResponse("qr.not.match")
		return
	}

	//dataMerchantDB := models.CallPlanVisitMerchantRes{}
	if dataStringBuilder.Tag00 == "01" {
		if dataStringBuilder.Tag6207 != "" && dataStringBuilder.Tag2601 != "" {
			//get phone number by API ROSE
			dataRose, err := rose.InquiryMerchant(dataStringBuilder.Tag6207, dataStringBuilder.Tag2601)
			if err != nil {
				res.Meta = utils.GetMetaResponse("qr.not.match")
				return
			}

			//GET DATA MERCHANT
			dataMerchantDB, errDB := postgres.GetMerchantByMerchantPhone(dataRose.PhoneNumber)
			if errDB != nil {
				res.Meta = utils.GetMetaResponse("merchant.not.found")
				return
			}
			dataMerchantDB.ClockIn = jodaTime.Format("yyyy-MM-dd HH:mm:ss", time.Now())

			res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)
			res.Data = dataMerchantDB
			return

		}
		res.Meta = utils.GetMetaResponse("qr.not.match")
		return

	}
	//else{
	res.Meta = utils.GetMetaResponse("qr.not.match")
	return
	//}

	//res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)
	//res.Data = dataMerchantDB
	//return
}
