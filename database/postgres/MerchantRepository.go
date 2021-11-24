package postgres

import (
	"fmt"
	"ottosfa-api-apk/database/dbmodels"
	"ottosfa-api-apk/models"
)

// MerchantRepository ..
type MerchantRepository struct {
}

// InitMerchantRepository ..
func InitMerchantRepository() *MerchantRepository {
	return &MerchantRepository{}
}

// CheckMerchantID ..
func (repo *MerchantRepository) CheckMerchantID(merchantID string) (dbmodels.Merchant, bool, error) {
	db := GetDbCon()
	var merchant dbmodels.Merchant
	err := db.Model(dbmodels.Merchant{}).Where("merchant_id = ?", merchantID).First(&merchant).Error
	if err != nil {
		return merchant, false, err
	}
	return merchant, true, err
}

// UpdateMerchant ..
func (repo *MerchantRepository) UpdateMerchant(req models.UpdateMerchantReq) (error) {
	fmt.Println(">>> UpdateMerchant - UpdateMerchant - Postgres <<<")

	merchant := dbmodels.Merchant{}
	Dbcon.Where("merchant_id = ?", req.MerchantId).First(&merchant)

	//if req.MerchantId != "" {
	//	merchant.MerchantId = req.MerchantId
	//}

	if req.Name != "" {
		merchant.Name = req.Name
	}

	if req.OwnerName != "" {
		merchant.OwnerName = req.OwnerName
	}

	if req.PhoneNumber != "" {
		merchant.PhoneNumber = req.PhoneNumber
	}

	if req.IdCard != "" {
		merchant.IdCard = req.IdCard
	}

	if req.ImageIdCard != "" {
		merchant.ImageIdCard = req.ImageIdCard
	}

	if req.ImageMerchantLocation != "" {
		merchant.ImageMerchantLocation = req.ImageMerchantLocation
	}

	if req.PhotoMerchantLocation != "" {
		merchant.PhotoMerchantLocation = req.PhotoMerchantLocation
	}

	if req.Signature != "" {
		merchant.Signature = req.Signature
	}

	if req.Longitude != "" {
		merchant.OwnerName = req.OwnerName
	}

	if req.Latitude != "" {
		merchant.Latitude = req.Latitude
	}

	if req.Address != "" {
		merchant.Address = req.Address
	}

	if req.BusinessLocation != "" {
		merchant.BusinessLocation = req.BusinessLocation
	}

	if req.MerchantLocation != 0 {
		merchant.MerchantLocation = req.MerchantLocation
	}

	if req.BusinessType != "" {
		merchant.BusinessType = req.BusinessType
	}

	if req.CategoryType != "" {
		merchant.CategoryType = req.CategoryType
	}

	if req.OperationHour != "" {
		merchant.OperationHour = req.OperationHour
	}

	if req.BestVisit != "" {
		merchant.BestVisit = req.BestVisit
	}

	if req.Dob != "" {
		merchant.Dob = req.Dob
	}

	if req.ProvinceId != 0 {
		merchant.ProvinceId = req.ProvinceId
	}

	if req.CityId != 0 {
		merchant.CityId = req.CityId
	}

	if req.DistrictId != 0 {
		merchant.DistrictId = req.DistrictId
	}

	if req.VillageId != 0 {
		merchant.VillageId = req.VillageId
	}

	if req.AddressBenchmark != "" {
		merchant.AddressBenchmark = req.AddressBenchmark
	}

	err := Dbcon.Save(&merchant).Error
	if err != nil {
		return err
	}
	return nil
}
