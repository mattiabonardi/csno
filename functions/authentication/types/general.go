package types

type ApplicationErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type LoginRequestDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponseDTO struct {
	Message      string   `json:"message"`
	AccessToken  TokenDTO `json:"accessToken"`
	RefreshToken TokenDTO `json:"refreshToken"`
}

type TokenDTO struct {
	Value      string `json:"value"`
	Expiration int64  `json:"expiration"`
}
