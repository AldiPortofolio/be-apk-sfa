package services

import (
	"fmt"
	"ottosfa-api-apk/constants"
	"ottosfa-api-apk/database/postgres"
	"ottosfa-api-apk/hosts/rose"
	"ottosfa-api-apk/models"
	redis "ottosfa-api-apk/redis"
	"ottosfa-api-apk/utils"
)

// UserCategory ..
func (svc *Service) UserCategory(bearer string, res *models.Response) {
	fmt.Println(">>> UserCategory - Service <<<")

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

	var data []models.UserCategoryRes
	dataUserCategoryRose, errRose := rose.UserCategory()
	if errRose != nil {
		res.Meta = utils.GetMetaResponse("default")
		return
	}
	for _, val := range dataUserCategoryRose {
		a := models.UserCategoryRes{
			Code:   val.Code,
			Name:   val.Notes,
			Logo:   val.Logo,
			Status: val.Status,
		}
		data = append(data, a)
	}

	res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)
	res.Data = data
	return
}
