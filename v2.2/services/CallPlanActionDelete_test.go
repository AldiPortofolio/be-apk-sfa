package services

import (
	"encoding/json"
	"log"
	ottologger "ottodigital.id/library/logger/v2"
	"ottosfa-api-apk/models"
	"testing"
)

func TestService_CallPlanActionDelete(t *testing.T) {
	var ottolog ottologger.OttologInterface
	var res *models.Response
	req := models.CallPlanActionDeleteReq{
		CallPlanActionId: 440,
	}
	go InitiateService(ottolog).CallPlanActionDelete("HVzwlFFaSfdfuJhmOEqFNLxkoqZQlRSX", req, res)

	byteRes, _ := json.Marshal(res)
	log.Println("res --> ", string(byteRes))
}
