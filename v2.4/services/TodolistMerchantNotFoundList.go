package services

import (
	"fmt"
	"ottosfa-api-apk/constants"
	"ottosfa-api-apk/database/postgres"
	"ottosfa-api-apk/models"
	redis "ottosfa-api-apk/redis"
	"ottosfa-api-apk/utils"
)

// TodolistMerchantNotFoundList ..
func (svc *Service) TodolistMerchantNotFoundList(bearer string, req models.TodolistMerchantNotFoundListReq, res *models.Response) {
	fmt.Println(">>> TodolistMerchantNotFoundList - Service <<<")

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

	data, err := postgres.MerchantNotFoundListV24(req)
	if err != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseFailed)
		return
	}

	res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)
	res.Data = data
	return
}
