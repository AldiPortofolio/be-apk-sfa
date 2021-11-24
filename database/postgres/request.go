package postgres

import (
	"fmt"
	ottologger "ottodigital.id/library/logger"
	"ottosfa-api-apk/database/dbmodels"
)

// UploadToTableRequests ..
func UploadToTableRequests(req dbmodels.Requests) error {
	fmt.Println(">>> ChangePinAndIdCardPhoto - UploadToTableRequests - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	err := Dbcon.Create(&req).Error
	if err != nil {
		sugarLogger.Error("Failed Insert to Table Requests")
		return err
	}
	return nil
}
