package services

import (
	"fmt"
	ottoutils "ottodigital.id/library/utils"
	"ottosfa-api-apk/constants"
	"ottosfa-api-apk/database/postgres"
	"ottosfa-api-apk/models"
	redis "ottosfa-api-apk/redis"
	"ottosfa-api-apk/utils"
	"time"
)

// AttendancePost ..
func (svc *Service) AttendancePost(bearer string, req models.AttendancePostReq, res *models.Response) {
	fmt.Println(">>> AttendancePost - Service <<<")

	token := utils.DecodeBearer(bearer)
	_, validateToken := redis.GetRedisKey(utils.RedisKeyAuth + token)
	if validateToken != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseInvalidToken)
		return
	}

	sales, errDB := postgres.CheckToken(bearer[7:])
	if errDB != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseInvalidToken)
		return
	}

	resPhotoSelfie, err := svc.SendMinio(req.PhotoSelfie, "selfie_attendance_"+fmt.Sprintf("%d", sales.ID), svc.General.SpanId)
	if err != nil {
		res.Meta = utils.GetMetaResponse("failed.upload.image")
		return
	}

	req.SalesID = int64(sales.ID)
	req.SalesPhone = sales.PhoneNumber
	req.SalesName = sales.FirstName + " " + sales.LastName
	req.PhotoSelfie = resPhotoSelfie.Url
	req.AttendanceCategoryType = utils.ChangeTypeCategoryAttendanceNameToCode(req.AttendanceCategoryType)

	// get status
	status := 0
	layout := "2006-01-02 15:04:05"
	t,_ := time.Parse(layout, req.Time)
	today := t.Format("2006-01-02")
	if req.AttendanceCategory == "Absen Harian" {
		data, err := postgres.CheckStatusAbsenHarian(req, today)
		if err != nil || len(data) == 0 {
			status = 1
		}
	}

	time1,_ := time.Parse(layout,today + ottoutils.GetEnv("ATTENDANCE_EVENT_TIME_1", " 13:00:00"))
	time2,_ := time.Parse(layout,today + ottoutils.GetEnv("ATTENDANCE_EVENT_TIME_2", " 13:59:59"))
	if req.AttendanceCategory == "Event" {
		data, err := postgres.CheckStatusEvent(req, today)
		if err != nil || len(data) == 0 {
			fmt.Println("t: ", t)
			fmt.Println("time1: ", time1)
			fmt.Println("time2: ", time2)
			if t.Unix() > time1.Unix() && t.Unix() < time2.Unix() {
				status = 1
			}
		}
	}

	fmt.Println("status: ", status)
	err = postgres.AttendancePostWithAccurationPhotoV23(req, status, t)
	if err != nil {
		res.Meta = utils.GetMetaResponse("attendande.failed.update.db")
		return
	}

	res.Data = []models.AttendancePostReq{}
	res.Meta = models.MetaData{
		Status:  true,
		Code:    200,
		Message: "Kamu Berhasil Mengisi Absensi " + req.AttendanceCategory,
	}
	return
}
