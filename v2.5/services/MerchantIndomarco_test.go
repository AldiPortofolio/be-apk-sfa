package services

import (
	"encoding/json"
	"log"
	"ottosfa-api-apk/models"
	"testing"

	ottologger "ottodigital.id/library/logger/v2"
)

func TestService_MerchantIndomarco(t *testing.T) {
	var ottolog ottologger.OttologInterface
	var res *models.Response
	req := models.UpdateMerchantIndomarcoReq{
		CustomerCode: "520052170070634329.0",
		MerchantID:   "081122330042",
		Phone:        "081122330042",
	}
	go InitiateService(ottolog).UpdateMerchantIndomarco(req, res)

	byteRes, _ := json.Marshal(res)
	log.Println("res --> ", string(byteRes))
}
