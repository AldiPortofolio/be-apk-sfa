package services

import (
	"fmt"
	"ottosfa-api-apk/constants"
	"ottosfa-api-apk/database/postgres"
	"ottosfa-api-apk/models"
	redis "ottosfa-api-apk/redis"
	"ottosfa-api-apk/utils"
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

	err = postgres.AttendancePostWithAccurationPhoto(req)
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
