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
	TodolistMerchantNotFoundList(string, models.TodolistMerchantNotFoundListReq, *models.Response)
	TodolistList(string, models.TodolistListReq, *models.Response)

	CallPlanMerchantList(string, models.CallPlanMerchantListReq, *models.Response)

	//CallPlan dan Todolist Rose
	CallPlanActionAddOrEdit(string, models.CallPlanActionAddOrEditReq, *models.Response)
	TodolistDetail(string, models.TodolistDetailReq, *models.Response)
	TodolistPost(string, models.TodolistPostV24Req, *models.Response)

	Login(models.LoginReq, *models.Response)
}

// InitiateService ..
func InitiateService(log ottologger.OttologInterface) ServiceInterface {
	return &Service{
		OttoLog:   log,
		SendMinio: utils.UploadImage2Minio,
		SendFDS:   fds.Send,
	}
}
