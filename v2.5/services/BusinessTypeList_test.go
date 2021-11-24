package services

import (
	"encoding/json"
	"log"
	"ottosfa-api-apk/models"
	"testing"

	ottologger "ottodigital.id/library/logger/v2"
)

func TestService_BusinessTypeList(t *testing.T) {
	var ottolog ottologger.OttologInterface
	var res *models.Response
	req := models.GetAcquisitionByName{
		Name: "",
		RoseMerchantGroup: "",
		RoseMerchantCategory: "",
		SalesRetailId: "",
	}
	go InitiateService(ottolog).BusinessTypeListByName(req, res)

	byteRes, _ := json.Marshal(res)
	log.Println("res --> ", string(byteRes))
}
