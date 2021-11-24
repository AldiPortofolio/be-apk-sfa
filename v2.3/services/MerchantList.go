package services

import (
	"fmt"
	"ottosfa-api-apk/constants"
	"ottosfa-api-apk/database/postgres"
	"ottosfa-api-apk/hosts/rose"
	"ottosfa-api-apk/models"
	"ottosfa-api-apk/redis"
	"ottosfa-api-apk/utils"
)

// MerchantList ..
func (svc *Service) MerchantList(bearer string, req models.MerchantListv23Req, res *models.Response) {
	fmt.Println(">>> MerchantList - Service <<<")

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
		res.Meta = utils.GetMetaResponse("merchant.not.found")
		return
	}

	village := []string{}
	for _, val := range villageId {
		village = append(village, val.VillageId)
	}

	dataRose, err := rose.PencapaianSales(sales.PhoneNumber, sales.SalesTypeId, village)
	if err != nil {
		res.Meta = utils.GetMetaResponse("merchant.not.found")
		return
	}

	merchantList := []models.MerchantListv23Res{}
	for _, val := range dataRose.AcquisitionData {
		a := models.MerchantListv23Res{
			ID:               val.ID,
			MerchantId:       val.MerchantID,
			Name:             val.Name,
			PhoneNumber:      val.PhoneNumber,
			ImageMerchant:    val.ImageMerchant,
			Address:          val.Address,
			//JoinAt:           jodaTime.Format("yyyy-MM-dd HH:mm:ss", val.JoinAt),
			MerchantCategory: val.MerchantCategory,
			MerchantStatus:   val.MerchantStatus,
		}
		merchantList = append(merchantList, a)
	}

	res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)
	res.Data = uniqueMerchantList(merchantList)
	return
}

// uniqueMerchantList ..
func uniqueMerchantList(intSlice []models.MerchantListv23Res) []models.MerchantListv23Res {
	keys := make(map[models.MerchantListv23Res]bool)
	list := []models.MerchantListv23Res{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

