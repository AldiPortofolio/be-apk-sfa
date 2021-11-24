package postgres

import (
	"fmt"
	ottologger "ottodigital.id/library/logger"
	"ottosfa-api-apk/database/dbmodels"
)

// SaveTodolistHistories ..
func SaveTodolistHistories(req dbmodels.TodoListHistories) error {
	fmt.Println(">>> CallPlanTodolistMerchantNotFound - SaveTodolistHistories - Postgres <<<")
	sugarLogger := ottologger.GetLogger()

	err := Dbcon.Save(&req).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when insert todolist histories")
		return err
	}
	return nil
}
