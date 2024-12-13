package Data

import (
	"fmt"
	"wan-api-verify-user/DTO"
	"wan-api-verify-user/Model"
	Interface "wan-api-verify-user/Service/User/Interafce"
	"wan-api-verify-user/Utils"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type UserData struct {
	DB_CONNECTION *gorm.DB
	DB_REDIS *redis.Client
}

func NewUserDataLayer(Conn *gorm.DB, ConnRedis *redis.Client) Interface.IUserData {
	return &UserData{
		DB_CONNECTION: Conn,
		DB_REDIS: ConnRedis,
	}
}

func (UserData *UserData) GetUserByUsername(username string) (*Model.UserProfile, error) {
	var db = UserData.DB_CONNECTION.Model(&Model.UserProfile{})
	var userprofileMode Model.UserProfile
	if err := db.Where(`"username" = ?`, username).First(&userprofileMode).Error; err != nil {
		return nil, err
	}
	return &userprofileMode, nil
}

func (UserData *UserData) GetUserByEmail(email string) (*Model.UserProfile, error) {
	var db = UserData.DB_CONNECTION.Model(&Model.UserProfile{})
	var userprofileMode Model.UserProfile
	if err := db.Where(`"email" = ?`, email).First(&userprofileMode).Error; err != nil {
		return nil, err
	}
	return &userprofileMode, nil
}

// * Check if the user already exists using Redis cache 
// ? We're going to the Filter Bloom pattern, using SET u:bit
func (UserData *UserData) CheckUserExists(username string, email string) (bool, error) {
	var db_redis = UserData.DB_REDIS

	// err := db_redis.SetBit(db_redis.Context(), "u:bit", 12345, 1).Err()
	// if err != nil {
	// 	return false, err
	// }

	// ? Get the value of the bit at offset x (to see if the user exists)
	_ , err := db_redis.GetBit(db_redis.Context(), "u:bit", 12345).Result()
	if err != nil {
		return false, err
	}

	return true, nil
}

func (UserData *UserData) CreateUser(params DTO.Param) (*Model.UserProfile, error) {
	var db = UserData.DB_CONNECTION.Model(&Model.UserProfile{})
	var userprofileModel Model.UserProfile
	

	userprofileModel.Username = Utils.ConvertInterface(params["Username"])
	userprofileModel.Email = Utils.ConvertInterface(params["Email"])
	userprofileModel.Password = Utils.ConvertInterface(params["Password"])
	userprofileModel.FirstName = Utils.ConvertInterface(params["FirstName"])
	userprofileModel.LastName = Utils.ConvertInterface(params["LastName"])
	userprofileModel.Active = true // ! Temporarily set active to true, in the future it will need email verification
	userprofileModel.Status = "ACTIVE"
	userprofileModel.CreatedBy = "SYSTEM"
	userprofileModel.CreatedDate = Utils.GetCurrentTime()

	fmt.Printf("UserprofileModel: %v", userprofileModel)

	
	if err := db.Create(&userprofileModel).Error; err != nil {
		return nil, err
	}
	return &userprofileModel, nil
}