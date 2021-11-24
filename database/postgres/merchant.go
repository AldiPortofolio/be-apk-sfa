package postgres

import (
	"fmt"
	"ottosfa-api-apk/database/dbmodels"
	"ottosfa-api-apk/models"

	ottologger "ottodigital.id/library/logger"
)

// ChangeMerchantPhone ..
func ChangeMerchantPhone(phone, merchantId string) error {
	sugarLogger := ottologger.GetLogger()
	merchantId = merchantId[8:]
	var err error
	err = Dbcon.Exec("update merchants set phone_number = ? where merchant_id = ?", phone, merchantId).Error
	if err != nil {
		sugarLogger.Error("Failed connect to database SFA when updating new phone")
		return err
	}
	return nil
}

// GetMerchantByMerchantPhone ..
func GetMerchantByMerchantPhone(merchantPhone string) (models.CallPlanVisitMerchantRes, error) {
	fmt.Println(">>> CallPlanVisitCheckMerchantPhone/CallPlanVisitCheckQRIS/MerchantDetailQRIS - GetMerchantByMerchantPhone - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	res := models.CallPlanVisitMerchantRes{}

	var err error
	//err = Dbcon.Where("merchant_id = ? or phone_number = ?", req.MerchantId, req.MerchantPhone).First(&res).Error
	err = Dbcon.Table("merchants a").
		Select("a.id as id_merchant, a.name merchant_name, a.merchant_id, a.address merchant_address, a.phone_number merchant_phone, a.merchant_type_id, a.mpan mpan, b.name merchant_type_name").
		Joins("LEFT JOIN merchant_types b ON b.id = a.merchant_type_id").
		Where("a.phone_number = ?", merchantPhone).First(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get merchant by merchant phone")
		return res, err
	}
	return res, nil
}

// MerchantByPhoneNumber ..
func MerchantByPhoneNumber(merchantPhone string, dataMerchant dbmodels.MerchantStatus) (dbmodels.MerchantStatus, error) {
	fmt.Println(">>> ChangeStatusMerchant - MerchantByPhoneNumber - Postgres <<<")
	err := Dbcon.Where("phone_number = ? ", merchantPhone).Find(&dataMerchant).Error
	return dataMerchant, err
}

// MerchantByMerchantType ..
func MerchantByMerchantType(merchantType string, dataMerchantType dbmodels.MerchantType) (dbmodels.MerchantType, error) {
	fmt.Println(">>> ChangeStatusMerchant - MerchantByMerchantType - Postgres <<<")
	err := Dbcon.Where("name = ? ", merchantType).Find(&dataMerchantType).Error
	return dataMerchantType, err
}

// UpdateStatusMerchant ..
func UpdateStatusMerchant(dataMerchant dbmodels.MerchantStatus) error {
	fmt.Println(">>> ChangeStatusMerchant - UpdateStatusMerchant - Postgres <<<")
	//return Dbcon.Exec("update merchants set status = ?, merchant_type_id = ? where phone_number = ?", dataMerchant.Status, dataMerchant.MerchantTypeID, dataMerchant.PhoneNumber).Error
	return Dbcon.Save(&dataMerchant).Error
}

// GetMerchantByMerchantID ..
func GetMerchantByMerchantID(merchantID string) (models.CallPlanVisitMerchantRes, error) {
	fmt.Println(">>> CallPlanActionAddOrEdit - GetMerchantByMerchantID - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	res := models.CallPlanVisitMerchantRes{}

	var err error
	//err = Dbcon.Where("merchant_id = ? or phone_number = ?", req.MerchantId, req.MerchantPhone).First(&res).Error
	err = Dbcon.Table("merchants a").
		Select("a.id as id_merchant, a.name merchant_name, a.merchant_id, a.address merchant_address, a.phone_number merchant_phone, a.merchant_type_id, a.mpan mpan, b.name merchant_type_name").
		Joins("LEFT JOIN merchant_types b ON b.id = a.merchant_type_id").
		Where("a.merchant_id = ?", merchantID).First(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get merchant by merchant phone")
		return res, err
	}
	return res, nil
}

// MerchantNotFoundListQRProblem ..
func MerchantNotFoundListQRProblem() (res []models.TodolistMerchantNotFoundListRes, err error) {
	fmt.Println(">>> MerchantNotFoundList - MerchantNotFoundListQRProblem - Postgres <<<")
	sugarLogger := ottologger.GetLogger()

	query := "select * from qr_problems where status = 'Active'"

	err = Dbcon.Raw(query).Scan(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get merchant not found list")
		return res, err
	}
	return res, nil
}

// AcquisitionsMerchantDetail ..
func AcquisitionsMerchantDetail(req models.AcquisitionsMerchantDetailReq) (res dbmodels.Merchant, err error) {
	fmt.Println(">>> AcquisitionsMerchantDetail - AcquisitionsMerchantDetail - Postgres <<<")
	err = Dbcon.Where("phone_number = ? ", req.MerchantPhone).Find(&res).Error
	return res, err
}

// GetMerchantByIDCard ..
func GetMerchantByIDCard(idCard string) (res models.CheckIdCardRes, err error) {
	fmt.Println(">>> CheckIdCard - GetMerchantByIDCard - Postgres <<<")
	query := "select count(*) as num_id_card from merchants where id_card = '" + idCard + "'"
	err = Dbcon.Raw(query).Scan(&res).Error
	//err = Dbcon.Where("id_card = ? ", idCard).Find(&res).Error
	return res, err
}

// GetMerchantByPhoneNumber ..
func GetMerchantByPhoneNumber(merchantPhone string) (merchant dbmodels.Merchant, err error) {
	fmt.Println(">>> IndomarcoCheckMerchantOttopay - GetMerchantByPhoneNumber - Postgres <<<")
	err = Dbcon.Where("phone_number = ? ", merchantPhone).Find(&merchant).Error
	return merchant, err
}

// GetMerchantLinkingIndomarco ..
func GetMerchantLinkingIndomarco(merchantID string) (res dbmodels.MerchantLinkingIndomarco, err error) {
	fmt.Println(">>> GetMerchantLinkingIndomarco - GetMerchantLinkingIndomarco - Postgres <<<")
	err = Dbcon.Where("merchant_id = ? ", merchantID).Find(&res).Error
	return res, err
}

// UpdateMerchantLinkingIndomarco ..
func UpdateMerchantLinkingIndomarco(dataMerchant dbmodels.MerchantLinkingIndomarco) error {
	fmt.Println(">>> ChangeDataMerchant - UpdateDataMerchant - Postgres <<<")
	return Dbcon.Save(&dataMerchant).Error
}
