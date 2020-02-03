package models

// RequestData stores all request related data for use in the request lifecycle
type RequestData struct {
	RequestQuery     map[string]string
	RequestPayload   []byte
}

// NewRequestData returns address of a blank RequestData struct
func NewRequestData() *RequestData {
	requestData := RequestData{}
	requestData.RequestQuery = make(map[string]string)
	return &requestData
}
