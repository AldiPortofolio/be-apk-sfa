package services

import (
	"encoding/json"
	"log"
	ottologger "ottodigital.id/library/logger/v2"
	"ottosfa-api-apk/models"
	"testing"
)

func TestService_CallPlanVisitAdd(t *testing.T) {
	var ottolog ottologger.OttologInterface
	var res *models.Response
	req := models.CallPlanVisitAddReq{
		CallPlanId:         14,
		IdMerchant:         864526,
		MerchantPhone:      "087885092997",
		MerchantAddress:    "Jakarta",
		MerchantTypeId:     4,
		MerchantStatus:     "Found - Open",
		MerchantId:         "OP1A01551805",
		MerchantName:       "Toko Testing Diana Dev",
		ClockIn:            "2020-10-22 15:52:01",
		Longitude:          "6.54321",
		Latitude:           "-6.54321",
		Status:             "Visited",
		CallPlanActionName: "Action 1",
		ActionId:           1,
		ActionName:         "Edukasi",
		ProductId:          5,
		ProductName:        "Otto Grosir",
		Description:        "Ini adalah edukasi Otto Grosir",
		Result:             true,
		MerchantAction:     "text dari FE",
		Amount:             90000,
		Reason:             "",
		Note:               "Edukasi Merchant",
		ActionStatus:       "Completed",
		ActionType:         "Visited",
	}
	go InitiateService(ottolog).CallPlanVisitAdd("HVzwlFFaSfdfuJhmOEqFNLxkoqZQlRSX", req, res)

	byteRes, _ := json.Marshal(res)
	log.Println("res --> ", string(byteRes))
}
