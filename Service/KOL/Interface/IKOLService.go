package Interface

import (
	"context"
	"wan-api-verify-user/DTO"
)

type IKOLService interface {
	UpdateKol(params DTO.AddedParam) (*DTO.KolDTO, error)
	UpdateClient(ctx context.Context)
}
