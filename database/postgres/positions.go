package postgres

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	ottologger "ottodigital.id/library/logger"
	"ottosfa-api-apk/database/dbmodels"
	"ottosfa-api-apk/models"
	"strconv"
	"strings"
)

// CheckPositions ..
func CheckPositions(salesmanID int) (*[]dbmodels.Positions, error) {
	fmt.Println(">>> MerchantList - CheckPositions - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	res := []dbmodels.Positions{}

	var err error
	err = Dbcon.Where("salesman_id = ?", salesmanID).Find(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when check position sales")
		return &res, err
	}
	return &res, nil
}

// CheckPositionsSalesSubarea ..
func CheckPositionsSalesSubarea(salesmanID int) (*[]dbmodels.Positions, error) {
	fmt.Println(">>> CallPlanList - CheckPositionsSalesSubarea - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	res := []dbmodels.Positions{}

	var err error
	err = Dbcon.Where("salesman_id = ? and sales_role_id = 4", salesmanID).Find(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when check position sales subarea")
		return &res, err
	}
	return &res, nil
}

// CheckTeamLeaderData ..
func CheckTeamLeaderData(salesmanID int) ([]models.TeamLeaderData, error) {
	fmt.Println(">>> CallPlanList - CheckTeamLeaderData - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	res := []models.TeamLeaderData{}

	var query = "select " +
		"(a.first_name || ' ' || a.last_name) team_leader_name, " +
		"e.branch_office as branch_office_name " +
		"from areas d " +
		"join positions b on b.regionable_id = d.id " +
		"join salesmen a on a.id = b.salesman_id " +
		"join branches e on e.id = d.branch_id " +
		"where " +
		"d.code in (select d.code from salesmen a " +
		"join positions b on b.salesman_id = a.id " +
		"join sub_areas c on c.id = b.regionable_id " +
		"join areas d on d.id = c.area_id " +
		"where a.id =  " + strconv.Itoa(salesmanID) + " ) " +
		"and b.sales_role_id = 3"

	err := Dbcon.Raw(query).Scan(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when check data team leader")
		return res, err
	}
	return res, nil
}

// GetMerchantListRegionWithFilter ..
func GetMerchantListRegionWithFilter(req models.MerchantListReq, salesId int) ([]dbmodels.Merchant, error) {
	resp := []dbmodels.Merchant{}

	var queryFilterVillage string
	if len(req.VillageId) != 0 {
		queryFilterVillage = "AND village_id in (" + strings.Join(req.VillageId[:], ",") + ") "
	}

	var queryFilterDistrict string
	if len(req.DistrictId) != 0 {
		queryFilterDistrict = "AND district_id in (" + strings.Join(req.DistrictId[:], ",") + ") "
	}

	var queryFilterCity string
	if len(req.CityId) != 0 {
		queryFilterCity = "AND city_id in (" + strings.Join(req.CityId[:], ",") + ") "
	}

	var queryFilterProvince string
	if len(req.ProvinceId) != 0 {
		queryFilterProvince = "AND province_id in (" + strings.Join(req.ProvinceId[:], ",") + ") "
	}

	var query = "SELECT * FROM merchants WHERE village_id in " +
		" (SELECT id FROM villages WHERE district_id in " +
		" (SELECT id FROM districts WHERE city_id in " +
		" (SELECT id FROM cities WHERE province_id in " +
		" (SELECT province_id FROM provinces_regions WHERE region_id in " +
		" (SELECT regionable_id FROM positions WHERE salesman_id = " + strconv.Itoa(salesId) + "))" +
		queryFilterProvince +
		")" +
		queryFilterCity +
		")" +
		queryFilterDistrict +
		")" +
		queryFilterVillage

	order := " ORDER BY name ASC "

	and := " AND "

	if req.Keyword != "" || req.MerchantCategory != "" {
		query = query
	}

	paramString := []string{}

	if req.Keyword != "" {
		paramString = append(paramString, " LOWER(name) like '%"+strings.ToLower(req.Keyword)+"%'")
	}

	if req.MerchantCategory != "" {
		switch req.MerchantCategory {
		case "OP":
			paramString = append(paramString, fmt.Sprintf(" (institution_id = '%s' OR institution_id = '' OR LOWER(institution_id) = 'idm' OR institution_id is NULL) ", req.MerchantCategory))
			break
		case "PGMI":
			paramString = append(paramString, fmt.Sprintf(" institution_id = '%s' ", req.MerchantCategory))
			break
		}
	}

	paramFilter := strings.Join(paramString[:], " AND ")

	if paramFilter != "" {
		paramFilter = and + paramFilter
	}

	page := fmt.Sprintf(" OFFSET %d LIMIT %d", (req.Page-1)*req.Limit, req.Limit)

	query = query + paramFilter + order + page

	sql := Dbcon.Raw(query).Scan(&resp)
	if sql.Error != nil {
		logs.Error("Failed")
		return resp, sql.Error
	}

	return resp, nil
}

// GetMerchantListBranchWithFilter ..
func GetMerchantListBranchWithFilter(req models.MerchantListReq, salesId int) ([]dbmodels.Merchant, error) {
	resp := []dbmodels.Merchant{}

	var queryFilterVillage string
	if len(req.VillageId) != 0 {
		queryFilterVillage = "AND village_id in (" + strings.Join(req.VillageId[:], ",") + ") "
	}

	var queryFilterDistrict string
	if len(req.DistrictId) != 0 {
		queryFilterDistrict = "AND district_id in (" + strings.Join(req.DistrictId[:], ",") + ") "
	}

	var queryFilterCity string
	if len(req.CityId) != 0 {
		queryFilterCity = "AND city_id in (" + strings.Join(req.CityId[:], ",") + ") "
	}

	var query = "SELECT * FROM merchants WHERE village_id in " +
		" (SELECT id FROM villages WHERE district_id in " +
		" (SELECT id FROM districts WHERE city_id in " +
		" (SELECT city_id FROM branches_cities WHERE branch_id in " +
		" (SELECT regionable_id FROM positions WHERE salesman_id = " + strconv.Itoa(salesId) + "))" +
		queryFilterCity +
		")" +
		queryFilterDistrict +
		")" +
		queryFilterVillage

	order := " ORDER BY name ASC "

	and := " AND "

	if req.Keyword != "" || req.MerchantCategory != "" {
		query = query
	}

	paramString := []string{}

	if req.Keyword != "" {
		paramString = append(paramString, " LOWER(name) like '%"+strings.ToLower(req.Keyword)+"%'")
	}

	if req.MerchantCategory != "" {
		switch req.MerchantCategory {
		case "OP":
			paramString = append(paramString, fmt.Sprintf(" (institution_id = '%s' OR institution_id = '' OR LOWER(institution_id) = 'idm' OR institution_id is NULL) ", req.MerchantCategory))
			break
		case "PGMI":
			paramString = append(paramString, fmt.Sprintf(" institution_id = '%s' ", req.MerchantCategory))
			break
		}
	}

	paramFilter := strings.Join(paramString[:], " AND ")

	if paramFilter != "" {
		paramFilter = and + paramFilter
	}

	page := fmt.Sprintf(" OFFSET %d LIMIT %d", (req.Page-1)*req.Limit, req.Limit)

	query = query + paramFilter + order + page

	sql := Dbcon.Raw(query).Scan(&resp)
	if sql.Error != nil {
		logs.Error("Failed")
		return resp, sql.Error
	}

	return resp, nil
}

// GetMerchantListAreaWithFilter ..
func GetMerchantListAreaWithFilter(req models.MerchantListReq, salesId int) ([]dbmodels.Merchant, error) {
	resp := []dbmodels.Merchant{}

	var queryFilterVillage string
	if len(req.VillageId) != 0 {
		queryFilterVillage = "AND village_id in (" + strings.Join(req.VillageId[:], ",") + ") "
	}

	var queryFilterDistrict string
	if len(req.DistrictId) != 0 {
		queryFilterDistrict = "AND district_id in (" + strings.Join(req.DistrictId[:], ",") + ") "
	}

	var query = "SELECT * FROM merchants WHERE village_id in " +
		" (SELECT id FROM villages WHERE district_id in " +
		" (SELECT district_id FROM areas_districts WHERE area_id in " +
		" (SELECT regionable_id FROM positions WHERE salesman_id = " + strconv.Itoa(salesId) + ") " +
		queryFilterDistrict +
		") ) " +
		queryFilterVillage

	order := " ORDER BY name ASC "

	and := " AND "

	if req.Keyword != "" || req.MerchantCategory != "" {
		query = query
	}

	paramString := []string{}

	if req.Keyword != "" {
		paramString = append(paramString, " LOWER(name) like '%"+strings.ToLower(req.Keyword)+"%'")
	}

	if req.MerchantCategory != "" {
		switch req.MerchantCategory {
		case "OP":
			paramString = append(paramString, fmt.Sprintf(" (institution_id = '%s' OR institution_id = '' OR LOWER(institution_id) = 'idm' OR institution_id is NULL) ", req.MerchantCategory))
			break
		case "PGMI":
			paramString = append(paramString, fmt.Sprintf(" institution_id = '%s' ", req.MerchantCategory))
			break
		}
	}

	paramFilter := strings.Join(paramString[:], " AND ")

	if paramFilter != "" {
		paramFilter = and + paramFilter
	}

	page := fmt.Sprintf(" OFFSET %d LIMIT %d", (req.Page-1)*req.Limit, req.Limit)

	query = query + paramFilter + order + page

	sql := Dbcon.Raw(query).Scan(&resp)
	if sql.Error != nil {
		logs.Error("Failed")
		return resp, sql.Error
	}

	return resp, nil
}

// GetMerchantListSubareaWithFilter ..
func GetMerchantListSubareaWithFilter(req models.MerchantListReq, salesId int) ([]dbmodels.Merchant, error) {
	resp := []dbmodels.Merchant{}

	var queryFilterVillage string
	if len(req.VillageId) != 0 {
		queryFilterVillage = "AND village_id in (" + strings.Join(req.VillageId[:], ",") + ") "
	}

	var query = "SELECT * FROM merchants " +
		"WHERE village_id in " +
		"(SELECT village_id FROM sub_areas_villages WHERE sub_area_id in " +
		"(SELECT regionable_id FROM positions WHERE salesman_id = " + strconv.Itoa(salesId) + "))" +
		queryFilterVillage

	order := " ORDER BY name ASC "

	and := " AND "

	if req.Keyword != "" || req.MerchantCategory != "" {
		query = query
	}

	paramString := []string{}

	if req.Keyword != "" {
		paramString = append(paramString, " LOWER(name) like '%"+strings.ToLower(req.Keyword)+"%'")
	}

	if req.MerchantCategory != "" {
		switch req.MerchantCategory {
		case "OP":
			paramString = append(paramString, fmt.Sprintf(" (institution_id = '%s' OR institution_id = '' OR LOWER(institution_id) = 'idm' OR institution_id is NULL) ", req.MerchantCategory))
			break
		case "PGMI":
			paramString = append(paramString, fmt.Sprintf(" institution_id = '%s' ", req.MerchantCategory))
			break
		}
	}

	paramFilter := strings.Join(paramString[:], " AND ")

	if paramFilter != "" {
		paramFilter = and + paramFilter
	}

	page := fmt.Sprintf(" OFFSET %d LIMIT %d", (req.Page-1)*req.Limit, req.Limit)

	query = query + paramFilter + order + page

	sql := Dbcon.Raw(query).Scan(&resp)
	if sql.Error != nil {
		logs.Error("Failed")
		return resp, sql.Error
	}

	return resp, nil
}
