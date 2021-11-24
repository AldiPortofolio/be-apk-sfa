package postgres

import (
	"fmt"
	ottologger "ottodigital.id/library/logger"
	"ottosfa-api-apk/database/dbmodels"
	"ottosfa-api-apk/models"
	"strconv"
)

// ListVillageByPositionSales ..
func ListVillageByPositionSales(salesId int) (res []models.TodolistVillageID, err error) {
	fmt.Println(">>> TodolistList/MerchantList - ListVillageByPositionSales - Postgres <<<")
	sugarLogger := ottologger.GetLogger()

	err = Dbcon.Table("positions p").
		Select("sav.village_id ").
		Joins("JOIN sub_areas_villages sav on p.regionable_id = sav.sub_area_id").
		Where("p.salesman_id = ? and p.regionable_type = 'SubArea'", salesId).Scan(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get village id by position sales")
		return res, err
	}
	return res, nil
}

// ListVillage ..
func ListVillage(districtId string) (*[]dbmodels.Villages, error) {
	fmt.Println(">>> Village - ListVillage - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	res := []dbmodels.Villages{}

	//err := Dbcon.Select("id, INITCAP(name) as name").Where("district_id = ?", districtId).Order("name asc").Find(&res).Error
	err := Dbcon.Select("id, name").Where("district_id = ?", districtId).Order("name asc").Find(&res).Error
	if err != nil {
		sugarLogger.Info("Failed to get data villages")
		return &res, err
	}
	return &res, nil
}

// ListVillageBySalesTypeSales ..
func ListVillageBySalesTypeSales(salesId int) (res []models.TodolistVillageID, err error) {
	fmt.Println(">>> TodolistList - ListVillageBySalesTypeSales - Postgres <<<")
	sugarLogger := ottologger.GetLogger()

	query := "SELECT village_id FROM sub_areas_villages WHERE sub_area_id in " +
		" (SELECT scs.sub_area_id FROM sales_area_channels_sub_areas scs LEFT JOIN sales_area_channels sac on sac.id = scs.sales_area_channel_id WHERE sac.sales_type_id in " + //" (SELECT sub_area_id FROM sales_area_channels_sub_areas WHERE sales_area_channel_id in " +
		" (SELECT st.id FROM sales_types st " + //" (SELECT st.id FROM sales_types_salesmen sts " +
		" LEFT JOIN sales_types_salesmen sts on st.id = sts.sales_type_id WHERE sts.salesman_id = " + strconv.Itoa(salesId) +//" LEFT JOIN sales_types st on st.id = sts.sales_type_id WHERE salesman_id = " + strconv.Itoa(salesId) +
		" ))"

	err = Dbcon.Raw(query).Scan(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get village id by sales type sales")
		return res, err
	}
	return res, nil
}
