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

// CallPlanTodolistList ..
func (svc *Service) CallPlanTodolistList(bearer string, req models.CallPlanTodolistListReq, res *models.Response) {
	fmt.Println(">>> CallPlanTodolistList - Service <<<")

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

	//GET DATA TODOLIST LIST
	dataTodolistDB, errDB := postgres.GetListTodolist(req)
	if errDB != nil {
		res.Meta = utils.GetMetaResponse("default")
		return
	}

	callPlanTodolistList := []models.CallPlanTodolistListRes{}
	//if len(dataTodolistDB) > 1 {
	for _, val := range dataTodolistDB {
		//GET DATA LONG LAT FROM FDS
		pendingTask := ""
		if jodaTime.Format("dd-MM-YYYY", val.PendingTaskDate) != "01-01-1" {
			pendingTask = jodaTime.Format("dd-MM-YYYY", val.PendingTaskDate)
		}

		var a = models.CallPlanTodolistListRes{
			MerchantName:          val.MerchantName,
			TaskDateString:        jodaTime.Format("dd-MM-YYYY", val.TaskDate),
			MerchantAddress:       val.MerchantAddress,
			MerchantID:            val.MerchantID,
			NameCategory:          val.NameCategory,
			Status:                val.Status,
			ID:                    val.ID,
			Reason:                val.Reason,
			PendingTaskDateString: pendingTask,
			VillageID:             val.VillageID,
		}
		callPlanTodolistList = append(callPlanTodolistList, a)
	}
	//}else{
	//	dataTodolisDetailtDB, errDB := postgres.GetListTodolist(req)
	//}

	res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)
	res.Data = models.TodoList{
		TodoList: callPlanTodolistList,
	}

	return
}
