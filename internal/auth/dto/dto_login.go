package dto

type LoginDTO struct {
	MsIsdn   string `json:"msisdn" binding:"required"`
	Password string `json:"password" binding:"required,gte=3,lte=20"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
