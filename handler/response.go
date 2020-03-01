package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"wechat-mall-web/defs"
)

func sendErrorResponse(w http.ResponseWriter, errResp defs.Response) {
	w.Header().Add("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(errResp.HttpSC)
	resStr, _ := json.Marshal(errResp.Error)
	_, _ = io.WriteString(w, string(resStr))
}

func sendNormalResponse(w http.ResponseWriter, result interface{}) {
	w.Header().Add("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	bytes, _ := json.Marshal(result)
	_, _ = io.WriteString(w, string(bytes))
}
