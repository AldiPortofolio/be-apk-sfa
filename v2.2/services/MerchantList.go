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
)

// MerchantList ..
func (svc *Service) MerchantList(bearer string, req models.MerchantListReq, res *models.Response) {
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

	dataPositionDB, errDB := postgres.CheckPositions(sales.ID)
	if errDB != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseInvalidToken)
		return
	}

	//Get Merchant List
	data := []dbmodels.Merchant{}
	for _, val := range *dataPositionDB {
		switch val.SalesRoleId {
		case 1:
			fmt.Println("==REGION==")
			merchant, err := postgres.GetMerchantListRegionWithFilter(req, sales.ID)
			if err != nil {
				res.Meta = utils.GetMetaResponse("merchant.not.found")
				return
			}
			for _, val := range merchant {
				data = append(data, val)
			}
			break
		case 2:
			fmt.Println("==BRANCH==")
			merchant, err := postgres.GetMerchantListBranchWithFilter(req, sales.ID)
			if err != nil {
				res.Meta = utils.GetMetaResponse("merchant.not.found")
				return
			}
			for _, val := range merchant {
				data = append(data, val)
			}
			break
		case 3:
			fmt.Println("==AREA==")
			merchant, err := postgres.GetMerchantListAreaWithFilter(req, sales.ID)
			if err != nil {
				res.Meta = utils.GetMetaResponse("merchant.not.found")
				return
			}
			for _, val := range merchant {
				data = append(data, val)
			}
			break
		case 4:
			fmt.Println("==SUBAREA==")
			merchant, err := postgres.GetMerchantListSubareaWithFilter(req, sales.ID)
			if err != nil {
				res.Meta = utils.GetMetaResponse("merchant.not.found")
				return
			}
			for _, val := range merchant {
				data = append(data, val)
			}
			break
		}
	}

	merchantList := []models.MerchantListRes{}
	for _, val := range data {
		merchantType := "OP"
		if val.InstitutionId == "PGMI" {
			merchantType = "PGMI"
		}

		a := models.MerchantListRes{
			ID:               val.ID,
			MerchantId:       val.MerchantId,
			Name:             val.Name,
			PhoneNumber:      val.PhoneNumber,
			ImageMerchant:    val.ImageMerchant,
			Address:          val.Address,
			JoinAt:           jodaTime.Format("yyyy-MM-dd HH:mm:ss", val.CreatedAt),
			MerchantCategory: merchantType,
			MerchantStatus:   val.Status,
		}
		merchantList = append(merchantList, a)
	}

	res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)
	res.Data = uniqueMerchantList(merchantList)
	return
}

func uniqueMerchantList(intSlice []models.MerchantListRes) []models.MerchantListRes {
	keys := make(map[models.MerchantListRes]bool)
	list := []models.MerchantListRes{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
