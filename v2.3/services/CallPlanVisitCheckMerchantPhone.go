package services

import (
	"fmt"
	"github.com/vjeantet/jodaTime"
	"ottosfa-api-apk/constants"
	"ottosfa-api-apk/database/postgres"
	"ottosfa-api-apk/database/postgresrose"
	"ottosfa-api-apk/models"
	redis "ottosfa-api-apk/redis"
	"ottosfa-api-apk/utils"
	"time"
)

// CallPlanVisitCheckMerchantPhone ..
func (svc *Service) CallPlanVisitCheckMerchantPhone(bearer string, req models.CallPlanVisitCheckMerchantPhoneReq, res *models.Response) {
	fmt.Println(">>> CallPlanVisitCheckMerchantPhone - Service <<<")

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

	//GET DATA MERCHANT
	dataMerchantDBRose, errDB := postgresrose.GetMerchantByMerchantPhone(req.MerchantPhone)
	if errDB != nil {
		res.Meta = utils.GetMetaResponse("merchant.not.found")
		return
	}

	dataMerchantDB, errDB := postgres.GetMerchantTypeById(dataMerchantDBRose.MerchantTypeId)
	if errDB != nil {
		res.Meta = utils.GetMetaResponse("merchant.not.found")
		return
	}

	dataMerchantDBRose.MerchantTypeName = dataMerchantDB.Name
	dataMerchantDBRose.ClockIn = jodaTime.Format("yyyy-MM-dd HH:mm:ss", time.Now())

	res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)
	res.Data = dataMerchantDBRose
	return
}
