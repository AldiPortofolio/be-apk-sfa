package services

import (
	"encoding/json"
	"log"
	ottologger "ottodigital.id/library/logger/v2"
	"ottosfa-api-apk/models"
	"testing"
)

func TestService_Village(t *testing.T) {
	var ottolog ottologger.OttologInterface
	var res *models.Response

	go InitiateService(ottolog).Village("1171020", res)

	byteRes, _ := json.Marshal(res)
	log.Println("res --> ", string(byteRes))
}
