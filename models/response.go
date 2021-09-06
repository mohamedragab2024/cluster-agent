package models

type Response struct {
	ResourceType string                 `json:"resourceType"`
	Data         map[string]interface{} `json:"data"`
	Message      string                 `json:"message"`
}
