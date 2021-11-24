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
	"strconv"
)

// TodolistList ..
func (svc *Service) TodolistList(bearer string, req models.TodolistListReq, res *models.Response) {
	fmt.Println(">>> TodolistList - Service <<<")

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

	dataTodolistDB, err := postgres.TodolistListV252(req, villageId, sales.PhoneNumber, strconv.Itoa(sales.SalesTypeId))
	if err != nil || len(dataTodolistDB) == 0 {
		res.Meta = utils.GetMetaResponse("todolist.data.not.found")
		return
	}

	List := []models.TodolistListDBResV24{}
	for _, val := range dataTodolistDB {
		var longitude, latitude, customerCode string
		if val.TodolistCategoryId == 6 {
			dataMerchant, _ := postgres.MerchantTodolistListV25(val)
			longitude = dataMerchant.Longitude
			latitude = dataMerchant.Latitude
			customerCode = dataMerchant.CustomerCode
		}else{
			longlatFDSRes := fdsmodels.LongLatMerchantRes{}
			longlatFDSReq := fdsmodels.LongLatMerchantReq{
				Phone: val.PhoneNumber,
			}
			dataFDS, _ := svc.SendFDS(longlatFDSReq, "GETLONGLATMERCHANT", svc.General.SpanId)
			_ = json.Unmarshal(dataFDS, &longlatFDSRes)

			longitude = fmt.Sprintf("%f", longlatFDSRes.LongLatData.RegisterLongLat.Longitude)
			latitude = fmt.Sprintf("%f", longlatFDSRes.LongLatData.RegisterLongLat.Latitude)
		}

		var a = models.TodolistListDBResV24{
			MerchantName:          val.MerchantName,
			TaskDateString:        jodaTime.Format("dd-MM-YYYY", val.TaskDate),
			MerchantAddress:       val.MerchantAddress,
			MerchantID:            val.MerchantID,
			TodolistCategoryId:    val.TodolistCategoryId,
			MerchantNewRecId:      val.MerchantNewRecId,
			NameCategory:          val.NameCategory,
			Status:                val.Status,
			CustomerCode:          customerCode,
			PhoneNumber:           val.PhoneNumber,
			ID:                    val.ID,
			Reason:                val.Reason,
			PendingTaskDateString: jodaTime.Format("dd-MM-YYYY", val.PendingTaskDate),
			Longitude:             longitude,
			Latitude:              latitude,
			VillageID:     			val.VillageID,
			SalesTypeId:   			val.SalesTypeId,
			SalesTypeName: 			val.SalesTypeName,
			AddressBenchmark:		val.AddressBenchmark,
			SalesPhone: 			val.SalesPhone,
		}
		List = append(List, a)
	}

	res.Data = models.TodolistListRes{
		TodoList: List,
	}
	res.Meta = utils.GetMetaResponse("success")
	return
}
