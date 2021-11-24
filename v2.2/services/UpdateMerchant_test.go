package services

import (
	"encoding/json"
	"log"
	ottologger "ottodigital.id/library/logger/v2"
	"ottosfa-api-apk/models"
	"testing"
)

func TestService_UpdateMerchant(t *testing.T) {
	var ottolog ottologger.OttologInterface
	var res *models.Response

	req := models.UpdateMerchantReq{
		Name: "testing nihhh",
		MerchantId: "OP1B00030143",
	}
	go InitiateService(ottolog).UpdateMerchant(req, res)

	byteRes, _ := json.Marshal(res)
	log.Println("res --> ", string(byteRes))
}
