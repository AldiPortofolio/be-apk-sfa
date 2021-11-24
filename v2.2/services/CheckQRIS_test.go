package services

import (
	"encoding/json"
	"log"
	ottologger "ottodigital.id/library/logger/v2"
	"ottosfa-api-apk/models"
	"testing"
)

func TestService_CheckQRIS(t *testing.T) {
	var ottolog ottologger.OttologInterface
	var res *models.Response
	req := models.CheckQRISReq{
		QRContent: "00020101021226630014ID.OTTOPAY.WWW01189360081100000000010212OP1A000002720303UMI52045814530336054032005802ID5920toko test qris fadil6015RT.5/RW.4, Kare61042224624301000512OP1A000002729912OP1A000002720703A016304854F",
	}
	go InitiateService(ottolog).CheckQRIS("geAgZLuqMRYBOszooflPvcVgdFPhsbry", req, res)

	byteRes, _ := json.Marshal(res)
	log.Println("res --> ", string(byteRes))
}
