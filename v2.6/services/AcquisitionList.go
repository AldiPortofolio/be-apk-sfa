package services

import (
	"fmt"
	"ottosfa-api-apk/database/postgres"
	"ottosfa-api-apk/models"
	"ottosfa-api-apk/utils"
	"strings"
)

func (svc *Service) AcquisitionListBySR(bearer string, req models.GetAcquisitionBySalesTypeID, res *models.Response) {

	//get Acquiosition
	acquisitionListSR, err := postgres.GetAcquisitionBySrV2(req)
	if err != nil || len(acquisitionListSR) == 0 {
		res.Meta = utils.GetMetaResponse("acquistion.data.not.found")
		return
	}

	fmt.Println(">>>>>>>>>>>> acquisitionListSR <<<<<<<<<<<<<<<<<", acquisitionListSR)

	ids := []int64{}
	for _, value := range acquisitionListSR {
		sr := strings.Split(value.SalesRetailId, ",")
		_, found := Find(sr, req.SalesTypeId)
		if found {
			ids = append(ids, value.Id)
		}
	}

	fmt.Println(">>>>>>>>>>>> IDS <<<<<<<<<<<<<<<<<", ids)

	//get Acquiosition
	acquisitionList, err := postgres.GetAcquisitionByID(ids)
	if err != nil || len(acquisitionList) == 0 {
		res.Meta = utils.GetMetaResponse("acquistion.data.not.found")
		return
	}

	res.Data = acquisitionList
	res.Meta = utils.GetMetaResponse("businesstype.success")
	return
}

// Find takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
