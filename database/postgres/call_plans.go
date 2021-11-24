package postgres

import (
	"fmt"
	"github.com/vjeantet/jodaTime"
	ottologger "ottodigital.id/library/logger"
	"ottosfa-api-apk/database/dbmodels"
	"ottosfa-api-apk/models"
	"strconv"
	"time"
)

// GetCallPlanList ..
func GetCallPlanList(salesId int) ([]dbmodels.CallPlans, error) {
	fmt.Println(">>> CallPlanList - GetCallPlanList - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	res := []dbmodels.CallPlans{}

	var query = "(select * from call_plans " +
		"where sales_id = " + strconv.Itoa(salesId) + " and call_plan_date < '" + jodaTime.Format("yyyy-MM-dd", time.Now()) + "' order by call_plan_date desc limit 3) " +
		"union " +
		"(select * from call_plans " +
		"where sales_id = " + strconv.Itoa(salesId) + " and call_plan_date = '" + jodaTime.Format("yyyy-MM-dd", time.Now()) + "' limit 1) " +
		"union " +
		"(select * from call_plans " +
		"where sales_id = " + strconv.Itoa(salesId) + "  and call_plan_date > '" + jodaTime.Format("yyyy-MM-dd", time.Now()) + "' order by call_plan_date asc limit 2) " +
		"order by call_plan_date asc"

	err := Dbcon.Raw(query).Scan(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get call plan list")
		return res, err
	}
	return res, nil
}

// GetCallPlanListWithFilter ..
func GetCallPlanListWithFilter(salesId int, req models.CallPlanListReq) ([]dbmodels.CallPlans, error) {
	fmt.Println(">>> CallPlanList - GetCallPlanList - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	res := []dbmodels.CallPlans{}

	var query = "select * from call_plans " +
		"where sales_id = " + strconv.Itoa(salesId) + "  and call_plan_date between '" + req.StartDate + "'   and '" + req.EndDate + "' order by call_plan_date asc"

	err := Dbcon.Raw(query).Scan(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get call plan list with filter")
		return res, err
	}
	return res, nil
}

// SubmitCallPlan ..
func SubmitCallPlan(req dbmodels.CallPlans) error {
	fmt.Println(">>> CallPlanActionAddOrEdit - SubmitCallPlan - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	var err error
	err = Dbcon.Exec("update call_plans set effective_call = ? ,  success_call = ? , updated_at = ? where id = ?",
		req.EffectiveCall, req.SuccessCall, req.UpdatedAt, req.Id).Error
	if err != nil {
		sugarLogger.Error("Failed connect to database SFA when updating call plans")
		return err
	}
	return nil
}
