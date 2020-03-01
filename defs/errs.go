package defs

import "net/http"

type Err struct {
	Code   int    `json:"error_code"`
	ErrMsg string `json:"msg"`
}

type Response struct {
	HttpSC int
	Error  Err
}

var (
	ErrorRequestBodyParseFailed = Response{HttpSC: http.StatusBadRequest, Error: Err{Code: 10001, ErrMsg: "Request body is not correct"}}
	ErrorNotAuthUser            = Response{HttpSC: http.StatusUnauthorized, Error: Err{Code: 10002, ErrMsg: "User authentication failed."}}
	ErrorRedisError             = Response{HttpSC: http.StatusInternalServerError, Error: Err{Code: 10003, ErrMsg: "Redis ops failed"}}
	ErrorInternalFaults         = Response{HttpSC: http.StatusInternalServerError, Error: Err{Code: 10004, ErrMsg: "Internal service error"}}
	ErrorTooManyRequests        = Response{HttpSC: http.StatusTooManyRequests, Error: Err{Code: 10005, ErrMsg: "Too many Request"}}
	ErrorParameterValidate      = Response{HttpSC: http.StatusBadRequest, Error: Err{Code: 10006, ErrMsg: "validate parameters failed"}}
	ErrorStorageError           = Response{HttpSC: http.StatusInternalServerError, Error: Err{Code: 10007, ErrMsg: "storage error"}}
)
