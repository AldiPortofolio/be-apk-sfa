package services

import (
	"fmt"
	"github.com/vjeantet/jodaTime"
	"ottosfa-api-apk/constants"
	"ottosfa-api-apk/database/postgres"
	"ottosfa-api-apk/models"
	redis "ottosfa-api-apk/redis"
	"ottosfa-api-apk/utils"
	"time"
)

// CallPlanActionMerchantUnknown ..
func (svc *Service) CallPlanActionMerchantUnknown(bearer string, req models.CallPlanActionMerchantUnknownReq, res *models.Response) {
	fmt.Println(">>> CallPlanActionMerchantUnknown - Service <<<")

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

	// BEGIN upload photo to minio
	dataMinio, errMinio := svc.SendMinio(req.PhotoLocation, fmt.Sprintf("%d", req.CallPlanMerchantId)+"-CallPlan_PhotoLocation", svc.General.SpanId)
	if errMinio != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseDefault)
		return
	}
	// END upload photo to minio

	//update clockIn merchant action
	req.PhotoLocation = dataMinio.Url
	req.ClockOut = jodaTime.Format("yyyy-MM-dd HH:mm:ss", time.Now())
	if postgres.UpdateCallPlanMerchantUnknown(req) != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseFailed)
	}

	res.Meta = utils.GetMetaResponse("photo.location.merchant.send.success")

	return
}
