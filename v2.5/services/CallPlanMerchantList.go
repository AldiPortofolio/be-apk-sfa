package services

import (
	"encoding/json"
	"fmt"
	"ottosfa-api-apk/constants"
	"ottosfa-api-apk/database/postgres"
	"ottosfa-api-apk/models"
	"ottosfa-api-apk/models/fdsmodels"
	redis "ottosfa-api-apk/redis"
	"ottosfa-api-apk/utils"
)

// CallPlanMerchantList ..
func (svc *Service) CallPlanMerchantList(bearer string, req models.CallPlanMerchantListReq, res *models.Response) {
	fmt.Println(">>> CallPlanMerchantList - Service <<<")

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

	//GET DATA CALL PLAN MERCHANT LIST
	dataCallPlanMerchantListDB, errDB := postgres.GetCallPlanMerchantListv25(req)
	if errDB != nil {
		res.Meta = utils.GetMetaResponse("default")
		return
	}

	callPlanMerchantList := []models.CallPlanMerchantListResV24{}
	for _, val := range dataCallPlanMerchantListDB {
		//GET DATA LONG LAT FROM FDS
		longlatFDSRes := fdsmodels.LongLatMerchantRes{}
		longlatFDSReq := fdsmodels.LongLatMerchantReq{
			Phone: val.MerchantPhone,
		}
		dataFDS, _ := svc.SendFDS(longlatFDSReq, "GETLONGLATMERCHANT", svc.General.SpanId)
		_ = json.Unmarshal(dataFDS, &longlatFDSRes)

		//dataFDS, err := fds.LongLatMerchant(val.MerchantId, val.MerchantPhone)
		//if err != nil || longlatFDSRes.ResponCode == "02" {
		//	res.Meta = utils.GetMetaResponse("fds.error")
		//	//return
		//}

		//GET DATA TODOLIST
		var TodolistStatus = 0
		dataTodolistStatusOpenLateDB, _ := postgres.GetTodolistByMerchantIDAndStatusOpenLate(val.MerchantId)
		if len(dataTodolistStatusOpenLateDB) > 0 {
			TodolistStatus = 1
		} else {
			dataTodolistStatusPendingDB, _ := postgres.GetTodolistByMerchantIDAndStatusPending(val.MerchantId)
			if len(dataTodolistStatusPendingDB) > 0 {
				TodolistStatus = 2
			}
		}

		a := models.CallPlanMerchantListResV24{
			CallPlanMerchantId:    val.CallPlanMerchantId,
			MerchantName:          val.MerchantName,
			MerchantPhone:         val.MerchantPhone,
			Longitude:             fmt.Sprintf("%f", longlatFDSRes.LongLatData.RegisterLongLat.Longitude),
			Latitude:              fmt.Sprintf("%f", longlatFDSRes.LongLatData.RegisterLongLat.Latitude),
			MerchantId:            val.MerchantId,
			MerchantTypeName:      val.MerchantTypeName,
			MerchantAddress:       val.MerchantAddress,
			Priority:              val.Priority,
			AmountAction:          len(postgres.GetCallPlanListActionsAll(val.CallPlanMerchantId)),
			AmountActionCompleted: len(postgres.GetCallPlanListActionsCompleted(val.CallPlanMerchantId)),
			MerchantStatus:        val.MerchantStatus,
			TodolistStatus:        TodolistStatus,
			SalesTypeId:           val.SalesTypeId,
			SalesTypeName:         val.SalesTypeName,
			AddressBenchmark: 	   val.AddressBenchmark,
		}

		callPlanMerchantList = append(callPlanMerchantList, a)

	}

	res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)
	res.Data = callPlanMerchantList
	return
}
