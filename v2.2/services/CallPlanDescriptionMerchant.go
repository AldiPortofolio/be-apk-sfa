package services

import (
	"fmt"
	"ottosfa-api-apk/constants"
	"ottosfa-api-apk/database/postgres"
	"ottosfa-api-apk/models"
	redis "ottosfa-api-apk/redis"
	"ottosfa-api-apk/utils"
)

// CallPlanDescriptionMerchant ..
func (svc *Service) CallPlanDescriptionMerchant(bearer string, req models.CallPlanDescriptionMerchantReq, res *models.Response) {
	fmt.Println(">>> CallPlanDescriptionMerchant - Service <<<")

	token := utils.DecodeBearer(bearer)
	_, validateToken := redis.GetRedisKey(utils.RedisKeyAuth + token)
	if validateToken != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseInvalidToken)
		return
	}

	//cek data sales (get data salesId by Token)
	_, errDB := postgres.CheckToken(bearer[7:])
	if errDB != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseInvalidToken)
		return
	}

	//GET DATA DESCRIPTION MERCHANT
	dataDescMerchantDB, errDB := postgres.GetCallPlanDescriptionMerchant(req)
	if errDB != nil {
		//res.Meta = utils.GetMetaResponse("default")
		//return
	}

	res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)
	res.Data = dataDescMerchantDB
	return
}
