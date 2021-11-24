package services

import (
	"encoding/json"
	"log"
	ottologger "ottodigital.id/library/logger/v2"
	"ottosfa-api-apk/models"
	"testing"
)

func TestService_Province(t *testing.T) {
	var ottolog ottologger.OttologInterface
	var res *models.Response

	go InitiateService(ottolog).Province("1", res)

	byteRes, _ := json.Marshal(res)
	log.Println("res --> ", string(byteRes))
}
