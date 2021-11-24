package services

import (
	"fmt"
	"ottosfa-api-apk/constants"
	"ottosfa-api-apk/database/dbmodels"
	"ottosfa-api-apk/database/postgres"
	"ottosfa-api-apk/models"
	redis "ottosfa-api-apk/redis"
	"ottosfa-api-apk/utils"
	"strconv"
	"strings"
)

// ChangePinAndIdCardPhoto ..
func (svc *Service) ChangePinAndIdCardPhoto(bearer string, req models.ChangePhotoPinKTPReq, res *models.Response) {
	fmt.Println(">>> ChangePinAndIdCardPhoto - Service <<<")

	token := utils.DecodeBearer(bearer)
	_, validateToken := redis.GetRedisKey(utils.RedisKeyAuth + token)
	if validateToken != nil {
		res.Meta = utils.GetMetaResponse("invalid-token")
		return
	}

	//cek data sales (get data salesId by Token)
	sales, errDB := postgres.CheckToken(bearer[7:])
	if errDB != nil {
		res.Meta = utils.GetMetaResponse("invalid-token")
		return
	}

	// hashing password
	newPin, _ := utils.HashPassword(req.NewPin)

	// BEGIN upload photo profil to minio
	dataPhotoProfilMinio, errMinio := svc.SendMinio(req.Photo, "PhotoProfil_"+strings.Replace(sales.FirstName, " ", "_", -1)+"-"+strconv.Itoa(sales.ID), svc.General.SpanId)
	if errMinio != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseDefault)
		return
	}
	// END upload photo profil to minio

	// BEGIN upload photo ktp to minio
	dataPhotoKTPMinio, errMinio := svc.SendMinio(req.IdCard, "KTP_"+strings.Replace(sales.FirstName, " ", "_", -1)+"-"+strconv.Itoa(sales.ID), svc.General.SpanId)
	if errMinio != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseDefault)
		return
	}
	// END upload photo ktp to minio

	changedAttr := map[string]interface{}{}
	changedAttr["photo"] = dataPhotoProfilMinio.Url
	changedAttr["id_card"] = dataPhotoKTPMinio.Url

	reqDB := dbmodels.Requests{
		RequestType:    3,
		Module:         0,
		Status:         0,
		ApprovableType: "Salesman",
		ApprovableId:   sales.ID,
		//MakerId:        sales.ID,
		ToBeChanged: changedAttr,
	}

	if postgres.UploadToTableRequests(reqDB) != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseDefault)
		return
	}

	//update to db
	errDB = postgres.UpdatePinSales(newPin, dataPhotoKTPMinio.Url, dataPhotoProfilMinio.Url, sales.ID)
	if errDB != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseDefault)
		return
	}

	res.Meta = utils.GetMetaResponse("success")
	res.Data = models.ChangePhotoPinKTPRes{
		ResponseCode:    "00",
		DescriptionCode: utils.GetMetaResponse("success.change.pin.idcard").Message,
	}

	return
}
