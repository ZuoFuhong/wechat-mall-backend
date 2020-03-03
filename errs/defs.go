package errs

import "net/http"

type Err struct {
	Code   int    `json:"error_code"`
	ErrMsg string `json:"msg"`
}

type HttpErr struct {
	HttpSC int
	Err
}

func (err Err) Error() string {
	return err.ErrMsg
}

var (
	ErrorRequestBodyParseFailed = HttpErr{HttpSC: http.StatusBadRequest, Err: Err{Code: 10001, ErrMsg: "Request body is not correct"}}
	ErrorNotAuthUser            = HttpErr{HttpSC: http.StatusUnauthorized, Err: Err{Code: 10002, ErrMsg: "User authentication failed."}}
	ErrorRedisError             = HttpErr{HttpSC: http.StatusInternalServerError, Err: Err{Code: 10003, ErrMsg: "Redis ops failed"}}
	ErrorInternalFaults         = HttpErr{HttpSC: http.StatusInternalServerError, Err: Err{Code: 10004, ErrMsg: "Internal service error"}}
	ErrorTooManyRequests        = HttpErr{HttpSC: http.StatusTooManyRequests, Err: Err{Code: 10005, ErrMsg: "Too many Request"}}
	ErrorParameterValidate      = HttpErr{HttpSC: http.StatusBadRequest, Err: Err{Code: 10006, ErrMsg: "validate parameters failed"}}
	ErrorStorageError           = HttpErr{HttpSC: http.StatusInternalServerError, Err: Err{Code: 10007, ErrMsg: "storage error"}}
	ErrorTokenInvalid           = HttpErr{HttpSC: http.StatusUnauthorized, Err: Err{Code: 10008, ErrMsg: "Token is invalid"}}
	ErrorValidateCodeInvalid    = HttpErr{HttpSC: http.StatusBadRequest, Err: Err{Code: 10009, ErrMsg: "Code is invalid"}}
	ErrorWechatError            = HttpErr{HttpSC: http.StatusInternalServerError, Err: Err{Code: 10010, ErrMsg: "wechat error"}}
)

func NewParameterError(errMsg string) HttpErr {
	return HttpErr{HttpSC: http.StatusBadRequest, Err: Err{Code: 10001, ErrMsg: errMsg}}
}
