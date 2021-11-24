package postgres

import (
	"fmt"
	ottologger "ottodigital.id/library/logger"
	"ottosfa-api-apk/database/dbmodels"
)

// GetCallPlanListActionsAll ..
func GetCallPlanListActionsAll(callPlanMerchantId int64) []dbmodels.CallPlanActions {
	fmt.Println(">>> CallPlanMerchantList - GetCallPlanListActionsAll - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	res := []dbmodels.CallPlanActions{}

	err := Dbcon.Model(dbmodels.CallPlanActions{}).Where("call_plan_merchant_id = ?", callPlanMerchantId).Find(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get call plan action")
		return res
	}
	return res
}

// GetCallPlanListActionsCompleted ..
func GetCallPlanListActionsCompleted(callPlanMerchantId int64) []dbmodels.CallPlanActions {
	fmt.Println(">>> CallPlanMerchantList/CallPlanActionUpdateClockInMerchant - GetCallPlanListActionsCompleted - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	res := []dbmodels.CallPlanActions{}

	err := Dbcon.Model(dbmodels.CallPlanActions{}).Where("call_plan_merchant_id = ? and status = 'Completed'", callPlanMerchantId).Find(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get call plan action")
		return res
	}
	return res
}

// AddCallPlanAction ..
func AddCallPlanAction(req dbmodels.CallPlanActions) error {
	fmt.Println(">>> CallPlanVisitAdd - AddCallPlanAction - Postgres <<<")
	sugarLogger := ottologger.GetLogger()

	err := Dbcon.Save(&req).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when save call plan action")
		return err
	}

	return err
}

// GetCallPlanListActionsAllWithError ..
func GetCallPlanListActionsAllWithError(callPlanMerchantId int64) ([]dbmodels.CallPlanActions, error) {
	fmt.Println(">>> CallPlanMerchantList - GetCallPlanListActionsAll - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	res := []dbmodels.CallPlanActions{}

	err := Dbcon.Model(dbmodels.CallPlanActions{}).Where("call_plan_merchant_id = ?", callPlanMerchantId).Order("id asc").Find(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get call plan action")
		return res, err
	}
	return res, err
}

// GetCallPlanActionByID ..
func GetCallPlanActionByID(id int64) (dbmodels.CallPlanActions, error) {
	fmt.Println(">>> CallPlanMerchantList - GetCallPlanListActionsAll - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	var res dbmodels.CallPlanActions

	err := Dbcon.Where(dbmodels.CallPlanActions{
		Id: id,
	}).First(&res).Error

	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get call plan action by id")
		return res, err
	}

	return res, err
}

// SaveCallPlanAction ..
func SaveCallPlanAction(req dbmodels.CallPlanActions) error {
	fmt.Println(">>> CallPlanActionAddOrEdit - SaveCallPlanAction - Postgres <<<")
	sugarLogger := ottologger.GetLogger()

	err := Dbcon.Save(&req).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when save call plan action (add or edit)")
		return err
	}
	return nil
}

// DeleteCallPlanAction ..
func DeleteCallPlanAction(req dbmodels.CallPlanActions) error {
	fmt.Println(">>> CallPlanActionDelete - DeleteCallPlanAction - Postgres <<<")
	sugarLogger := ottologger.GetLogger()

	//err := Dbcon.Delete(dbmodels.CallPlanActions{Id: req.Id, ActionType: req.ActionType}).Error
	err := Dbcon.Where("id = ? and action_type != 'Default' ", req.Id).Delete(dbmodels.CallPlanActions{}).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when delete call plan action")
		return err
	}
	return nil
}
