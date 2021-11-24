package services

import (
	"encoding/json"
	"log"
	ottologger "ottodigital.id/library/logger/v2"
	"ottosfa-api-apk/models"
	"testing"
)

func TestService_ScanAndUpdateQRIS(t *testing.T) {
	var ottolog ottologger.OttologInterface
	var res *models.Response
	req := models.ScanAndUpdateQRISReq{
		QRContent:     "",
		MerchantId:    "",
		MerchantPhone: "",
		Mpan:          "",
	}
	go InitiateService(ottolog).ScanAndUpdateQRIS("geAgZLuqMRYBOszooflPvcVgdFPhsbry", req, res)

	byteRes, _ := json.Marshal(res)
	log.Println("res --> ", string(byteRes))
}
