package services

import (
	"fmt"
	"ottosfa-api-apk/constants"
	"ottosfa-api-apk/database/postgres"
	"ottosfa-api-apk/models"
	"ottosfa-api-apk/utils"
)

// ClearSession ..
func (svc *Service) ClearSession(res *models.Response) {
	fmt.Println(">>> ClearSession - Service <<<")

	//clear sesssion token
	errDB := postgres.ClearSession()
	if errDB != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseInvalidToken)
		return
	}

	res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)
	return
}
