package postgres

import (
	"fmt"
	"ottosfa-api-apk/database/dbmodels"
)

// GetMerchantTypeById ..
func GetMerchantTypeById(merchantTypeId int) (res dbmodels.MerchantTypes, err error) {
	fmt.Println(">>> CallPlanVisitCheckMerchantPhone - GetMerchantTypeById - Postgres <<<")
	err = Dbcon.Where("id = ?", merchantTypeId).First(&res).Error
	return res, err
}