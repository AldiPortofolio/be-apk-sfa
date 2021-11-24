package services

import (
	"encoding/json"
	"log"
	ottologger "ottodigital.id/library/logger/v2"
	"ottosfa-api-apk/models"
	"testing"
)

func TestService_CallPlanActionCheckMerchantPhone(t *testing.T) {
	var ottolog ottologger.OttologInterface
	var res *models.Response
	req := models.CallPlanActionCheckMerchantPhoneReq{
		CallPlanMerchantId: 37,
		MerchantPhone:      "0878850929821",
		Longitude:          "7.54321",
		Latitude:           "-7.54321",
	}
	go InitiateService(ottolog).CallPlanActionCheckMerchantPhone("HVzwlFFaSfdfuJhmOEqFNLxkoqZQlRSX", req, res)

	byteRes, _ := json.Marshal(res)
	log.Println("res --> ", string(byteRes))
}
