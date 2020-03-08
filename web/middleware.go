package web

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"runtime"
	"strings"
	"time"
	"wechat-mall-backend/errs"
	"wechat-mall-backend/utils"
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

func (m Middleware) ValidateAuthToken(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		uri := r.URL.RequestURI()
		if strings.HasPrefix(uri, "/cms") {
			if uri == "/cms/login" {
				goto nextHandler
			}
			authorization := r.Header.Get("Authorization")
			if authorization == "" {
				panic(errs.ErrorTokenInvalid)
			}
			if !strings.HasPrefix(authorization, "Bearer ") {
				panic(errs.ErrorTokenInvalid)
			}
			tmpArr := strings.Split(authorization, " ")
			if len(tmpArr) != 2 {
				panic(errs.ErrorTokenInvalid)
			}
			refreshToken := tmpArr[1]
			if !utils.ValidateToken(refreshToken) {
				panic(errs.ErrorTokenInvalid)
			}
		}
	nextHandler:
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func (m Middleware) RecoverPanic(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				var httpErr errs.HttpErr
				log.Printf("Recover from panic:%+v", err)
				printStack()
				switch err.(type) {
				case errs.HttpErr:
					httpErr = err.(errs.HttpErr)
				default:
					httpErr = errs.ErrorInternalFaults
				}

				w.Header().Add("Content-Type", "application/json;charset=UTF-8")
				w.WriteHeader(httpErr.HttpSC)
				resStr, _ := json.Marshal(httpErr.Err)
				_, _ = io.WriteString(w, string(resStr))
			}
		}()
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	log.Print(string(buf[:n]))
}
