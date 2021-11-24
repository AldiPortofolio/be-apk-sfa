package postgres

import (
	"fmt"
	"ottosfa-api-apk/database/dbmodels"

	ottologger "ottodigital.id/library/logger"
)

// GetBusinessCategoryList ..
func GetBusinessCategoryList() ([]dbmodels.BusinessCategory, error) {
	fmt.Println(">>> GetBusinessCategoryList - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	res := []dbmodels.BusinessCategory{}

	err := Dbcon.Find(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get business category")
		return res, err
	}
	return res, err
}
