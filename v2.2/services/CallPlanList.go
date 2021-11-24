package services

import (
	"fmt"
	"github.com/vjeantet/jodaTime"
	"ottosfa-api-apk/constants"
	"ottosfa-api-apk/database/dbmodels"
	"ottosfa-api-apk/database/postgres"
	"ottosfa-api-apk/models"
	redis "ottosfa-api-apk/redis"
	"ottosfa-api-apk/utils"
	"strings"
	"time"
)

// CallPlanList ..
func (svc *Service) CallPlanList(bearer string, req models.CallPlanListReq, res *models.Response) {
	fmt.Println(">>> CallPlanList - Service <<<")

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

	//GET DATA POSITIONS SALES
	dataPositionsDB, errDB := postgres.CheckPositionsSalesSubarea(sales.ID)
	if errDB != nil {
		res.Meta = utils.GetMetaResponse("callplan.not.found")
		return
	}
	if len(*dataPositionsDB) == 0 {
		res.Meta = utils.GetMetaResponse("callplan.not.found")
		return
	}

	//GET DATA TEAM LEADER
	dataTeamLeaderDB, errDB := postgres.CheckTeamLeaderData(sales.ID)
	if errDB != nil {
		//res.Meta = utils.GetMetaResponse("callplan.not.found")
		//return
	}

	dataTeamLeaderTrim := uniqueTeamLeaderName(dataTeamLeaderDB)
	teamLeaderNameArr := []string{}
	branchOfficeNameArr := []string{}
	for _, val := range dataTeamLeaderTrim {
		teamLeaderNameArr = append(teamLeaderNameArr, val.TeamLeaderName)
		branchOfficeNameArr = append(branchOfficeNameArr, val.BranchOfficeName)
	}
	teamLeaderName := strings.Join(teamLeaderNameArr[:], ", ")
	branchOfficeName := strings.Join(utils.UniqueString(branchOfficeNameArr[:]), ", ")

	//GET DATA CALL PLAN LIST
	dataCallPlanListDB := []dbmodels.CallPlans{}
	fmt.Println(dataCallPlanListDB)
	if req.StartDate != "" && req.EndDate != "" {
		dataCallPlanListDB, errDB = postgres.GetCallPlanListWithFilter(sales.ID, req)
		if errDB != nil {
			res.Meta = utils.GetMetaResponse("callplan.not.found")
			return
		}
	} else {
		dataCallPlanListDB, errDB = postgres.GetCallPlanListv23(sales.ID)
		if errDB != nil {
			res.Meta = utils.GetMetaResponse("callplan.not.found")
			return
		}
	}

	callPlanList := []models.ListCallPlan{}
	for _, val := range dataCallPlanListDB {
		a := models.ListCallPlan{
			CallPlanId:           val.Id,
			Date:                 jodaTime.Format("yyyy-MM-dd", val.CallPlanDate),
			VillageName:          val.ClusterCoverageArea,
			AmountTasksCompleted: len(postgres.GetCallPlanListMerchantsCompleted(val.Id)),
			AmountTasks:          len(postgres.GetCallPlanListMerchantsAll(val.Id)),
		}

		if jodaTime.Format("yyyy-MM-dd", val.CallPlanDate) < jodaTime.Format("yyyy-MM-dd", time.Now()) {
			a.StatusDate = 0
		} else if jodaTime.Format("yyyy-MM-dd", val.CallPlanDate) == jodaTime.Format("yyyy-MM-dd", time.Now()) {
			a.StatusDate = 1
		} else if jodaTime.Format("yyyy-MM-dd", val.CallPlanDate) > jodaTime.Format("yyyy-MM-dd", time.Now()) {
			a.StatusDate = 2
		}

		callPlanList = append(callPlanList, a)
	}

	response := models.CallPlanListRes{
		TeamLeaderName:   teamLeaderName,
		BranchOfficeName: branchOfficeName,
		DateToday:        utils.ConvertTime(time.Now()),
		CallPlanList:     callPlanList,
	}

	res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)
	res.Data = response
	return
}

func uniqueTeamLeaderName(intSlice []models.TeamLeaderData) []models.TeamLeaderData {
	keys := make(map[models.TeamLeaderData]bool)
	list := []models.TeamLeaderData{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
