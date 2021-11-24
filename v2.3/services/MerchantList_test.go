package services

import (
	"encoding/json"
	"log"
	ottologger "ottodigital.id/library/logger/v2"
	"ottosfa-api-apk/models"
	"testing"
)

func TestService_MerchantList(t *testing.T) {
	var ottolog ottologger.OttologInterface
	var res *models.Response
	req := models.MerchantListv23Req{
		ProvinceId:       nil,
		CityId:           nil,
		DistrictId:       nil,
		VillageId:        nil,
	}
	go InitiateService(ottolog).MerchantList("geAgZLuqMRYBOszooflPvcVgdFPhsbry", req, res)

	byteRes, _ := json.Marshal(res)
	log.Println("res --> ", string(byteRes))
}
