package postgres

import (
	"fmt"
	ottologger "ottodigital.id/library/logger"
	"ottosfa-api-apk/database/dbmodels"
	"ottosfa-api-apk/models"
)

// GetDataSalesAndLocation ..
func GetDataSalesAndLocation(phoneNumber string) (models.GetDataSalesLocationDB, error) {
	fmt.Println(">>> UpdatePhotoProfilSales - GetDataSalesAndLocation - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	var res models.GetDataSalesLocationDB

	//var query = " select * from salesmen s left join location_areas ls on ls.salesman_id = s.id where phone_number = '"+ phoneNumber +"' "
	var query = " select * from salesmen s left join location_areas ls on ls.salesman_id = s.id where phone_number = '080808080806' "
	err := Dbcon.Raw(query).Scan(&res).Error
	if err != nil {
		sugarLogger.Error("Failed to get data salesmen and location")
		return res, err
	}
	return res, nil
}

// UpdatePhotoProfile ..
func UpdatePhotoProfile(namePhoto, phoneNumber string) error {
	fmt.Println(">>> UpdatePhotoProfilSales - UpdatePhotoProfile - Postgres <<<")
	sugarLogger := ottologger.GetLogger()

	err := Dbcon.Exec("update salesmen set photo = ? where phone_number = ?", namePhoto, phoneNumber).Error
	if err != nil {
		sugarLogger.Error("Failed to update photo profil sales")
		return err
	}
	return nil
}

// GetDataSales ..
func GetDataSales(phoneNumber string) (*dbmodels.Salesmen, error) {
	fmt.Println(">>> ChangePasswordSales - GetDataSales - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	res := dbmodels.Salesmen{
		PhoneNumber: phoneNumber,
	}

	err := Dbcon.Where("phone_number = ?", phoneNumber).First(&res).Error
	if err != nil {
		sugarLogger.Info("Failed to get data salesmen by phone number")
		return &res, err
	}
	return &res, nil
}

// UpdatePassword ..
func UpdatePassword(id int, pin string) error {
	fmt.Println(">>> ChangePasswordSales - UpdatePhotoProfile - Postgres <<<")
	sugarLogger := ottologger.GetLogger()

	//res := dbmodels.Salesmen{
	//	PasswordDigest: pin,
	//}
	err := Dbcon.Exec("update salesmen set password_digest = ? where id = ?", pin, id).Error
	if err != nil {
		sugarLogger.Error("Failed connect to database SFA when updating pin")
		return err
	}
	return nil
}

// CheckToken ..
func CheckToken(token string) (*dbmodels.Salesmen, error) {
	fmt.Println(">>> CheckProfilSales - CheckToken - Postgres <<<")
	sugarLogger := ottologger.GetLogger()

	res := dbmodels.Salesmen{}
	err := Dbcon.Where("session_token = ?", token).First(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when check token sales")
		return &res, err
	}
	return &res, nil
}

// GetPosition ..
func GetPosition(salesmanID int) (*dbmodels.Positions, error) {
	fmt.Println(">>> CheckProfilSales - GetPositions - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	res := dbmodels.Positions{}

	var err error
	err = Dbcon.Where("salesman_id = ?", salesmanID).First(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get position sales")
		return &res, err
	}
	return &res, nil
}

// InnerJoinSubArea ..
func InnerJoinSubArea(id int) ([]models.ListLocationSales, error) {
	fmt.Println(">>> CheckProfilSales - InnerJoinSubArea - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	resp := []models.ListLocationSales{}
	//query := fmt.Sprintf(
	//	" select a.id, a.first_name, a.last_name,"+
	//		" b.role_name, b.regionable_id, b.regionable_type,"+
	//		" c.id as sub_area_id, c.area_id as area_id, c.name as sub_area_name,"+
	//		" d.id as area_id2, d.branch_id as branch_id, d.name as area_name,"+
	//		" e.id as branch_id2, e.region_id as region_id, e.name as branch_name,"+
	//		" f.id as region_id2, f.name as region_name"+
	//		" from salesmen a "+
	//		" join positions b on b.salesman_id = a.id "+
	//		" join sub_areas c on c.id = b.regionable_id"+
	//		" join areas d on d.id = c.area_id"+
	//		" join branches e on e.id = d.branch_id"+
	//		" join regions f on f.id = e.region_id"+
	//		" where a.id = %d", id)

	query := fmt.Sprintf(
		" select "+
			"a.first_name, a.last_name, "+
			"b.id id, b.sales_role_id role_id, b.role_name role_name, "+
			"(c.code || ' - ' || c.name) sub_area_name, c.id sub_area_id, "+
			"(d.code || ' - ' || d.name) area_name, d.id area_id, "+
			"(e.code || ' - ' || e.name) branch_name, e.branch_office, e.id branch_id, "+
			"(f.code || ' - ' || f.name) region_name, f.id region_id "+
			"from salesmen a "+
			"join positions b on b.salesman_id = a.id "+
			"join sub_areas c on c.id = b.regionable_id "+
			"join areas d on d.id = c.area_id "+
			"join branches e on e.id = d.branch_id "+
			"join regions f on f.id = e.region_id "+
			"where a.id = %d "+
			"group by a.id, b.id, "+
			"c.id, c.code, c.name, "+
			"d.id, d.code, d.name, "+
			"e.id, e.code, e.name, "+
			"f.id, f.code, f.name", id)

	sql := Dbcon.Raw(query).Scan(&resp)
	fmt.Println("=========================================")
	fmt.Println("Position SubArea : ", resp)
	fmt.Println("=========================================")
	if sql.Error != nil {
		sugarLogger.Error("Failed connect to database SFA when get position sales")
	}

	return resp, nil
}

// InnerJoinRegion ..
func InnerJoinRegion(id int) ([]models.ListLocationSales, error) {
	fmt.Println(">>> CheckProfilSales - InnerJoinRegion - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	resp := []models.ListLocationSales{}
	//query := fmt.Sprintf(
	//	"select a.id, a.first_name, a.last_name, " +
	//		" b.role_name, b.regionable_id, b.regionable_type, "+
	//		" c.id as region_id, c.name as region_name, "+
	//		" d.id as branch_id, d.name as branch_name, "+
	//		" e.id as area_id, e.name as area_name, "+
	//		" f.id as sub_area_id, f.name as sub_area_name"+
	//		" from salesmen a"+
	//		" left join positions b on b.salesman_id = a.id"+
	//		" left join regions c on b.regionable_id = c.id"+
	//		" left join branches d on c.id = d.region_id"+
	//		" left join areas e on d.id = e.branch_id"+
	//		" left join sub_areas f on e.id = f.area_id"+
	//		" where a.id = %d;", id)

	query := fmt.Sprintf(
		" select "+
			"a.first_name, a.last_name, "+
			"b.id id, b.sales_role_id role_id, b.role_name role_name, "+
			"(f.code || ' - ' || f.name) region_name, f.id region_id "+
			"from salesmen a "+
			"join positions b on b.salesman_id = a.id "+
			"join regions f on f.id = b.regionable_id  "+
			"where a.id = %d "+
			"group by a.id, b.id, "+
			"f.id, f.code, f.name", id)

	sql := Dbcon.Raw(query).Scan(&resp)
	fmt.Println("=========================================")
	fmt.Println("Position Region : ", resp)
	fmt.Println("=========================================")
	if sql.Error != nil {
		sugarLogger.Error("Failed connect to database SFA when get position sales")
	}

	return resp, nil
}

// InnerJoinArea ..
func InnerJoinArea(id int) ([]models.ListLocationSales, error) {
	fmt.Println(">>> CheckProfilSales - InnerJoinArea - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	resp := []models.ListLocationSales{}
	//query := fmt.Sprintf(
	//	"select a.id, a.first_name, a.last_name, b.role_name, b.regionable_id, b.regionable_type, "+
	//		" e.id as region_id, e.name as region_name, "+
	//		" d.id as branch_id, d.name as branch_name, "+
	//		" c.id as area_id, c.name as area_name, "+
	//		" f.id as sub_area_id, f.name as sub_area_name"+
	//		" from salesmen a"+
	//		" left join positions b on b.salesman_id = a.id"+
	//		" left join areas c on b.regionable_id = c.id"+
	//		" left join branches d on c.branch_id = d.id"+
	//		" left join regions e on d.region_id = e.id"+
	//		" left join sub_areas f on f.area_id = c.id"+
	//		" where a.id = %d;", id)

	query := fmt.Sprintf(
		" select "+
			"a.first_name, a.last_name, "+
			"b.id id, b.sales_role_id role_id, b.role_name role_name, "+
			"(d.code || ' - ' || d.name) area_name, d.id area_id, "+
			"(e.code || ' - ' || e.name) branch_name, e.branch_office, e.id branch_id, "+
			"(f.code || ' - ' || f.name) region_name, f.id region_id "+
			"from salesmen a "+
			"join positions b on b.salesman_id = a.id "+
			"join areas d on d.id = b.regionable_id "+
			"join branches e on e.id = d.branch_id "+
			"join regions f on f.id = e.region_id "+
			"where a.id = %d "+
			"group by a.id, b.id, "+
			"d.id, d.code, d.name, "+
			"e.id, e.code, e.name, "+
			"f.id, f.code, f.name", id)

	sql := Dbcon.Raw(query).Scan(&resp)
	fmt.Println("=========================================")
	fmt.Println("Position Area : ", resp)
	fmt.Println("=========================================")
	if sql.Error != nil {
		sugarLogger.Error("Failed connect to database SFA when get position sales")
	}

	return resp, nil
}

// InnerJoinBranch ..
func InnerJoinBranch(id int) ([]models.ListLocationSales, error) {
	fmt.Println(">>> CheckProfilSales - InnerJoinBranch - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	resp := []models.ListLocationSales{}
	//query := fmt.Sprintf(
	//	"select a.id, a.first_name, a.last_name,"+
	//		" b.role_name, b.regionable_id, b.regionable_type,"+
	//		" d.id as branch_id, d.name as branch_name,"+
	//		" e.id as region_id, e.name as region_name"+
	//		" from salesmen a "+
	//		" join positions b on b.salesman_id = a.id"+
	//		" join branches d on d.id = b.regionable_id"+
	//		" join regions e on e.id = d.region_id"+
	//		" where a.id = %d", id)

	query := fmt.Sprintf(
		" select "+
			"a.first_name, a.last_name, "+
			"b.id id, b.sales_role_id role_id, b.role_name role_name, "+
			"(e.code || ' - ' || e.name) branch_name, e.branch_office, e.id branch_id, "+
			"(f.code || ' - ' || f.name) region_name, f.id region_id "+
			"from salesmen a "+
			"join positions b on b.salesman_id = a.id "+
			"join branches e on e.id = b.regionable_id "+
			"join regions f on f.id = e.region_id "+
			"where a.id = %d "+
			"group by a.id, b.id, "+
			"e.id, e.code, e.name, "+
			"f.id, f.code, f.name", id)

	sql := Dbcon.Raw(query).Scan(&resp)
	fmt.Println("=========================================")
	fmt.Println("Position Branch : ", resp)
	fmt.Println("=========================================")
	if sql.Error != nil {
		sugarLogger.Error("Failed connect to database SFA when get position sales")
	}

	return resp, nil
}

// UpdatePinSales ..
func UpdatePinSales(NewPin, IdCard, photo string, id int) error {
	fmt.Println(">>> ChangePinAndIdCardPhoto - UpdatePinSales - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	//err := Dbcon.Exec("update salesmen set password_digest = ?, id_card = ?, photo = ?, status = 4 where id = ?", NewPin, IdCard, photo, id).Error
	err := Dbcon.Exec("update salesmen set password_digest = ?, status = 4 where id = ?", NewPin, id).Error
	if err != nil {
		sugarLogger.Error("Failed connect to database SFA when Update Photo Profil Pin and Photo KTP")
		return err
	}

	return nil
}

// UpdateDeviceLogin is func for updating device info when sales succes login
func UpdateDeviceLogin(id int, deviceID string, deviceToken string, token string, firebaseToken string) error {
	fmt.Println(">>> LoginIndomarco/Login - UpdateDeviceLogin - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	sales := dbmodels.Salesmen{ID: id}
	err := Dbcon.Model(&sales).Updates(dbmodels.Salesmen{DeviceID: deviceID, DeviceToken: deviceToken, SessionToken: token, FirebaseToken: firebaseToken}).Error
	if err != nil {
		sugarLogger.Error(fmt.Sprintf("Failed connect to database SFA when updating sales device info: ", err))
		return err
	}
	return err
}

// GetDataSalesByPhoneNumber ..
func GetDataSalesByPhoneNumber(IdNumber int) (dbmodels.Salesmen, error) {
	fmt.Println(">>> Login - GetDataSalesByPhoneNumber - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	res := dbmodels.Salesmen{}
	var query = fmt.Sprintf("select *, to_char(functional_position_id, '9') functional_position from salesmen where phone_number = '%d'", IdNumber)
	err := Dbcon.Raw(query).First(&res).Error
	if err != nil {
		sugarLogger.Error("Failed to get data Sales by phone number")
		return res, err
	}

	return res, nil
}

// GetDataSalesBySalesId ..
func GetDataSalesBySalesId(salesId string) (dbmodels.Salesmen, error) {
	fmt.Println(">>> LoginIndomarco/Login - GetDataSalesBySalesId - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	res := dbmodels.Salesmen{}
	var query = fmt.Sprintf("select *, to_char(functional_position_id, '9') functional_position from salesmen where sales_id = '%s'", salesId)
	err := Dbcon.Raw(query).First(&res).Error
	if err != nil {
		sugarLogger.Error("Failed to get data Sales by sales id")
		return res, err
	}

	return res, nil
}

// ClearSession ..
func ClearSession() error {
	fmt.Println(">>> ClearSession - ClearSession - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	sales := dbmodels.Salesmen{}
	err := Dbcon.Model(&sales).Updates(dbmodels.Salesmen{SessionToken: "-"}).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when clear session sales")
		return err
	}
	return nil
}

// GetDataSalesTypeSalesmen ..
func GetDataSalesTypeSalesmen(salesId int) ([]dbmodels.SalesType, error) {
	fmt.Println(">>> Login - GetDataSalesTypeSalesmen - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	res := []dbmodels.SalesType{}
	var query = fmt.Sprintf("select st.* from sales_types_salesmen sts left join sales_types st on st.id = sts.sales_type_id where salesman_id = %d", salesId)
	err := Dbcon.Raw(query).Scan(&res).Error
	if err != nil {
		sugarLogger.Error("Failed to get data Sales Type by sales id")
		return res, err
	}
	fmt.Println("res", res)
	return res, nil
}

// GetDataSalesBySalesIdv24 ..
func GetDataSalesBySalesIdv24(salesId string) (dbmodels.Salesmen, error) {
	fmt.Println(">>> Login - GetDataSalesBySalesIdv24 - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	res := dbmodels.Salesmen{}
	var query = fmt.Sprintf("select *, to_char(functional_position_id, '9') functional_position from salesmen where sales_id = '%s'", salesId)
	err := Dbcon.Raw(query).First(&res).Error
	if err != nil {
		sugarLogger.Error("Failed to get data Sales by sales id")
		return res, err
	}

	return res, nil
}

// PositionSubArea ..
func PositionSubArea(id int) (res models.PostionSales, err error) {
	fmt.Println(">>> Login - PositionSubArea - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	query := fmt.Sprintf(
		" select b.id position_id, b.sales_role_id role_id, b.role_name role_name," +
			" c.id sub_area_id, c.code sub_area_code, c.name sub_area_name, " +
			" d.id area_id, d.code area_code, d.name area_name, " +
			" e.id branch_id, e.code branch_code, e.name branch_name, e.branch_office, " +
			" f.id region_id, f.code region_code, f.name region_name, " +
			" g.id sales_type_id, g.name sales_type_name " +
			"from salesmen a "+
			"join positions b on b.salesman_id = a.id "+
			"join sub_areas c on c.id = b.regionable_id "+
			"join areas d on d.id = c.area_id "+
			"join branches e on e.id = d.branch_id "+
			"join regions f on f.id = e.region_id "+
			"join sales_types g on g.id = a.sales_type_id " +
			"where a.id = %d ", id)
	err = Dbcon.Raw(query).Scan(&res).Error
	if err != nil {
		sugarLogger.Error("Failed connect to database SFA when get position sales")
		return res, err
	}
	return res, nil
}

// PositionArea ..
func PositionArea(id int) (res models.PostionSales, err error) {
	fmt.Println(">>> Login - PositionArea - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	query := fmt.Sprintf(
		" select b.id position_id, b.sales_role_id role_id, b.role_name role_name," +
			//" c.id sub_area_id, c.code sub_area_code, c.name sub_area_name, " +
			" d.id area_id, d.code area_code, d.name area_name, " +
			" e.id branch_id, e.code branch_code, e.name branch_name, e.branch_office, " +
			" f.id region_id, f.code region_code, f.name region_name, " +
			" g.id sales_type_id, g.name sales_type_name " +
			"from salesmen a "+
			"join positions b on b.salesman_id = a.id "+
			"join areas d on d.id = b.regionable_id "+
			"join branches e on e.id = d.branch_id "+
			"join regions f on f.id = e.region_id "+
			"join sales_types g on g.id = a.sales_type_id " +
			"where a.id = %d ", id)
	err = Dbcon.Raw(query).Scan(&res).Error
	if err != nil {
		sugarLogger.Error("Failed connect to database SFA when get position sales")
		return res, err
	}
	return res, nil
}

// PositionBranch ..
func PositionBranch(id int) (res models.PostionSales, err error) {
	fmt.Println(">>> Login - PositionBranch - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	query := fmt.Sprintf(
		" select b.id position_id, b.sales_role_id role_id, b.role_name role_name," +
			//" c.id sub_area_id, c.code sub_area_code, c.name sub_area_name, " +
			//" d.id area_id, d.code area_code, d.name area_name, " +
			" e.id branch_id, e.code branch_code, e.name branch_name, e.branch_office, " +
			" f.id region_id, f.code region_code, f.name region_name, " +
			" g.id sales_type_id, g.name sales_type_name " +
			"from salesmen a "+
			"join positions b on b.salesman_id = a.id "+
			"join branches e on e.id = b.regionable_id "+
			"join regions f on f.id = e.region_id "+
			"join sales_types g on g.id = a.sales_type_id " +
			"where a.id = %d ", id)
	err = Dbcon.Raw(query).Scan(&res).Error
	if err != nil {
		sugarLogger.Error("Failed connect to database SFA when get position sales")
		return res, err
	}
	return res, nil
}

// PositionRegion ..
func PositionRegion(id int) (res models.PostionSales, err error) {
	fmt.Println(">>> Login - PositionRegion - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	query := fmt.Sprintf(
		" select b.id position_id, b.sales_role_id role_id, b.role_name role_name," +
		//" c.id sub_area_id, c.code sub_area_code, c.name sub_area_name, " +
		//" d.id area_id, d.code area_code, d.name area_name, " +
		//	" e.id branch_id, e.code branch_code, e.name branch_name, e.branch_office, " +
			" f.id region_id, f.code region_code, f.name region_name, " +
			" g.id sales_type_id, g.name sales_type_name " +
			"from salesmen a "+
			"join positions b on b.salesman_id = a.id "+
			"join regions f on f.id = b.regionable_id "+
			"join sales_types g on g.id = a.sales_type_id " +
			"where a.id = %d ", id)
	err = Dbcon.Raw(query).Scan(&res).Error
	if err != nil {
		sugarLogger.Error("Failed connect to database SFA when get position sales")
		return res, err
	}
	return res, nil
}
