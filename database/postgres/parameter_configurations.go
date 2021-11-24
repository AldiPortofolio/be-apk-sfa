package postgres

import (
	"fmt"
	"ottosfa-api-apk/database/dbmodels"
)

// GetNumDuplicateIdCard ..
func GetNumDuplicateIdCard() (res dbmodels.ParameterConfigurations, err error) {
	fmt.Println(">>> CheckIdCard - GetPercentageAcquisitions - Postgres <<<")
	err = Dbcon.Where("name = 'num_duplicate_id'" ).Find(&res).Error
	return res, err
}