package services

import (
	"encoding/json"
	"log"
	ottologger "ottodigital.id/library/logger/v2"
	"ottosfa-api-apk/models"
	"testing"
)

func TestService_TodolistList(t *testing.T) {
	var ottolog ottologger.OttologInterface
	var res *models.Response
	req := models.TodolistListReq{
		Keyword:       "",
		TaskDateStart: "",
		TaskDateEnd:   "",
		Status:        nil,
		CategoryID:    nil,
		VillageID:     nil,
		ClusterID:     nil,
		Page:          0,
		Limit:         0,
	}
	go InitiateService(ottolog).TodolistList("HVzwlFFaSfdfuJhmOEqFNLxkoqZQlRSX", req, res)

	byteRes, _ := json.Marshal(res)
	log.Println("res --> ", string(byteRes))
}
