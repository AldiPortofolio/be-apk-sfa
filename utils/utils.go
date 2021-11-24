package utils

import (
	"fmt"
	"github.com/astaxie/beego"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"ottosfa-api-apk/hosts/minio"
	"ottosfa-api-apk/models"
	"ottosfa-api-apk/models/miniomodels"
	"strconv"
	"strings"
	"time"
)

var (
	// RedisKeyAuth ..
	RedisKeyAuth string
)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func init() {
	RedisKeyAuth = beego.AppConfig.DefaultString("redis.key.auth", "OTTO-SFA-TOKEN :")
}

// DecodeBearer ..
func DecodeBearer(str string) string {
	token := strings.Replace(str, "Bearer ", "", 1)
	return token
}

// Response ..
func Response(key string) models.Response {
	return models.Response{
		Meta: GetMetaResponse(key),
	}
}

// StatusAccount ..
func StatusAccount(Status int) string {
	var statusacc string

	switch Status {
	case 0:
		statusacc = "Unregistered"
		break
	case 1:
		statusacc = "Registered"
		break
	case 2:
		statusacc = "Verified"
		break
	case 3:
		statusacc = "Inactive"
		break
	case 4:
		statusacc = "Pending"
		break
	}

	return statusacc
}

// Rand64String ..
func Rand64String(n int) string {
	//todo bisa di pindahkan di global variable
	var src = rand.NewSource(time.Now().UnixNano())

	b := make([]byte, n)
	l := len(letterBytes)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < l {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

// HashPassword ..
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// UploadImage2Minio ..
func UploadImage2Minio(imageBase64 string, nameFile string, spanID string) (miniomodels.UploadRes, error) {
	random := rand.Intn(100000000)

	req := miniomodels.UploadReq{
		BucketName:  "ottodigital",
		Data:        imageBase64,
		NameFile:    nameFile + "-" + strconv.Itoa(random) + ".jpeg",
		ContentType: "image/jpeg",
	}

	res, errMinio := minio.Send(req, spanID)
	if errMinio != nil {
		fmt.Println("Failed to connect to minio:", errMinio)
		return res, errMinio
	}
	return res, nil
}

// UniqueString ..
func UniqueString(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// ConvertDateFormat ..
func ConvertDateFormat(reqDate string) string {
	reqDateLayout := "2020-10-16"
	reqDate = "2020-10-16T00:00:00Z"

	date, _ := time.Parse(reqDateLayout, reqDate)

	return date.Format("2006-01-02")
}

// ConvertTime ..
func ConvertTime(t time.Time) string {
	return t.Format("2006-01-02")
}

// ReformatReq ..
func ReformatReq(imageReq string) string {
	imageReq = fmt.Sprintf("%d", len(imageReq))
	return imageReq
}

// ChangeTypeCategoryAttendanceNameToCode ..
func ChangeTypeCategoryAttendanceNameToCode(name string) string {
	code := ""
	if name == "All" {
		code = "0"
	} else if name == "In" {
		code = "1"
	} else if name == "Out" {
		code = "2"
	}
	return code
}

// GetMessageFailedError ..
func GetMessageFailedError(code int, err error) models.MetaData {
	return models.MetaData{
		Status:  false,
		Code:    code,
		Message: err.Error(),
	}
}
