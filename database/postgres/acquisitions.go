package postgres

import (
	"fmt"
	"ottosfa-api-apk/database/dbmodels"
	"ottosfa-api-apk/models"

	ottologger "ottodigital.id/library/logger"
)

// GetAcquisitionBySR ..
func GetAcquisitionBySR(req models.GetAcquisitionBySalesTypeID) ([]dbmodels.Acquisitions, error) {
	fmt.Println(">>> AcquisitionBySR - GetAcquisitionBySR - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	res := []dbmodels.Acquisitions{}

	err := Dbcon.Where("sales_retails like '%" + req.SalesTypeId + "%' and show_in_app = 'Active'").Order("sequence ASC").Find(&res).Error
	// err := Dbcon.Find(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get business type by sales retail")
		return res, err
	}
	return res, err
}

// GetAcquisitionBySR ..
func GetAcquisitionBySrV2(req models.GetAcquisitionBySalesTypeID) ([]dbmodels.AcquisitionsSR, error) {
	fmt.Println(">>> AcquisitionBySR - GetAcquisitionBySR - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	res := []dbmodels.AcquisitionsSR{}

	err := Dbcon.Where("sales_retails like '%" + req.SalesTypeId + "%' and show_in_app = 'Active'").Order("sequence ASC").Find(&res).Error
	// err := Dbcon.Find(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get business type by sales retail")
		return res, err
	}
	return res, err
}

// GetAcquisitionByID ..
func GetAcquisitionByID(ids []int64) ([]dbmodels.Acquisitions, error) {
	fmt.Println(">>> AcquisitionBySR - GetAcquisitionBySR - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	res := []dbmodels.Acquisitions{}

	err := Dbcon.Where("id in (?) and show_in_app = 'Active'", ids).Order("sequence ASC").Find(&res).Error
	// err := Dbcon.Find(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get business type by sales retail")
		return res, err
	}
	return res, err
}

// GetBussinessTypeByName..
func GetBusinessTypeByName(req models.GetAcquisitionByName) ([]dbmodels.Acquisitions, error) {
	fmt.Println(">>> AcquisitionBySR - GetAcquisitionBySR - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	res := []dbmodels.Acquisitions{}

	where := "1=1"
	if req.Name != "" {
		where += " and lower(name) = lower('" + req.Name + "')"
	}

	if req.RoseMerchantGroup != "" {
		where += " and lower(rose_merchant_group) like lower('%" + req.RoseMerchantGroup + "%')"
	}

	if req.RoseMerchantCategory != "" {
		where += " and lower(rose_merchant_category) like lower('%" + req.RoseMerchantCategory + "%')"
	}

	if req.SalesRetailId != "" {
		where += " and lower(sales_retails) like lower('%" + req.SalesRetailId + "%')"
	}

	err := Dbcon.Where(where).Find(&res).Error
	// err := Dbcon.Find(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get business type by sales retail")
		return res, err
	}
	return res, err
}

// GetBusinessTypeByCode..
func GetBusinessTypeByCode(req []string) ([]dbmodels.MerchantBusinessTypes, error) {
	fmt.Println(">>> BusinessTypeListByName - GetBusinessTypeByCode - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	res := []dbmodels.MerchantBusinessTypes{}

	err := Dbcon.Where("code in (?)", req).Find(&res).Error
	// err := Dbcon.Find(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get business type by sales retail")
		return res, err
	}
	return res, err
}
