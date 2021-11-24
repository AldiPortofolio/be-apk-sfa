package services

import (
	"fmt"
	"ottosfa-api-apk/constants"
	"ottosfa-api-apk/database/postgres"
	"ottosfa-api-apk/hosts/rose"
	"ottosfa-api-apk/models"
	redis "ottosfa-api-apk/redis"
	"ottosfa-api-apk/utils"
)

// MerchantDetailQRIS ..
func (svc *Service) MerchantDetailQRIS(bearer string, req models.MerchantDetailQRISReq, res *models.Response) {
	fmt.Println(">>> MerchantDetailQRIS - Service <<<")

	// if merchant_phone is empty, do not continue
	if req.MerchantPhone == "" {
		res.Meta = utils.GetMetaResponse("default")
		return
	}

	token := utils.DecodeBearer(bearer)
	_, validateToken := redis.GetRedisKey(utils.RedisKeyAuth + token)
	if validateToken != nil {
		res.Meta = utils.GetMetaResponse("invalid-token")
		return
	}

	//GET DATA MERCHANT
	dataMerchantDB, errDB := postgres.GetMerchantByMerchantPhone(req.MerchantPhone)
	if errDB != nil {
		res.Meta = utils.GetMetaResponse("qris.merchant.not.found.db")
		return
	}

	dataRose, getDataRoseErr := rose.FindByMid(dataMerchantDB.MerchantId)
	if getDataRoseErr != nil {
		res.Meta = utils.GetMetaResponse("qris.merchant.not.found.rose")
		return
	}

	res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)
	res.Data = models.MerchantDetailQRISRes{
		MerchantId:    dataMerchantDB.MerchantId,
		Mpan:          dataMerchantDB.Mpan,
		MerchantPhone: dataMerchantDB.MerchantPhone, //dataRose.PhoneNumber,
		MerchantName:  dataRose.MerchantName,
	}
	return
}
