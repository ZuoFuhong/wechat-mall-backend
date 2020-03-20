package web

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"runtime"
	"strings"
	"time"
	"wechat-mall-backend/defs"
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
			if uri == "/cms/user/login" {
				goto nextHandler
			}
			if uri == "/cms/user/refresh" {
				goto nextHandler
			}
			payload := parseTokenAndValidate(r)
			// Inject the uid into the context
			ctx := context.WithValue(r.Context(), defs.ContextKey, payload.Uid)
			r = r.WithContext(ctx)
		}
		if strings.HasPrefix(uri, "/api") {
			if uri == "/api/wxapp/login" {
				goto nextHandler
			}
			payload := parseTokenAndValidate(r)
			// Inject the uid into the context
			ctx := context.WithValue(r.Context(), defs.ContextKey, payload.Uid)
			r = r.WithContext(ctx)
		}
	nextHandler:
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func (m Middleware) CORSHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Set("Access-Control-Allow-Origin", "*")
		header.Set("Access-Control-Allow-Headers", "*")
		header.Set("Access-Control-Allow-Credentials", "true")
		header.Set("Access-Control-Allow-Methods", "POST,GET,DELETE,PUT,OPTIONS")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func parseTokenAndValidate(r *http.Request) *utils.Payload {
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
	token := tmpArr[1]
	if !utils.ValidateToken(token) {
		panic(errs.ErrorTokenInvalid)
	}
	payload, err := utils.ParseToken(token)
	if err != nil {
		panic(errs.ErrorTokenInvalid)
	}
	return payload
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
