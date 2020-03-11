package defs

import (
	"encoding/json"
	"io"
	"net/http"
)

func SendNormalResponse(w http.ResponseWriter, result interface{}) {
	w.Header().Add("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	bytes, _ := json.Marshal(result)
	_, _ = io.WriteString(w, string(bytes))
}
