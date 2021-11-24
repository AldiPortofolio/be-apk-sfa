package services

import (
	"fmt"
	"ottosfa-api-apk/database/postgres"
	"ottosfa-api-apk/hosts/notif"
	"ottosfa-api-apk/models"
	"ottosfa-api-apk/utils"
	"strconv"
)

// PushNotif ..
func (svc *Service) PushNotif(req models.PushNotifReq, res *models.Response) {
	fmt.Println(">>> PushNotif - Service <<<")

	phoneNumber, _ := strconv.Atoi(req.PhoneNumber)
	dataSalesDB, errDB := postgres.GetDataSalesByPhoneNumber(phoneNumber)
	if errDB != nil {
		res.Meta = utils.GetMetaResponse("phone.number.not.found")
		return
	}

	notifReq := notif.NotifRequest{
		CollapseKey: "type_a",
		Notification: notif.NotificationData{
			Title: req.Title,
			Body:  req.Body,
			Sound: "default",
		},
		Data: notif.DataNotif{
			Title:  req.Title,
			Body:   req.Body,
			Target: req.Target,
		},
	}

	fmt.Println("FirebaseToken: ", dataSalesDB.FirebaseToken)
	errNotif := notif.Send(dataSalesDB.FirebaseToken, notifReq)
	if errNotif != nil {
		res.Meta = utils.GetMetaResponse("send.push.notif.failed")
		return
	}

	res.Meta = utils.GetMetaResponse("success")
	return
}
