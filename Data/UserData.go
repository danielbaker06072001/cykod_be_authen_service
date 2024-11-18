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

func (UserData *UserData) GetUserByUsername(username string) (*Model.Userprofile, error) {
	var db = UserData.DB_CONNECTION.Model(&Model.Userprofile{})
	var userprofileMode Model.Userprofile
	if err := db.Where(`"username" = ?`, username).First(&userprofileMode).Error; err != nil {
		return nil, err
	}
	return &userprofileMode, nil
}

func (UserData *UserData) GetUserByEmail(email string) (*Model.Userprofile, error) {
	var db = UserData.DB_CONNECTION.Model(&Model.Userprofile{})
	var userprofileMode Model.Userprofile
	if err := db.Where(`"email" = ?`, email).First(&userprofileMode).Error; err != nil {
		return nil, err
	}
	return &userprofileMode, nil
}

func (UserData *UserData) CheckUserExists(username string, email string) (bool, error) {
	var db = UserData.DB_CONNECTION.Model(&Model.Userprofile{})
	var userprofileMode Model.Userprofile
	if err := db.Where(`"username" = ? OR "email" = ?`, username, email).First(&userprofileMode).Error; err != nil {
		return false, err
	}
	return true, nil
}