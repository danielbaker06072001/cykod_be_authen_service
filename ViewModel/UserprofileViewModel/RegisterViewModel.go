package ViewModel

import (
	DTO "wan-api-verify-user/DTO/RegisterDTO"
	"wan-api-verify-user/ViewModel"
)

type RegisterViewModel struct {
	ViewModel.CommonResponse
	Data *DTO.RegisterInputDTO `json:"data"`
	Meta *struct { 
		NextStep []string `json:"next_steps"`
	} `json:"meta"`
}