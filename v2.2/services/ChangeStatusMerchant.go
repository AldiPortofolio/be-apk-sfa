package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	ottologger "ottodigital.id/library/logger"
	"ottosfa-api-apk/database/dbmodels"
	"ottosfa-api-apk/database/postgres"
	"ottosfa-api-apk/models"
	"ottosfa-api-apk/utils"
	"strings"
)

// ChangeStatusMerchant for change status merchant ottopay
func (svc *Service) ChangeStatusMerchant(req models.ChangeStatusMerchantReq, res *models.Response) {
	fmt.Println(">>> ChangeStatusMerchant - Service <<<")

	sugarLogger := ottologger.GetLogger()

	var dataErrors []models.DataMerchantErrorByRow
	var dataMerchant dbmodels.MerchantStatus
	var errorMessages []string
	var rowError models.DataMerchantErrorByRow
	var dataMerchantType dbmodels.MerchantType

	merchantPhone := req.PhoneNumber
	status := req.Status
	merchantType := req.MerchantType

	if merchantPhone == "" || status == "" || merchantType == "" {
		var msg []string
		if merchantPhone == "" {
			msg = append(msg, "No HP kosong")
		}
		if status == "" {
			msg = append(msg, "Is Merchant Active kosong")
		}
		if merchantType == "" {
			msg = append(msg, "Merchant Type kosong")
		}
		data, _ := json.Marshal(req)
		json.Unmarshal(data, &rowError)
		rowError.ErrorMessages = strings.Join(msg, ", ")
		//dataErrors = append(dataErrors, rowError)

		//return errors.New(rowError.ErrorMessages)
		res.Meta = utils.GetMessageFailedError(422, errors.New("Failed to Upadate Status merchants : "+rowError.ErrorMessages))
		return
	}

	dataMerchant, merchErr := postgres.MerchantByPhoneNumber(merchantPhone, dataMerchant)

	if merchErr != nil {
		log.Println("Failed to get merchant : "+merchantPhone, merchErr)
		sugarLogger.Error(fmt.Sprintf("Failed to get merchant : "+merchantPhone, merchErr))
		errorMessages = append(errorMessages, "No HP Merchant tidak ditemukan")
		res.Meta = utils.GetMessageFailedError(422, errors.New("No HP Merchant tidak ditemukan"))
		return
	}

	if status != "Active" && status != "Inactive" && status != "Dormant" {
		err := errors.New("Opsi yang di masukan anatara 'Active', 'Inactive' atau 'Dormant'")
		log.Println("Failed to update status : ", err)
		sugarLogger.Error(fmt.Sprintf("Failed to update status : %v", err))
		errorMessages = append(errorMessages, "Opsi yang di masukan anatara 'Active', 'Inactive' atau 'Dormant'")

		res.Meta = utils.GetMessageFailedError(422, errors.New("Opsi yang di masukan anatara 'Active', 'Inactive' atau 'Dormant'"))
		return
	}

	dataMerchantType, mtErr := postgres.MerchantByMerchantType(merchantType, dataMerchantType)

	if mtErr != nil {
		log.Println("Failed to get merchant type: "+merchantType, merchErr)
		sugarLogger.Error(fmt.Sprintf("Failed to get merchant type : "+merchantType, merchErr))
		errorMessages = append(errorMessages, "Merchant type tidak ditemukan")
		res.Meta = utils.GetMessageFailedError(422, errors.New("Merchant type tidak ditemukan"))
		return
	}

	if len(errorMessages) > 0 {
		data, _ := json.Marshal(req)
		json.Unmarshal(data, &rowError)
		rowError.ErrorMessages = strings.Join(errorMessages, "|")
		dataErrors = append(dataErrors, rowError)
	}

	dataMerchant.Status = status
	dataMerchant.MerchantTypeID = dataMerchantType.ID
	mErr := postgres.UpdateStatusMerchant(dataMerchant)
	if mErr != nil {
		log.Println("Failed to update merchant : ", mErr)
		sugarLogger.Error(fmt.Sprintf("Failed to update merchant : %v", mErr))
		data, _ := json.Marshal(req)
		json.Unmarshal(data, &rowError)
		rowError.ErrorMessages = "Failed to update merchant : " + mErr.Error()
		dataErrors = append(dataErrors, rowError)
		fmt.Println(dataErrors)
		res.Meta = utils.GetMessageFailedError(422, errors.New("Failed to update merchant : "+mErr.Error()))
		return
	}

	res.Meta = utils.GetMetaResponse("change.status.merchant.success")
	return
}
