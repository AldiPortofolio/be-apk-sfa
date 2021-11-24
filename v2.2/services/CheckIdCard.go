package services

import (
	"fmt"
	"ottosfa-api-apk/database/postgres"
	"ottosfa-api-apk/models"
	redis "ottosfa-api-apk/redis"
	"ottosfa-api-apk/utils"
	"strconv"
)

// CheckIdCard ..
func (svc *Service) CheckIdCard(bearer string, req models.CheckIdCardReq, res *models.Response) {
	fmt.Println(">>> CheckIdCard - Service <<<")

	//log := svc.OttoLog

	// if qr content is empty, do not continue
	if req.IdCard == "" {
		res.Meta = utils.GetMetaResponse("default")
		return
	}

	token := utils.DecodeBearer(bearer)
	_, validateToken := redis.GetRedisKey(utils.RedisKeyAuth + token)
	if validateToken != nil {
		res.Meta = utils.GetMetaResponse("invalid-token")
		return
	}

	merchantDB, errDB := postgres.GetMerchantByIDCard(req.IdCard)
	if errDB != nil {
		res.Meta = utils.GetMetaResponse("default")
		return
	}

	dataDB, errDB := postgres.GetNumDuplicateIdCard()
	if errDB != nil {
		res.Meta = utils.GetMetaResponse("default")
		return
	}

	//sumDataMerchant := len(merchantDB)
	countIdCard,_ := strconv.Atoi(merchantDB.NumIdCard)
	maxDuplicate,_ := strconv.Atoi(dataDB.ParamValue)

	//fmt.Println(sumDataMerchant)
	fmt.Println(maxDuplicate)

	if (countIdCard < maxDuplicate){
		res.Meta = models.MetaData{
			Status:  true,
			Code:    200,
			//Message: "Nomor KTP telah digunakan sebelumnya sebanyak " + strconv.Itoa(len(merchantDB)),
			Message: "Nomor KTP telah digunakan sebelumnya sebanyak " + merchantDB.NumIdCard,
		}
		return
	}

	res.Meta = utils.GetMetaResponse("idcard_exceeded")
	return
}

