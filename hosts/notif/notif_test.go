package notif_test

import (
	"context"
	"fmt"
	"ottosfa-api-apk/hosts/notif"
	"ottosfa-api-apk/models/protomodels"
	"testing"
)

func TestNotifGRPC(t *testing.T) {
	notifReq := &protomodels.NotifRequest{
		IssuerId:      "OTTOPAY",
		FirebaseToken: "Tests123",
		TypeMessage:   "transaksi.sukses",
		Data:          []byte{},
	}
	notifClient := notif.GrpcNotifClient()
	_, err := notifClient.SendNotif(context.Background(), notifReq)
	if err != nil {
		fmt.Print(err)
	}
}
