package services

import (
	"encoding/json"
	"log"
	ottologger "ottodigital.id/library/logger/v2"
	"ottosfa-api-apk/models"
	"testing"
)

func TestService_TodolistPost(t *testing.T) {
	var ottolog ottologger.OttologInterface
	var res *models.Response
	req := models.TodolistPostV24Req{
		TaskID:              nil,
		Label:               nil,
		ContentType:         nil,
		Body:                nil,
		LabelPhotoMerchant1: "",
		LabelPhotoMerchant2: "",
		LabelPhotoMerchant3: "",
		PhotoMerchant1:      "",
		PhotoMerchant2:      "",
		PhotoMerchant3:      "",
		NewTaskDate:         "",
		Reason:              "",
		TodolistID:          0,
		OldTaskDate:         "",
		Village:             "",
		District:            "",
		City:                "",
		Province:            "",
		Longitude:           "",
		Latitude:            "",
		Alamat:              "",
		Patokan:             "",
		Status:              "",
		Long:                "",
		Lat:                 "",
		VersionID:           0,
	}
	go InitiateService(ottolog).TodolistPost("HVzwlFFaSfdfuJhmOEqFNLxkoqZQlRSX", req, res)

	byteRes, _ := json.Marshal(res)
	log.Println("res --> ", string(byteRes))
}
