package ViewModel

import (
	DTOLogin "wan-api-verify-user/DTO/LoginDTO"
	"wan-api-verify-user/ViewModel"
)

type LoginViewModel struct {
	ViewModel.CommonResponse
	Data *DTOLogin.LoginOutputDTO `json:"data"`
	Meta *struct { 
		NextStep []string `json:"next_steps"`
	} `json:"meta"`
}