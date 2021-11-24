package services

import (
	"encoding/json"
	"fmt"
	"ottosfa-api-apk/constants"
	"ottosfa-api-apk/database/postgres"
	"ottosfa-api-apk/models"
	"ottosfa-api-apk/utils"
	"strings"
)

// IndomarcoCheckMerchantOttopay ..
func (svc *Service) IndomarcoCheckMerchantOttopay(req models.CheckMerchantOttopayReq, res *models.Response) {
	fmt.Println(">>> IndomarcoCheckMerchantOttopay - Service <<<")

	// if merchant_phone is empty, do not continue
	if req.Phone == "" {
		res.Meta = utils.GetMetaResponse("default")
		return
	}

	//GET DATA MERCHANT From FDS
	dataFDS, errFDS := svc.SendFDS(req,"CHECKMERCHANT", svc.General.SpanId)
	if errFDS != nil {
		res.Meta = utils.GetMetaResponse("fds.error")
		return
	}

	dataRes := models.CheckMerchantOttopayFDSRes{}
	if err := json.Unmarshal(dataFDS, &dataRes); err != nil {
		res.Meta = utils.GetMetaResponse("fds.error")
		return
	}

	if dataRes.ResponCode == "01" {
		//GET DATA MERCHANT From DB SFA
		dataDB, errDB := postgres.GetMerchantByPhoneNumber(req.Phone)
		if errDB != nil {
			res.Meta = utils.GetMetaResponse("default")
			return
		}

		if dataDB.PhoneNumber == req.Phone {
			fmt.Println("nomor merchant sudah dipakai")
			if len(strings.TrimSpace(dataRes.AccontNumber)) < 2 {
				res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)
				res.Data = models.CheckMerchantOttopayAlreadyExistRes{
					DescriptionCode: dataRes.DescriptionCode,
					MerchantID:      dataRes.MerchantID,
					ResponCode:      dataRes.ResponCode,
				}
				return
			}
			res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)
			res.Data = models.CheckMerchantOttopayRes{
				Nickname:        dataRes.Nickname,
				Gender:          dataRes.Gender,
				DescriptionCode: dataRes.DescriptionCode,
				BankFavorite:    dataRes.BankFavorite,
				AccontNumber:    dataRes.AccontNumber,
				Address:         dataRes.Address,
				OwnerName:       dataRes.OwnerName,
				Status:          dataRes.Status,
				MerchantID:      dataRes.MerchantID,
				BirthDate:       dataRes.BirthDate,
				ResponCode:      dataRes.ResponCode,
				CustomerID:      dataRes.CustomerID,
				Nama:            dataRes.Nama,
				VirtualAccount:  dataRes.VirtualAccount,
				Email:           dataRes.Email,
				IDCardNumber:    req.IdCard,
				VerifyStatus:    dataRes.VerifyStatus,
			}
			return
		}
	}

	res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)
	res.Data = models.CheckMerchantOttopayNotExistRes{
		ResponCode:      dataRes.ResponCode,
		DescriptionCode: dataRes.DescriptionCode,
	}
	return
}
