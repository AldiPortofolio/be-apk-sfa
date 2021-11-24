package postgres

import (
	"fmt"
	"ottosfa-api-apk/database/dbmodels"

	ottologger "ottodigital.id/library/logger"
)

// GetBusinessTypeList ..
func GetBusinessTypeList() ([]dbmodels.BusinessType, error) {
	fmt.Println(">>> GetBusinessTypeList - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	res := []dbmodels.BusinessType{}

	err := Dbcon.Find(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get business type")
		return res, err
	}
	return res, err
}

// GetBusinessTypeListBySR ..
func GetBusinessTypeListBySR(salesTypeId []int) ([]dbmodels.BusinessType, error) {
	fmt.Println(">>> BusinessTypeListBySR - GetBusinessTypeListBySR - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	res := []dbmodels.BusinessType{}

	err := Dbcon.Where("sales_type_id in (?)", salesTypeId).Find(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get business type by sales retail")
		return res, err
	}
	return res, err
}
