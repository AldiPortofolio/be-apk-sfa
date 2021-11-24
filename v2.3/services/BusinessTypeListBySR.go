package services

import (
	"fmt"
	"ottosfa-api-apk/constants"
	"ottosfa-api-apk/database/postgres"
	"ottosfa-api-apk/models"
	redis "ottosfa-api-apk/redis"
	"ottosfa-api-apk/utils"
)

// BusinessTypeListBySR ..
func (svc *Service) BusinessTypeListBySR(bearer string, req models.BusinessTypeListBySRReq, res *models.Response) {
	fmt.Println(">>> BusinessTypeListBySR - Service <<<")

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

	var businessType models.BusinessTypeListRes
	data, err := postgres.GetBusinessTypeListBySR(req.SalesTypeId)
	if err != nil {
		res.Meta = utils.GetMetaResponse("default")
		return
	}
	for _, val := range data {
		a := models.DataModule{
			Code: val.Code,
			Name: val.Name,
		}
		businessType.TipeBisnis = append(businessType.TipeBisnis, a)
	}

	res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)
	res.Data = businessType
	return
}
