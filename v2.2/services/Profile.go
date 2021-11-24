package services

import (
	"encoding/json"
	"fmt"
	ottoutils "ottodigital.id/library/utils"
	"ottosfa-api-apk/constants"
	"ottosfa-api-apk/database/postgres"
	"ottosfa-api-apk/hosts/fds"
	"ottosfa-api-apk/models"
	"ottosfa-api-apk/models/fdsmodels"
	redis "ottosfa-api-apk/redis"
	"ottosfa-api-apk/utils"
	"strings"
)

// Profile ..
func (svc *Service) Profile(bearer string, res *models.Response) {
	fmt.Println(">>> Profile - Service <<<")

	token := utils.DecodeBearer(bearer)
	_, validateToken := redis.GetRedisKey(utils.RedisKeyAuth + token)
	if validateToken != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseInvalidToken)
		return
	}

	//cek data sales (get data salesId by Token)
	sales, errDB := postgres.CheckToken(bearer[7:])
	if errDB != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseInvalidToken)
		return
	}

	req := models.CheckProfilSalesFDSReq{
		Phone: "0" + sales.PhoneNumber,
	}

	// BEGIN check profil sales by phonenumber to fds
	checkProfilDataFDS, errFDS := fds.Send(req, "CHECKPROFILSALES", "")
	if errFDS != nil {
		fmt.Println("Failed to connect to fds:", errFDS)
		res.Meta = utils.GetMetaResponse(constants.KeyResponseFailed)
		return
	}

	dataFds := fdsmodels.CheckProfilSalesRes{}
	if err := json.Unmarshal(checkProfilDataFDS, &dataFds); err != nil {
		fmt.Println("Failed to unmarshal json response from fds:", err)
		res.Meta = utils.GetMetaResponse(constants.KeyResponseFailed)
		return
	}

	lenPencapaian := 0
	for _, val := range dataFds.AcquisitionData {
		if val.Status == "3" {
			lenPencapaian++
		}
	}

	//cek data sales (get data salesId by Token)
	dataPositionDB, errDB := postgres.GetPosition(sales.ID)
	if errDB != nil {
		fmt.Println("Failed to connect to db:", errDB)
		//res.Meta = utils.GetMetaResponse(constants.KeyResponseFailed)
		//return
	}

	dataPositionSalesDB := []models.ListLocationSales{}
	switch dataPositionDB.RegionableType {
	case "SubArea":
		dataPositionSalesDB, _ = postgres.InnerJoinSubArea(dataPositionDB.SalesmenId)
		break
	case "Area":
		dataPositionSalesDB, _ = postgres.InnerJoinArea(dataPositionDB.SalesmenId)
		break
	case "Branch":
		dataPositionSalesDB, _ = postgres.InnerJoinBranch(dataPositionDB.SalesmenId)
		break
	case "Region":
		dataPositionSalesDB, _ = postgres.InnerJoinRegion(dataPositionDB.SalesmenId)
		break
	}

	id := []string{}
	roleId := []string{}
	role := []string{}
	regionId := []string{}
	regionName := []string{}
	branchId := []string{}
	branchName := []string{}
	areaId := []string{}
	areaName := []string{}
	subAreaID := []string{}
	subAreaName := []string{}
	for _, val := range dataPositionSalesDB {

		//fmt.Printf("id->", val.Id)
		id = append(id, fmt.Sprintf("%d", val.Id))
		roleId = append(roleId, fmt.Sprintf("%d", val.RoleId))
		regionId = append(regionId, fmt.Sprintf("%d", val.RegionId))
		branchId = append(branchId, fmt.Sprintf("%d", val.BranchId))
		areaId = append(areaId, fmt.Sprintf("%d", val.AreaId))
		subAreaID = append(subAreaID, fmt.Sprintf("%d", val.SubAreaID))
		role = append(role, val.RoleName)
		regionName = append(regionName, val.RegionName)
		branchName = append(branchName, val.BranchName)
		areaName = append(areaName, val.AreaName)
		subAreaName = append(subAreaName, val.SubAreaName)
	}

	var region, branch string
	for _, val := range dataPositionSalesDB {
		region = val.RegionName
		branch = val.BranchName
	}

	areaAquitionSales := models.ProfileAreaAquisitions{
		Provinces: region,
		City:      branch,
		Locations: nil,
	}

	for _, value := range dataPositionSalesDB {
		temp := models.LocationsRes{
			Village:  value.SubAreaName,
			District: value.AreaName,
		}
		areaAquitionSales.Locations = append(areaAquitionSales.Locations, temp)
	}

	var desc string
	if sales.Status == 4 { //status = pending
		desc = ottoutils.GetEnv("OTTOSFA_MESSAGE_DESC_SALES_PENDING", "")
	}

	var branhOffice string
	if len(dataPositionSalesDB) > 0 {
		branhOffice = dataPositionSalesDB[0].BranchOffice
	}

	dataRes := models.CheckProfilSalesRes{
		ResponseCode:    "00",
		SalesName:       sales.FirstName + " " + sales.LastName,
		Email:           sales.Email,
		DescriptionCode: "Login Berhasil",
		Phone:           "0" + sales.PhoneNumber,
		Photo:           sales.Photo,
		SessionToken:    sales.SessionToken,
		Status:          utils.StatusAccount(sales.Status),
		SalesId:         sales.SalesId,
		SFAID:           sales.SfaID,
		SumMerchant:     len(dataFds.AcquisitionData),
		SumVerified:     lenPencapaian,
		AreaAquisitions: areaAquitionSales,
		Position: models.ProfilePosition{
			ID:           strings.Join(id[:], ", "),
			RoleID:       strings.Join(roleId[:], ", "),
			Role:         strings.Join(role[:], ", "),
			RegionID:     strings.Join(regionId[:], ", "),
			Region:       strings.Join(regionName[:], ", "),
			BranchID:     strings.Join(branchId[:], ", "),
			Branch:       strings.Join(branchName[:], ", "),
			BranchOffice: branhOffice,
			AreaID:       strings.Join(areaId[:], ", "),
			Area:         strings.Join(areaName[:], ", "),
			SubAreaID:    strings.Join(subAreaID[:], ", "),
			SubArea:      strings.Join(subAreaName[:], ", "),
		},
		Description: desc,
	}

	res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)
	res.Data = dataRes
	return
}
