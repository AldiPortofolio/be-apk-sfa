package postgres

import (
	"fmt"
	ottologger "ottodigital.id/library/logger"
	"ottosfa-api-apk/database/dbmodels"
)

// ListProvince ..
func ListProvince(countryId string) (*[]dbmodels.Provinces, error) {
	fmt.Println(">>> Province - ListProvince - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	res := []dbmodels.Provinces{}

	//err := Dbcon.Select("id, INITCAP(name) as name").Order("name asc").Find(&res).Error
	err := Dbcon.Select("id, name").Order("name asc").Find(&res).Error
	if err != nil {
		sugarLogger.Info("Failed to get data provinces")
		return &res, err
	}
	return &res, nil
}
