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

// CallPlanActionSubmit ..
func (svc *Service) CallPlanActionSubmit(bearer string, req models.CallPlanActionSubmitReq, res *models.Response) {
	fmt.Println(">>> CallPlanActionSubmit - Service <<<")

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

	//GET DATA ACTION LIST
	callPlanActionListDB, errDB := postgres.GetCallPlanListActionsAllWithError(req.CallPlanMerchantId)
	if errDB != nil {
		res.Meta = utils.GetMetaResponse("default")
		return
	}

	if req.Status == "Visited" {
		res.Meta = utils.GetMetaResponse("call.plan.submit.success")
		return
	}

	//update call plan merchant
	var amount float32
	var status = "Incompleted"
	for _, val := range callPlanActionListDB {
		amount = amount + val.Amount
		if strings.ToLower(val.Status) == "completed" {
			status = "Completed"
		}
	}

	if status == "Incompleted" {
		res.Meta = utils.GetMetaResponse("action.call.plan.incompleted")
		return
	}

	effectiveCall := false
	if amount > 0 {
		effectiveCall = true
	}

	reqCallPlanMerchant := dbmodels.CallPlanMerchants{
		Id:             req.CallPlanMerchantId,
		EffectiveCall:  effectiveCall,
		Amount:         amount,
		UpdatedAt:      time.Now(),
		Status:         status,
		MerchantStatus: req.MerchantStatus,
		Notes:          req.Notes,
		ClockOut:       jodaTime.Format("yyyy-MM-dd HH:mm:ss", time.Now()),
	}
	if errDB = postgres.SubmitCallPlanMerchant(reqCallPlanMerchant); errDB != nil {
		res.Meta = utils.GetMetaResponse("default")
		return
	}

	callPlanMerchantListDB, errDB := postgres.GetCallPlanMerchantListById(req.CallPlanMerchantId)
	if errDB != nil {
		res.Meta = utils.GetMetaResponse("default")
		return
	}

	var countStatusCompleted int
	var countEffectiveCall int
	for _, val := range callPlanMerchantListDB {
		if strings.ToLower(val.Status) == "completed" {
			countStatusCompleted = countStatusCompleted + 1
		}

		if val.EffectiveCall == true {
			countEffectiveCall = countEffectiveCall + 1
		}
	}

	successCallAttr := float32(countStatusCompleted) / float32(len(callPlanMerchantListDB)) * 100
	effectiveCallAttr := float32(countEffectiveCall) / float32(len(callPlanMerchantListDB)) * 100

	//update call plan
	reqCallPlan := dbmodels.CallPlans{
		Id:            callPlanMerchantListDB[0].CallPlanId,
		SuccessCall:   fmt.Sprintf("%0.2f", successCallAttr),
		EffectiveCall: fmt.Sprintf("%0.2f", effectiveCallAttr),
		UpdatedAt:     time.Now(),
	}
	if errDB = postgres.SubmitCallPlan(reqCallPlan); errDB != nil {
		res.Meta = utils.GetMetaResponse("default")
		return
	}

	res.Meta = utils.GetMetaResponse("call.plan.submit.success")

	return
}
