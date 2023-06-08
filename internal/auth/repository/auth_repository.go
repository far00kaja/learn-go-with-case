package repository

import (
	"github.com/far00kaja/learn-go-with-case/internal/auth/dto"
	"github.com/far00kaja/learn-go-with-case/internal/auth/models"
	"github.com/far00kaja/learn-go-with-case/lib"
	"gorm.io/gorm"
)

type authRepository struct {
	connection *gorm.DB
}

type AuthRepository interface {
	FindAuthByUsernameOrMsisdn(username string, msisdn string) (models.Auth, error)
	FindAuthByMsisdn(msisdn string) (models.Auth, error)
	Save(dto.RegisterDTO) (string, error)
}

func NewAuthServiceRepository(db *gorm.DB) *authRepository {
	return &authRepository{
		connection: db,
	}
}

func (r *authRepository) FindAuthByUsernameOrMsisdn(username string, msisdn string) (models.Auth, error) {
	var auth models.Auth
	result := r.connection.Where("username= ?", username).Or("msisdn= ?", msisdn).First(&auth)

	if result.Error != nil {
		return auth, result.Error
	}
	return auth, nil
}

func (r *authRepository) FindAuthByMsisdn(msisdn string) (models.Auth, error) {
	var auth models.Auth
	result := r.connection.Where("msisdn= ?", msisdn).First(&auth)

	if result.Error != nil {
		return auth, result.Error
	}
	return auth, nil
}

func (r *authRepository) Save(register dto.RegisterDTO) (string, error) {
	newPassword, salt, err := lib.PasswordHash(register.Password)
	if err != nil {
		return "", err
	}

	auth := &models.Auth{
		Msisdn:   register.Msisdn,
		Username: register.Username,
		Password: newPassword,
		Salt:     salt,
		Name:     register.Name,
	}

	result := r.connection.Create(&auth)
	if result.Error != nil {
		return "", err
	}
	return "success register", nil
}
