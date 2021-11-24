package logging

import (
	"encoding/json"
	"fmt"
	"log"

	"ottosfa-api-apk/logging/kafka"
	"ottosfa-api-apk/logging/kafka/services"
	"ottosfa-api-apk/logging/models"

	"github.com/astaxie/beego/logs"
)

var (
	appname     string
	prod        bool
	procedure   kafka.KafkaProducer
	kafkaClient services.ServicePush
)

// Info is function for show log with status Info
func Info(req models.LogRequest) {
	req.Services = appname
	req.Level = "Info"
	logs.Info(fmt.Sprintf("[%sReq: %s | Res: %s | Message: %s | At: %s%s", req.State, req.RequestData, req.ResponseData, req.RawMessage, req.Packages, req.Function))
	if !prod {
		return
	}
	formInBytes, _ := json.Marshal(req)
	go kafkaClient.Push(formInBytes)
}

// Warn is function for show log with status Warn
func Warn(req models.LogRequest) {
	req.Services = appname
	req.Level = "Warn"
	logs.Warn(fmt.Sprintf("[%sReq: %s | Res: %s | Message: %s | At: %s%s", req.State, req.RequestData, req.ResponseData, req.RawMessage, req.Packages, req.Function))
	if !prod {
		return
	}
	formInBytes, _ := json.Marshal(req)
	go kafkaClient.Push(formInBytes)
}

// Trace is function for show log with status Trace
func Trace(req models.LogRequest) {
	req.Services = appname
	req.Level = "Trace"
	logs.Trace(fmt.Sprintf("[%sReq: %s | Res: %s | Message: %s | At: %s%s", req.State, req.RequestData, req.ResponseData, req.RawMessage, req.Packages, req.Function))
	if !prod {
		return
	}
	formInBytes, _ := json.Marshal(req)
	go kafkaClient.Push(formInBytes)
}

// Error is function for show log with status Error
func Error(req models.LogRequest) {
	req.Services = appname
	req.Level = "Error"
	logs.Error(fmt.Sprintf("[%sReq: %s | Res: %s | Message: %s | At: %s%s", req.State, req.RequestData, req.ResponseData, req.RawMessage, req.Packages, req.Function))
	if !prod {
		return
	}
	formInBytes, _ := json.Marshal(req)
	go kafkaClient.Push(formInBytes)

}

// Debug is function for show log with status Debug
func Debug(req models.LogRequest) {
	req.Services = appname
	req.Level = "Debug"
	logs.Debug(fmt.Sprintf("[%sReq: %s | Res: %s | Message: %s | At: %s%s", req.State, req.RequestData, req.ResponseData, req.RawMessage, req.Packages, req.Function))
	if !prod {
		return
	}
	formInBytes, _ := json.Marshal(req)
	go kafkaClient.Push(formInBytes)
}

// Fatal is function for show log with status Fatal
func Fatal(req models.LogRequest) {
	req.Services = appname
	req.Level = "Fatal"
	log.Fatal(fmt.Sprintf("[%sReq: %s | Res: %s | Message: %s | At: %s%s", req.State, req.RequestData, req.ResponseData, req.RawMessage, req.Packages, req.Function))
	if !prod {
		return
	}
	formInBytes, _ := json.Marshal(req)
	go kafkaClient.Push(formInBytes)
}

// Panic is function for show log with status Panic
func Panic(req models.LogRequest) {
	req.Services = appname
	req.Level = "Panic"
	log.Panic(fmt.Sprintf("[%sReq: %s | Res: %s | Message: %s | At: %s%s", req.State, req.RequestData, req.ResponseData, req.RawMessage, req.Packages, req.Function))
	if !prod {
		return
	}
	formInBytes, _ := json.Marshal(req)
	go kafkaClient.Push(formInBytes)
}
