package services

import (
	"encoding/json"
	"log"
	ottologger "ottodigital.id/library/logger/v2"
	"ottosfa-api-apk/models"
	"testing"
)

func TestService_CallPlanTodolistList(t *testing.T) {
	var ottolog ottologger.OttologInterface
	var res *models.Response
	status := []string{}
	status = append(status, "Open")
	status = append(status, "Late")
	req := models.CallPlanTodolistListReq{
		MerchantPhone: "0878850929821",
		Status:        status,
		Page:          1,
		Limit:         25,
	}
	go InitiateService(ottolog).CallPlanTodolistList("HVzwlFFaSfdfuJhmOEqFNLxkoqZQlRSX", req, res)

	byteRes, _ := json.Marshal(res)
	log.Println("res --> ", string(byteRes))
}
