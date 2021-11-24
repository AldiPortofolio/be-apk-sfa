package services

import (
	"encoding/json"
	"log"
	ottologger "ottodigital.id/library/logger/v2"
	"ottosfa-api-apk/models"
	"testing"
)

func TestService_IDMTodolistDetail(t *testing.T) {
	var ottolog ottologger.OttologInterface
	var res *models.Response
	req := models.TodolistDetailReq{
		TodolistId:         "213",
		TodolistCategoryId: 6,
	}
	go InitiateService(ottolog).IDMTodolistDetail("HVzwlFFaSfdfuJhmOEqFNLxkoqZQlRSX", req, res)

	byteRes, _ := json.Marshal(res)
	log.Println("res --> ", string(byteRes))
}
