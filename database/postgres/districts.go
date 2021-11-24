package postgres

import (
	"fmt"
	ottologger "ottodigital.id/library/logger"
	"ottosfa-api-apk/database/dbmodels"
)

// ListDistrict ..
func ListDistrict(cityId string) (*[]dbmodels.Districts, error) {
	fmt.Println(">>> District - ListDistrict - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	res := []dbmodels.Districts{}

	//err := Dbcon.Select("id, INITCAP(name) as name").Where("city_id = ?", cityId).Order("name asc").Find(&res).Error
	err := Dbcon.Select("id, name").Where("city_id = ?", cityId).Order("name asc").Find(&res).Error
	if err != nil {
		sugarLogger.Info("Failed to get data districts")
		return &res, err
	}
	return &res, nil
}
