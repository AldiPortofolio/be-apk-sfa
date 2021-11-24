package services

import (
	"fmt"
	"ottosfa-api-apk/constants"
	"ottosfa-api-apk/database/dbmodels"
	"ottosfa-api-apk/database/postgres"
	"ottosfa-api-apk/models"
	redis "ottosfa-api-apk/redis"
	"ottosfa-api-apk/utils"
)

// CallPlanActionDelete ..
func (svc *Service) CallPlanActionDelete(bearer string, req models.CallPlanActionDeleteReq, res *models.Response) {
	fmt.Println(">>> CallPlanActionDelete - Service <<<")

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

	reqCallPlanAction := dbmodels.CallPlanActions{
		Id: req.CallPlanActionId,
		//ActionType: "Additional",
	}

	//delete call plan action
	if errDB = postgres.DeleteCallPlanAction(reqCallPlanAction); errDB != nil {
		res.Meta = utils.GetMetaResponse("default")
		return
	}

	res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)

	return
}
