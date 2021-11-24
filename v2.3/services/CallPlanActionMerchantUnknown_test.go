package services

import (
	"encoding/json"
	"log"
	ottologger "ottodigital.id/library/logger/v2"
	"ottosfa-api-apk/models"
	"testing"
)

func TestService_CallPlanActionMerchantUnknown(t *testing.T) {
	var ottolog ottologger.OttologInterface
	var res *models.Response
	req := models.CallPlanActionMerchantUnknownReq{
		CallPlanMerchantId: 38,
		Longitude:          "7.54321",
		Latitude:           "-7.54321",
		MerchantStatus:     "Not Found",
		Status:             "Completed",
		PhotoLocation:      "data:image/jpg;base64,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",
	}
	go InitiateService(ottolog).CallPlanActionMerchantUnknown("HVzwlFFaSfdfuJhmOEqFNLxkoqZQlRSX", req, res)

	byteRes, _ := json.Marshal(res)
	log.Println("res --> ", string(byteRes))
}
