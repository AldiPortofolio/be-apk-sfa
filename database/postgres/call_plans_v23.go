package postgres

import (
	"fmt"
	"github.com/vjeantet/jodaTime"
	ottologger "ottodigital.id/library/logger"
	"ottosfa-api-apk/database/dbmodels"
	"strconv"
	"time"
)

// GetCallPlanListv23 ..
func GetCallPlanListv23(salesId int) ([]dbmodels.CallPlans, error) {
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
