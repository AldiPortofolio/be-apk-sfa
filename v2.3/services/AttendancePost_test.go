package services

import (
	"encoding/json"
	"log"
	ottologger "ottodigital.id/library/logger/v2"
	"ottosfa-api-apk/models"
	"testing"
)

func TestService_AttendancePost(t *testing.T) {
	var ottolog ottologger.OttologInterface
	var res *models.Response
	req := models.AttendancePostReq{
		SalesID:                0,
		SalesPhone:             "",
		SalesName:              "",
		AttendanceCategory:     "",
		AttendanceCategoryType: "",
		TypeAttendance:         "",
		Notes:                  "",
		PhotoSelfie:            "",
		Location:               "",
		Longitude:              "",
		Latitude:               "",
		Time:                   "",
		TypeTimezone:           "",
		AccurationPhoto:        "",
	}
	go InitiateService(ottolog).AttendancePost("HVzwlFFaSfdfuJhmOEqFNLxkoqZQlRSX", req, res)

	byteRes, _ := json.Marshal(res)
	log.Println("res --> ", string(byteRes))
}
