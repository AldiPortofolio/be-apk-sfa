package postgres

import (
	"fmt"
	ottologger "ottodigital.id/library/logger"
	"ottosfa-api-apk/database/dbmodels"
	"ottosfa-api-apk/models"
)

// GetCallPlanDescriptionMerchant ..
func GetCallPlanDescriptionMerchant(req models.CallPlanDescriptionMerchantReq) (dbmodels.CallPlanDescriptions, error) {
	fmt.Println(">>> CallPlanDescriptionMerchant - GetCallPlanDescriptionMerchant - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	res := dbmodels.CallPlanDescriptions{}

	err := Dbcon.Model(dbmodels.CallPlanDescriptions{}).Where("action_merchant_id = ? and product_merchant_id = ?", req.ActionMerchantId, req.ProductMerchantId).Find(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get call plan description merchant")
		return res, err
	}
	return res, err
}
