package postgres

import (
	"fmt"
	ottologger "ottodigital.id/library/logger"
	"ottosfa-api-apk/database/dbmodels"
	"ottosfa-api-apk/models"
)

// GetCallPlanMerchantList ..
func GetCallPlanMerchantList(req models.CallPlanMerchantListReq) ([]models.CallPlanMerchantListRes, error) {
	fmt.Println(">>> CallPlanMerchantList - GetCallPlanMerchantList - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	res := []models.CallPlanMerchantListRes{}

	var query = "SELECT " +
		"cpm.id call_plan_merchant_id, cpm.merchant_name, cpm.mid as merchant_id, cpm.merchant_status, cpm.merchant_phone, " +
		"mt.name merchant_type_name, merchant_address, mt.priority " +
		"FROM call_plan_merchants cpm " +
		"LEFT JOIN merchant_types mt on mt.id = cpm.merchant_type_id	" +
		"WHERE " +
		"cpm.call_plan_id = " + fmt.Sprintf("%d", req.CallPlanId) +
		"AND lower(cpm.status) = lower('" + req.Status + "')"

	page := fmt.Sprintf(" OFFSET %d LIMIT %d", (req.Page-1)*req.Limit, req.Limit)

	query = query + page

	err := Dbcon.Raw(query).Scan(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get call plan merchant list")
		return res, err
	}
	return res, nil
}

// GetCallPlanListMerchantsAll ..
func GetCallPlanListMerchantsAll(callPlanId int64) []dbmodels.CallPlanMerchants {
	fmt.Println(">>> CallPlanList - GetCallPlanListMerchantsAll - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	res := []dbmodels.CallPlanMerchants{}

	err := Dbcon.Model(dbmodels.CallPlanMerchants{}).Where("call_plan_id = ? and status != 'Visited'", callPlanId).Find(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get call plan merchant")
		return res
	}
	return res
}

// GetCallPlanListMerchantsCompleted ..
func GetCallPlanListMerchantsCompleted(callPlanId int64) []dbmodels.CallPlanMerchants {
	fmt.Println(">>> CallPlanList - GetCallPlanListMerchantsAllCompleted - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	res := []dbmodels.CallPlanMerchants{}

	err := Dbcon.Model(dbmodels.CallPlanMerchants{}).Where("call_plan_id = ? and status = 'Completed'", callPlanId).Find(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get call plan merchant")
		return res
	}
	return res
}

// AddCallPlanMerchant ..
func AddCallPlanMerchant(req dbmodels.CallPlanMerchants) (dbmodels.CallPlanMerchants, error) {
	fmt.Println(">>> CallPlanVisitAdd - AddCallPlanMerchant - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	res := dbmodels.CallPlanMerchants{
		Id: 0,
	}

	err := Dbcon.Save(&req).Scan(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when save call plan merchant")
		return res, err
	}

	return res, err
}

// UpdateClockInCallPlanMerchant ..
func UpdateClockInCallPlanMerchant(callPlanMerchantId int64) error {
	fmt.Println(">>> CallPlanActionUpdateClockInMerchant - UpdateClockInCallPlanMerchant - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	var err error
	err = Dbcon.Exec("update call_plan_merchants set action_date = now() where id = ?", callPlanMerchantId).Error
	if err != nil {
		sugarLogger.Error("Failed connect to database SFA when updating new phone")
		return err
	}
	return nil
}

// UpdateLongLatCallPlanMerchant ..
func UpdateLongLatCallPlanMerchant(callPlanMerchantId int64, longitude string, latitude string) error {
	fmt.Println(">>> CallPlanActionCheckQRIS - UpdateLongLatCallPlanMerchant - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	var err error
	err = Dbcon.Exec("update call_plan_merchants set longitude = ? , latitude = ? where id = ?", longitude, latitude, callPlanMerchantId).Error
	if err != nil {
		sugarLogger.Error("Failed connect to database SFA when updating long lat call plan merchant")
		return err
	}
	return nil
}

// GetCallPlanMerchantsByCallPlanMerchantId ..
func GetCallPlanMerchantsByCallPlanMerchantId(callPlanMerchantId int64) (dbmodels.CallPlanMerchants, error) {
	fmt.Println(">>> CallPlanList - GetCallPlanMerchantsByCallPlanMerchantId - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	res := dbmodels.CallPlanMerchants{}

	err := Dbcon.Model(dbmodels.CallPlanMerchants{}).Where("id = ? ", callPlanMerchantId).Find(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get call plan merchant by id")
		return res, err
	}
	return res, err
}

// UpdateCallPlanMerchantUnknown ..
func UpdateCallPlanMerchantUnknown(req models.CallPlanActionMerchantUnknownReq) error {
	fmt.Println(">>> CallPlanActionMerchantUnknown - UpdateCallPlanMerchantUnknown - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	var err error
	err = Dbcon.Exec("update call_plan_merchants set status = ? , merchant_status = ? , clock_time = ? , longitude = ? , latitude = ? , photo_location = ? where id = ?",
		req.Status, req.MerchantStatus, req.ClockOut, req.Longitude, req.Latitude, req.PhotoLocation, req.CallPlanMerchantId).Error
	if err != nil {
		sugarLogger.Error("Failed connect to database SFA when updating call plan merchant unknown")
		return err
	}
	return nil
}

// SubmitCallPlanMerchant ..
func SubmitCallPlanMerchant(req dbmodels.CallPlanMerchants) error {
	fmt.Println(">>> CallPlanActionAddOrEdit - SubmitCallPlanMerchant - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	var err error
	err = Dbcon.Exec("update call_plan_merchants set amount = ? , effective_call = ? , updated_at = ? , status = ? , merchant_status = ? , clock_time = ?, notes = ? where id = ?",
		req.Amount, req.EffectiveCall, req.UpdatedAt, req.Status, req.MerchantStatus, req.ClockOut, req.Notes, req.Id).Error
	if err != nil {
		sugarLogger.Error("Failed connect to database SFA when updating call plan merchant")
		return err
	}
	return nil
}

// GetCallPlanMerchantListById ..
func GetCallPlanMerchantListById(id int64) ([]dbmodels.CallPlanMerchants, error) {
	fmt.Println(">>> CallPlanActionSubmit - GetCallPlanMerchantListById - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	res := []dbmodels.CallPlanMerchants{}

	var query = "SELECT * FROM call_plan_merchants " +
		"WHERE " +
		"call_plan_id =  (select call_plan_id from call_plan_merchants where id = " + fmt.Sprintf("%d", id) + ") " +
		"AND status != 'Visited'"

	err := Dbcon.Raw(query).Scan(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get call plan merchant list by id")
		return res, err
	}
	return res, nil
}

// GetCallPlanMerchantsByMerchantPhone ..
func GetCallPlanMerchantsByMerchantPhone(MerchantPhone string, callPlanId int64) (dbmodels.CallPlanMerchants, error) {
	fmt.Println(">>> CallPlanList - GetCallPlanMerchantsByCallPlanMerchantId - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	res := dbmodels.CallPlanMerchants{}

	err := Dbcon.Model(dbmodels.CallPlanMerchants{}).Where("merchant_phone = ? and call_plan_id = ? and status = 'Visited'", MerchantPhone, callPlanId).Find(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get call plan merchant by merchant phone")
		return res, err
	}
	return res, err
}

// UpdateCallPlanMerchantUnknownV23 ..
func UpdateCallPlanMerchantUnknownV23(req models.CallPlanActionMerchantUnknownReq) error {
	fmt.Println(">>> CallPlanActionMerchantUnknown - UpdateCallPlanMerchantUnknownV23 - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	var err error
	err = Dbcon.Exec("update call_plan_merchants set status = ? , merchant_status = ? , clock_time = ? , longitude = ? , latitude = ? , photo_location = ? , notes = ? where id = ?",
		req.Status, req.MerchantStatus, req.ClockOut, req.Longitude, req.Latitude, req.PhotoLocation, req.Notes, req.CallPlanMerchantId).Error
	if err != nil {
		sugarLogger.Error("Failed connect to database SFA when updating call plan merchant unknown")
		return err
	}
	return nil
}

// GetCallPlanMerchantListv23 ..
func GetCallPlanMerchantListv23(req models.CallPlanMerchantListReq) ([]models.CallPlanMerchantListResV23, error) {
	fmt.Println(">>> CallPlanMerchantList - GetCallPlanMerchantListv23 - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	res := []models.CallPlanMerchantListResV23{}

	var query = "SELECT " +
		"cpm.id call_plan_merchant_id, cpm.merchant_name, cpm.mid as merchant_id, cpm.merchant_status, cpm.merchant_phone, " +
		"sac.sales_type_id, sac.name as sales_type_name, " +
		"mt.name merchant_type_name, merchant_address, mt.priority " +
		"FROM call_plan_merchants cpm " +
		"LEFT JOIN merchant_types mt on mt.id = cpm.merchant_type_id	" +
		"LEFT JOIN sales_area_channels sac on sac.sales_type_id = cpm.sales_type_id	" +
		"WHERE " +
		"cpm.call_plan_id = " + fmt.Sprintf("%d", req.CallPlanId) +
		"AND lower(cpm.status) = lower('" + req.Status + "')"

	page := fmt.Sprintf(" OFFSET %d LIMIT %d", (req.Page-1)*req.Limit, req.Limit)

	query = query + page

	err := Dbcon.Raw(query).Scan(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get call plan merchant list")
		return res, err
	}
	return res, nil
}

// GetCallPlanMerchantListv24 ..
func GetCallPlanMerchantListv24(req models.CallPlanMerchantListReq) ([]models.CallPlanMerchantListResV24, error) {
	fmt.Println(">>> CallPlanMerchantList - GetCallPlanMerchantListv24 - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	res := []models.CallPlanMerchantListResV24{}

	var query = "SELECT " +
		"cpm.id call_plan_merchant_id, cpm.merchant_name, cpm.mid as merchant_id, cpm.merchant_status, cpm.merchant_phone, " +
		"sac.sales_type_id, sac.name as sales_type_name, " +
		"mt.name merchant_type_name, merchant_address, mt.priority, " +
		"m.address_benchmark " +
		"FROM call_plan_merchants cpm " +
		"LEFT JOIN merchant_types mt on mt.id = cpm.merchant_type_id	" +
		"LEFT JOIN sales_area_channels sac on sac.sales_type_id = cpm.sales_type_id	" +
		"LEFT JOIN merchants m on m.merchant_id = cpm.mid " +
		"WHERE " +
		"cpm.call_plan_id = " + fmt.Sprintf("%d", req.CallPlanId) +
		"AND lower(cpm.status) = lower('" + req.Status + "')"

	page := fmt.Sprintf(" OFFSET %d LIMIT %d", (req.Page-1)*req.Limit, req.Limit)

	query = query + page

	err := Dbcon.Raw(query).Scan(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get call plan merchant list")
		return res, err
	}
	return res, nil
}

// GetCallPlanMerchantListv25 ..
func GetCallPlanMerchantListv25(req models.CallPlanMerchantListReq) ([]models.CallPlanMerchantListResV24, error) {
	fmt.Println(">>> CallPlanMerchantList - GetCallPlanMerchantListv25 - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	res := []models.CallPlanMerchantListResV24{}

	var query = "SELECT " +
		"cpm.id call_plan_merchant_id, cpm.merchant_name, cpm.mid as merchant_id, cpm.merchant_status, cpm.merchant_phone, " +
		" cpm.merchant_address, cpm.address_benchmark, " +
		" sac.sales_type_id, sac.name as sales_type_name, " +
		" mt.name merchant_type_name,  mt.priority " +
		"FROM call_plan_merchants cpm " +
		"LEFT JOIN merchant_types mt on mt.id = cpm.merchant_type_id	" +
		"LEFT JOIN sales_area_channels sac on sac.sales_type_id = cpm.sales_type_id	" +
		"WHERE " +
		"cpm.call_plan_id = " + fmt.Sprintf("%d", req.CallPlanId) +
		"AND lower(cpm.status) = lower('" + req.Status + "')"

	page := fmt.Sprintf(" OFFSET %d LIMIT %d", (req.Page-1)*req.Limit, req.Limit)

	query = query + page

	err := Dbcon.Raw(query).Scan(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get call plan merchant list")
		return res, err
	}
	return res, nil
}

// GetMerchantCallPlanByMerchantPhone ..
func GetMerchantCallPlanByMerchantPhone(merchantPhone string) (models.CallPlanVisitMerchantRes, error) {
	fmt.Println(">>> CallPlanVisitCheckMerchantPhone/CallPlanVisitCheckQRIS/MerchantDetailQRIS/CallPlanActionAddOrEdit - GetMerchantCallPlanByMerchantPhone - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	res := models.CallPlanVisitMerchantRes{}

	var err error
	err = Dbcon.Table("call_plan_merchants cpm").
		Select("cpm.merchant_id as id_merchant, cpm.merchant_name, cpm.mid as merchant_id, cpm.merchant_address, cpm.merchant_phone, cpm.merchant_type_id, cpm.mpan, cpm.merchant_type_name").
		Where("cpm.merchant_phone = ?", merchantPhone).First(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get merchant by merchant phone")
		return res, err
	}
	return res, nil
}