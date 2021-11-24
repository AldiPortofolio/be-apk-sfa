package services

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"ottosfa-api-apk/constants"
	"ottosfa-api-apk/database/dbmodels"
	"ottosfa-api-apk/database/postgres"
	"ottosfa-api-apk/models"
	redis "ottosfa-api-apk/redis"
	"ottosfa-api-apk/utils"
	"strconv"
	"strings"
)

// TodolistPost ..
func (svc *Service) TodolistPost(bearer string, req models.TodolistPostReq, res *models.Response) {
	fmt.Println(">>> TodolistPost - Service <<<")

	//check version id
	ver, _ := redis.GetRedisKey("SFA:ANDROID-VERSION")
	codeRedis, _ := strconv.Atoi(ver)
	codeReq := req.VersionID
	switch {
	case codeReq >= int64(codeRedis):
		break
	case codeReq < int64(codeRedis):
		res.Meta = utils.GetMetaResponse(constants.KeyResponseInvalidToken)
		return
		break
	default:
		break
	}

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

	data := []dbmodels.FollowUps{}
	for i := 0; i < len(req.Label); i++ {
		if req.Body[i] != "" {

			if req.ContentType[i] == "Signature" || req.ContentType[i] == "Image" {
				//upload Image or Signature
				resUpload, err := svc.SendMinio(req.Body[i], "Task", svc.General.SpanId)
				if err != nil {
					res.Meta = utils.GetMetaResponse("failed.upload.image")
					return
				}
				req.Body[i] = resUpload.Url
			}

			if req.ContentType[i] == "ImageMerchant" {
				//upload photo 1
				resPhoto1, err := svc.SendMinio(req.PhotoMerchant1, strings.Replace(req.LabelPhotoMerchant1, " ", "_", -1), svc.General.SpanId)
				if err != nil {
					res.Meta = utils.GetMetaResponse("failed.upload.image")
					return
				}
				request := models.PostTodolistToDBReq{
					TaskID:      req.TaskID[i],
					Label:       req.LabelPhotoMerchant1,
					ContentType: "Image",
					Body:        resPhoto1.Url,
				}
				err = postgres.TodolistPost(request)
				if err != nil {
					res.Meta = utils.GetMetaResponse("todolist.failed.update.db")
					return
				}

				//upload photo 2
				resPhoto2, err := svc.SendMinio(req.PhotoMerchant2, strings.Replace(req.LabelPhotoMerchant2, " ", "_", -1), svc.General.SpanId)
				if err != nil {
					res.Meta = utils.GetMetaResponse("failed.upload.image")
					return
				}
				request = models.PostTodolistToDBReq{
					TaskID:      req.TaskID[i],
					Label:       req.LabelPhotoMerchant2,
					ContentType: "Image",
					Body:        resPhoto2.Url,
				}
				err = postgres.TodolistPost(request)
				if err != nil {
					res.Meta = utils.GetMetaResponse("todolist.failed.update.db")
					return
				}

				//upload photo 3
				resPhoto3, err := svc.SendMinio(req.PhotoMerchant3, strings.Replace(req.LabelPhotoMerchant3, " ", "_", -1), svc.General.SpanId)
				if err != nil {
					res.Meta = utils.GetMetaResponse("failed.upload.image")
					return
				}
				request = models.PostTodolistToDBReq{
					TaskID:      req.TaskID[i],
					Label:       req.LabelPhotoMerchant3,
					ContentType: "Image",
					Body:        resPhoto3.Url,
				}
				err = postgres.TodolistPost(request)
				if err != nil {
					res.Meta = utils.GetMetaResponse("todolist.failed.update.db")
					return
				}
			}

			if req.ContentType[i] == "MAP" {
				req.Body[i] = req.Alamat + ";" + req.Province + ";" + req.City + ";" + req.District + ";" + req.Village + ";" + req.Longitude + ";" + req.Latitude + ";" + req.Patokan
			}

			request := models.PostTodolistToDBReq{
				TaskID:      req.TaskID[i],
				Label:       req.Label[i],
				ContentType: req.ContentType[i],
				Body:        req.Body[i],
			}

			if req.ContentType[i] != "ImageMerchant" {
				err := postgres.TodolistPost(request)
				if err != nil {
					res.Meta = utils.GetMetaResponse("todolist.failed.update.db")
					return
				}
			}

			err := postgres.UpdateTasks(sales.PhoneNumber, request)
			if err != nil {
				res.Meta = utils.GetMetaResponse("todolist.failed.update.db")
				return
			}
			logs.Info(data)
		}
	}

	status := req.Status
	if req.Reason == "" {
		status = "Done"

		err := postgres.UpdateStatusDoneTodolist(sales.PhoneNumber, "", status, req.TodolistID, req.Long, req.Lat)
		if err != nil {
			res.Meta = utils.GetMetaResponse("todolist.failed.update.status.db")
			return
		}

		res.Meta = utils.GetMetaResponse("todolist.success")
		return
	}

	//insert ke table history jika pending
	request2 := models.PostHistoryTodolistReq{
		TodolistID:  req.TodolistID,
		OldTaskDate: req.OldTaskDate,
		NewTaskDate: req.NewTaskDate,
		Reason:      req.Reason,
		Longitude:   req.Long,
		Latitude:    req.Lat,
	}
	erry := postgres.PostHistoryTodolist(request2)
	if erry != nil {
		res.Meta = utils.GetMetaResponse("todolist.failed.update.status.db")
		return
	}

	if req.Status != "Late" {
		status = "Pending"
	}
	err := postgres.UpdateStatusPendingTodolist(sales.PhoneNumber, req.NewTaskDate, status, req.TodolistID, req.Long, req.Lat)
	if err != nil {
		res.Meta = utils.GetMetaResponse("todolist.failed.update.status.db")
		return
	}

	res.Meta = utils.GetMetaResponse("todolist.success")
	return
}
