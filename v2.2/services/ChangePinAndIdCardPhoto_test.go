package services

import (
	"encoding/json"
	"log"
	ottologger "ottodigital.id/library/logger/v2"
	"ottosfa-api-apk/models"
	"testing"
)

func TestService_ChangePinAndIdCardPhoto(t *testing.T) {
	var ottolog ottologger.OttologInterface
	var res *models.Response
	req := models.ChangePhotoPinKTPReq{
		NewPin: "",
		IdCard: "",
		Photo:  "",
	}
	go InitiateService(ottolog).ChangePinAndIdCardPhoto("geAgZLuqMRYBOszooflPvcVgdFPhsbry", req, res)

	byteRes, _ := json.Marshal(res)
	log.Println("res --> ", string(byteRes))
}
