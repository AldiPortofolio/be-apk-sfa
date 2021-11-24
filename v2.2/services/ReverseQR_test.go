package services

import (
	"encoding/json"
	"log"
	ottologger "ottodigital.id/library/logger/v2"
	"ottosfa-api-apk/models"
	"testing"
)

func TestService_ReverseQR(t *testing.T) {
	var ottolog ottologger.OttologInterface
	var res *models.Response
	req := models.CheckQRISReq{
		QRContent: "00020101021126610014COM.GO-JEK.WWW01189360091435456007810210G5456007810303UMI51440014ID.CO.QRIS.WWW0215ID10190000023280303UMI5204581253033605802ID5916Kantin Ibu Lilik6013Jakarta Pusat61051031062070703A0163044C6B",
	}
	go InitiateService(ottolog).ReverseQR("HVzwlFFaSfdfuJhmOEqFNLxkoqZQlRSX", req, res)

	byteRes, _ := json.Marshal(res)
	log.Println("res --> ", string(byteRes))
}
