package services

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"ottosfa-api-apk/constants"
	"ottosfa-api-apk/database/postgres"
	"ottosfa-api-apk/models"
	redis "ottosfa-api-apk/redis"
	"ottosfa-api-apk/utils"
	"strings"
)

// TodolistTaskBySubCategory ..
func (svc *Service) TodolistTaskBySubCategory(bearer string, req models.TodolistTaskBySubCategoryReq, res *models.Response) {
	fmt.Println(">>> TodolistTaskBySubCategory - Service <<<")

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

	dataDB, err := postgres.GetTaskTodolistBySubCategory(req)
	if err != nil {
		res.Meta = utils.GetMetaResponse("task.todolist.not.found")
		return
	}

	data := []models.TodolistTaskBySubCategoryRes{}
	for _, val := range dataDB {
		x := strings.Split(val.Name, ";")

		var y []string
		for i := 0; i < len(x); i++ {
			if strings.HasSuffix(x[i], "[Text]") {
				x[i] = strings.TrimSuffix(x[i], "[Text]")
			}
			if strings.Contains(x[i], "[supplier]") {
				logs.Info("val.SupplierName ", val.SupplierName)
				x[i] = strings.Replace(x[i], "[supplier]", val.SupplierName, -1)
			}
			y = append(y, x[i])
		}

		a := models.TodolistTaskBySubCategoryRes{
			ID:            val.ID,
			Name:          y,
			LabelType:     val.LabelType,
			SubCategoryID: val.SubCategoryID,
			CreatedAt:     val.CreatedAt,
			UpdatedAt:     val.UpdatedAt,
			Condition:     val.Condition,
			Step:          val.Step,
		}
		data = append(data, a)
	}

	res.Meta = utils.GetMetaResponse("success")
	res.Data = data
	return
}
