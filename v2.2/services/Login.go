package services

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"ottosfa-api-apk/database/dbmodels"
	"ottosfa-api-apk/database/postgres"
	"ottosfa-api-apk/models"
	redis "ottosfa-api-apk/redis"
	//redis "ottosfa-api-apk/redis"
	"ottosfa-api-apk/utils"
	"strconv"
	"strings"
)

// Login ..
func (svc *Service) Login(req models.LoginReq, res *models.Response) {
	fmt.Println(">>> Login - Service <<<")

	//check version
	var version bool
	versionCode, _ := strconv.Atoi(req.VersionCode)
	if strings.ToLower(req.Role) == "sfa" {
		versionRedis, _ := redis.GetRedisKey("SFA:ANDROID-VERSION")
		versionCodeRedis, _ := strconv.Atoi(versionRedis)

		switch {
		case versionCode >= versionCodeRedis:
			version = false
			break
		case versionCode < versionCodeRedis:
			version = true
		default:
			version = false
			break
		}
	} else if strings.ToLower(req.Role) == "indomarco" {
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
	} else if strings.ToLower(req.Role) == "" {
		res.Meta = utils.GetMetaResponse("update.apps")
		return
	}

	// Login by Phone Number
	var errDB error
	dataSalesDB := dbmodels.Salesmen{}
	phoneNumber, _ := strconv.Atoi(req.PhoneNumber)
	if phoneNumber == 0 {
		dataSalesDB, errDB = postgres.GetDataSalesBySalesId(req.SalesID)
		if errDB != nil {
			res.Meta = utils.GetMetaResponse("id.sales.not.found")
			return
		}
	} else {
		dataSalesDB, errDB = postgres.GetDataSalesByPhoneNumber(phoneNumber)
		if errDB != nil {
			res.Meta = utils.GetMetaResponse("no.handphone.not.found")
			return
		}
	}

	//get data sales type
	dataSalesTypeDB, errDB := postgres.GetDataSalesTypeSalesmen(dataSalesDB.ID)
	if errDB != nil {
		//res.Meta = utils.GetMetaResponse("id.number.not.found")
		//return
	}

	//generate token
	generateToken := utils.Rand64String(32)

	//encrypt password
	if bcrypt.CompareHashAndPassword([]byte(dataSalesDB.PasswordDigest), []byte(req.Pin)) != nil {
		res.Meta = utils.GetMetaResponse("wrong.pin")
		return
	}

	//save data sales login
	go SaveDataSalesLogin(dataSalesDB, generateToken, req, dataSalesDB.ID)

	res.Meta = utils.GetMetaResponse("success")
	res.Data = models.LoginRes{
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
		SalesType:    dataSalesTypeDB,
	}

	return
}

// SaveDataSalesLogin ..
func SaveDataSalesLogin(data dbmodels.Salesmen, generateToken string, req models.LoginReq, id int) {
	jsonData, _ := json.Marshal(data)
	go redis.SaveRedis(utils.RedisKeyAuth+generateToken, string(jsonData))

	//redisKeyExpToken := ottoutils.GetEnv("REDIS_KEY_EXP_TOKEN", "2h")
	//go redis.SaveRedisExp(utils.RedisKeyAuth+generateToken, redisKeyExpToken, string(jsonData))
	go postgres.UpdateDeviceLogin(id, req.DeviceID, req.DeviceToken, generateToken, req.FirebaseToken)
}
