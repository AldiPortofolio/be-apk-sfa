package postgres

import (
	"fmt"
	ottologger "ottodigital.id/library/logger"
)

// UpdateStatusSuccessMerchantNewRec ..
func UpdateStatusSuccessMerchantNewRec(idMerchant int64) error {
	fmt.Println(">>> TodolistPost - UpdateStatusSuccessMerchantNewRec - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	err := Dbcon.Exec("update merchant_new_recruitments set status = 'Registered' , updated_at = now() where id = ? ", idMerchant).Error
	if err != nil {
		sugarLogger.Error("Failed Update Status Merchant New Recruitments (Registered)")
		return err
	}
	return nil
}
