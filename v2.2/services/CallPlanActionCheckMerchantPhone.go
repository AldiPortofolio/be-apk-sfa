package services

import (
	"fmt"
	"ottosfa-api-apk/constants"
	"ottosfa-api-apk/database/postgres"
	"ottosfa-api-apk/models"
	redis "ottosfa-api-apk/redis"
	"ottosfa-api-apk/utils"
)

// CallPlanActionCheckMerchantPhone ..
func (svc *Service) CallPlanActionCheckMerchantPhone(bearer string, req models.CallPlanActionCheckMerchantPhoneReq, res *models.Response) {
	fmt.Println(">>> CallPlanActionCheckMerchantPhone - Service <<<")

	// if request null
	if req.MerchantPhone == "" && req.Latitude == "" && req.Longitude == "" {
		res.Meta = utils.GetMetaResponse("default")
		return
	}

	token := utils.DecodeBearer(bearer)
	_, validateToken := redis.GetRedisKey(utils.RedisKeyAuth + token)
	if validateToken != nil {
		res.Meta = utils.GetMetaResponse("invalid-token")
		return
	}

	callPlanMerchantDB, errDB := postgres.GetCallPlanMerchantsByCallPlanMerchantId(req.CallPlanMerchantId)
	if errDB != nil {
		res.Meta = utils.GetMetaResponse("default")
		return
	}

	if callPlanMerchantDB.MerchantPhone == req.MerchantPhone {
		//UPDATE LONG LAT CALL PLAN MERCHANT
		if postgres.UpdateLongLatCallPlanMerchant(req.CallPlanMerchantId, req.Longitude, req.Latitude) != nil {
			res.Meta = utils.GetMetaResponse("merchant.phone.not.match")
			return
		}

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

	res.Meta = utils.GetMetaResponse("merchant.phone.not.match")
	return

}
