package dto

type RegisterDTO struct {
	Msisdn   string `json:"msisdn" binding:"required,gte=5,lte=20"`
	Name     string `json:"name" binding:"required,gte=3,lte=100"`
	Username string `json:"username" binding:"required,gte=8,lte=15"`
	Password string `json:"password" binding:"required,gte=8,lte=25"`
}
