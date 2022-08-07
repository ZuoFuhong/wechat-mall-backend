package web

import (
	"fmt"
	"net/http"
	"wechat-mall-backend/app/interfaces"
	"wechat-mall-backend/pkg/config"
	"wechat-mall-backend/pkg/log"
)

type Server struct {
	name   string
	addr   string
	router *Router
}

func NewServer() *Server {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("load config fail: %v" + err.Error())
	}
	config.SetGlobalConfig(cfg)

	return &Server{
		name:   cfg.Server.Name,
		addr:   fmt.Sprintf("%s:%d", cfg.Server.Addr, cfg.Server.Port),
		router: NewRouter(),
	}
}

func (s *Server) Register(services *interfaces.MallHttpServiceImpl) {
	s.router.registerHandler(services)
}

func (s *Server) Serve() error {
	log.Debugf("%s runs on http://%s", s.name, s.addr)
	return http.ListenAndServe(s.addr, s.router)
}
