package DTOLogin

import "wan-api-verify-user/DTO"

type LoginInputDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginOutputDTO struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Role      string `json:"role"`
	Token     DTO.TokenDTO `json:"token"`
}