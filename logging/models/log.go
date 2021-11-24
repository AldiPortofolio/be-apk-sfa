package models

// LogRequest ..
type LogRequest struct {
	ID           string `json:"id"`
	Services     string `json:"services"`
	Level        string `json:"level"`
	State        string `json:"state"`
	Packages     string `json:"packages"`
	Function     string `json:"function"`
	Query        string `json:"query,omitempty"`
	RequestData  string `json:"requestData,omitempty"`
	ResponseData string `json:"responseData,omitempty"`
	RawMessage   string `json:"rawMessage,omitempty"`
}
