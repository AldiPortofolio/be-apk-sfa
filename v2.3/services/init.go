package services

import (
	"ottosfa-api-apk/database/postgres"
	"ottosfa-api-apk/hosts/fds"
	"ottosfa-api-apk/models"
	"ottosfa-api-apk/models/miniomodels"
	"ottosfa-api-apk/utils"

	ottologger "ottodigital.id/library/logger/v2"
)

// Service ..
type Service struct {
	General            models.GeneralModel
	OttoLog            ottologger.OttologInterface
	MerchantRepository *postgres.MerchantRepository
	SendMinio          func(imageBase64 string, nameFile string, spanID string) (miniomodels.UploadRes, error)
	SendFDS            func(msgReq interface{}, typeTrans string, spanID string) ([]byte, error)
}

// ServiceInterface ..
type ServiceInterface interface {
	//sfa
	CallPlanList(string, models.CallPlanListReq, *models.Response)
	CallPlanActionScanQR(string, models.CallPlanActionScanQRReq, *models.Response)
	CallPlanActionAddOrEdit(string, models.CallPlanActionAddOrEditReq, *models.Response)
	CallPlanActionMerchantUnknown(string, models.CallPlanActionMerchantUnknownReq, *models.Response)
	CallPlanActionSubmit(string, models.CallPlanActionSubmitReq, *models.Response)
	CallPlanMerchantList(string, models.CallPlanMerchantListReq, *models.Response)
	Login(models.LoginReq, *models.Response)
	MerchantList(string, models.MerchantListv23Req, *models.Response)

	//callplan rose
	CallPlanVisitCheckMerchantPhone(string, models.CallPlanVisitCheckMerchantPhoneReq, *models.Response)
	CallPlanVisitCheckQRIS(string, models.CheckQRISReq, *models.Response)
	CallPlanTodolistList(string, models.CallPlanTodolistListReq, *models.Response)

	//indomarco
	IDMTodolistList(string, models.TodolistListReq, *models.Response)
	IDMTodolistDetail(string, models.TodolistDetailReq, *models.Response)

	//merchants
	CheckIdCard(string, models.CheckIdCardReq, *models.Response)

	//toko Ottopay
	MerchantDetailTokoOttopay(string, models.MerchantDetailQRISReq, *models.Response)

	TodolistList(string, models.TodolistListReq, *models.Response)
	TodolistDetail(string, models.TodolistDetailReq, *models.Response)
	TodolistPost(string, models.TodolistPostV23Req, *models.Response)
	TodolistMerchantNotFoundList(string, models.TodolistMerchantNotFoundListReq, *models.Response)
	BusinessCategoryList(string, *models.Response)
	BusinessTypeList(string, *models.Response)
	BusinessTypeListBySR(string, models.BusinessTypeListBySRReq, *models.Response)
	AttendancePost(string, models.AttendancePostReq, *models.Response)
}

// InitiateService ..
func InitiateService(log ottologger.OttologInterface) ServiceInterface {
	return &Service{
		OttoLog:   log,
		SendMinio: utils.UploadImage2Minio,
		SendFDS:   fds.Send,
	}
}
