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

// UploadPhotoProfilSales ..
func (svc *Service) UploadPhotoProfilSales(bearer string, req models.UploadPhotoProfilSalesReq, res *models.Response) {
	fmt.Println(">>> UploadPhotoProfilSales - Service <<<")

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

	// BEGIN upload photo profil to minio
	dataPhotoProfilMinio, errMinio := svc.SendMinio(req.PhotoProfil, "PhotoProfil_"+strings.ReplaceAll(sales.FirstName, " ", "_")+"-"+strconv.Itoa(sales.ID), svc.General.SpanId)
	if errMinio != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseDefault)
		return
	}
	// END upload photo profil to minio

	changedAttr := map[string]interface{}{}
	changedAttr["photo"] = dataPhotoProfilMinio.Url

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

	res.Meta = utils.GetMetaResponse("success")

	return
}
