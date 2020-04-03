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
	ErrorInternalFaults      = HttpErr{HttpSC: http.StatusInternalServerError, Err: Err{Code: 10004, ErrMsg: "Internal service error"}}
	ErrorWechatError         = HttpErr{HttpSC: http.StatusInternalServerError, Err: Err{Code: 10007, ErrMsg: "微信内部异常！"}}
	ErrorTokenInvalid        = HttpErr{HttpSC: http.StatusUnauthorized, Err: Err{Code: 10008, ErrMsg: "Token is invalid"}}
	ErrorRefreshTokenInvalid = HttpErr{HttpSC: http.StatusUnauthorized, Err: Err{Code: 10009, ErrMsg: "Refresh token is invalid"}}
	ErrorParameterValidate   = HttpErr{HttpSC: http.StatusOK, Err: Err{Code: 10010, ErrMsg: "参数验证失败！"}}
	ErrorBannerNotExist      = HttpErr{HttpSC: http.StatusOK, Err: Err{Code: 10011, ErrMsg: "Banner不存在！"}}
	ErrorCategory            = HttpErr{HttpSC: http.StatusOK, Err: Err{Code: 10012, ErrMsg: "分类不存在！"}}
	ErrorGridCategory        = HttpErr{HttpSC: http.StatusOK, Err: Err{Code: 10013, ErrMsg: "宫格不存在！"}}
	ErrorSpecification       = HttpErr{HttpSC: http.StatusOK, Err: Err{Code: 10014, ErrMsg: "规格不存在！"}}
	ErrorSpecificationAttr   = HttpErr{HttpSC: http.StatusOK, Err: Err{Code: 10015, ErrMsg: "规格属性不存在！"}}
	ErrorGoods               = HttpErr{HttpSC: http.StatusOK, Err: Err{Code: 10016, ErrMsg: "商品不存在！"}}
	ErrorSKU                 = HttpErr{HttpSC: http.StatusOK, Err: Err{Code: 10017, ErrMsg: "SKU不存在！"}}
	ErrorCoupon              = HttpErr{HttpSC: http.StatusOK, Err: Err{Code: 10019, ErrMsg: "优惠券不存在！"}}
	ErrorAddress             = HttpErr{HttpSC: http.StatusOK, Err: Err{Code: 10020, ErrMsg: "收货地址不存在！"}}
	ErrorOrder               = HttpErr{HttpSC: http.StatusOK, Err: Err{Code: 10021, ErrMsg: "订单不存在！"}}
	ErrorGoodsCart           = HttpErr{HttpSC: http.StatusOK, Err: Err{Code: 10022, ErrMsg: "购物车商品不存在！"}}
	ErrorGroup               = HttpErr{HttpSC: http.StatusOK, Err: Err{Code: 10023, ErrMsg: "分组不存在！"}}
	ErrorCMSUser             = HttpErr{HttpSC: http.StatusOK, Err: Err{Code: 10024, ErrMsg: "用户不存在！"}}
	ErrorModulePage          = HttpErr{HttpSC: http.StatusOK, Err: Err{Code: 10025, ErrMsg: "页面不存在！"}}
	ErrorMiniappUser         = HttpErr{HttpSC: http.StatusOK, Err: Err{Code: 10026, ErrMsg: "用户不存在"}}
	ErrorOrderRefund         = HttpErr{HttpSC: http.StatusOK, Err: Err{Code: 10027, ErrMsg: "退款单不存在"}}
)

func NewParameterError(errMsg string) HttpErr {
	return HttpErr{HttpSC: http.StatusOK, Err: Err{Code: 10010, ErrMsg: errMsg}}
}

func NewCategoryError(errMsg string) HttpErr {
	return HttpErr{HttpSC: http.StatusOK, Err: Err{Code: 10012, ErrMsg: errMsg}}
}

func NewGridCategoryError(errMsg string) HttpErr {
	return HttpErr{HttpSC: http.StatusOK, Err: Err{Code: 10013, ErrMsg: errMsg}}
}

func NewSpecificationError(errMsg string) HttpErr {
	return HttpErr{HttpSC: http.StatusOK, Err: Err{Code: 10014, ErrMsg: errMsg}}
}

func NewSpecificationAttr(errMsg string) HttpErr {
	return HttpErr{HttpSC: http.StatusOK, Err: Err{Code: 10015, ErrMsg: errMsg}}
}

func NewErrorGoods(errMsg string) HttpErr {
	return HttpErr{HttpSC: http.StatusOK, Err: Err{Code: 10016, ErrMsg: errMsg}}
}

func NewErrorSKU(errMsg string) HttpErr {
	return HttpErr{HttpSC: http.StatusOK, Err: Err{Code: 10017, ErrMsg: errMsg}}
}

func NewErrorCoupon(errMsg string) HttpErr {
	return HttpErr{HttpSC: http.StatusOK, Err: Err{Code: 10019, ErrMsg: errMsg}}
}

func NewErrorAddress(errMsg string) HttpErr {
	return HttpErr{HttpSC: http.StatusOK, Err: Err{Code: 10020, ErrMsg: errMsg}}
}

func NewErrorOrder(errMsg string) HttpErr {
	return HttpErr{HttpSC: http.StatusOK, Err: Err{Code: 10021, ErrMsg: errMsg}}
}

func NewErrorGoodsCart(errMsg string) HttpErr {
	return HttpErr{HttpSC: http.StatusOK, Err: Err{Code: 10022, ErrMsg: errMsg}}
}

func NewErrorGroup(errMsg string) HttpErr {
	return HttpErr{HttpSC: http.StatusOK, Err: Err{Code: 10023, ErrMsg: errMsg}}
}

func NewErrorCMSUser(errMsg string) HttpErr {
	return HttpErr{HttpSC: http.StatusOK, Err: Err{Code: 10024, ErrMsg: errMsg}}
}

func NewErrorOrderRefund(errMsg string) HttpErr {
	return HttpErr{HttpSC: http.StatusOK, Err: Err{Code: 10027, ErrMsg: errMsg}}
}
