package services

import (
	"fmt"
	"ottosfa-api-apk/database/postgres"
	"ottosfa-api-apk/hosts/indomarco"
	"ottosfa-api-apk/hosts/rose"
	"ottosfa-api-apk/models"
	"ottosfa-api-apk/utils"
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

	fmt.Println("======= RESPONSE INDOMARCO ========", resIndomarco)

	merchant, errMer := postgres.GetMerchantLinkingIndomarco(req.MerchantID)
	if errMer != nil {
		res.Meta = models.MetaData{
			Status:  false,
			Code:    400,
			Message: errMer.Error(),
		}
		return
	}

	merchant.CustomerCode = req.CustomerCode
	merchant.IndomarcoStatus = true
	merchant.BusinessType = "INDOMARCO"
	merchant.SalesTypeID = "1"

	errMer = postgres.UpdateMerchantLinkingIndomarco(merchant)
	if errMer != nil {
		res.Meta = models.MetaData{
			Status:  false,
			Code:    400,
			Message: errMer.Error(),
		}
		return
	}

	roseReq := rose.UpdateDataMerchantV2Req{}
	roseReq.PartnerCustomerID = req.CustomerCode
	roseReq.MerchantGroupID = req.MerchantGroupID
	roseReq.UserCategoryCode = req.MerchantCategoryID
	roseReq.MID = req.MerchantID
	roseReq.SrID = "1"
	roseReq.TipeBisnis = "INDOMARCO"
	roseReq.PartnerCode = "idm"

	resRose, err := rose.UpdateDataMerchantV2(roseReq)
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
