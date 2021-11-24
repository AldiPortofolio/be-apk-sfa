package services

import (
	"ottosfa-api-apk/database/postgres"
	"ottosfa-api-apk/models"
	"ottosfa-api-apk/utils"
)

func (svc *Service) AcquisitionListBySR(bearer string, req models.GetAcquisitionBySalesTypeID, res *models.Response) {
	//get Acquiosition
	businessList, err := postgres.GetAcquisitionBySR(req)
	if err != nil || len(businessList) == 0 {
		res.Meta = utils.GetMetaResponse("acquistion.data.not.found")
		return
	}

	res.Data = businessList
	res.Meta = utils.GetMetaResponse("businesstype.success")
	return
}