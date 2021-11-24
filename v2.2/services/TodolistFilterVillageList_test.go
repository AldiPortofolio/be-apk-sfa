package services

import (
	"encoding/json"
	"log"
	ottologger "ottodigital.id/library/logger/v2"
	"ottosfa-api-apk/models"
	"testing"
)

func TestService_TodolistFilterVillageList(t *testing.T) {
	var ottolog ottologger.OttologInterface
	var res *models.Response
	req := models.TodolistFilterVillageListReq{
		Keyword: "",
	}
	go InitiateService(ottolog).TodolistFilterVillageList("HVzwlFFaSfdfuJhmOEqFNLxkoqZQlRSX", req, res)

	byteRes, _ := json.Marshal(res)
	log.Println("res --> ", string(byteRes))
}
