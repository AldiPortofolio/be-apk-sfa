package services

import (
	"fmt"
	"ottosfa-api-apk/constants"
	"ottosfa-api-apk/database/postgres"
	"ottosfa-api-apk/hosts/rose"
	"ottosfa-api-apk/hosts/stringbuilder"
	"ottosfa-api-apk/models"
	redis "ottosfa-api-apk/redis"
	"ottosfa-api-apk/utils"
)

// CallPlanActionCheckQRIS ..
func (svc *Service) CallPlanActionCheckQRIS(bearer string, req models.CallPlanActionCheckQRISReq, res *models.Response) {
	fmt.Println(">>> CallPlanActionCheckQRIS - Service <<<")

	//log := svc.OttoLog

	// if qr content is empty, do not continue
	if req.QRContent == "" && req.Latitude == "" && req.Longitude == "" {
		res.Meta = utils.GetMetaResponse("default")
		return
	}

	token := utils.DecodeBearer(bearer)
	_, validateToken := redis.GetRedisKey(utils.RedisKeyAuth + token)
	if validateToken != nil {
		res.Meta = utils.GetMetaResponse("invalid-token")
		return
	}

	//connect to string builder
	dataStringBuilder, getDataStringBuilderErr := stringbuilder.ReverseQR(req.QRContent)
	if getDataStringBuilderErr != nil {
		res.Meta = utils.GetMetaResponse("qr.not.match")
		return
	}

	callPlanMerchantDB, errDB := postgres.GetCallPlanMerchantsByCallPlanMerchantId(req.CallPlanMerchantId)
	if errDB != nil {
		res.Meta = utils.GetMetaResponse("default")
		return
	}

	if dataStringBuilder.Tag00 == "01" {
		if dataStringBuilder.Tag6207 != "" && dataStringBuilder.Tag2601 != "" {
			//get phone number by API ROSE
			dataRose, err := rose.InquiryMerchant(dataStringBuilder.Tag6207, dataStringBuilder.Tag2601)
			if err != nil {
				res.Meta = utils.GetMetaResponse("qr.not.match")
				return
			}

			if dataRose.PhoneNumber == callPlanMerchantDB.MerchantPhone {
				//UPDATE LONG LAT CALL PLAN MERCHANT
				if postgres.UpdateLongLatCallPlanMerchant(req.CallPlanMerchantId, req.Longitude, req.Latitude) != nil {
					res.Meta = utils.GetMetaResponse("qr.not.match")
					return
				}

				//GET DATA ACTION LIST
				callPlanListActionDB, errDB := postgres.GetCallPlanListActionsAllWithError(req.CallPlanMerchantId)
				if errDB != nil {
					res.Meta = utils.GetMetaResponse("qr.not.match")
					return
				}

				res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)
				res.Data = callPlanListActionDB
				return
			}

			res.Meta = utils.GetMetaResponse("qr.not.match")
			return
		}
		res.Meta = utils.GetMetaResponse("qr.not.match")
		return
	}

	res.Meta = utils.GetMetaResponse("qr.not.match")
	return

}
