package web

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
	"wechat-mall-web/defs"
)

type Middleware struct {
}

func (m Middleware) LoggingHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("[%s] %q %v", r.Method, r.URL.String(), time.Now().Sub(startTime))
	}
	return http.HandlerFunc(fn)
}

func (m Middleware) RecoverPanic(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recover from panic:%+v", err)

				w.Header().Add("Content-Type", "application/json;charset=UTF-8")
				w.WriteHeader(http.StatusInternalServerError)
				resStr, _ := json.Marshal(defs.Err{Code: 10004, ErrMsg: err.(string)})

				_, _ = io.WriteString(w, string(resStr))
			}
		}()
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
