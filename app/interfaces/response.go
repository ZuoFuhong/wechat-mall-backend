package interfaces

import (
	"encoding/json"
	"io"
	"net/http"
)

// Result 统一的回包结构
type Result struct {
	Retcode int         `json:"retcode"`
	Errmsg  string      `json:"errmsg"`
	Data    interface{} `json:"data"`
}

// Ok 响应
func Ok(w http.ResponseWriter, data interface{}) {
	result := &Result{
		Data: data,
	}
	writeToResponse(w, result)
}

// Error 响应
func Error(w http.ResponseWriter, errcode int, errmsg string) {
	result := &Result{
		Retcode: errcode,
		Errmsg:  errmsg,
	}
	writeToResponse(w, result)
}

func writeToResponse(w http.ResponseWriter, result *Result) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	// 理论上不会失败
	retstr, _ := json.Marshal(&result)
	_, _ = io.WriteString(w, string(retstr))
}
