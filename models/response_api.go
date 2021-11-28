package models

type APIResp struct {
	ResponseCode   string `json:"response_code"`
	ResponseStatus string `json:"response_status"`
	Message        string `json:"response_description"`
}
