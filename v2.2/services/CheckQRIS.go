package services

import (
	"fmt"
	"ottosfa-api-apk/constants"
	"ottosfa-api-apk/hosts/stringbuilder"
	"ottosfa-api-apk/models"
	redis "ottosfa-api-apk/redis"
	"ottosfa-api-apk/utils"
)

// CheckQRIS ..
func (svc *Service) CheckQRIS(bearer string, req models.CheckQRISReq, res *models.Response) {
	fmt.Println(">>> CheckQRIS - Service <<<")

	//log := svc.OttoLog

	// if qr content is empty, do not continue
	if req.QRContent == "" {
		res.Meta = utils.GetMetaResponse("default")
		return
	}

	token := utils.DecodeBearer(bearer)
	_, validateToken := redis.GetRedisKey(utils.RedisKeyAuth + token)
	if validateToken != nil {
		res.Meta = utils.GetMetaResponse("invalid-token")
		return
	}

	dataRes := models.CheckQRISRes{}

	//connect to string builder
	dataStringBuilder, getDataStringBuilderErr := stringbuilder.ReverseQR(req.QRContent)
	if getDataStringBuilderErr != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseDefault)
		return
	}

	if dataStringBuilder.Tag00 == "01" && dataStringBuilder.Tag5102 != "" {
		if dataStringBuilder.Tag2601 != "" && dataStringBuilder.Tag2602 != "" && dataStringBuilder.Tag5102[0:2] == "ID" {
			//go redis.SaveRedis("SFA:QR:"+ dataStringBuilder.Tag5102, req.QRContent)
			//check MID to DB SFA
			if _, ok, _ := svc.MerchantRepository.CheckMerchantID(dataStringBuilder.Tag2602); ok {
				res.Meta = utils.GetMetaResponse("qr.exist")
				return
			}
			dataRes.MPAN = dataStringBuilder.Tag2601
			dataRes.MID = dataStringBuilder.Tag2602
			dataRes.NMID = dataStringBuilder.Tag5102
			dataRes.StoreNamePreprinted = dataStringBuilder.Tag59

		} else {
			res.Meta = utils.GetMetaResponse("qr.unknown")
			return
		}
	} else {
		res.Meta = utils.GetMetaResponse("qr.unknown")
		return
	}

	res.Meta = utils.GetMetaResponse("qr.success")
	res.Data = dataRes
	return
}
