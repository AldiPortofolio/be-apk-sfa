package services

import (
	"fmt"
	"ottosfa-api-apk/constants"
	"ottosfa-api-apk/database/postgres"
	"ottosfa-api-apk/models"
	"ottosfa-api-apk/utils"
)

// Village ..
func (svc *Service) Village(districtId string, res *models.Response) {
	fmt.Println(">>> Village - Service <<<")

	villageListDB, errDB := postgres.ListVillage(districtId)
	if errDB != nil {
		res.Meta = utils.GetMetaResponse("default")
		return
	}

	data := models.Village{
		Village: villageListDB,
	}

	res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)
	res.Data = data

	return
}
