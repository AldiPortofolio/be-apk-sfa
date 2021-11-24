package services

import (
	"fmt"
	"ottosfa-api-apk/hosts/rose"
	"ottosfa-api-apk/models"
	redis "ottosfa-api-apk/redis"
	"ottosfa-api-apk/utils"
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

	dataRose, errRose := rose.CheckIdCard(req.IdCard)
	if errRose != nil {
		res.Meta = utils.GetMetaResponse("default")
		return
	}

	if dataRose.Rc == "01"{
		res.Meta = models.MetaData{
			Status:  false,
			Code:    422,
			Message: dataRose.Msg,
		}
		return
	}

	res.Meta = utils.GetMetaResponse("success")
	return
}

