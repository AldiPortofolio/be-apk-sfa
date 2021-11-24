package services

import (
	"fmt"
	"github.com/vjeantet/jodaTime"
	"ottosfa-api-apk/constants"
	"ottosfa-api-apk/database/postgres"
	"ottosfa-api-apk/models"
	redis "ottosfa-api-apk/redis"
	"ottosfa-api-apk/utils"
)

// TodolistDetail ..
func (svc *Service) TodolistDetail(bearer string, req models.TodolistDetailReq, res *models.Response) {
	fmt.Println(">>> TodolistDetail - Service <<<")

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

	dataDB, erry := postgres.GetDetailTodolist(req)
	if erry != nil || len(dataDB) == 0 {
		if req.TodolistCategoryId == 6 {
			res.Meta = utils.GetMetaResponse("todolist.detail.not.found.new.rec")
			return
		}
		res.Meta = utils.GetMetaResponse("todolist.detail.not.found")
		return
	}

	taskTodolist := []models.TaskTodolistRes{}
	for _, val := range dataDB {
		data, _ := postgres.CekStatusTaskTodolist(val.TaskID)
		status := false
		if data.TaskID == val.TaskID {
			status = true
		}
		var a = models.TaskTodolistRes{
			TodolistSubCategoryID: val.TodolistSubCategoryID,
			Code:                  val.Code,
			Name:                  val.Name,
			TaskID:                int64(val.TaskID),
			Status:                status,
			LabelType:             val.LabelType,
			Link:                  val.Link,
		}
		taskTodolist = append(taskTodolist, a)
	}

	data := models.TodolistDetailRes{
		TodolistID:         dataDB[0].TodolistID,
		IdMerchant:         dataDB[0].IdMerchant,
		MerchantName:       dataDB[0].MerchantName,
		TaskDateString:     jodaTime.Format("dd-MM-YYYY", dataDB[0].TaskDate),
		CreatedAtString:    jodaTime.Format("dd-MM-YYYY", dataDB[0].CreatedAt),
		MerchantAddress:    dataDB[0].MerchantAddress,
		MerchantID:         dataDB[0].MerchantID,
		MerchantPhone:      dataDB[0].MerchantPhone,
		CustomerCode:       dataDB[0].CustomerCode,
		NameCategory:       dataDB[0].NameCategory,
		TodolistCategoryId: dataDB[0].TodolistCategoryId,
		Status:             dataDB[0].Status,
		Reason:             dataDB[0].Reason,
		Task:               taskTodolist,
	}

	res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)
	res.Data = data
	return
}
