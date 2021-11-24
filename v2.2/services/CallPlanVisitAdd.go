package services

import (
	"fmt"
	"github.com/vjeantet/jodaTime"
	"ottosfa-api-apk/constants"
	"ottosfa-api-apk/database/dbmodels"
	"ottosfa-api-apk/database/postgres"
	"ottosfa-api-apk/models"
	redis "ottosfa-api-apk/redis"
	"ottosfa-api-apk/utils"
	"time"
)

// CallPlanVisitAdd ..
func (svc *Service) CallPlanVisitAdd(bearer string, req models.CallPlanVisitAddReq, res *models.Response) {
	fmt.Println(">>> CallPlanVisitAdd - Service <<<")

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

	effectiveCall := false
	if req.Amount > 0 {
		effectiveCall = true
	}
	//INSERT DATA CALL PLAN MERCHANT
	callPlanMerchant := dbmodels.CallPlanMerchants{
		CallPlanId:      req.CallPlanId,
		MerchantId:      req.IdMerchant,
		MerchantPhone:   req.MerchantPhone,
		MerchantAddress: req.MerchantAddress,
		MerchantTypeId:  req.MerchantTypeId,
		Status:          "Visited",
		MerchantStatus:  req.MerchantStatus, //Found - Open
		EffectiveCall:   effectiveCall,
		Amount:          req.Amount,
		ClockIn:         req.ClockIn,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		MID:             req.MerchantId,
		MerchantName:    req.MerchantName,
		ClockOut:        jodaTime.Format("yyyy-MM-dd HH:mm:ss", time.Now()),
		Longitude:       req.Longitude,
		Latitude:        req.Latitude,
		PhotoLocation:   "", //kosong kalo dari visited
	}

	callPlanMerchantDB, errDB := postgres.GetCallPlanMerchantsByMerchantPhone(req.MerchantPhone, req.CallPlanId)
	//if errDB != nil {
	//	res.Meta = utils.GetMetaResponse("default")
	//	return
	//}

	if callPlanMerchantDB.MerchantPhone != req.MerchantPhone && errDB != nil {
		_, errDB = postgres.AddCallPlanMerchant(callPlanMerchant)
		if errDB != nil {
			res.Meta = utils.GetMetaResponse("default")
			return
		}

		callPlanMerchantDB, errDB = postgres.GetCallPlanMerchantsByMerchantPhone(req.MerchantPhone, req.CallPlanId)
		if errDB != nil {
			res.Meta = utils.GetMetaResponse("default")
			return
		}
	}

	//INSERT DATA CALL PLAN ACTION
	callPlanAction := dbmodels.CallPlanActions{
		CallPlanMerchantId: callPlanMerchantDB.Id,
		Name:               req.CallPlanActionName,
		Action:             req.ActionName,
		ActionType:         "Visited",
		Product:            req.ProductName,
		Description:        req.Description,
		MerchantAction:     req.MerchantAction,
		Result:             req.Result,
		Amount:             req.Amount,
		Reason:             req.Reason,
		Note:               req.Note,
		Status:             "Completed",
		CreatedAt:          jodaTime.Format("yyyy-MM-dd HH:mm:ss", time.Now()),
		UpdatedAt:          jodaTime.Format("yyyy-MM-dd HH:mm:ss", time.Now()),
	}

	errDB = postgres.AddCallPlanAction(callPlanAction)
	if errDB != nil {
		res.Meta = utils.GetMetaResponse("default")
		return
	}

	//GET DATA ACTION LIST
	callPlanListActionDB, errDB := postgres.GetCallPlanListActionsAllWithError(callPlanMerchantDB.Id)
	if errDB != nil {
		res.Meta = utils.GetMetaResponse("default")
		return
	}

	res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)
	res.Data = callPlanListActionDB

	return
}
