package ViewModel

import (
	"wan-api-verify-user/DTO"
)

type UpdateKolViewModel struct {
	CommonUpdateResponse
	KolId int64       `json:"KolId"`
	Kol   *DTO.KolDTO `json:"Kol"`
}
