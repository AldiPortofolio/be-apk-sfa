package postgresrose

import (
	"fmt"
	ottologger "ottodigital.id/library/logger"
	"ottosfa-api-apk/models"
)

// GetMerchantByMerchantPhone ..
func GetMerchantByMerchantPhone(merchantPhone string) (models.CallPlanVisitMerchantRes, error) {
	fmt.Println(">>> CallPlanVisitCheckMerchantPhone/CallPlanVisitCheckQRIS - GetMerchantByMerchantPhone - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	res := models.CallPlanVisitMerchantRes{}

	var err error
	//err = Dbcon.Where("merchant_id = ? or phone_number = ?", req.MerchantId, req.MerchantPhone).First(&res).Error
	err = Dbcon.Table("merchant a").
		Select("a.id as id_merchant, a.store_name merchant_name, a.merchant_outlet_id merchant_id, a.alamat merchant_address, a.store_phone_number merchant_phone, b.merchant_type_id, a.merchant_pan mpan").
		Joins("LEFT JOIN merchant_sfa b ON b.merchant_id = a.id").
		Where("a.store_phone_number = ?", merchantPhone).First(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database ROSE when get merchant by merchant phone")
		return res, err
	}
	return res, nil
}
