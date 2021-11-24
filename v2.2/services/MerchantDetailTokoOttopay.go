package services

import (
	"fmt"
	"ottosfa-api-apk/constants"
	"ottosfa-api-apk/hosts/rose"
	"ottosfa-api-apk/models"
	redis "ottosfa-api-apk/redis"
	"ottosfa-api-apk/utils"
)

// MerchantDetailTokoOttopay ..
func (svc *Service) MerchantDetailTokoOttopay(bearer string, req models.MerchantDetailQRISReq, res *models.Response) {
	fmt.Println(">>> MerchantDetailTokoOttopay - Service <<<")

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
	dataRose, getDataRoseErr := rose.FindByPhoneNumber(req.MerchantPhone)
	if getDataRoseErr != nil {
		res.Meta = utils.GetMetaResponse("merchant.not.found")
		return
	}

	res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)
	res.Data = models.MerchantDetailTokoOttopayRes{
		MerchantId:    dataRose.MID,
		MerchantPhone: req.MerchantPhone,
		MerchantName:  dataRose.StoreName,
		OwnerName:     dataRose.OwnerName,
		Address:       dataRose.Address,
	}
	return
}
