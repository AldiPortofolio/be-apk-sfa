package services

import (
	"encoding/json"
	"log"
	ottologger "ottodigital.id/library/logger/v2"
	"ottosfa-api-apk/models"
	"testing"
)

func TestService_Login(t *testing.T) {
	var ottolog ottologger.OttologInterface
	var res *models.Response
	req := models.LoginReq{
		PhoneNumber: "089634679074",
		Pin:         "123123",
		DeviceID:    "44375834754",
		DeviceToken: "dsfjk842374j3y",
		SalesID:     "282",
		VersionCode: "",
		Role:        "sfa",
	}
	go InitiateService(ottolog).Login(req, res)

	byteRes, _ := json.Marshal(res)
	log.Println("res --> ", string(byteRes))
}
