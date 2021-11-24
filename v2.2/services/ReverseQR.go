package services

import (
	"fmt"
	"ottosfa-api-apk/constants"
	"ottosfa-api-apk/hosts/stringbuilder"
	"ottosfa-api-apk/models"
	//"ottosfa-api-apk/rediscluster"
	redis "ottosfa-api-apk/redis"
	"ottosfa-api-apk/utils"
	"strings"
)

// ReverseQR ..
func (svc *Service) ReverseQR(bearer string, req models.CheckQRISReq, res *models.Response) {
	fmt.Println(">>> ReverseQR - Service <<<")

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

	//connect to string builder
	dataStringBuilder, getDataStringBuilderErr := stringbuilder.ReverseQR(req.QRContent)
	if getDataStringBuilderErr != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseDefault)
		return
	}

	if dataStringBuilder.Tag00 == "01" && dataStringBuilder.Tag5102 != "" {
		if dataStringBuilder.Tag5102[0:2] == "ID" && !strings.HasPrefix(dataStringBuilder.Tag2601, "93600811") { //dataStringBuilder.Tag2601[0:9] == "93600811"{
			go redis.SaveRedis("SFA:QR:"+dataStringBuilder.Tag5102, req.QRContent)
			//redisKeyExpQR := ottoutils.GetEnv("REDIS_KEY_EXP_QR", "1")
			//go redis.SaveRedisExp("SFA:QR:"+dataStringBuilder.Tag5102, redisKeyExpQR, req.QRContent)
			res.Meta = utils.GetMetaResponse("qr.success")
		} else {
			res.Meta = utils.GetMetaResponse("qr.unknown")
		}
	} else {
		res.Meta = utils.GetMetaResponse("qr.unknown")
	}

	data := models.ReverseQRRes{
		NMID: dataStringBuilder.Tag5102,
	}

	//res.Meta = utils.GetMetaResponse("qr.success")
	res.Data = data
	return
}
