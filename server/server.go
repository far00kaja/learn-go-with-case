package server

import (
	"fmt"
	"log"
	"time"

	"github.com/far00kaja/learn-go-with-case/auth-service/config"
	"github.com/far00kaja/learn-go-with-case/auth-service/db"
	"github.com/far00kaja/learn-go-with-case/auth-service/internal/auth/controller"
	"github.com/far00kaja/learn-go-with-case/auth-service/internal/auth/repository"
	"github.com/far00kaja/learn-go-with-case/auth-service/internal/auth/service"
	middleware "github.com/far00kaja/learn-go-with-case/auth-service/middlewares"
	"github.com/far00kaja/learn-go-with-case/auth-service/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init() *gin.Engine {

	// initialization db
	db, err := db.ConnectDB()
	if err != nil {
		panic("failed connect DB")
	}

	// initialization redis
	_, msg, err := config.RedisConnectInit()
	if err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	} else {
		fmt.Println(msg)
		log.Println(msg)
	}

	// initalization
	var (
		authRepository repository.AuthRepository = repository.NewAuthServiceRepository(db)
		authService    service.AuthService       = service.NewAuthService(authRepository)
		authController controller.AuthController = controller.NewAuthController(authService)
		authRouter     routes.AuthRouter         = routes.NewAuthRouter(authController)
	)
	middleware.SetupLogger()
	r := gin.New()
	r.Use(cors.Default())
	r.GET("/auth-service/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Use(gin.Recovery(), middleware.Logger())
	r.GET("/ping", Ping)
	api := r.Group("/api")
	authRouter.RouterAuthV1(api)

	return r
}

// @BasePath

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags tags
// @Accept application/json
// @Produce json
// @Success 200
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "PING",
		"data":    time.Now().UnixNano(),
	})
}
