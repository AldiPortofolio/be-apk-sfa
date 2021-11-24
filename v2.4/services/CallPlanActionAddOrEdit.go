package services

import (
	"fmt"
	"github.com/vjeantet/jodaTime"
	"ottosfa-api-apk/constants"
	"ottosfa-api-apk/database/dbmodels"
	"ottosfa-api-apk/database/postgres"
	"ottosfa-api-apk/hosts/rose"
	"ottosfa-api-apk/models"
	redis "ottosfa-api-apk/redis"
	"ottosfa-api-apk/utils"
	"strings"
	"time"
)

// CallPlanActionAddOrEdit ..
func (svc *Service) CallPlanActionAddOrEdit(bearer string, req models.CallPlanActionAddOrEditReq, res *models.Response) {
	fmt.Println(">>> CallPlanActionAddOrEdit - Service <<<")

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

	var data dbmodels.CallPlanActions
	if req.Id > 0 {
		//data, errDB = postgres.GetCallPlanActionByID(req.Id)
		//if errDB != nil {
		//	res.Meta = utils.GetMetaResponse("default")
		//	return
		//}
		//edit
		data = dbmodels.CallPlanActions{
			Id:                 req.Id,
			CallPlanMerchantId: req.CallPlanMerchantId,
			Name:               req.Name,
			Action:             req.Action,
			ActionType:         req.ActionType,
			Product:            req.Product,
			Description:        req.Description,
			MerchantAction:     req.MerchantAction, //Action 4
			Result:             req.Result,
			Amount:             req.Amount,
			Reason:             req.Reason,
			Note:               req.Note,
			Status:             req.Status,
			CreatedAt:          req.CreatedAt,
			UpdatedAt:          jodaTime.Format("yyyy-MM-dd HH:mm:ss", time.Now()),
		}
	} else {
		//add
		data = dbmodels.CallPlanActions{
			Id:                 0,
			CallPlanMerchantId: req.CallPlanMerchantId,
			Name:               req.Name,
			Action:             req.Action,
			ActionType:         req.ActionType,
			Product:            req.Product,
			Description:        req.Description,
			MerchantAction:     req.MerchantAction, //Action 4
			Result:             req.Result,
			Amount:             req.Amount,
			Reason:             req.Reason,
			Note:               req.Note,
			Status:             req.Status,
			CreatedAt:          jodaTime.Format("yyyy-MM-dd HH:mm:ss", time.Now()),
			UpdatedAt:          jodaTime.Format("yyyy-MM-dd HH:mm:ss", time.Now()),
		}
	}

	//update call plan action
	if errDB = postgres.SaveCallPlanAction(data); errDB != nil {
		res.Meta = utils.GetMetaResponse("default")
		return
	}

	if strings.ToLower(req.Action) == "update" && strings.ToLower(req.Product) == "qris" {
		fmt.Println(">>> CallPlanActionAddOrEdit - Service - Update QRIS <<<")
		dataMerchantDB, errDB := postgres.GetMerchantCallPlanByMerchantPhone(req.MerchantPhone)
		if errDB != nil {
			res.Meta = utils.GetMetaResponse("qris.merchant.not.found.db")
			return
		}

		_, getDataRoseErr := rose.FindByMid(req.MerchantId)
		if getDataRoseErr != nil {
			res.Meta = utils.GetMetaResponse("qris.merchant.not.found.rose")
			return
		}

		roseReq := rose.ReplaceMidMpanRequest{
			Mid:                 dataMerchantDB.MerchantId,    //dari DB SFA table CallPlanMerchant
			Mpan:                dataMerchantDB.Mpan,          //dari DB SFA table CallPlanMerchant
			Nmid:                req.NMID,                     //dari Scan QR
			StoreNamePreprinted: req.StoreNamePreprinted,      //dari Scan QR
			StorePhoneNumber:    dataMerchantDB.MerchantPhone, //dari DB SFA table CallPlanMerchant
		}

		getDataRoseErr = rose.ReplaceMidMpan(roseReq)
		if getDataRoseErr != nil {
			res.Meta = utils.GetMetaResponse("qris.merchant.not.found.rose")
			return
		}

		//GET DATA ACTION LIST
		callPlanListActionDB, errDB := postgres.GetCallPlanListActionsAllWithError(req.CallPlanMerchantId)
		if errDB != nil {
			res.Meta = utils.GetMetaResponse("default")
			return
		}

		res.Meta.Message = utils.GetMetaResponse("qris.merchant.not.found.rose").Message + req.StoreNamePreprinted
		res.Meta.Code = utils.GetMetaResponse("qris.merchant.not.found.rose").Code
		res.Meta.Status = utils.GetMetaResponse("qris.merchant.not.found.rose").Status
		res.Data = callPlanListActionDB

		return
	}

	//GET DATA ACTION LIST
	callPlanListActionDB, errDB := postgres.GetCallPlanListActionsAllWithError(req.CallPlanMerchantId)
	if errDB != nil {
		res.Meta = utils.GetMetaResponse("default")
		return
	}

	res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)
	res.Data = callPlanListActionDB

	return
}
