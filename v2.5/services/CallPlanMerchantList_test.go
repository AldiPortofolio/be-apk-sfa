package services

import (
	"encoding/json"
	"log"
	ottologger "ottodigital.id/library/logger/v2"
	"ottosfa-api-apk/models"
	"testing"
)

func TestService_CallPlanMerchantList(t *testing.T) {
	var ottolog ottologger.OttologInterface
	var res *models.Response
	req := models.CallPlanMerchantListReq{
		CallPlanId: 14,
		Status:     "Incompleted",
		Limit:      1,
		Page:       25,
	}
	go InitiateService(ottolog).CallPlanMerchantList("geAgZLuqMRYBOszooflPvcVgdFPhsbry", req, res)

	byteRes, _ := json.Marshal(res)
	log.Println("res --> ", string(byteRes))
}
