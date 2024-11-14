package Data

import (
	"wan-api-verify-user/Model"
	"wan-api-verify-user/Service/KOL/Interface"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type KOLDataLayer struct {
	DB_CONNECTION *gorm.DB
}

func NewKolDataLayer(Conn *gorm.DB) Interface.IKOLDataLayer {
	return &KOLDataLayer{
		DB_CONNECTION: Conn,
	}
}

func (KOLDataLayer *KOLDataLayer) GetKolByID(KolID int64) (*Model.KOL, error) {
	var err error
	var db = KOLDataLayer.DB_CONNECTION.Model(&Model.KOL{})
	var kolModel Model.KOL
	if err = db.Where(`"KolID" = ?`, KolID).Find(&kolModel).Error; err != nil {
		return nil, err
	}
	return &kolModel, nil
}

func (KolDataLayer *KOLDataLayer) UpdateKol(KolID int64, model *Model.KOL) (*Model.KOL, error) {
	var kolModel Model.KOL
	var err error
	var db = KolDataLayer.DB_CONNECTION.Model(&Model.KOL{})

	if err = db.Where(`"KolID" = ?`, KolID).Save(&model).Find(&kolModel).Error; err != nil {
		return nil, err
	}

	return &kolModel, nil
}

func (KolDataLayer *KOLDataLayer) UpdateClient(ctx context.Context) {
	return
}
