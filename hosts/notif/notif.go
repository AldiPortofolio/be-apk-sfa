package notif

import (
	"errors"
	"fmt"
	"log"
	"ottosfa-api-apk/models/protomodels"
	"github.com/appleboy/go-fcm"
	"ottosfa-api-apk/logging"
	logmodels "ottosfa-api-apk/logging/models"
	"google.golang.org/grpc"
	ottoutils "ottodigital.id/library/utils"
)

var (
	firebasekey      string
	grpcNotifAddress string
)

func init() {
	firebasekey = ottoutils.GetEnv("NOTIF_KEY", "")
	grpcNotifAddress = ottoutils.GetEnv("NOTIF_GRPC_ADDRESS", "")
	//firebasekey = beego.AppConfig.DefaultString("notif.key", "AAAAU4tck2Q:APA91bFXkMQOmQQxBiTq6VUrB52af9Q2oyaIdKCNtbRqVF8gJ-_P7XR8luW0c9FwtCDc8lgHwTrgBt4iLNK8zlvifQLhAL-sS3M0Wq3gTvBMqLXoItP5M_weYfsKksDxzXZrrLt4ZHuC")
	//grpcNotifAddress = beego.AppConfig.DefaultString("notif.grpc.address", "18.138.101.49:9191")
}

// Send ..
func Send(firebaseToken string, req NotifRequest) error {
	if firebaseToken == "" {
		return errors.New("firebase token is nil")
	}
	notif := fcm.Notification{
		Title: req.Notification.Title,
		Body:  req.Notification.Body,
		Sound: req.Notification.Sound,
	}
	// Create the message to be sent.
	msg := &fcm.Message{
		To:           firebaseToken,
		CollapseKey:  req.CollapseKey,
		Notification: &notif,
		Data: map[string]interface{}{
			"title":  req.Data.Title,
			"body":   req.Data.Body,
			"target": req.Data.Target,
		},
	}

	// Create a FCM client to send the message.
	client, err := fcm.NewClient(firebasekey)
	if err != nil {

		logging.Error(logmodels.LogRequest{
			State:       "[ERROR-FIREBASEKEY]",
			Packages:    "[hosts]",
			Function:    "[Send]",
			RequestData: fmt.Sprintf("Firebasekey Error = %v", client),
			RawMessage:  fmt.Sprintf("Error %v", err),
		})
		return err
	}

	// Send the message and receive the response without retries.
	response, err := client.Send(msg)
	if err != nil {

		logging.Error(logmodels.LogRequest{
			State:       "[ERROR-MESSAGE]",
			Packages:    "[hosts]",
			Function:    "[Send]",
			RequestData: fmt.Sprintf("Message Error = %v", response),
			RawMessage:  fmt.Sprintf("Error %v", err),
		})
		return err
	}

	log.Printf("%#v\n", response)
	return nil
}

// GrpNotifClient ...
func GrpcNotifClient() protomodels.NotifsClient {
	port := grpcNotifAddress
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {

		logging.Error(logmodels.LogRequest{
			State:       "[ErrorConnecting]",
			Packages:    "[hosts]",
			Function:    "[GrpcNotifClient]",
			RequestData: fmt.Sprintf("Connection failed = %v", conn),
			RawMessage:  fmt.Sprintf("could not connect to %v %v", port, err),
		})
	}
	return protomodels.NewNotifsClient(conn)
}

