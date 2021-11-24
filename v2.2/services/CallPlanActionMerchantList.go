package services

import (
	"fmt"
	"ottosfa-api-apk/constants"
	"ottosfa-api-apk/database/postgres"
	"ottosfa-api-apk/models"
	redis "ottosfa-api-apk/redis"
	"ottosfa-api-apk/utils"
)

// CallPlanActionMerchantList ..
func (svc *Service) CallPlanActionMerchantList(bearer string, res *models.Response) {
	fmt.Println(">>> CallPlanActionMerchantList - Service <<<")

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

	//GET DATA ACTION MERCHANT LIST
	dataActionMerchantListDB, errDB := postgres.GetCallPlanActionMerchantList()
	if errDB != nil {
		res.Meta = utils.GetMetaResponse("default")
		return
	}

	res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)
	res.Data = dataActionMerchantListDB
	return
}
