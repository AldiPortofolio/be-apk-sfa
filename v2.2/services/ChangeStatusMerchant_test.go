package services

import (
	"encoding/json"
	"log"
	ottologger "ottodigital.id/library/logger/v2"
	"ottosfa-api-apk/models"
	"testing"
)

func TestService_ChangeStatusMerchant(t *testing.T) {
	var ottolog ottologger.OttologInterface
	var res *models.Response
	req := models.ChangeStatusMerchantReq{
		Name:         "",
		MID:          "",
		PhoneNumber:  "",
		Status:       "",
		MerchantType: "",
	}
	go InitiateService(ottolog).ChangeStatusMerchant(req, res)

	byteRes, _ := json.Marshal(res)
	log.Println("res --> ", string(byteRes))
}
