package services

import (
	"fmt"
	ottoutils "ottodigital.id/library/utils"
	"ottosfa-api-apk/constants"
	"ottosfa-api-apk/database/postgres"
	"ottosfa-api-apk/hosts/rose"
	"ottosfa-api-apk/models"
	redis "ottosfa-api-apk/redis"
	"ottosfa-api-apk/utils"
)

// AcquitisionsDropdownList ..
func (svc *Service) AcquitisionsDropdownList(bearer string, res *models.Response) {
	fmt.Println(">>> AcquitisionsDropdownList - Service <<<")

	token := utils.DecodeBearer(bearer)
	_, validateToken := redis.GetRedisKey(utils.RedisKeyAuth + token)
	if validateToken != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseInvalidToken)
		return
	}

	//cek data sales (get data salesId by Token)
	_, errDB := postgres.CheckToken(bearer[7:])
	if errDB != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseInvalidToken)
		return
	}

	var data models.AcquitisionsDropdownListRes
	lokasiBisnis := ottoutils.GetEnv("LOOKUP_LOKASI_BISNIS", "JENIS_USAHA")
	dataLokasiBisnisRose, errRose := rose.LookUpGroup(lokasiBisnis)
	if errRose != nil {
		res.Meta = utils.GetMetaResponse("default")
		return
	}
	for _, val := range dataLokasiBisnisRose {
		a := models.DataModule{
			Code: val.Code,
			Name: val.Name,
		}
		data.LokasiBisnis = append(data.LokasiBisnis, a)
	}

	jenisLokasiBisnis := ottoutils.GetEnv("LOOKUP_JENIS_LOKASI_BISNIS", "JENIS_USAHA")
	dataJenisLokasiBisnisRose, errRose := rose.LookUpGroup(jenisLokasiBisnis)
	if errRose != nil {
		res.Meta = utils.GetMetaResponse("default")
		return
	}
	for _, val := range dataJenisLokasiBisnisRose {
		a := models.DataModule{
			Code: val.Code,
			Name: val.Name,
		}
		data.JenisLokasiBisnis = append(data.JenisLokasiBisnis, a)
	}

	jamOperasional := ottoutils.GetEnv("LOOKUP_JAM_OPERASIONAL", "JENIS_USAHA")
	dataJamOperasionalRose, errRose := rose.LookUpGroup(jamOperasional)
	if errRose != nil {
		res.Meta = utils.GetMetaResponse("default")
		return
	}
	for _, val := range dataJamOperasionalRose {
		a := models.DataModule{
			Code: val.Code,
			Name: val.Name,
		}
		data.JamOperational = append(data.JamOperational, a)
	}

	jamKunjunganTerbaik := ottoutils.GetEnv("LOOKUP_JAM_KUNJUNGAN_TERBAIK", "JENIS_USAHA")
	dataJamKunjunganTerbaikRose, errRose := rose.LookUpGroup(jamKunjunganTerbaik)
	if errRose != nil {
		res.Meta = utils.GetMetaResponse("default")
		return
	}
	for _, val := range dataJamKunjunganTerbaikRose {
		a := models.DataModule{
			Code: val.Code,
			Name: val.Name,
		}
		data.JamKunjunganTerbaik = append(data.JamKunjunganTerbaik, a)
	}

	res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)
	res.Data = data
	return
}
