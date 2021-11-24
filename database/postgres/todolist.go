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

// GetTodolistByMerchantIDAndStatusOpenLate ..
func GetTodolistByMerchantIDAndStatusOpenLate(merchantId string) ([]dbmodels.TodoLists, error) {
	fmt.Println(">>> CallPlanMerchantList - GetTodolistByMerchantIDAndStatusOpenLate - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	res := []dbmodels.TodoLists{}

	var query = "select * from todolists where mid = '" + merchantId + "' and status in ('Open', 'Late')"

	err := Dbcon.Raw(query).Scan(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get todolist merchant status Open Late")
		return res, err
	}
	return res, nil
}

// GetTodolistByMerchantIDAndStatusPending ..
func GetTodolistByMerchantIDAndStatusPending(merchantId string) ([]dbmodels.TodoLists, error) {
	fmt.Println(">>> CallPlanMerchantList - GetTodolistByMerchantIDAndStatusPending - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	res := []dbmodels.TodoLists{}

	var query = "select * from todolists where mid = '" + merchantId + "' and status in ('Pending')"

	err := Dbcon.Raw(query).Scan(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get todolist merchant status Pending")
		return res, err
	}
	return res, nil
}

// GetListTodolist ..
func GetListTodolist(req models.CallPlanTodolistListReq) ([]models.CallPlanTodolistListRes, error) {
	fmt.Println(">>> CallPlanMerchantList - GetListTodolist - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	resp := []models.CallPlanTodolistListRes{}
	var query = "select distinct t.id as id, t.village_id, m.longitude, m.latitude, m.phone_number, " +
		" t.merchant_name, t.task_date , t.pending_task_date, m.address as merchant_address, t.mid as merchant_id, tc.name as name_category, t.status, " +
		" (SELECT description FROM todolist_histories where todolist_id = t.id ORDER BY id desc limit 1) as reason " +
		" from todolists t " +
		" left join todolist_categories tc on tc.id = t.todolist_category_id " +
		" left join merchants m on m.merchant_id = t.mid " +
		" left join todolist_histories th on th.todolist_id = t.id " +
		" where " +
		" m.phone_number = '" + req.MerchantPhone + "' "

	order := " ORDER BY t.task_date desc "

	and := " AND "

	if req.Status[0] != "" {
		query = query
	}

	paramString := []string{}

	if req.Status[0] != "" {

		arr := []string{}
		for _, val := range req.Status {
			val = "'" + val + "'"
			arr = append(arr, val)
		}
		str := strings.Join(arr[:], ",")

		fmt.Println("statusStr : ", str)

		q := "t.status in ( " + str + " )"
		paramString = append(paramString, fmt.Sprintf(q))
	}

	paramFilter := strings.Join(paramString[:], " AND ")

	if paramFilter != "" {
		paramFilter = and + paramFilter
	}

	page := fmt.Sprintf(" OFFSET %d LIMIT %d", (req.Page-1)*req.Limit, req.Limit)

	query = query + paramFilter + order + page

	sql := Dbcon.Raw(query).Scan(&resp)
	if sql.Error != nil {
		sugarLogger.Info("Failed connect to database SFA when get list todolist")
		return resp, sql.Error
	}

	return resp, nil
}

// GetDetailTodolist ..
func GetDetailTodolist(req models.TodolistDetailReq) (res []models.TodolistDetailDBRes, err error) {
	fmt.Println(">>> TodolistDetail - GetDetailTodolist - Postgres <<<")
	sugarLogger := ottologger.GetLogger()

	var query = " select distinct ts.code, ts.id as todolist_sub_category_id, ts.name, " +
		" t.todolist_category_id, t.merchant_name, t.task_date, t.created_at, t.updated_at, t.mid as merchant_id, t.status, t.id as todolist_id, " +
		" m.address as merchant_address, m.id as id_merchant, m.phone_number as merchant_phone, " +
		" tc.id as todolist_category_id, tc.name as name_category,  " +
		" ta.id as task_id, ta.file_edukasi as link, " +
		" t.notes as reason, " +
		" lt.label_type " +
		" from todolists t " +
		" left join merchants m on m.merchant_id = t.mid " +
		" left join todolist_categories tc on tc.id = t.todolist_category_id " +
		" left join tasks ta on ta.todolist_id = t.id " +
		" left join todolist_sub_categories ts on ts.id = ta.todolist_sub_category_id " +
		" left join label_tasks lt on lt.sub_category_id = ts.id " +
		" left join todolist_histories th on th.todolist_id = t.id " +
		" where m.phone_number = '" + req.MerchantPhone + "'" +
		" and t.id = " + req.TodolistId +
		" and lt.step = 1 "

	if req.TodolistCategoryId == 6 {
		query = " select distinct ts.code, ts.id as todolist_sub_category_id, ts.name, " +
			" t.todolist_category_id, t.merchant_name, t.task_date, t.created_at, t.updated_at, t.mid as merchant_id, t.status, t.id as todolist_id," +
			" m.address as merchant_address, m.id as id_merchant, m.customer_code, m.phone_number as merchant_phone, " +
			" tc.id as todolist_category_id, tc.name as name_category,  " +
			" ta.id as task_id, ta.file_edukasi as link, " +
			" t.notes as reason, " +
			" lt.label_type " +
			" from todolists t " +
			" left join merchant_new_recruitments m on m.id = t.merchant_new_recruitment_id " +
			" left join todolist_categories tc on tc.id = t.todolist_category_id " +
			" left join tasks ta on ta.todolist_id = t.id " +
			" left join todolist_sub_categories ts on ts.id = ta.todolist_sub_category_id " +
			" left join label_tasks lt on lt.sub_category_id = ts.id " +
			" left join todolist_histories th on th.todolist_id = t.id " +
			" where (m.phone_number = '" + req.MerchantPhone + "' or m.customer_code = '" + req.CustomerCode + "')" +
			" and t.id = " + req.TodolistId +
			" and lt.step = 1 "
		//if req.MerchantPhone != "" {
		//	query = " select distinct ts.code, ts.id as todolist_sub_category_id, ts.name, " +
		//		" t.todolist_category_id, t.merchant_name, t.task_date, t.created_at, t.updated_at, t.mid as merchant_id, t.status, t.id as todolist_id, t.merchant_phone, t.customer_code," +
		//		" m.address as merchant_address, m.id as id_merchant, " +
		//		" tc.id as todolist_category_id, tc.name as name_category,  " +
		//		" ta.id as task_id, ta.file_edukasi as link, " +
		//		" t.notes as reason, " +
		//		" lt.label_type " +
		//		" from todolists t " +
		//		" left join merchant_new_recruitments m on m.phone_number = t.merchant_phone " +
		//		" left join todolist_categories tc on tc.id = t.todolist_category_id " +
		//		" left join tasks ta on ta.todolist_id = t.id " +
		//		" left join todolist_sub_categories ts on ts.id = ta.todolist_sub_category_id " +
		//		" left join label_tasks lt on lt.sub_category_id = ts.id " +
		//		" left join todolist_histories th on th.todolist_id = t.id " +
		//		" where m.phone_number = '" + req.MerchantPhone + "'" +
		//		" and t.id = " + req.TodolistId +
		//		" and lt.step = 1 "
		//}else {
		//	query = " select distinct ts.code, ts.id as todolist_sub_category_id, ts.name, " +
		//		" t.todolist_category_id, t.merchant_name, t.task_date, t.created_at, t.updated_at, t.mid as merchant_id, t.status, t.id as todolist_id, t.merchant_phone, t.customer_code," +
		//		" m.address as merchant_address, m.id as id_merchant, " +
		//		" tc.id as todolist_category_id, tc.name as name_category,  " +
		//		" ta.id as task_id, ta.file_edukasi as link, " +
		//		" t.notes as reason, " +
		//		" lt.label_type " +
		//		" from todolists t " +
		//		" left join merchant_new_recruitments m on m.customer_code = t.customer_code  " +
		//		" left join todolist_categories tc on tc.id = t.todolist_category_id " +
		//		" left join tasks ta on ta.todolist_id = t.id " +
		//		" left join todolist_sub_categories ts on ts.id = ta.todolist_sub_category_id " +
		//		" left join label_tasks lt on lt.sub_category_id = ts.id " +
		//		" left join todolist_histories th on th.todolist_id = t.id " +
		//		" where m.customer_code  = '" + req.CustomerCode + "'" +
		//		" and t.id = " + req.TodolistId +
		//		" and lt.step = 1 "
		//}
	}

	sql := Dbcon.Raw(query).Scan(&res)
	if sql.Error != nil {
		sugarLogger.Info("Failed connect to database SFA when get detail todolist")
		return res, sql.Error
	}

	return res, nil
}

// UpdateTodolistMerchantNotFound ..
func UpdateTodolistMerchantNotFound(req dbmodels.TodoLists) error {
	fmt.Println(">>> CallPlanTodolistMerchantNotFound - UpdateTodolistMerchantNotFound - Postgres <<<")
	sugarLogger := ottologger.GetLogger()

	err := Dbcon.Exec("update todolists set status = ? , updated_at = ? , action_date = ? , sales_phone = ? , longitude = ? , latitude = ?  where id = ?",
		req.Status, req.UpdatedAt, req.ActionDate, req.SalesPhone, req.Longitude, req.Latitude, req.ID).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when Update Todolist Merchant Not Found")
		return err
	}
	return nil
}

// GetListTodolistByMerchantId ..
func GetListTodolistByMerchantId(req models.CallPlanTodolistListReq) ([]models.CallPlanTodolistListRes, error) {
	fmt.Println(">>> CallPlanMerchantList - GetListTodolist - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	resp := []models.CallPlanTodolistListRes{}
	var query = "select distinct t.id as id, t.village_id, m.longitude, m.latitude, m.phone_number, " +
		" t.merchant_name, t.task_date , t.pending_task_date, m.address as merchant_address, t.mid as merchant_id, tc.name as name_category, t.status, " +
		" (SELECT description FROM todolist_histories where todolist_id = t.id ORDER BY id desc limit 1) as reason " +
		" from todolists t " +
		" left join todolist_categories tc on tc.id = t.todolist_category_id " +
		" left join merchants m on m.merchant_id = t.mid " +
		" left join todolist_histories th on th.todolist_id = t.id " +
		" where " +
		" m.phone_number = '" + req.MerchantPhone + "' "

	order := " ORDER BY t.task_date desc "

	and := " AND "

	if req.Status[0] != "" {
		query = query
	}

	paramString := []string{}

	if req.Status[0] != "" {

		arr := []string{}
		for _, val := range req.Status {
			val = "'" + val + "'"
			arr = append(arr, val)
		}
		str := strings.Join(arr[:], ",")

		fmt.Println("statusStr : ", str)

		q := "t.status in ( " + str + " )"
		paramString = append(paramString, fmt.Sprintf(q))
	}

	paramFilter := strings.Join(paramString[:], " AND ")

	if paramFilter != "" {
		paramFilter = and + paramFilter
	}

	page := fmt.Sprintf(" OFFSET %d LIMIT %d", (req.Page-1)*req.Limit, req.Limit)

	query = query + paramFilter + order + page

	sql := Dbcon.Raw(query).Scan(&resp)
	if sql.Error != nil {
		sugarLogger.Info("Failed connect to database SFA when get list todolist")
		return resp, sql.Error
	}

	return resp, nil
}

// TodolistPost ..
func TodolistPost(req models.PostTodolistToDBReq) error {
	fmt.Println(">>> TodolistPost - TodolistPost - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	err := Dbcon.Exec("INSERT INTO follow_ups (label, content_type, body, task_id, created_at, updated_at) "+
		"VALUES ( ? , ? , ? , ? , now(), now());",
		req.Label, req.ContentType, req.Body, req.TaskID).Error
	if err != nil {
		sugarLogger.Error("Failed Post Todolist")
		return err
	}
	return nil
}

// UpdateStatusDoneTodolist ..
func UpdateStatusDoneTodolist(salesPhone string, taskDate string, status string, todolistID int64, long string, lat string) error {
	fmt.Println(">>> TodolistPost - UpdateStatusDoneTodolist - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	err := Dbcon.Exec("update todolists set status = ? , action_date = now(), sales_phone = ?, longitude = ?, latitude = ?  where id = ? ", status, salesPhone, long, lat, todolistID).Error
	if err != nil {
		sugarLogger.Error("Failed Update Status Todolist (Done)")
		return err
	}
	return nil
}

// UpdateStatusPendingTodolist ..
func UpdateStatusPendingTodolist(salesPhone string, taskDate string, status string, todolistID int64, long string, lat string) error {
	fmt.Println(">>> TodolistPost - UpdateStatusPendingTodolist - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	err := Dbcon.Exec("update todolists set status = ? , action_date = now(), pending_task_date = ?, sales_phone = ?, longitude = ?, latitude = ?  where id = ? ", status, taskDate, salesPhone, long, lat, todolistID).Error
	if err != nil {
		sugarLogger.Error("Failed Update Status Todolist (Pending)")
		return err
	}
	return nil
}

// PostHistoryTodolist ..
func PostHistoryTodolist(req models.PostHistoryTodolistReq) error {
	fmt.Println(">>> TodolistPost - PostHistoryTodolist - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	err := Dbcon.Exec("INSERT INTO todolist_histories (todolist_id, description, status, new_task_date, old_task_date, created_at, updated_at, longitude, latitude) "+
		"VALUES ( ? , ? , 'Pending' , ? , ?, now(), now(), ?, ?);",
		req.TodolistID, req.Reason, req.NewTaskDate, req.OldTaskDate, req.Longitude, req.Latitude).Error
	if err != nil {
		sugarLogger.Error("Failed Post Todolist")
		return err
	}
	return nil
}

// MerchantNotFoundList ..
func MerchantNotFoundList(req models.TodolistMerchantNotFoundListReq) (res []models.TodolistMerchantNotFoundListRes, err error) {
	fmt.Println(">>> CallPlanMerchantList - MerchantNotFoundList - Postgres <<<")
	sugarLogger := ottologger.GetLogger()

	query := "select * from qr_problems"
	if req.TodolistCategoryId == 6 {
		query = "select * from merchant_new_recruitment_not_found_lists"
	}

	err = Dbcon.Raw(query).Scan(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get merchant not found list")
		return res, err
	}
	return res, nil
}

// TodolistList ..
func TodolistList(req models.TodolistListReq, villageId []models.TodolistVillageID, phoneNumberSales string) (res []models.TodolistListDBRes, err error) {
	fmt.Println(">>> TodolistList - TodolistList - Postgres <<<")
	sugarLogger := ottologger.GetLogger()

	villArr := []string{}
	for _, val := range villageId {
		villArr = append(villArr, val.VillageId)
	}
	villStr := strings.Join(villArr[:], ",")

	query := "SELECT distinct t.id as id, " +
		" t.mid as merchant_id, t.merchant_name, t.sales_phone, t.village_id, t.task_date, t.pending_task_date, t.status, t.merchant_new_recruitment_id as merchant_new_rec_id, " +
		//" m.longitude, m.latitude, m.phone_number, m.address as merchant_address, " +
		" tc.id as todolist_category_id, tc.name as name_category, " +
		" (SELECT description FROM todolist_histories where todolist_id = t.id ORDER BY id desc limit 1) AS reason " +
		" FROM todolists t " +
		" LEFT JOIN todolist_categories tc on tc.id = t.todolist_category_id " +
		//" LEFT JOIN merchants m on m.merchant_id = t.mid " +
		" LEFT JOIN todolist_histories th on th.todolist_id = t.id "
	where := " WHERE "
	order := " ORDER BY t.task_date desc "
	and := " AND "
	con1 := " t.village_id in (" + villStr + ") and t.sales_phone = '' "
	con2 := " t.sales_phone = '" + phoneNumberSales + "' "

	if req.Status[0] != "" || req.TaskDateStart != "" || req.TaskDateEnd != "" || req.VillageID[0] != "" || req.CategoryID[0] != "" || req.ClusterID[0] != "" || req.Keyword != "" {
		query = query + where
	}

	paramString := []string{}
	if req.Status[0] != "" {
		arr := []string{}
		for _, val := range req.Status {
			val = "'" + val + "'"
			arr = append(arr, val)
		}
		str := strings.Join(arr[:], ",")
		paramString = append(paramString, fmt.Sprintf("t.status in ( "+str+" )"))
	}

	if req.TaskDateStart != "" && req.TaskDateEnd != "" {
		paramString = append(paramString, fmt.Sprintf(" t.task_date between '%s' and '%s'", req.TaskDateStart+" 00:00:00", req.TaskDateEnd+" 23:59:59"))
	}

	if req.VillageID[0] != "" && req.ClusterID[0] == "" {
		arr := []string{}
		for _, val := range req.VillageID {
			arr = append(arr, val)
		}
		str := strings.Join(arr[:], ",")
		paramString = append(paramString, fmt.Sprintf(" t.village_id in ( %s ) ", str))
	}

	if req.CategoryID[0] != "" {
		arr := []string{}
		for _, val := range req.CategoryID {
			arr = append(arr, val)
		}
		str := strings.Join(arr[:], ",")
		paramString = append(paramString, fmt.Sprintf(" t.todolist_category_id in ( %s ) ", str))
	}

	paramFilter := strings.Join(paramString[:], " AND ")

	if paramFilter != "" {
		paramFilter = and + paramFilter
	}

	if req.ClusterID[0] != "" {
		arr := []string{}
		for _, val := range req.ClusterID {
			arr = append(arr, fmt.Sprintf("'"+val+"'"))
		}
		str := strings.Join(arr[:], ",")
		paramFilter = paramFilter + and + fmt.Sprintf("m.cluster_id in ( %s ) ", str)
	}

	conKeyword := ""
	if req.Keyword != "" {
		conKeyword = and + " LOWER(t.merchant_name) like '%" + strings.ToLower(req.Keyword) + "%'"
	}

	page := fmt.Sprintf(" OFFSET %d LIMIT %d", (req.Page-1)*req.Limit, req.Limit)

	query = query +
		"(" +
		" ( " + con1 + paramFilter + ") " +
		" or " +
		" ( " + con2 + paramFilter + ") " +
		") " + conKeyword +
		order + page

	err = Dbcon.Raw(query).Scan(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get todolist list")
		return res, err
	}
	return res, nil
}

// MerchantTodolistList ..
func MerchantTodolistList(req models.TodolistListDBRes) (res models.MerchantTodolistListDBRes, err error) {
	fmt.Println(">>> TodolistList - MerchantTodolistList - Postgres <<<")
	sugarLogger := ottologger.GetLogger()

	query := "SELECT longitude, latitude, phone_number, address FROM merchants WHERE merchant_id = '" + req.MerchantID + "'"
	if req.TodolistCategoryId == 6 {
		query = "SELECT longitude, latitude, phone_number, customer_code, address FROM merchant_new_recruitments WHERE id = '" + fmt.Sprintf("%d", req.MerchantNewRecId) + "'"
	}

	err = Dbcon.Raw(query).Scan(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get merchant not found list")
		return res, err
	}
	return res, nil
}

// CekStatusTaskTodolist ..
func CekStatusTaskTodolist(taskId int64) (res models.CheckStatusTaskTodolistRes, err error) {
	fmt.Println(">>> TodolistDetail - CekStatusTaskTodolist - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	var query = " select distinct task_id from follow_ups where task_id = " + strconv.Itoa(int(taskId))
	err = Dbcon.Raw(query).Scan(&res).Error
	if err != nil {
		sugarLogger.Info("Failed check status todolist")
		return res, err
	}
	return res, nil
}

// IDMMerchantTodolistList ..
func IDMMerchantTodolistList(req models.TodolistListDBRes) (res models.MerchantTodolistListDBRes, err error) {
	fmt.Println(">>> IDMTodolistList - IDMMerchantTodolistList - Postgres <<<")
	sugarLogger := ottologger.GetLogger()

	query := "SELECT longitude, latitude, phone_number, customer_code, address FROM merchants WHERE merchant_id = '" + req.MerchantID + "'"
	if req.TodolistCategoryId == 6 {
		query = "SELECT longitude, latitude, phone_number, customer_code, address FROM merchant_new_recruitments WHERE id = '" + fmt.Sprintf("%d", req.MerchantNewRecId) + "'"
	}

	err = Dbcon.Raw(query).Scan(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get merchant not found list")
		return res, err
	}
	return res, nil
}

// IDMGetDetailTodolist ..
func IDMGetDetailTodolist(req models.TodolistDetailReq) (res []models.TodolistDetailDBRes, err error) {
	fmt.Println(">>> IDMTodolistDetail - IDMGetDetailTodolist - Postgres <<<")
	sugarLogger := ottologger.GetLogger()

	var query = " select distinct ts.code, ts.id as todolist_sub_category_id, ts.name, " +
		" t.todolist_category_id, t.merchant_name, t.task_date, t.created_at, t.updated_at, t.mid as merchant_id, t.status, t.id as todolist_id, t.notes as notes, " +
		" m.address as merchant_address, m.id as id_merchant, m.phone_number as merchant_phone, m.id_card, " +
		" tc.id as todolist_category_id, tc.name as name_category,  " +
		" ta.id as task_id, ta.file_edukasi as link, " +
		" t.notes as reason, " +
		" lt.label_type " +
		" from todolists t " +
		" left join merchants m on m.merchant_id = t.mid " +
		" left join todolist_categories tc on tc.id = t.todolist_category_id " +
		" left join tasks ta on ta.todolist_id = t.id " +
		" left join todolist_sub_categories ts on ts.id = ta.todolist_sub_category_id " +
		" left join label_tasks lt on lt.sub_category_id = ts.id " +
		" left join todolist_histories th on th.todolist_id = t.id " +
		" where t.id = " + req.TodolistId +
		" and lt.step = 1 "

	if req.TodolistCategoryId == 6 {
		query = " select distinct ts.code, ts.id as todolist_sub_category_id, ts.name, " +
			" t.todolist_category_id, t.merchant_name, t.task_date, t.created_at, t.updated_at, t.mid as merchant_id, t.status, t.id as todolist_id, t.notes as notes, " +
			" m.address as merchant_address, m.id as id_merchant, m.customer_code, m.phone_number as merchant_phone, m.id_card, " +
			" tc.id as todolist_category_id, tc.name as name_category,  " +
			" ta.id as task_id, ta.file_edukasi as link, " +
			" t.notes as reason, " +
			" lt.label_type " +
			" from todolists t " +
			" left join merchant_new_recruitments m on m.id = t.merchant_new_recruitment_id " +
			" left join todolist_categories tc on tc.id = t.todolist_category_id " +
			" left join tasks ta on ta.todolist_id = t.id " +
			" left join todolist_sub_categories ts on ts.id = ta.todolist_sub_category_id " +
			" left join label_tasks lt on lt.sub_category_id = ts.id " +
			" left join todolist_histories th on th.todolist_id = t.id " +
			" where t.id = " + req.TodolistId +
			" and lt.step = 1 "
	}

	sql := Dbcon.Raw(query).Scan(&res)
	if sql.Error != nil {
		sugarLogger.Info("Failed connect to database SFA when get detail todolist")
		return res, sql.Error
	}

	return res, nil
}

// MerchantNotFoundListV24 ..
func MerchantNotFoundListV24(req models.TodolistMerchantNotFoundListReq) (res []models.TodolistMerchantNotFoundListRes, err error) {
	fmt.Println(">>> CallPlanMerchantList - MerchantNotFoundListV24 - Postgres <<<")
	sugarLogger := ottologger.GetLogger()

	query := "select * from qr_problems where status = 'Active'"
	if req.TodolistCategoryId == 6 {
		query = "select * from merchant_new_recruitment_not_found_lists where status = 'Active'"
	}

	err = Dbcon.Raw(query).Scan(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get merchant not found list")
		return res, err
	}
	return res, nil
}

// GetVillageListForTodolistFilter ..
func GetVillageListForTodolistFilter(salesID int, salesPhone string, keyword string) (res []dbmodels.Villages, err error) {
	fmt.Println(">>> TodolistFilterVillageList - GetVillageListForTodolistFilter - Postgres <<<")
	sugarLogger := ottologger.GetLogger()

	query := " SELECT distinct v.* " +
		" FROM villages v " +
		" LEFT JOIN todolists t on t.village_id = v.id " +
		" WHERE t.village_id in " +
		" (SELECT village_id FROM sub_areas_villages WHERE sub_area_id in " +
		" (SELECT scs.sub_area_id FROM sales_area_channels_sub_areas scs LEFT JOIN sales_area_channels sac on sac.id = scs.sales_area_channel_id WHERE sac.sales_type_id in " + //" (SELECT sub_area_id FROM sales_area_channels_sub_areas WHERE sales_area_channel_id in " +
		" (SELECT st.id FROM sales_types st " +//" (SELECT st.id FROM sales_types_salesmen sts " +
		" LEFT JOIN sales_types_salesmen sts on st.id = sts.sales_type_id WHERE sts.salesman_id = " + strconv.Itoa(salesID) + //" LEFT JOIN sales_types st on st.id = sts.sales_type_id WHERE salesman_id = " + strconv.Itoa(salesID) +
		" )" +
		" )" +
		" UNION " +
		" SELECT distinct v.id " +
		" FROM todolists t " +
		" LEFT JOIN villages v on v.id = t.village_id " +
		" WHERE t.sales_phone = '" + salesPhone + "'" +
		" ) "

	if keyword != "" {
		query = query + " AND lower(v.name) like '%" + strings.ToLower(keyword) + "%'"
	}

	err = Dbcon.Raw(query).Scan(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get village list")
		return res, err
	}
	return res, nil
}

// TodolistListV24 ..
func TodolistListV24(req models.TodolistListReq, villageId []models.TodolistVillageID, phoneNumberSales string) (res []models.TodolistListDBResV24, err error) {
	fmt.Println(">>> TodolistList - TodolistListV24 - Postgres <<<")
	sugarLogger := ottologger.GetLogger()

	villArr := []string{}
	for _, val := range villageId {
		villArr = append(villArr, val.VillageId)
	}
	villStr := strings.Join(villArr[:], ",")

	//query := "SELECT distinct t.id as id, " +
	//	" t.mid as merchant_id, t.merchant_name, t.sales_phone, t.village_id, t.task_date, t.pending_task_date, t.status, t.merchant_new_recruitment_id as merchant_new_rec_id, " +
	//	" sac.sales_type_id, sac.name as sales_type_name, " +
	//	" tc.id as todolist_category_id, tc.name as name_category, " +
	//	" (SELECT description FROM todolist_histories where todolist_id = t.id ORDER BY id desc limit 1) AS reason " +
	//	" FROM todolists t " +
	//	" LEFT JOIN merchants m on tc.id = t.todolist_category_id " +
	//	" LEFT JOIN todolist_categories tc on tc.id = t.todolist_category_id " +
	//	" LEFT JOIN todolist_histories th on th.todolist_id = t.id "
	//where := " WHERE "
	//order := " ORDER BY t.task_date desc "
	//and := " AND "
	//con1 := " t.village_id in (" + villStr + ") and t.sales_phone = '' and m.sales_type_id = " + strconv.Itoa(salesTypeId)
	//con2 := " t.sales_phone = '" + phoneNumberSales + "' "

	query := "SELECT distinct t.id as id, " +
		" t.mid as merchant_id, t.merchant_name, t.sales_phone, t.village_id, t.task_date, t.pending_task_date, t.status, t.merchant_new_recruitment_id as merchant_new_rec_id, t.sales_phone, " +
	//" m.longitude, m.latitude, m.phone_number, m.address as merchant_address, " +
		" tc.id as todolist_category_id, tc.name as name_category, " +
		" (SELECT description FROM todolist_histories where todolist_id = t.id ORDER BY id desc limit 1) AS reason " +
		" FROM todolists t " +
		" LEFT JOIN todolist_categories tc on tc.id = t.todolist_category_id " +
	//" LEFT JOIN merchants m on m.merchant_id = t.mid " +
		" LEFT JOIN todolist_histories th on th.todolist_id = t.id "
	where := " WHERE "
	order := " ORDER BY t.task_date desc "
	and := " AND "
	con1 := " t.village_id in (" + villStr + ") and t.sales_phone = '' "
	con2 := " t.sales_phone = '" + phoneNumberSales + "' "
	if req.Status[0] != "" || req.TaskDateStart != "" || req.TaskDateEnd != "" || req.VillageID[0] != "" || req.CategoryID[0] != "" || req.ClusterID[0] != "" || req.Keyword != "" {
		query = query + where
	}

	paramString := []string{}
	if req.Status[0] != "" {
		arr := []string{}
		for _, val := range req.Status {
			val = "'" + val + "'"
			arr = append(arr, val)
		}
		str := strings.Join(arr[:], ",")
		paramString = append(paramString, fmt.Sprintf("t.status in ( "+str+" )"))
	}

	if req.TaskDateStart != "" && req.TaskDateEnd != "" {
		paramString = append(paramString, fmt.Sprintf(" t.task_date between '%s' and '%s'", req.TaskDateStart+" 00:00:00", req.TaskDateEnd+" 23:59:59"))
	}

	if req.VillageID[0] != "" && req.ClusterID[0] == "" {
		arr := []string{}
		for _, val := range req.VillageID {
			arr = append(arr, val)
		}
		str := strings.Join(arr[:], ",")
		paramString = append(paramString, fmt.Sprintf(" t.village_id in ( %s ) ", str))
	}

	if req.CategoryID[0] != "" {
		arr := []string{}
		for _, val := range req.CategoryID {
			arr = append(arr, val)
		}
		str := strings.Join(arr[:], ",")
		paramString = append(paramString, fmt.Sprintf(" t.todolist_category_id in ( %s ) ", str))
	}

	paramFilter := strings.Join(paramString[:], " AND ")

	if paramFilter != "" {
		paramFilter = and + paramFilter
	}

	if req.ClusterID[0] != "" {
		arr := []string{}
		for _, val := range req.ClusterID {
			arr = append(arr, fmt.Sprintf("'"+val+"'"))
		}
		str := strings.Join(arr[:], ",")
		paramFilter = paramFilter + and + fmt.Sprintf("m.cluster_id in ( %s ) ", str)
	}

	conKeyword := ""
	if req.Keyword != "" {
		conKeyword = and + " LOWER(t.merchant_name) like '%" + strings.ToLower(req.Keyword) + "%'"
	}

	page := fmt.Sprintf(" OFFSET %d LIMIT %d", (req.Page-1)*req.Limit, req.Limit)

	query = query +
		"(" +
		" ( " + con1 + paramFilter + ") " +
		" or " +
		" ( " + con2 + paramFilter + ") " +
		") " + conKeyword +
		order + page

	err = Dbcon.Raw(query).Scan(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get todolist list")
		return res, err
	}
	return res, nil
}

// MerchantTodolistListV24 ..
func MerchantTodolistListV24(req models.TodolistListDBResV24) (res models.MerchantTodolistListDBResV24, err error) {
	fmt.Println(">>> TodolistList - MerchantTodolistListV24 - Postgres <<<")
	sugarLogger := ottologger.GetLogger()

	query := "SELECT longitude, latitude, phone_number, address, business_type, address_benchmark, sales_type_id " +
		" FROM merchants " +
		//" LEFT JOIN merchant_business_types mbt on code " +
		" WHERE merchant_id = '" + req.MerchantID + "'"
	if req.TodolistCategoryId == 6 {
		query = "SELECT longitude, latitude, phone_number, customer_code, address, sales_type_id FROM merchant_new_recruitments WHERE id = '" + fmt.Sprintf("%d", req.MerchantNewRecId) + "'"
	}

	err = Dbcon.Raw(query).Scan(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get merchant not found list")
		return res, err
	}
	return res, nil
}

// GetCountTodolist ..
func GetCountTodolist(villageId []models.TodolistVillageID, phoneSales, salesTypeId string) (models.Count, error) {
	resp := models.Count{}

	villArr := []string{}
	for _, val := range villageId {
		villArr = append(villArr, val.VillageId)
	}
	villStr := strings.Join(villArr[:], ",")
	var query = " SELECT SUM(a) as count FROM " +
				" ( " +
		" (SELECT  count (distinct t.*) a FROM todolists t " + "" +
		" LEFT JOIN merchants m on m.merchant_id = t.mid " +
		" WHERE " +
		" (t.sales_phone = '" + phoneSales + "' AND t.status in ( 'Open', 'Pending', 'Late' ) and t.todolist_category_id != 6)" +
		" OR" +
		" (t.village_id in (" + villStr + ") AND t.sales_phone = ''  AND t.status in ( 'Open', 'Pending', 'Late' ) and t.todolist_category_id != 6 and m.sales_type_id = "+ salesTypeId +" )" +
		" )" +
		" UNION ALL " +
		" ( SELECT count (distinct t.*) a FROM todolists t " +
		" LEFT JOIN merchant_new_recruitments m on m.id = t.merchant_new_recruitment_id" +
		" WHERE" +
		" (t.sales_phone = '" + phoneSales + "' AND t.status in ( 'Open', 'Pending', 'Late' ) AND t.todolist_category_id = 6 )" +
		" OR" +
		" (t.village_id in (" + villStr + ") AND t.sales_phone = ''  AND t.status in ( 'Open', 'Pending', 'Late' ) AND t.todolist_category_id = 6 and m.sales_type_id = "+ salesTypeId +" )" +
		" ) " +
		" ) x"
	sql := Dbcon.Raw(query).Scan(&resp)
	if sql.Error != nil {
		logs.Error("Failed")
		return resp, sql.Error
	}

	return resp, nil
}

// CallPlanTodolistList ..
func CallPlanTodolistList(req models.CallPlanTodolistListReq) ([]models.CallPlanTodolistListResv23, error) {
	fmt.Println(">>> CallPlanTodolistList - CallPlanTodolistList - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	resp := []models.CallPlanTodolistListResv23{}
	var query = "select distinct t.id as id, t.village_id, " +
		" t.merchant_phone phone_number, t.address as merchant_address, t.address_benchmark, " +
		" t.merchant_name, t.task_date , t.pending_task_date, t.mid as merchant_id, tc.name as name_category, t.status, " +
		" (SELECT description FROM todolist_histories where todolist_id = t.id ORDER BY id desc limit 1) as reason " +
		" from todolists t " +
		" left join todolist_categories tc on tc.id = t.todolist_category_id " +
		//" left join merchants m on m.merchant_id = t.mid " +
		" left join todolist_histories th on th.todolist_id = t.id " +
		" where " +
		" t.merchant_phone = '" + req.MerchantPhone + "' "

	order := " ORDER BY t.task_date desc "

	and := " AND "

	if req.Status[0] != "" {
		query = query
	}

	paramString := []string{}

	if req.Status[0] != "" {

		arr := []string{}
		for _, val := range req.Status {
			val = "'" + val + "'"
			arr = append(arr, val)
		}
		str := strings.Join(arr[:], ",")

		fmt.Println("statusStr : ", str)

		q := "t.status in ( " + str + " )"
		paramString = append(paramString, fmt.Sprintf(q))
	}

	paramFilter := strings.Join(paramString[:], " AND ")

	if paramFilter != "" {
		paramFilter = and + paramFilter
	}

	page := fmt.Sprintf(" OFFSET %d LIMIT %d", (req.Page-1)*req.Limit, req.Limit)

	query = query + paramFilter + order + page

	sql := Dbcon.Raw(query).Scan(&resp)
	if sql.Error != nil {
		sugarLogger.Info("Failed connect to database SFA when get list todolist")
		return resp, sql.Error
	}

	return resp, nil
}

// TodolistListV25 ..
func TodolistListV25(req models.TodolistListReq, villageId []models.TodolistVillageID, phoneNumberSales string) (res []models.TodolistListDBResV24, err error) {
	fmt.Println(">>> TodolistList - TodolistListV25 - Postgres <<<")
	sugarLogger := ottologger.GetLogger()

	villArr := []string{}
	for _, val := range villageId {
		villArr = append(villArr, val.VillageId)
	}
	villStr := strings.Join(villArr[:], ",")


	query := "SELECT distinct t.id as id, " +
		" t.mid as merchant_id, t.merchant_name, t.sales_phone, t.village_id, t.task_date, t.pending_task_date, t.status, t.merchant_new_recruitment_id as merchant_new_rec_id, t.sales_phone, " +
	//" m.longitude, m.latitude, " + longitude, latitude, phone_number, address, business_type, address_benchmark, sales_type_id
		" t.merchant_phone phone_number, t.address as merchant_address, t.longitude, t.latitude, t.address_benchmark, t.sales_type_id, " +
		" tc.id as todolist_category_id, tc.name as name_category, " +
		" (SELECT description FROM todolist_histories where todolist_id = t.id ORDER BY id desc limit 1) AS reason " +
		" FROM todolists t " +
		" LEFT JOIN todolist_categories tc on tc.id = t.todolist_category_id " +
	//" LEFT JOIN merchants m on m.merchant_id = t.mid " +
		" LEFT JOIN todolist_histories th on th.todolist_id = t.id "
	where := " WHERE "
	order := " ORDER BY t.task_date desc "
	and := " AND "
	con1 := " t.village_id in (" + villStr + ") and t.sales_phone = '' "
	con2 := " t.sales_phone = '" + phoneNumberSales + "' "
	if req.Status[0] != "" || req.TaskDateStart != "" || req.TaskDateEnd != "" || req.VillageID[0] != "" || req.CategoryID[0] != "" || req.ClusterID[0] != "" || req.Keyword != "" {
		query = query + where
	}

	paramString := []string{}
	if req.Status[0] != "" {
		arr := []string{}
		for _, val := range req.Status {
			val = "'" + val + "'"
			arr = append(arr, val)
		}
		str := strings.Join(arr[:], ",")
		paramString = append(paramString, fmt.Sprintf("t.status in ( "+str+" )"))
	}

	if req.TaskDateStart != "" && req.TaskDateEnd != "" {
		paramString = append(paramString, fmt.Sprintf(" t.task_date between '%s' and '%s'", req.TaskDateStart+" 00:00:00", req.TaskDateEnd+" 23:59:59"))
	}

	if req.VillageID[0] != "" && req.ClusterID[0] == "" {
		arr := []string{}
		for _, val := range req.VillageID {
			arr = append(arr, val)
		}
		str := strings.Join(arr[:], ",")
		paramString = append(paramString, fmt.Sprintf(" t.village_id in ( %s ) ", str))
	}

	if req.CategoryID[0] != "" {
		arr := []string{}
		for _, val := range req.CategoryID {
			arr = append(arr, val)
		}
		str := strings.Join(arr[:], ",")
		paramString = append(paramString, fmt.Sprintf(" t.todolist_category_id in ( %s ) ", str))
	}

	paramFilter := strings.Join(paramString[:], " AND ")

	if paramFilter != "" {
		paramFilter = and + paramFilter
	}

	if req.ClusterID[0] != "" {
		arr := []string{}
		for _, val := range req.ClusterID {
			arr = append(arr, fmt.Sprintf("'"+val+"'"))
		}
		str := strings.Join(arr[:], ",")
		paramFilter = paramFilter + and + fmt.Sprintf("m.cluster_id in ( %s ) ", str)
	}

	conKeyword := ""
	if req.Keyword != "" {
		conKeyword = and + " LOWER(t.merchant_name) like '%" + strings.ToLower(req.Keyword) + "%'"
	}

	page := fmt.Sprintf(" OFFSET %d LIMIT %d", (req.Page-1)*req.Limit, req.Limit)

	query = query +
		"(" +
		" ( " + con1 + paramFilter + ") " +
		" or " +
		" ( " + con2 + paramFilter + ") " +
		") " + conKeyword +
		order + page

	err = Dbcon.Raw(query).Scan(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get todolist list")
		return res, err
	}
	return res, nil
}

// TodolistListV252 ..
func TodolistListV252(req models.TodolistListReq, villageId []models.TodolistVillageID, phoneNumberSales, salesTypeId string) (res []models.TodolistListDBResV24, err error) {
	fmt.Println(">>> TodolistList - TodolistListV25 - Postgres <<<")
	sugarLogger := ottologger.GetLogger()

	villArr := []string{}
	for _, val := range villageId {
		villArr = append(villArr, val.VillageId)
	}
	villStr := strings.Join(villArr[:], ",")

	queryMerchantRose := "SELECT distinct t.id as id, " +
		" t.mid as merchant_id, t.merchant_name, t.sales_phone, t.village_id, t.task_date, t.pending_task_date, " +
		" t.status, t.merchant_new_recruitment_id as merchant_new_rec_id, t.longitude, t.latitude, " +
		" tc.id as todolist_category_id, tc.name as name_category,  " +
		" t.address as merchant_address, t.merchant_phone phone_number, t.sales_type_id, t.address_benchmark, st.name as sales_type_name, " +
		" (SELECT description FROM todolist_histories where todolist_id = t.id ORDER BY id desc limit 1) AS reason " +
		" FROM todolists t " +
		" LEFT JOIN todolist_categories tc on tc.id = t.todolist_category_id " +
		" LEFT JOIN sales_types st on st.id = t.sales_type_id  " +
		" LEFT JOIN todolist_histories th on th.todolist_id = t.id "

	queryMerchantNewRec := "SELECT distinct t.id as id, " +
		" t.mid as merchant_id, t.merchant_name, t.sales_phone, t.village_id, t.task_date, t.pending_task_date, " +
		" t.status, t.merchant_new_recruitment_id as merchant_new_rec_id, mnr.longitude, mnr.latitude, " +
		" tc.id as todolist_category_id, tc.name as name_category, " +
		" mnr.address as merchant_address, mnr.phone_number phone_number, mnr.sales_type_id, t.address_benchmark, st.name as sales_type_name, " +
		" (SELECT description FROM todolist_histories where todolist_id = t.id ORDER BY id desc limit 1) AS reason " +
		" FROM todolists t " +
		" LEFT JOIN merchant_new_recruitments mnr on mnr.id = t.merchant_new_recruitment_id " +
		" LEFT JOIN todolist_categories tc on tc.id = t.todolist_category_id " +
		" LEFT JOIN sales_types st on st.id = mnr.sales_type_id  " +
		" LEFT JOIN todolist_histories th on th.todolist_id = t.id "

	where := " WHERE "
	order := " ORDER BY c.task_date desc "
	and := " AND "
	con1 := " t.village_id in (" + villStr + ") and t.sales_phone = '' "
	con2 := " t.sales_phone = '" + phoneNumberSales + "' "
	if req.Status[0] != "" || req.TaskDateStart != "" || req.TaskDateEnd != "" || req.VillageID[0] != "" || req.CategoryID[0] != "" || req.ClusterID[0] != "" || req.Keyword != "" {
		queryMerchantRose = queryMerchantRose + where
		queryMerchantNewRec = queryMerchantNewRec + where
	}

	paramString := []string{}
	if req.Status[0] != "" {
		arr := []string{}
		for _, val := range req.Status {
			val = "'" + val + "'"
			arr = append(arr, val)
		}
		str := strings.Join(arr[:], ",")
		paramString = append(paramString, fmt.Sprintf("t.status in ( "+str+" )"))
	}

	if req.TaskDateStart != "" && req.TaskDateEnd != "" {
		paramString = append(paramString, fmt.Sprintf(" t.task_date between '%s' and '%s'", req.TaskDateStart+" 00:00:00", req.TaskDateEnd+" 23:59:59"))
	}

	if req.VillageID[0] != "" && req.ClusterID[0] == "" {
		arr := []string{}
		for _, val := range req.VillageID {
			arr = append(arr, val)
		}
		str := strings.Join(arr[:], ",")
		paramString = append(paramString, fmt.Sprintf(" t.village_id in ( %s ) ", str))
	}

	if req.CategoryID[0] != "" {
		arr := []string{}
		for _, val := range req.CategoryID {
			arr = append(arr, val)
		}
		str := strings.Join(arr[:], ",")
		paramString = append(paramString, fmt.Sprintf(" t.todolist_category_id in ( %s ) ", str))
	}

	paramFilter := strings.Join(paramString[:], " AND ")

	if paramFilter != "" {
		paramFilter = and + paramFilter
	}

	//if req.ClusterID[0] != "" {
	//	arr := []string{}
	//	for _, val := range req.ClusterID {
	//		arr = append(arr, fmt.Sprintf("'"+val+"'"))
	//	}
	//	str := strings.Join(arr[:], ",")
	//	paramFilter = paramFilter + and + fmt.Sprintf("m.cluster_id in ( %s ) ", str)
	//}

	conKeyword := ""
	if req.Keyword != "" {
		conKeyword = where + " LOWER(c.merchant_name) like '%" + strings.ToLower(req.Keyword) + "%'"
	}

	page := fmt.Sprintf(" OFFSET %d LIMIT %d", (req.Page-1)*req.Limit, req.Limit)

	query := " SELECT * FROM  ( " +
			"( "+ queryMerchantRose +
			" ( " + con1 + " AND t.todolist_category_id != 6 and t.sales_type_id = " + salesTypeId + " " + paramFilter + ") " +
			" OR " +
			" ( " + con2 + " AND t.todolist_category_id != 6 " + paramFilter + ") " +
			")" +
			" UNION " +
			"( "+ queryMerchantNewRec +
			" ( " + con1 + " AND t.todolist_category_id = 6 and mnr.sales_type_id = " + salesTypeId + " " + paramFilter + ") " +
			" OR " +
			" ( " + con2 + " AND t.todolist_category_id = 6 " + paramFilter + ") " +
			")" +
			" ) c" +
			conKeyword +
			order + page

	err = Dbcon.Raw(query).Scan(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get todolist list")
		return res, err
	}
	return res, nil
}

// MerchantTodolistListV25 ..
func MerchantTodolistListV25(req models.TodolistListDBResV24) (res models.MerchantTodolistListDBResV24, err error) {
	fmt.Println(">>> TodolistList - MerchantTodolistListV25 - Postgres <<<")
	sugarLogger := ottologger.GetLogger()

	//query := "SELECT longitude, latitude, phone_number, address, business_type, address_benchmark, sales_type_id " +
	//	" FROM merchants " +
	////" LEFT JOIN merchant_business_types mbt on code " +
	//	" WHERE merchant_id = '" + req.MerchantID + "'"

	query := "SELECT longitude, latitude, phone_number, customer_code, address, sales_type_id FROM merchant_new_recruitments WHERE id = '" + fmt.Sprintf("%d", req.MerchantNewRecId) + "'"
	err = Dbcon.Raw(query).Scan(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get merchant not found list")
		return res, err
	}
	return res, nil
}

// GetDetailTodolistv24 ..
func GetDetailTodolistv24(req models.TodolistDetailReq) (res []models.TodolistDetailDBRes, err error) {
	fmt.Println(">>> TodolistDetail - GetDetailTodolistv24 - Postgres <<<")
	sugarLogger := ottologger.GetLogger()

	var query = " select distinct ts.code, ts.id as todolist_sub_category_id, ts.name, " +
		" t.todolist_category_id, t.merchant_name, t.task_date, t.created_at, t.updated_at, t.mid as merchant_id, t.status, t.id as todolist_id, " +
		" t.address as merchant_address, t.merchant_id as id_merchant, t.merchant_phone, " +
		" tc.id as todolist_category_id, tc.name as name_category,  " +
		" ta.id as task_id, ta.file_edukasi as link, " +
		" t.notes as reason, " +
		" lt.label_type " +
		" from todolists t " +
		//" left join merchants m on m.merchant_id = t.mid " +
		" left join todolist_categories tc on tc.id = t.todolist_category_id " +
		" left join tasks ta on ta.todolist_id = t.id " +
		" left join todolist_sub_categories ts on ts.id = ta.todolist_sub_category_id " +
		" left join label_tasks lt on lt.sub_category_id = ts.id " +
		" left join todolist_histories th on th.todolist_id = t.id " +
		" where t.merchant_phone = '" + req.MerchantPhone + "'" +
		" and t.id = " + req.TodolistId +
		" and lt.step = 1 "

	if req.TodolistCategoryId == 6 {
		query = " select distinct ts.code, ts.id as todolist_sub_category_id, ts.name, " +
			" t.todolist_category_id, t.merchant_name, t.task_date, t.created_at, t.updated_at, t.mid as merchant_id, t.status, t.id as todolist_id," +
			" m.address as merchant_address, m.id as id_merchant, m.customer_code, m.phone_number as merchant_phone, " +
			" tc.id as todolist_category_id, tc.name as name_category,  " +
			" ta.id as task_id, ta.file_edukasi as link, " +
			" t.notes as reason, " +
			" lt.label_type " +
			" from todolists t " +
			" left join merchant_new_recruitments m on m.id = t.merchant_new_recruitment_id " +
			" left join todolist_categories tc on tc.id = t.todolist_category_id " +
			" left join tasks ta on ta.todolist_id = t.id " +
			" left join todolist_sub_categories ts on ts.id = ta.todolist_sub_category_id " +
			" left join label_tasks lt on lt.sub_category_id = ts.id " +
			" left join todolist_histories th on th.todolist_id = t.id " +
			" where (m.phone_number = '" + req.MerchantPhone + "' or m.customer_code = '" + req.CustomerCode + "')" +
			" and t.id = " + req.TodolistId +
			" and lt.step = 1 "
	}

	sql := Dbcon.Raw(query).Scan(&res)
	if sql.Error != nil {
		sugarLogger.Info("Failed connect to database SFA when get detail todolist")
		return res, sql.Error
	}

	return res, nil
}

// GetTaskTodolistBySubCategory ..
func GetTaskTodolistBySubCategory(req models.TodolistTaskBySubCategoryReq) (res []dbmodels.LabelTasks, err error) {
	fmt.Println(">>> TodolistTaskBySubCategory - GetTaskTodolistBySubCategory - Postgres <<<")
	sugarLogger := ottologger.GetLogger()

	var query = "select lt.*, ta.supplier_name" +
		" from todolists t " +
		//" left join merchants m on m.merchant_id = t.mid " +
		" left join tasks ta on ta.todolist_id = t.id " +
		" left join todolist_sub_categories ts on ts.id = ta.todolist_sub_category_id " +
		" left join label_tasks lt on lt.sub_category_id = ts.id" +
		" where t.merchant_phone = '" + req.MerchantPhone + "'" +
		" and t.id = " + req.TodolistId +
		" and lt.sub_category_id = " + req.SubCategoryId

	sql := Dbcon.Raw(query).Scan(&res)
	if sql.Error != nil {
		sugarLogger.Info("Failed connect to database SFA when get task todolist")
		return res, sql.Error
	}
	return res, nil
}