package Interface

import (
	"context"
	"wan-api-verify-user/Model"
)

type IKOLDataLayer interface {
	GetKolByID(KolID int64) (*Model.KOL, error)
	UpdateKol(KolID int64, model *Model.KOL) (*Model.KOL, error)
	UpdateClient(ctx context.Context)
}
