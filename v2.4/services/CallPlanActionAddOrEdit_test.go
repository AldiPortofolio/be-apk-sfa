package services

import (
	"encoding/json"
	"log"
	ottologger "ottodigital.id/library/logger/v2"
	"ottosfa-api-apk/models"
	"testing"
)

func TestService_CallPlanActionAddOrEdit(t *testing.T) {
	var ottolog ottologger.OttologInterface
	var res *models.Response
	req := models.CallPlanActionAddOrEditReq{
		Id:                  1237,
		CallPlanMerchantId:  37,
		Name:                "Action 5",
		Action:              "Transaksi",
		ActionType:          "Additional",
		Product:             "Topup",
		Description:         "Ini Adalah Transaksi Top up",
		MerchantAction:      "merchant_action",
		Result:              true,
		Amount:              10000,
		Reason:              "",
		Note:                "",
		Status:              "Completed",
		CreatedAt:           "2020-10-15T11:38:18.05956Z",
		UpdatedAt:           "",
		MerchantId:          "",
		MerchantPhone:       "",
		NMID:                "",
		StoreNamePreprinted: "",
	}
	go InitiateService(ottolog).CallPlanActionAddOrEdit("HVzwlFFaSfdfuJhmOEqFNLxkoqZQlRSX", req, res)

	byteRes, _ := json.Marshal(res)
	log.Println("res --> ", string(byteRes))
}
