package services

import (
	"fmt"
	ottoutils "ottodigital.id/library/utils"
	"ottosfa-api-apk/constants"
	"ottosfa-api-apk/hosts/rose"
	"ottosfa-api-apk/database/postgres"
	"ottosfa-api-apk/models"
	"ottosfa-api-apk/utils"
	"regexp"
	"strconv"
	"strings"
)

// AcquisitionsMerchantDetail ..
func (svc *Service) AcquisitionsMerchantDetail(req models.AcquisitionsMerchantDetailReq, res *models.Response) {
	fmt.Println(">>> AcquisitionsMerchantDetail - Service <<<")

	dataDB, errDB := postgres.AcquisitionsMerchantDetail(req)
	if errDB != nil {
		res.Meta = utils.GetMetaResponse("qris.merchant.not.found.db")
		return
	}

	urlImage := ottoutils.GetEnv("URL_IMAGE", "")

	postalCode := getKodePos(dataDB.Address)
	merchantlocation := getLokasiBisnisFromRose(strconv.Itoa(dataDB.MerchantLocation))

	level := "silver"
	if dataDB.Salesman == true {
		level = "sales_basic"
	}

	priorityLevel := "-"
	merchantGroupName := "Ottopay"
	if strings.ToUpper(dataDB.InstitutionId) == "IDM" {
		priorityLevel = "VIP"
		merchantGroupName = "INDOMARCO"
	}

	sales, _ := postgres.GetDataSalesBySalesId(strconv.Itoa(dataDB.SalesmanId))

	data := models.AcquisitionsMerchantDetailRes{
		PhotoKTP:                 dataDB.ImageIdCard,
		PhotoLocation:            dataDB.PhotoMerchantLocation,
		PhotoLocation2:           dataDB.ImageMerchantLocation,
		PhotoSelfie:              dataDB.ImageMerchant,
		PhotoSign:                dataDB.Signature,
		MerchantGroupName:        merchantGroupName,
		PriorityLevel:            priorityLevel,
		StoreName:                dataDB.Name,
		StoreJenisUsaha:          dataDB.BusinessType,
		StoreAlamat:              dataDB.Address,
		StoreKelurahan:           strconv.Itoa(dataDB.VillageId),
		StoreKecamatan:           strconv.Itoa(dataDB.DistrictId),
		StoreJamOperasional:      dataDB.OperationHour,
		StoreJenisLokasiBisnis:   dataDB.BusinessLocation,
		StoreKabupatenKota:       strconv.Itoa(dataDB.CityId),
		StorePostalCode:          postalCode,
		StoreProvince:            strconv.Itoa(dataDB.ProvinceId),
		StoreLatitude:            dataDB.Latitude,
		StoreLongitude:           dataDB.Longitude,
		StoreLokasiBisnis:        merchantlocation,
		StorePhoneNumber:         dataDB.PhoneNumber,
		AgentId:                  strconv.Itoa(dataDB.SalesmanId),
		AgentName:                sales.FirstName + " " + sales.LastName,
		AgentCompanyID:           sales.CompanyCode,
		AgentPhoneNumber:         sales.PhoneNumber,
		OwnerAddress:             "",
		OwnerFirstName:           dataDB.OwnerName,
		OwnerJenisKelamin:        "",
		OwnerKabupatenKota:       "",
		OwnerKecamatan:           "",
		OwnerKelurahan:           "",
		OwnerKodePos:             "",
		OwnerLastName:            "",
		OwnerNamaGadisIbuKandung: "",
		OwnerNoId:                dataDB.IdCard,
		OwnerNoTelf:              "",
		OwnerNoTelfLainnya:       "",
		OwnerPekerjaan:           "",
		OwnerProvinsi:            "",
		OwnerRt:                  "",
		OwnerRW:                  "",
		OwnerTanggalLahir:        "",
		OwnerTempatTinggal:       "",
		OwnerTglWxpiredId:        "",
		OwnerTipeId:              "",
		OwnerTitle:               "",
		DeviceType:               "",
		MetodePembayaran:         "",
		DeviceGroup:              "",
		DeviceBrand:              "",
		OutletId:                 "",
		TerminalPhoneNumber:      "",
		TerminalProvider:         "",
		InstitutionId:            dataDB.InstitutionId,
		Notes:                    "",
		MPAN:                     dataDB.MPAN,
		MerchantPAN:              dataDB.MPAN,
		MerchantOutletId:         dataDB.MerchantId,
		SalesID:                  "",
		KategoriBisnis:           dataDB.CategoryType,
	NMID:                     "",
		Level:                    level,
	ExistingQRValue:          "",
		Category:                 strings.ToLower(dataDB.InstitutionId),
	StoreNamePreprinted:      "",
		PhotoLocationLeft:        urlImage + dataDB.PhoneNumber + "_PhotoLocationLeft.jpeg",
		PhotoLocationRight:       urlImage + dataDB.PhoneNumber + "_PhotoLocationRight.jpeg",
		PhotoQrPreprinted:        urlImage + dataDB.PhoneNumber + "_PhotoQrPreprinted.jpeg",
		PartnerCustomerId:        dataDB.CustomerCode,
		Patokan:                  dataDB.AddressBenchmark,
		AcquitisionDate: 		  dataDB.CreatedAt,
		SrId:                     strconv.Itoa(dataDB.SalesTypeId),
	}

	res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)
	res.Data = data
	return
}

// getKodePos ..
func getKodePos(alamat string) string   {
	var regex, _ = regexp.Compile(`[0-9]+`)
	var str = regex. FindAllString(alamat, -1)
	postalCode := "00000"
	for i := 0; i < len(str); i++ {
		if len(str[i]) >= 5 {
			postalCode = str[i]
		}
	}
	return postalCode
}

// getLokasiBisnisFromRose ..
func getLokasiBisnisFromRose(lokasiBisnisId string) string {
	lokasiBisnis := ottoutils.GetEnv("LOOKUP_LOKASI_BISNIS", "")
	var lokasiBisnisName string
	dataLokasiBisnisRose, errRose := rose.LookUpGroup(lokasiBisnis)
	if errRose != nil {
		return lokasiBisnisName
	}
	for _, val := range dataLokasiBisnisRose {
		if val.Code == lokasiBisnisId {
			lokasiBisnisName = val.Name
		}
	}

	return lokasiBisnisName
}