package handlers

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/kube-carbonara/cluster-agent/models"
)

type RequestHandler struct {
}

func (r RequestHandler) Handle(request models.ServerRequest) models.Response {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	if request.Verb == "GET" {
		return get(request)
	} else {
		return post(request)
	}

}

func get(request models.ServerRequest) models.Response {
	fmt.Printf("%s \n", request.Path)
	url := fmt.Sprintf("http://localhost:1323/%s", request.Path)
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed to call %s \n", url)
		return models.Response{
			Type:    "error",
			Status:  500,
			Message: fmt.Sprintf("Failed to call %s %s", url, err.Error()),
		}
	}
	defer res.Body.Close()
	if res.StatusCode >= 200 && res.StatusCode < 300 {
		var result models.Response
		body, _ := ioutil.ReadAll(res.Body)
		json.Unmarshal(body, &result)
		fmt.Print()
		return result
	}
	return models.Response{
		Type:    "error",
		Status:  res.StatusCode,
		Message: res.Status,
	}
}

func post(request models.ServerRequest) models.Response {
	url := fmt.Sprintf("http://localhost:1323/%s", request.Path)
	payload, _ := json.Marshal(request.PayLoad)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		fmt.Printf("Failed to call %s", url)
		return models.Response{
			Type:    "error",
			Status:  500,
			Message: fmt.Sprintf("Failed to call %s %s", url, err.Error()),
		}
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("Failed to call %s", url)
		return models.Response{
			Type:    "error",
			Status:  500,
			Message: fmt.Sprintf("Failed to call %s %s", url, err.Error()),
		}
	}
	defer res.Body.Close()
	if res.StatusCode >= 200 && res.StatusCode < 300 {
		var result models.Response
		body, _ := ioutil.ReadAll(res.Body)
		json.Unmarshal(body, &result.Data)
		result.Status = 200
		result.ResourceType = request.ResourceType
		return result
	}
	return models.Response{
		Type:    "error",
		Status:  res.StatusCode,
		Message: res.Status,
	}

}
