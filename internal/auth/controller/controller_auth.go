package controller

import (
	"net/http"

	"github.com/far00kaja/learn-go-with-case/auth-service/internal/auth/dto"
	"github.com/far00kaja/learn-go-with-case/auth-service/internal/auth/service"
	"github.com/far00kaja/learn-go-with-case/auth-service/lib/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type authController struct {
	service service.AuthService
}

type AuthController interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	Token(c *gin.Context)
}

func NewAuthController(service service.AuthService) *authController {
	return &authController{
		service: service,
	}
}

// Register godoc
// @Summary Create user for auth
// @Schemes
// @Param tags body dto.RegisterDTO true "register auth"
// @Description Save tags data in Db
// @Tags tags
// @Accept application/json
// @Produce json
// @Success 201
// @Router /api/v1/register [post]
func (c *authController) Register(ctx *gin.Context) {
	validate := validator.New()

	var register dto.RegisterDTO

	err := ctx.ShouldBindJSON(&register)

	if err != nil {
		response.BadRequest(ctx, err, register)
		return
	}

	valErr := validate.Struct(&register)
	if valErr != nil {
		ctx.JSON(400, response.ResponseBadRequest{
			Code:    0,
			Message: response.ErrorValidationResponse(valErr),
		})
		return
	}

	_, err = c.service.RegisterService(register)
	if err != nil {
		ctx.JSON(400, response.ResponseBadRequest{
			Code:    0,
			Message: err.Error(),
		})
	}

	ctx.JSON(http.StatusCreated, response.ResponseNoData{
		Code:    1,
		Message: "Success",
	})
}

// Login godoc
// @Summary Login User
// @Schemes
// @Description check account into db
// @Tags login
// @Param tags body dto.LoginDTO true "Login"
// @Accept application/json
// @Produce json
// @Success 200
// @Router /api/v1/login [post]
func (c *authController) Login(ctx *gin.Context) {
	validate := validator.New()

	var login dto.LoginDTO

	err := ctx.ShouldBindJSON(&login)

	if err != nil {
		response.BadRequest(ctx, err, login)
		return
	}

	valErr := validate.Struct(&login)
	if valErr != nil {
		ctx.JSON(400, response.ResponseBadRequest{
			Code:    0,
			Message: response.ErrorValidationResponse(valErr),
		})
		return
	}

	result, err := c.service.LoginService(login)
	if err != nil {
		ctx.JSON(400, response.ResponseBadRequest{
			Code:    0,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, response.ResponseData{
		Code:    1,
		Message: "Success",
		Data:    result,
	})
}

// Token godoc
// @Summary check user token for authorization
// @Schemes
// @Description Save see detail behind jwt
// @Param Authorization header string true "With the basic started"
// @Tags tags
// @Accept application/json
// @Produce json
// @Success 200
// @Router /api/v1/token [get]
func (c *authController) Token(ctx *gin.Context) {
	validate := validator.New()

	var token dto.Tokens

	err := ctx.ShouldBindHeader(&token)

	if err != nil {
		response.BadRequest(ctx, err, token)
		return
	}

	valErr := validate.Struct(&token)
	if valErr != nil {
		ctx.JSON(400, response.ResponseBadRequest{
			Code:    0,
			Message: response.ErrorValidationResponse(valErr),
		})
		return
	}

	result, err := c.service.TokenService(token)
	if err != nil {
		ctx.JSON(400, response.ResponseBadRequest{
			Code:    0,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, response.ResponseData{
		Code:    1,
		Message: "Success",
		Data:    result,
	})
}
