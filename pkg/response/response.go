package response

import (
	"errors"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ResponseMessage struct {
	Message string `json:"message"`
}

type ResponseSucess struct {
	Success bool `json:"success"`
}

type ResponseStatus struct {
	Status string `json:"status"`
}

// APIError
// @Description validation error details
type ApiError struct {
	Field string   `json:"field"`           // the field cause the error
	Msg   string   `json:"msg"`             // error message
	Enums []string `json:"enums,omitempty"` // available options incase of field's payload is enums
}

type IDataResponse[T any] interface{}

type Response[T any] struct {
	Data         IDataResponse[T] `json:"data"`
	Error        string           `json:"error,omitempty"`
	ErrorDetails []ApiError       `json:"errors,omitempty"`
}

type DataResponse[T any] struct {
	Pagination *PaginationResponse `json:"metadata,omitempty"`
	Data       T                   `json:"data"`
}

type ErrorResponse struct {
	Error        string     `json:"error"`
	ErrorDetails []ApiError `json:"errors"`
}

type ResponseString struct {
	Data string `json:"data"`
}

func CreateResponse[T any](data T, paging *PaginationResponse, err error, payload any) Response[T] {
	resp := Response[T]{
		Data: data,
	}

	if paging != nil {
		resp.Data = DataResponse[T]{
			Data:       data,
			Pagination: paging,
		}
	}

	var ve validator.ValidationErrors
	if err != nil {
		resp.Error = err.Error()
	}

	if err != nil && errors.As(err, &ve) {
		errs := make([]ApiError, len(ve))
		for i, fe := range ve {
			var msg string
			var enums []string
			if payload != nil {
				field, ok := reflect.TypeOf(payload).FieldByName(fe.StructField())
				if ok {
					msg = field.Tag.Get("msg")
					if len(field.Tag.Get("enums")) > 0 {
						enums = strings.Split(field.Tag.Get("enums"), ",")
					}
				}
			}
			if msg == "" {
				splitErrMsg := strings.Split(fe.Error(), "Error:")
				if len(splitErrMsg) > 1 {
					msg = splitErrMsg[1]
				} else {
					msg = splitErrMsg[0]
				}
			}
			errs[i] = ApiError{Field: fe.Field(), Msg: msg, Enums: enums}
		}
		resp.ErrorDetails = errs
	}

	return resp
}
