package baseerrorresponse

import "fmt"

type Response struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Result  any    `json:"result"`
}

func NewError(Error error) Response {
	return Response{Error: Error.Error(), Success: false, Result: nil}
}

func NewSErrorf(Error string, arg ...any) Response {
	return Response{Error: fmt.Errorf(Error, arg...).Error(), Success: false, Result: nil}
}

func New(result any) Response {
	return Response{Result: result, Success: true, Error: ""}
}
