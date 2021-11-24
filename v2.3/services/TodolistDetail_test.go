package services

import (
	"encoding/json"
	"log"
	ottologger "ottodigital.id/library/logger/v2"
	"ottosfa-api-apk/models"
	"testing"
)

func TestService_TodolistDetail(t *testing.T) {
	var ottolog ottologger.OttologInterface
	var res *models.Response
	req := models.TodolistDetailReq{
		MerchantPhone: "089898988897",
		CustomerCode:  "543543422556.0",
		TodolistId:    "213",
	}
	go InitiateService(ottolog).TodolistDetail("HVzwlFFaSfdfuJhmOEqFNLxkoqZQlRSX", req, res)

	byteRes, _ := json.Marshal(res)
	log.Println("res --> ", string(byteRes))
}
