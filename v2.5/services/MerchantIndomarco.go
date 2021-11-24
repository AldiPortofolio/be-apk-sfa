package services

import (
	"ottosfa-api-apk/database/postgres"
	"ottosfa-api-apk/hosts/indomarco"
	"ottosfa-api-apk/hosts/rose"
	"ottosfa-api-apk/models"
	"ottosfa-api-apk/utils"
	"strconv"
)

func (svc *Service) UpdateMerchantIndomarco(req models.UpdateMerchantIndomarcoReq, res *models.Response) {
	//get UpdateMerchantIndomarco
	resIndomarco, err := indomarco.SendV2(req, "UPDATEMERCHANT", req.Phone)
	if err != nil {
		res.Meta = models.MetaData{
			Status:  false,
			Code:    400,
			Message: err.Error(),
		}
		return
	}

	merchant, errMer := postgres.GetMerchantByMerchantID(req.MerchantID)
	if errMer != nil {
		res.Meta = models.MetaData{
			Status:  false,
			Code:    400,
			Message: errMer.Error(),
		}
		return
	}

	roseReq := rose.UpdateDataMerchantReq{}
	roseReq.PartnerCustomerID = req.CustomerCode
	roseReq.MerchantGroupID = resIndomarco.CustomerGroup
	roseReq.UserCategoryCode = strconv.Itoa(merchant.MerchantTypeId)
	roseReq.MID = req.MerchantID

	resRose, err := rose.UpdateDataMerchant(roseReq)
	if err != nil {
		res.Meta = models.MetaData{
			Status:  false,
			Code:    400,
			Message: err.Error(),
		}
		return
	}

	res.Data = resRose
	res.Meta = utils.GetMetaResponse("update.merchant.success")

	return
}
