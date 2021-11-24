package services

import (
	"fmt"
	"ottosfa-api-apk/constants"
	"ottosfa-api-apk/database/postgres"
	"ottosfa-api-apk/models"
	redis "ottosfa-api-apk/redis"
	"ottosfa-api-apk/utils"
)

// TodolistFilterVillageList ..
func (svc *Service) TodolistFilterVillageList(bearer string, req models.TodolistFilterVillageListReq, res *models.Response) {
	fmt.Println(">>> TodolistFilterVillageList - Service <<<")

	token := utils.DecodeBearer(bearer)
	_, validateToken := redis.GetRedisKey(utils.RedisKeyAuth + token)
	if validateToken != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseInvalidToken)
		return
	}

	//cek data sales (get data salesId by Token)
	sales, errDB := postgres.CheckToken(bearer[7:])
	if errDB != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseInvalidToken)
		return
	}

	data, err := postgres.GetVillageListForTodolistFilter(sales.ID, sales.PhoneNumber, req.Keyword)
	if err != nil {
		res.Meta = utils.GetMetaResponse("todolist.village.list.not.found")
		return
	}

	res.Meta = utils.GetMetaResponse("success")
	res.Data = data
	return
}
