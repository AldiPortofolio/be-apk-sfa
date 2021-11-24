package services

import (
	"fmt"
	"ottosfa-api-apk/constants"
	"ottosfa-api-apk/database/postgres"
	"ottosfa-api-apk/models"
	"ottosfa-api-apk/utils"
)

// Province ..
func (svc *Service) Province(countryId string, res *models.Response) {
	fmt.Println(">>> Province - Service <<<")

	provinceListDB, errDB := postgres.ListProvince(countryId)
	if errDB != nil {
		res.Meta = utils.GetMetaResponse("default")
		return
	}

	data := models.Province{
		Provinces: provinceListDB,
	}

	res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)
	res.Data = data

	return
}
