package services

import (
	"fmt"
	"ottosfa-api-apk/constants"
	"ottosfa-api-apk/models"
	"ottosfa-api-apk/utils"
)

// UpdateMerchant ..
func (svc *Service) UpdateMerchant(req models.UpdateMerchantReq, res *models.Response) {
	fmt.Println(">>> UpdateMerchant - Service <<<")

	err := svc.MerchantRepository.UpdateMerchant(req)
	if err != nil{
		res.Meta = utils.GetMetaResponse("default")
		return
	}

	res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)
	return
}
