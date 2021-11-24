package services

import (
	"encoding/json"
	"log"
	ottologger "ottodigital.id/library/logger/v2"
	"ottosfa-api-apk/models"
	"testing"
)

func TestService_CallPlanVisitCheckQRIS(t *testing.T) {
	var ottolog ottologger.OttologInterface
	var res *models.Response
	req := models.CheckQRISReq{
		QRContent: "00020101021126640015ID.OTTOCASH.WWW01189360081101000164420212OP1A015518050303UMI51380014ID.CO.QRIS.WWW02099999999990303UMI5204999953033605802ID5922Toko Testing Diana Dev6007Jakarta610512950623905121145001203199612OP1A015518050703A0163040A7C",
	}
	go InitiateService(ottolog).CallPlanVisitCheckQRIS("HVzwlFFaSfdfuJhmOEqFNLxkoqZQlRSX", req, res)

	byteRes, _ := json.Marshal(res)
	log.Println("res --> ", string(byteRes))
}
