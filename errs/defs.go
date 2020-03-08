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
	ErrorParameterValidate      = HttpErr{HttpSC: http.StatusBadRequest, Err: Err{Code: 10006, ErrMsg: "validate parameters failed"}}
	ErrorTokenInvalid           = HttpErr{HttpSC: http.StatusUnauthorized, Err: Err{Code: 10008, ErrMsg: "Token is invalid"}}
	ErrorValidateCodeInvalid    = HttpErr{HttpSC: http.StatusBadRequest, Err: Err{Code: 10009, ErrMsg: "Code is invalid"}}
	ErrorWechatError            = HttpErr{HttpSC: http.StatusInternalServerError, Err: Err{Code: 10010, ErrMsg: "wechat error"}}
	ErrorBannerNotExist         = HttpErr{HttpSC: http.StatusBadRequest, Err: Err{Code: 10011, ErrMsg: "Banner does not exist"}}
	ErrorCategory               = HttpErr{HttpSC: http.StatusBadRequest, Err: Err{Code: 10012, ErrMsg: "Category does not exist"}}
	ErrorGridCategory           = HttpErr{HttpSC: http.StatusBadRequest, Err: Err{Code: 10013, ErrMsg: "GridCategory does not exist"}}
	ErrorSpecification          = HttpErr{HttpSC: http.StatusBadRequest, Err: Err{Code: 10014, ErrMsg: "Specification does not exist"}}
	ErrorSpecificationAttr      = HttpErr{HttpSC: http.StatusBadRequest, Err: Err{Code: 10015, ErrMsg: "Specification attr does not exist"}}
	ErrorSPU                    = HttpErr{HttpSC: http.StatusBadRequest, Err: Err{Code: 10016, ErrMsg: "SPU does not exist"}}
	ErrorSKU                    = HttpErr{HttpSC: http.StatusBadRequest, Err: Err{Code: 10017, ErrMsg: "SKU does not exist"}}
	ErrorActivity               = HttpErr{HttpSC: http.StatusBadRequest, Err: Err{Code: 10018, ErrMsg: "Activity does not exist"}}
	ErrorCoupon                 = HttpErr{HttpSC: http.StatusBadRequest, Err: Err{Code: 10019, ErrMsg: "Coupon does not exist"}}
)

func NewAuthUserError(errMsg string) HttpErr {
	return HttpErr{HttpSC: http.StatusUnauthorized, Err: Err{Code: 10002, ErrMsg: errMsg}}
}

func NewParameterError(errMsg string) HttpErr {
	return HttpErr{HttpSC: http.StatusBadRequest, Err: Err{Code: 10001, ErrMsg: errMsg}}
}

func NewCategoryError(errMsg string) HttpErr {
	return HttpErr{HttpSC: http.StatusBadRequest, Err: Err{Code: 10012, ErrMsg: errMsg}}
}

func NewGridCategoryError(errMsg string) HttpErr {
	return HttpErr{HttpSC: http.StatusBadRequest, Err: Err{Code: 10013, ErrMsg: errMsg}}
}

func NewSpecificationError(errMsg string) HttpErr {
	return HttpErr{HttpSC: http.StatusBadRequest, Err: Err{Code: 10014, ErrMsg: errMsg}}
}

func NewSpecificationAttr(errMsg string) HttpErr {
	return HttpErr{HttpSC: http.StatusBadRequest, Err: Err{Code: 10015, ErrMsg: errMsg}}
}

func NewErrorSPU(errMsg string) HttpErr {
	return HttpErr{HttpSC: http.StatusBadRequest, Err: Err{Code: 10016, ErrMsg: errMsg}}
}

func NewErrorSKU(errMsg string) HttpErr {
	return HttpErr{HttpSC: http.StatusBadRequest, Err: Err{Code: 10017, ErrMsg: errMsg}}
}

func NewErrorActivity(errMsg string) HttpErr {
	return HttpErr{HttpSC: http.StatusBadRequest, Err: Err{Code: 10018, ErrMsg: errMsg}}
}

func NewErrorCoupon(errMsg string) HttpErr {
	return HttpErr{HttpSC: http.StatusBadRequest, Err: Err{Code: 10019, ErrMsg: errMsg}}
}
