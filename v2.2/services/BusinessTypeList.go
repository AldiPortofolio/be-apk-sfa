package services

import (
	"fmt"
	ottoutils "ottodigital.id/library/utils"
	"ottosfa-api-apk/constants"
	"ottosfa-api-apk/database/postgres"
	"ottosfa-api-apk/hosts/rose"
	"ottosfa-api-apk/models"
	redis "ottosfa-api-apk/redis"
	"ottosfa-api-apk/utils"
)

// BusinessTypeList ..
func (svc *Service) BusinessTypeList(bearer string, res *models.Response) {
	fmt.Println(">>> BusinessTypeList - Service <<<")

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

	var data models.BusinessTypeListRes
	tipeBisnis := ottoutils.GetEnv("LOOKUP_TIPE_BISNIS", "")
	dataTipeBisnisRose, errRose := rose.LookUpGroup(tipeBisnis)
	if errRose != nil {
		res.Meta = utils.GetMetaResponse("default")
		return
	}
	for _, val := range dataTipeBisnisRose {
		a := models.DataModule{
			Code: val.Code,
			Name: val.Name,
		}
		data.TipeBisnis = append(data.TipeBisnis, a)
	}

	res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)
	res.Data = data
	return
}
