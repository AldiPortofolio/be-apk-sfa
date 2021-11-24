package postgres

import (
	"fmt"
	ottologger "ottodigital.id/library/logger"
	"ottosfa-api-apk/database/dbmodels"
)

// ListCity ..
func ListCity(provinceId string) (*[]dbmodels.Cities, error) {
	fmt.Println(">>> City - ListCity - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	res := []dbmodels.Cities{}

	//err := Dbcon.Select("id, INITCAP(name) as name").Where("province_id = ?", provinceId).Order("name asc").Find(&res).Error
	err := Dbcon.Select("id, name").Where("province_id = ?", provinceId).Order("name asc").Find(&res).Error
	if err != nil {
		sugarLogger.Info("Failed to get data cities")
		return &res, err
	}
	return &res, nil
}
