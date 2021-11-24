package services

import (
	"fmt"
	"ottosfa-api-apk/constants"
	"ottosfa-api-apk/database/postgres"
	"ottosfa-api-apk/models"
	"ottosfa-api-apk/utils"
)

// District ..
func (svc *Service) District(cityId string, res *models.Response) {
	fmt.Println(">>> District - Service <<<")

	districtListDB, errDB := postgres.ListDistrict(cityId)
	if errDB != nil {
		res.Meta = utils.GetMetaResponse("default")
		return
	}

	data := models.District{
		District: districtListDB,
	}

	res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)
	res.Data = data

	return
}
