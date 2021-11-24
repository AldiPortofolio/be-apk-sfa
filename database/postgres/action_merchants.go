package postgres

import (
	"fmt"
	ottologger "ottodigital.id/library/logger"
	"ottosfa-api-apk/database/dbmodels"
)

// GetCallPlanActionMerchantList ..
func GetCallPlanActionMerchantList() ([]dbmodels.ActionMerchants, error) {
	fmt.Println(">>> CallPlanActionMerchantList - GetCallPlanActionMerchantList - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	res := []dbmodels.ActionMerchants{}

	//err := Dbcon.Find(&res).Error
	err := Dbcon.Model(dbmodels.ActionMerchants{}).Find(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get call plan product merchant list")
		return res, err
	}
	return res, err
}
