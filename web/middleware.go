package web

import (
	"context"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"strings"
	"time"
	"wechat-mall-backend/consts"
	"wechat-mall-backend/pkg/utils"
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
			payload, err := parseTokenAndValidate(r)
			if err != nil {
				return
			}
			// Inject the uid into the context
			ctx := context.WithValue(r.Context(), consts.ContextKey, payload.Uid)
			r = r.WithContext(ctx)
		}
		if strings.HasPrefix(uri, "/api") {
			if strings.HasPrefix(uri, "/api/wxapp/login") {
				goto nextHandler
			}
			payload, err := parseTokenAndValidate(r)
			if err != nil {
				return
			}
			// Inject the uid into the context
			ctx := context.WithValue(r.Context(), consts.ContextKey, payload.Uid)
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

func parseTokenAndValidate(r *http.Request) (*utils.Payload, error) {
	authorization := r.Header.Get("Authorization")
	if authorization == "" {
		return nil, errors.New("token is invalid")
	}
	if !strings.HasPrefix(authorization, "Bearer ") {
		return nil, errors.New("token is invalid")
	}
	tmpArr := strings.Split(authorization, " ")
	if len(tmpArr) != 2 {
		return nil, errors.New("token is invalid")
	}
	token := tmpArr[1]
	if !utils.ValidateToken(token) {
		return nil, errors.New("token is invalid")
	}
	payload, err := utils.ParseToken(token)
	if err != nil {
		return nil, errors.New("token is invalid")
	}
	return payload, nil
}
