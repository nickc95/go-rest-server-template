package errors

import "encoding/json"

type error struct {
	Message     string
	Description string
}

// Client error definitions
var (
	InternalServerError = error{
		Message:     "error",
		Description: "error",
	}
)

// GetErrorJSON returns json encoding representation of lr client error
func GetErrorJSON(errorStruct error) []byte {
	jsonString, _ := json.Marshal(errorStruct)
	return jsonString
}

// GetErrorJSONWithCustomDescription returns json encoding representation of lr client error with custom description set
func GetErrorJSONWithCustomDescription(errorStruct error, customDescription string) []byte {
	tempErrorStruct := errorStruct
	tempErrorStruct.Description = customDescription
	jsonString, _ := json.Marshal(tempErrorStruct)
	return jsonString
}
