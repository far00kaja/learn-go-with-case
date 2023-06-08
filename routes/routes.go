package routes

import (
	"github.com/far00kaja/learn-go-with-case/auth-service/internal/auth/controller"
	"github.com/far00kaja/learn-go-with-case/auth-service/lib"
	"github.com/gin-gonic/gin"
)

type authRouter struct {
	router     gin.Engine
	controller controller.AuthController
}
type AuthRouter interface {
	RouterAuthV1(r *gin.RouterGroup)
}

func NewAuthRouter(controller controller.AuthController) *authRouter {
	return &authRouter{
		router:     gin.Engine{},
		controller: controller,
	}
}

func (a *authRouter) RouterAuthV1(r *gin.RouterGroup) {
	authRouterGroup := r.Group("/v1")
	authRouterGroup.POST("/register", a.controller.Register)
	authRouterGroup.POST("/login", a.controller.Login)
	authRouterGroup.GET("/token", lib.VerifyJWT(), a.controller.Token)
	// authRouterGroup.GET("/token", a.controller.Token)
}
