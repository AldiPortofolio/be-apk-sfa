package services

import (
	"encoding/json"
	"log"
	ottologger "ottodigital.id/library/logger/v2"
	"ottosfa-api-apk/models"
	"testing"
)

func TestService_CheckIndomarcoAccount(t *testing.T) {
	var ottolog ottologger.OttologInterface
	var res *models.Response
	req := models.CheckIndomarcoAccountReq{
		CustomerCode: "520052130001604.0",
	}
	go InitiateService(ottolog).CheckIndomarcoAccount("geAgZLuqMRYBOszooflPvcVgdFPhsbry", req, res)

	byteRes, _ := json.Marshal(res)
	log.Println("res --> ", string(byteRes))
}
