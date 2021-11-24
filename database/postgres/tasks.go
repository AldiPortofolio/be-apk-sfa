package postgres

import (
	"fmt"
	ottologger "ottodigital.id/library/logger"
	"ottosfa-api-apk/models"
)

// UpdateTasks ..
func UpdateTasks(salesPhone string, req models.PostTodolistToDBReq) error {
	fmt.Println(">>> TodolistPost - UpdateTasks - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	err := Dbcon.Exec("update tasks set action_by = ? , updated_at = now(), action_date = now()  where id = ?", salesPhone, req.TaskID).Error
	if err != nil {
		sugarLogger.Error("Failed Update Table Task")
		return err
	}
	return nil
}
