package services

import (
	"encoding/json"
	"fmt"
	"ottosfa-api-apk/models"
	redis "ottosfa-api-apk/redis"
	"ottosfa-api-apk/utils"
)

// CheckIndomarcoAccount ..
func (svc *Service) CheckIndomarcoAccount(bearer string, req models.CheckIndomarcoAccountReq, res *models.Response) {
	fmt.Println(">>> CheckIndomarcoAccount - Service <<<")

	token := utils.DecodeBearer(bearer)
	_, validateToken := redis.GetRedisKey(utils.RedisKeyAuth + token)
	if validateToken != nil {
		res.Meta = utils.GetMetaResponse("invalid-token")
		return
	}

	if len(req.CustomerCode) < 13 {
		res.Meta = utils.GetMetaResponse("customer.code.not.valid")
		return
	}

	dbyteIDM, errIDM := svc.SendIDM(req, "CHECKACCOUNT", svc.General.SpanId)
	if errIDM != nil {
		res.Meta = utils.GetMetaResponse("default")
		return
	}

	data := models.CheckAccountIndomarcoRes{}
	err := json.Unmarshal(dbyteIDM, &data)
	if err != nil {
		res.Meta = utils.GetMetaResponse("default")
		return
	}

	res.Meta = data.Meta
	res.Data = data.Data
	return
}
