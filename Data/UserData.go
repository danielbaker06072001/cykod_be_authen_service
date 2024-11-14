package Data

import (
	"wan-api-verify-user/Model"
	Interface "wan-api-verify-user/Service/User/Interafce"

	"gorm.io/gorm"
)

type UserData struct {
	DB_CONNECTION *gorm.DB
}

func NewUserDataLayer(Conn *gorm.DB) Interface.IUserData {
	return &UserData{
		DB_CONNECTION: Conn,
	}
}

func (UserData *UserData) GetUserByUsername(username string) (*Model.KOL, error) {
	panic("implement me")
}