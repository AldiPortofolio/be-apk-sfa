package services

import (
	"encoding/json"
	"log"
	"ottosfa-api-apk/models"
	"testing"

	ottologger "ottodigital.id/library/logger/v2"
)

func TestService_AcquisitionList(t *testing.T) {
	var ottolog ottologger.OttologInterface
	var res *models.Response
	req := models.GetAcquisitionBySalesTypeID{
		SalesTypeId: "",
	}
	go InitiateService(ottolog).AcquisitionListBySR("HVzwlFFaSfdfuJhmOEqFNLxkoqZQlRSX", req, res)

	byteRes, _ := json.Marshal(res)
	log.Println("res --> ", string(byteRes))
}
