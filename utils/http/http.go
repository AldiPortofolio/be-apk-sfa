package http

import (
	"crypto/tls"
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"net/http"

	"time"

	"github.com/astaxie/beego"
	"github.com/parnurzeal/gorequest"
	hcredismodels "ottodigital.id/library/healthcheck/models/redismodels"
	redis "ottosfa-api-apk/redis"
)

var (
	debugClientHTTP bool
	instuition      string
	timeout         string
	retrybad        int
)

func init() {
	debugClientHTTP = beego.AppConfig.DefaultBool("debugClientHTTP", true)
	timeout = beego.AppConfig.DefaultString("timeout", "60s")
	retrybad = beego.AppConfig.DefaultInt("retrybad", 1)
}

// HTTPGet func
func HTTPGet(url string, header http.Header) ([]byte, error) {
	request := gorequest.New()
	request.SetDebug(debugClientHTTP)
	timeout, _ := time.ParseDuration(timeout)
	//_ := errors.New("Connection Problem")
	// if url[:5] == "https" {
	// 	request.TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	// }
	reqagent := request.Get(url)
	reqagent.Header = header
	_, body, errs := reqagent.
		Timeout(timeout).
		Retry(retrybad, time.Second, http.StatusInternalServerError).
		End()
	if errs != nil {
		return []byte(body), errs[0]
	}
	return []byte(body), nil
}

// HTTPPost func
func HTTPPost(url string, jsondata interface{}) ([]byte, error) {
	request := gorequest.New()
	request.SetDebug(debugClientHTTP)
	timeout, _ := time.ParseDuration(timeout)
	//_ := errors.New("Connection Problem")
	if url[:5] == "https" {
		request.TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	reqagent := request.Post(url)
	reqagent.Header.Set("Content-Type", "application/json")
	_, body, errs := reqagent.
		Send(jsondata).
		Timeout(timeout).
		Retry(retrybad, time.Second, http.StatusInternalServerError).
		End()
	if errs != nil {
		return []byte(body), errs[0]
	}
	return []byte(body), nil
}

// HTTPPostWithHeader func
func HTTPPostWithHeader(url string, jsondata interface{}, header http.Header) ([]byte, error) {
	request := gorequest.New()
	request.SetDebug(debugClientHTTP)
	timeout, _ := time.ParseDuration(timeout)
	//_ := errors.New("Connection Problem")
	// if url[:5] == "https" {
	// 	request.TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	// }
	reqagent := request.Post(url)
	reqagent.Header = header
	_, body, errs := reqagent.
		Send(jsondata).
		Timeout(timeout).
		Retry(retrybad, time.Second, http.StatusInternalServerError).
		End()
	if errs != nil {
		return []byte(body), errs[0]
	}
	return []byte(body), nil
}

// HTTPPutWithHeader func
func HTTPPutWithHeader(url string, jsondata interface{}, header http.Header) ([]byte, error) {
	request := gorequest.New()
	request.SetDebug(debugClientHTTP)
	timeout, _ := time.ParseDuration(timeout)
	//_ := errors.New("Connection Problem")
	if url[:5] == "https" {
		request.TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	reqagent := request.Put(url)
	reqagent.Header = header
	_, body, errs := reqagent.
		Send(jsondata).
		Timeout(timeout).
		Retry(retrybad, time.Second, http.StatusInternalServerError).
		End()
	if errs != nil {
		return []byte(body), errs[0]
	}
	return []byte(body), nil
}

// HTTPDeleteWithHeader func
func HTTPDeleteWithHeader(url string, jsondata interface{}, header http.Header) ([]byte, error) {
	request := gorequest.New()
	request.SetDebug(debugClientHTTP)
	timeout, _ := time.ParseDuration(timeout)
	//_ := errors.New("Connection Problem")
	if url[:5] == "https" {
		request.TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	reqagent := request.Delete(url)
	reqagent.Header = header
	_, body, errs := reqagent.
		Send(jsondata).
		Timeout(timeout).
		Retry(retrybad, time.Second, http.StatusInternalServerError).
		End()
	if errs != nil {
		return []byte(body), errs[0]
	}
	return []byte(body), nil
}

// HTTPPostXForm func
func HTTPPostXForm(url string, jsondata map[string]string, key string) ([]byte, error) {
	request := gorequest.New()
	request.SetDebug(debugClientHTTP)
	timeout, _ := time.ParseDuration(timeout)
	//_ := errors.New("Connection Problem")
	reqagent := request.Post(url)
	reqagent.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, body, errs := reqagent.
		Send(jsondata).
		Timeout(timeout).
		Retry(retrybad, time.Second, http.StatusInternalServerError).
		End()
	healthCheckData, _ := json.Marshal(hcredismodels.ServiceHealthCheckRedis{
		StatusCode: resp.StatusCode,
		UpdatedAt:  time.Now().UTC(),
	})

	go redis.SaveRedis(key, healthCheckData)
	if errs != nil {
		logs.Error("Error Sending ", errs)
		return nil, errs[0]
	}
	return []byte(body), nil
}

// HTTPPostJson func
func HTTPPostJson(url string, jsondata map[string]string, key string) ([]byte, error) {
	request := gorequest.New()
	request.SetDebug(debugClientHTTP)
	timeout, _ := time.ParseDuration(timeout)
	//_ := errors.New("Connection Problem")
	reqagent := request.Post(url)
	reqagent.Header.Set("Content-Type", "application/json")
	resp, body, errs := reqagent.
		Send(jsondata).
		Timeout(timeout).
		Retry(retrybad, time.Second, http.StatusInternalServerError).
		End()
	healthCheckData, _ := json.Marshal(hcredismodels.ServiceHealthCheckRedis{
		StatusCode: resp.StatusCode,
		UpdatedAt:  time.Now().UTC(),
	})

	go redis.SaveRedis(key, healthCheckData)
	if errs != nil {
		logs.Error("Error Sending ", errs)
		return nil, errs[0]
	}
	return []byte(body), nil
}

// HTTPUploadMinio ..
func HTTPUploadMinio(url string, jsondata interface{}, key string) ([]byte, error) {
	request := gorequest.New()
	request.SetDebug(debugClientHTTP)
	timeout, _ := time.ParseDuration(timeout)
	reqagent := request.Post(url)
	reqagent.Header.Set("Content-Type", "application/json")
	resp, body, errs := reqagent.
		Timeout(timeout).
		Retry(retrybad, time.Second, http.StatusInternalServerError).
		Type("form-data").
		Send(jsondata).
		End()
	healthCheckData, _ := json.Marshal(hcredismodels.ServiceHealthCheckRedis{
		StatusCode: resp.StatusCode,
		UpdatedAt:  time.Now().UTC(),
	})

	go redis.SaveRedis(key, healthCheckData)
	if errs != nil {
		logs.Error("Error Sending ", errs)
		return nil, errs[0]
	}

	return []byte(body), nil
}
