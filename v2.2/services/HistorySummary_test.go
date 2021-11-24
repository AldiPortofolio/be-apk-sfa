package services

import (
	"encoding/json"
	"log"
	ottologger "ottodigital.id/library/logger/v2"
	"ottosfa-api-apk/models"
	"testing"
)

func TestService_HistorySummary(t *testing.T) {
	var ottolog ottologger.OttologInterface
	var res *models.Response
	req := models.ReportBySalesReq{
		Phone:    "",
		DateFrom: "",
		DateTo:   "",
	}
	go InitiateService(ottolog).HistorySummary("HVzwlFFaSfdfuJhmOEqFNLxkoqZQlRSX", *req, res)

	byteRes, _ := json.Marshal(res)
	log.Println("res --> ", string(byteRes))
}
