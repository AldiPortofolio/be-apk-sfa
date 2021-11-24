package services

import (
	"fmt"
	"ottosfa-api-apk/constants"
	"ottosfa-api-apk/database/dbmodels"
	"ottosfa-api-apk/database/postgres"
	"ottosfa-api-apk/models"
	redis "ottosfa-api-apk/redis"
	"ottosfa-api-apk/utils"
	"strings"
	"time"
)

// CallPlanTodolistMerchantNotFound ..
func (svc *Service) CallPlanTodolistMerchantNotFound(bearer string, req models.CallPlanTodolistMerchantNotFoundReq, res *models.Response) {
	fmt.Println(">>> CallPlanTodolistMerchantNotFound - Service <<<")

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

	listIssue := []string{}
	for _, val := range req.Issue {
		listIssue = append(listIssue, val)
	}
	issue := strings.Join(listIssue[:], "&")

	// BEGIN upload photo to minio
	dataMinio, errMinio := svc.SendMinio(req.MerchantImage, "image-location", svc.General.SpanId)
	if errMinio != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseDefault)
		return
	}
	// END upload photo to minio

	//GET DATA TODOLIST LIST
	dataTodolistDB, errDB := postgres.GetTodolistByMerchantIDAndStatusOpenLate(req.MerchantId)
	if errDB != nil {
		res.Meta = utils.GetMetaResponse("default")
		return
	}

	for _, val := range dataTodolistDB {
		//UPDATE TODOLIST
		todolistReq := dbmodels.TodoLists{
			ID:         val.ID,
			ActionDate: time.Now(),
			SalesPhone: sales.PhoneNumber,
			Status:     "Not Exist",
			Longitude:  req.Longitude,
			Latitude:   req.Latitude,
			UpdatedAt:  time.Now(),
		}
		if postgres.UpdateTodolistMerchantNotFound(todolistReq) != nil {
			res.Meta = utils.GetMetaResponse("default")
			return
		}

		//INSERT TO HISTORY TODOLIST
		todoListHistoriesReq := dbmodels.TodoListHistories{
			TodoListId:   val.ID,
			Description:  issue,
			Status:       "Not Exist",
			NewTaskDate:  time.Now(),
			OldTaskDate:  time.Now(),
			FotoLocation: dataMinio.Url,
			Longitude:    req.Longitude,
			Latitude:     req.Latitude,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}
		if postgres.SaveTodolistHistories(todoListHistoriesReq) != nil {
			res.Meta = utils.GetMetaResponse("default")
			return
		}
	}

	res.Meta = utils.GetMetaResponse("todolist.merchant.not.found")
	return
}
