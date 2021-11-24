package services

import (
	"fmt"
	"ottosfa-api-apk/constants"
	"ottosfa-api-apk/database/postgres"
	"ottosfa-api-apk/models"
	redis "ottosfa-api-apk/redis"
	"ottosfa-api-apk/utils"
)

// CallPlanDetail ..
func (svc *Service) CallPlanDetail(bearer string, req models.CallPlanDetailReq, res *models.Response) {
	fmt.Println(">>> CallPlanDetail - Service <<<")

	token := utils.DecodeBearer(bearer)
	_, validateToken := redis.GetRedisKey(utils.RedisKeyAuth + token)
	if validateToken != nil {
		res.Meta = utils.GetMetaResponse("invalid-token")
		return
	}

	//callPlanMerchantDB, errDB := postgres.GetCallPlanMerchantsByCallPlanMerchantId(req.CallPlanMerchantId)
	//if errDB != nil {
	//	res.Meta = utils.GetMetaResponse("default")
	//	return
	//}

	//GET DATA ACTION LIST
	callPlanListActionDB, errDB := postgres.GetCallPlanListActionsAllWithError(req.CallPlanMerchantId)
	if errDB != nil {
		res.Meta = utils.GetMetaResponse("merchant.phone.not.match")
		return
	}

	res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)
	res.Data = callPlanListActionDB
	return
}
