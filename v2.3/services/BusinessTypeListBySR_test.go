package services

import (
	"encoding/json"
	"log"
	"ottosfa-api-apk/models"
	"testing"

	ottologger "ottodigital.id/library/logger/v2"
)

func TestService_BusinessTypeListBySR(t *testing.T) {
	var ottolog ottologger.OttologInterface
	var res *models.Response
	req := models.BusinessTypeListBySRReq{
		SalesTypeId: []int{1},
	}

	go InitiateService(ottolog).BusinessTypeListBySR("HVzwlFFaSfdfuJhmOEqFNLxkoqZQlRSX", req, res)

	byteRes, _ := json.Marshal(res)
	log.Println("res --> ", string(byteRes))
}
