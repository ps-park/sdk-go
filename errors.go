package sdk_go

import "fmt"

type ResponseValidationError struct {
	Message string
	Code    int
	Data    interface{}
}

func (e *ResponseValidationError) Error() string {
	return fmt.Sprintf("%s (Code: %d, Data: %v)", e.Message, e.Code, e.Data)
}
