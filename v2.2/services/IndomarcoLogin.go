package services

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"ottosfa-api-apk/constants"
	"ottosfa-api-apk/database/dbmodels"
	"ottosfa-api-apk/database/postgres"
	"ottosfa-api-apk/models"
	redis "ottosfa-api-apk/redis"
	"ottosfa-api-apk/utils"
	"strconv"
	"strings"
)

// LoginIndomarco ..
func (svc *Service) LoginIndomarco(req models.IndomarcoLoginReq, res *models.Response) {
	fmt.Println(">>> LoginIndomarco - Service <<<")

	//check version
	var version bool
	versionCode, _ := strconv.Atoi(req.VersionCode)
	versionRedis, _ := redis.GetRedisKey("INDOMARCO:ANDROID-VERSION")
	versionCodeRedis, _ := strconv.Atoi(versionRedis)

	switch {
	case versionCode >= versionCodeRedis:
		version = false
		break
	case versionCode < versionCodeRedis:
		version = true
		break
	}

	// Login by Sales Id
	dataSalesDB, err := postgres.GetDataSalesBySalesId(req.SalesID)
	if err != nil {
		res.Meta = utils.GetMetaResponse("id.number.not.found")
		return
	}

	//generate token
	generateToken := utils.Rand64String(32)

	//encrypt password
	if bcrypt.CompareHashAndPassword([]byte(dataSalesDB.PasswordDigest), []byte(req.Pin)) != nil {
		res.Meta = utils.GetMetaResponse("wrong.pin")
		return
	}

	//save data sales
	go SaveDataSales(dataSalesDB, generateToken, req, dataSalesDB.ID)

	res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)
	res.Data = models.IndomarcoLoginRes{
		ResponseCode: "00",
		SalesName:    fmt.Sprintf("%s %s", dataSalesDB.FirstName, dataSalesDB.LastName),
		Email:        dataSalesDB.Email,
		Description:  "Login Berhasil",
		Phone:        fmt.Sprintf("0%s", dataSalesDB.PhoneNumber),
		Photo:        dataSalesDB.Photo,
		SessionToken: generateToken,
		Status:       utils.StatusAccount(dataSalesDB.Status),
		SalesId:      dataSalesDB.SalesId,
		SFAID:        dataSalesDB.SfaID,
		ForceUpdate:  version,
		FunctionalId: strings.Trim(dataSalesDB.FunctionalPosition, " "),
	}

	return
}

// SaveDataSales ..
func SaveDataSales(data dbmodels.Salesmen, generateToken string, req models.IndomarcoLoginReq, id int) {
	jsonData, _ := json.Marshal(data)
	go redis.SaveRedis(utils.RedisKeyAuth+generateToken, string(jsonData))
	go postgres.UpdateDeviceLogin(id, req.DeviceID, req.DeviceToken, generateToken, req.FirebaseToken)
}
