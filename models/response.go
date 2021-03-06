package models

// Response ..
type Response struct {
	Data interface{} `json:"data,omitempty"`
	Meta MetaData    `json:"meta"`
}

// MetaData ..
type MetaData struct {
	Status  bool   `json:"status" example:"true"`
	Code    int    `json:"code"  example:"200"`
	Message string `json:"message"  example:"Ok"`
}

// MappingErrorCodes models
type MappingErrorCodes struct {
	Key     string           `json:"key"`
	Content ContentErrorCode `json:"content"`
}

// ContentErrorCode models
type ContentErrorCode struct {
	Status  bool   `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}
