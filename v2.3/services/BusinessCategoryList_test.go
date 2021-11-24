package services

import (
	"encoding/json"
	"log"
	"ottosfa-api-apk/models"
	"testing"

	ottologger "ottodigital.id/library/logger/v2"
)

func TestService_BusinessCategoryList(t *testing.T) {
	var ottolog ottologger.OttologInterface
	var res *models.Response

	go InitiateService(ottolog).BusinessCategoryList("HVzwlFFaSfdfuJhmOEqFNLxkoqZQlRSX", res)

	byteRes, _ := json.Marshal(res)
	log.Println("res --> ", string(byteRes))
}
