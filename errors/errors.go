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

// GetErrorJSON returns json encoding representation of an error struct
func GetErrorJSON(errorStruct error) []byte {
	jsonString, _ := json.Marshal(errorStruct)
	return jsonString
}
