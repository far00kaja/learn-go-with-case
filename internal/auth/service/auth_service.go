package service

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/far00kaja/learn-go-with-case/internal/auth/dto"
	"github.com/far00kaja/learn-go-with-case/internal/auth/repository"
	"github.com/far00kaja/learn-go-with-case/lib"
)

type authService struct {
	authRepository repository.AuthRepository
}

type AuthService interface {
	RegisterService(register dto.RegisterDTO) (string, error)
	LoginService(login dto.LoginDTO) (dto.LoginResponse, error)
	TokenService(login dto.Tokens) (dto.TokensResponse, error)
}

func NewAuthService(authRepository repository.AuthRepository) *authService {
	return &authService{
		authRepository: authRepository,
	}
}
func (s *authService) RegisterService(register dto.RegisterDTO) (string, error) {
	// check username and msisdn
	_, err := s.authRepository.FindAuthByUsernameOrMsisdn(register.Username, register.Msisdn)
	if err != nil {
		result, err := s.authRepository.Save(register)
		if err != nil {
			return "", nil
		}
		return result, nil
	}

	return "", errors.New("username or msisdn was used")
}

func (s *authService) LoginService(login dto.LoginDTO) (dto.LoginResponse, error) {
	// check username and msisdn
	result, err := s.authRepository.FindAuthByMsisdn(login.MsIsdn)
	if err != nil {
		return dto.LoginResponse{}, errors.New("msisdn or password invalid")
	}
	match, err := lib.VerifyPassword(login.Password, result.Password)

	if !match || err != nil {
		return dto.LoginResponse{}, errors.New("msisdn or password invalid")
	}

	token, err := lib.GenerateToken(result.ID)
	if err != nil {
		return dto.LoginResponse{}, errors.New("failed generate token")
	}

	// set redis
	_, err = lib.SetRedis(token, result.Msisdn)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	response := dto.LoginResponse{
		Token: token,
	}

	if err != nil {
		return response, err
	}

	return response, nil

}
func (s *authService) TokenService(login dto.Tokens) (dto.TokensResponse, error) {
	// check username and msisdn
	headerToken := strings.Split(login.Authorization, " ")

	_, err := lib.GetRedisFromKeyStrValue(headerToken[1])
	if err != nil {
		return dto.TokensResponse{}, err
	}
	fmt.Println("redis connected")

	token, err := jwt.ParseWithClaims(headerToken[1], &dto.TokensResponse{}, func(token *jwt.Token) (interface{}, error) {
		jwtSecret := os.Getenv("JWT_SECRET")
		return []byte(jwtSecret), nil
	})

	if claims, ok := token.Claims.(*dto.TokensResponse); ok && token.Valid {
		fmt.Println(claims.Token)
		return dto.TokensResponse{
			Token:          headerToken[1],
			ID:             claims.ID,
			StandardClaims: claims.StandardClaims,
		}, nil
	} else {

		return dto.TokensResponse{}, err
	}

}
