package services

import (
	ottologger "ottodigital.id/library/logger/v2"
	"ottosfa-api-apk/database/dbmodels"
	"ottosfa-api-apk/database/postgres"
	"ottosfa-api-apk/hosts/fds"
	"ottosfa-api-apk/models"
	"ottosfa-api-apk/models/miniomodels"
	"ottosfa-api-apk/utils"
	"ottosfa-api-apk/hosts/indomarco"
)

// Service ..
type Service struct {
	General            models.GeneralModel
	OttoLog            ottologger.OttologInterface
	MerchantRepository *postgres.MerchantRepository
	SendMinio          func(imageBase64 string, nameFile string, spanID string) (miniomodels.UploadRes, error)
	SendFDS            func(msgReq interface{}, typeTrans string, spanID string) ([]byte, error)
	SendIDM            func(msgReq interface{}, typeTrans string, spanID string) ([]byte, error)
}

// ServiceInterface ..
type ServiceInterface interface {
	//merchants
	CheckQRIS(string, models.CheckQRISReq, *models.Response)
	CheckIdCard(string, models.CheckIdCardReq, *models.Response)
	ReverseQR(string, models.CheckQRISReq, *models.Response)
	BusinessTypeList(string, *models.Response)
	AcquitisionsDropdownList(string, *models.Response)
	UpdateMerchant(models.UpdateMerchantReq, *models.Response)
	IndomarcoCheckMerchantOttopay(models.CheckMerchantOttopayReq, *models.Response)

	//To Rose
	AcquisitionsMerchantDetail(models.AcquisitionsMerchantDetailReq, *models.Response)

	//sales
	Profile(string, *models.Response)

	//sfa
	MerchantList(string, models.MerchantListReq, *models.Response)
	CallPlanList(string, models.CallPlanListReq, *models.Response)
	CallPlanMerchantList(string, models.CallPlanMerchantListReq, *models.Response)
	CallPlanTodolistList(string, models.CallPlanTodolistListReq, *models.Response)
	CallPlanTodolistMerchantNotFound(string, models.CallPlanTodolistMerchantNotFoundReq, *models.Response)
	CallPlanActionMerchantList(string, *models.Response)
	CallPlanProductMerchantList(string, *models.Response)
	CallPlanDescriptionMerchant(string, models.CallPlanDescriptionMerchantReq, *models.Response)
	CallPlanVisitCheckMerchantPhone(string, models.CallPlanVisitCheckMerchantPhoneReq, *models.Response)
	CallPlanVisitCheckQRIS(string, models.CheckQRISReq, *models.Response)
	CallPlanVisitAdd(string, models.CallPlanVisitAddReq, *models.Response)
	CallPlanDetail(string, models.CallPlanDetailReq, *models.Response)
	CallPlanActionCheckMerchantPhone(string, models.CallPlanActionCheckMerchantPhoneReq, *models.Response)
	CallPlanActionCheckQRIS(string, models.CallPlanActionCheckQRISReq, *models.Response)
	CallPlanActionUpdateClockInMerchant(string, models.CallPlanActionUpdateClockInMerchantReq, *models.Response)
	CallPlanActionMerchantUnknown(string, models.CallPlanActionMerchantUnknownReq, *models.Response)
	CallPlanActionDetail(string, models.CallPlanActionDetailReq, *models.Response)
	CallPlanActionAddOrEdit(string, dbmodels.CallPlanActions, *models.Response)
	CallPlanActionSubmit(string, models.CallPlanActionSubmitReq, *models.Response)
	CallPlanActionDelete(string, models.CallPlanActionDeleteReq, *models.Response)

	//update QRIS
	MerchantDetailQRIS(string, models.MerchantDetailQRISReq, *models.Response)
	ScanAndUpdateQRIS(string, models.ScanAndUpdateQRISReq, *models.Response)

	//toko Ottopay
	MerchantDetailTokoOttopay(string, models.MerchantDetailQRISReq, *models.Response)

	//todolist
	TodolistPost(string, models.TodolistPostReq, *models.Response) //tidak dipakai
	TodolistFilterVillageList(string, models.TodolistFilterVillageListReq, *models.Response)
	TodolistCount(string, *models.Response)

	//todolist rose
	TodolistTaskBySubCategory(string, models.TodolistTaskBySubCategoryReq, *models.Response)

	//indomarco
	CheckIndomarcoAccount(string, models.CheckIndomarcoAccountReq, *models.Response)

	LoginIndomarco(models.IndomarcoLoginReq, *models.Response)
	Login(models.LoginReq, *models.Response)
	ChangePinAndIdCardPhoto(string, models.ChangePhotoPinKTPReq, *models.Response)
	AttendancePost(string, models.AttendancePostReq, *models.Response)
	UploadPhotoProfilSales(string, models.UploadPhotoProfilSalesReq, *models.Response)
	UserCategory(string, *models.Response)
	ClearSession(*models.Response)
	PushNotif(models.PushNotifReq, *models.Response)
	MerchantNotFoundList(string, *models.Response)

	Province(string, *models.Response)
	City(string, *models.Response)
	District(string, *models.Response)
	Village(string, *models.Response)
	ChangeStatusMerchant(models.ChangeStatusMerchantReq, *models.Response)

	ReportBySales(string, models.ReportBySalesReq, *models.Response)
	HistorySummary(string, models.ReportBySalesReq, *models.Response)
	HistoryDetail(string, models.ReportHistoryDetailReq, *models.Response)
	HealthCheck(*models.Response)
}

// InitiateService ..
func InitiateService(log ottologger.OttologInterface) ServiceInterface {
	return &Service{
		OttoLog:   log,
		SendMinio: utils.UploadImage2Minio,
		SendFDS:   fds.Send,
		SendIDM:   indomarco.Send,
	}
}
