package postgres

import (
	"fmt"
	ottologger "ottodigital.id/library/logger"
	"ottosfa-api-apk/database/dbmodels"
)

// GetCallPlanProductMerchantList ..
func GetCallPlanProductMerchantList() ([]dbmodels.ProductMerchants, error) {
	fmt.Println(">>> CallPlanProductMerchantList - GetCallPlanProductMerchantList - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	res := []dbmodels.ProductMerchants{}

	//err := Dbcon.Find(&res).Error
	err := Dbcon.Model(dbmodels.ProductMerchants{}).Find(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get call plan product merchant list")
		return res, err
	}
	return res, err
}
