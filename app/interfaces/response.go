package interfaces

import (
	"encoding/json"
	"io"
	"net/http"
)

// Result 统一的回包结构
type Result struct {
	Retcode int    `json:"error_code"`
	Errmsg  string `json:"msg"`
}

// Ok 响应
func Ok(w http.ResponseWriter, data interface{}) {
	writeToResponse(w, data)
}

// Error 响应
func Error(w http.ResponseWriter, errcode int, errmsg string) {
	result := &Result{
		Retcode: errcode,
		Errmsg:  errmsg,
	}
	writeToResponse(w, result)
}

func writeToResponse(w http.ResponseWriter, rspBody interface{}) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	// 理论上不会失败
	bodyBytes, _ := json.Marshal(&rspBody)
	_, _ = io.WriteString(w, string(bodyBytes))
}
