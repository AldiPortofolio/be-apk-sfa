package services

import (
	"ottosfa-api-apk/database/postgres"
	"ottosfa-api-apk/models"
	"ottosfa-api-apk/utils"
	"strings"
)

func (svc *Service) BusinessTypeListByName(req models.GetAcquisitionByName, res *models.Response) {
	//get businessList
	acquisitions, err := postgres.GetBusinessTypeByName(req)
	if err != nil || len(acquisitions) <= 0 {
		res.Meta = utils.GetMetaResponse("businesstype.data.not.found")
		return
	}

	code := "" 
	for i, obj := range acquisitions {
		if i > 0 {
			code += ","
		}
        code += obj.BusinessTypes
    }

	codeTrimSpace := strings.TrimSpace(code)
	codes := strings.Split(codeTrimSpace, ",")

	keys := make(map[string]bool)
	list := []string{}
    for _, entry := range codes {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            list = append(list, entry)
        }
    }
	
	business_type, err := postgres.GetBusinessTypeByCode(list)

	res.Data = business_type
	res.Meta = utils.GetMetaResponse("businesstype.success")
	return
}
