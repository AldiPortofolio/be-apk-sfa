package services

import (
	"fmt"
	"ottosfa-api-apk/constants"
	"ottosfa-api-apk/database/postgres"
	"ottosfa-api-apk/models"
	"ottosfa-api-apk/utils"
)

// City ..
func (svc *Service) City(provinceId string, res *models.Response) {
	fmt.Println(">>> City - Service <<<")

	cityListDB, errDB := postgres.ListCity(provinceId)
	if errDB != nil {
		res.Meta = utils.GetMetaResponse("default")
		return
	}

	data := models.City{
		Cities: cityListDB,
	}

	res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)
	res.Data = data

	return
}
