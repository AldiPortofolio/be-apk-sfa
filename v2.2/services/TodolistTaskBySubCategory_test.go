package services

import (
	"encoding/json"
	"log"
	ottologger "ottodigital.id/library/logger/v2"
	"ottosfa-api-apk/models"
	"testing"
)

func TestService_TodolistTaskBySubCategory(t *testing.T) {
	var ottolog ottologger.OttologInterface
	var res *models.Response
	req := models.TodolistTaskBySubCategoryReq{
		MerchantPhone: "",
		TodolistId:    "",
		SubCategoryId: "",
	}
	go InitiateService(ottolog).TodolistTaskBySubCategory("HVzwlFFaSfdfuJhmOEqFNLxkoqZQlRSX", req, res)

	byteRes, _ := json.Marshal(res)
	log.Println("res --> ", string(byteRes))
}
