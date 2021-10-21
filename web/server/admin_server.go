package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"study-go/web/api"
	"study-go/web/common"
)

type AdminServer struct {
	router *gin.Engine
	server *http.Server
	api    *api.API
}

func NewAdminServer() (*AdminServer, error) {
	router := gin.New()
	server := &http.Server{
		Addr:    ":8008",
		Handler: router,
	}
	return &AdminServer{
		router: router,
		server: server,
	}, nil
}

func (s *AdminServer) Run() {
	err := s.server.ListenAndServe()
	if err != nil {
		println(err.Error())
	}
}

func (s *AdminServer) SetAPI(api *api.API) {
	s.api = api
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
