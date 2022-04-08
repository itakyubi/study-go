package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"study-go/web/api"
	"study-go/web/common"
	"study-go/web/config"
	"study-go/web/log"
)

type AdminServer struct {
	cfg    *config.Config
	router *gin.Engine
	server *http.Server
	api    *api.API
	log    *log.Logger
}

func NewAdminServer(cfg *config.Config) (*AdminServer, error) {
	router := gin.New()
	server := &http.Server{
		Addr:           cfg.AdminServer.Port,
		Handler:        router,
		ReadTimeout:    cfg.AdminServer.ReadTimeout,
		WriteTimeout:   cfg.AdminServer.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	return &AdminServer{
		cfg:    cfg,
		router: router,
		server: server,
		log:    log.L().With(log.Any("server", "AdminServer")),
	}, nil
}

func (s *AdminServer) Run() {
	err := s.server.ListenAndServe()
	if err != nil {
		log.L().Info("admin server stopped", log.Error(err))
	}
}

func (s *AdminServer) SetAPI(api *api.API) {
	s.api = api
}

func (s *AdminServer) Close() {
	ctx, _ := context.WithTimeout(context.Background(), s.cfg.AdminServer.ShutdownTime)
	s.server.Shutdown(ctx)
}

func (s *AdminServer) InitRoute() {
	s.router.NoRoute(NoRouteHandler)
	s.router.NoMethod(NoMethodHandler)
	s.router.GET("/health", Health)

	// 类似于拦截器，加requestID、记录日志、鉴权
	/*s.router.Use(RequestIDHandler)
	s.router.Use(LoggerHandler)
	s.router.Use(s.AuthHandler)
	s.router.Use(s.ExternalHandlers...)*/

	v1 := s.router.Group("v1")
	{
		users := v1.Group("/users")
		users.POST("", common.Wrapper(s.api.AddUser))
		users.DELETE("/:userId", common.Wrapper(s.api.DeleteUser))
		users.PUT("/:userId", common.Wrapper(s.api.UpdateUser))
		users.GET("/:userId", common.Wrapper(s.api.GetUser))
		users.GET("", common.Wrapper(s.api.GetUserByName))
	}
}
