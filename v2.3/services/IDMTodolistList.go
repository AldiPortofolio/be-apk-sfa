package services

import (
	"encoding/json"
	"fmt"
	"github.com/vjeantet/jodaTime"
	"ottosfa-api-apk/constants"
	"ottosfa-api-apk/database/postgres"
	"ottosfa-api-apk/models"
	"ottosfa-api-apk/models/fdsmodels"
	redis "ottosfa-api-apk/redis"
	"ottosfa-api-apk/utils"
)

// IDMTodolistList ..
func (svc *Service) IDMTodolistList(bearer string, req models.TodolistListReq, res *models.Response) {
	fmt.Println(">>> IDMTodolistList - Service <<<")

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

	dataTodolistDB, err := postgres.TodolistList(req, villageId, sales.PhoneNumber)
	if err != nil || len(dataTodolistDB) == 0 {
		res.Meta = utils.GetMetaResponse("todolist.data.not.found")
		return
	}

	List := []models.TodolistListDBRes{}
	for _, val := range dataTodolistDB {

		fmt.Println("id merchant new rec", val.MerchantNewRecId)

		dataMerchant, _ := postgres.IDMMerchantTodolistList(val)

		longlatFDSRes := fdsmodels.LongLatMerchantRes{}
		longlatFDSReq := fdsmodels.LongLatMerchantReq{
			Phone: dataMerchant.PhoneNumber,
		}
		dataFDS, _ := svc.SendFDS(longlatFDSReq, "GETLONGLATMERCHANT", svc.General.SpanId)
		_ = json.Unmarshal(dataFDS, &longlatFDSRes)

		longitude := fmt.Sprintf("%f", longlatFDSRes.LongLatData.RegisterLongLat.Longitude)
		latitude := fmt.Sprintf("%f", longlatFDSRes.LongLatData.RegisterLongLat.Latitude)
		if val.TodolistCategoryId == 6 {
			longitude = dataMerchant.Longitude
			latitude = dataMerchant.Latitude
		}

		var a = models.TodolistListDBRes{
			MerchantName:          val.MerchantName,
			TaskDateString:        jodaTime.Format("dd-MM-YYYY", val.TaskDate),
			MerchantAddress:       dataMerchant.Address,
			MerchantID:            val.MerchantID,
			TodolistCategoryId:    val.TodolistCategoryId,
			MerchantNewRecId:      val.MerchantNewRecId,
			NameCategory:          val.NameCategory,
			Status:                val.Status,
			CustomerCode:          dataMerchant.CustomerCode,
			PhoneNumber:           dataMerchant.PhoneNumber,
			ID:                    val.ID,
			Reason:                val.Reason,
			PendingTaskDateString: jodaTime.Format("dd-MM-YYYY", val.PendingTaskDate),
			Longitude:             longitude, //val.Longitude,
			Latitude:              latitude,  //val.Latitude,
			//Longitude: fmt.Sprintf("%f", longlatFDSRes.LongLatData.RegisterLongLat.Longitude),
			//Latitude: fmt.Sprintf("%f", longlatFDSRes.LongLatData.RegisterLongLat.Latitude),
			VillageID: val.VillageID,
		}
		List = append(List, a)

	}

	res.Data = models.TodolistListRes{
		TodoList: List,
	}
	res.Meta = utils.GetMetaResponse("todolist.success")
	return
}
