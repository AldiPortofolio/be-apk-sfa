package services

import (
	"fmt"
	"ottosfa-api-apk/constants"
	"ottosfa-api-apk/database/postgres"
	"ottosfa-api-apk/models"
	redis "ottosfa-api-apk/redis"
	"ottosfa-api-apk/utils"
	"strconv"
)

// TodolistCount ..
func (svc *Service) TodolistCount(bearer string, res *models.Response) {
	fmt.Println(">>> TodolistCount - Service <<<")

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

	//get village_id by sales_id
	villageId, err := postgres.ListVillageByPositionSales(sales.ID)
	if err != nil || len(villageId) == 0 {
		res.Meta = utils.GetMetaResponse("todolist.data.not.found")
		return
	}

	//get count todolist
	countList, erry := postgres.GetCountTodolist(villageId, sales.PhoneNumber, strconv.Itoa(sales.SalesTypeId))
	if erry != nil {
		res.Meta = utils.GetMetaResponse("todolist.data.not.found")
		return
	}

	res.Meta = utils.GetMetaResponse("success")
	res.Data = countList
	return
}
