package Data

import (
	"fmt"
	"hash/crc32"
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
// @return error (if the user already exists return error)
func (UserData *UserData) CheckUserExists(username string, email string) (error) {
	var db_redis = UserData.DB_REDIS

	// ? Step 1: Hash the username and email using CRC32 (no duplicate should appear, if it does, we can let user know to recreate new username or email)
	usernameHash := crc32.ChecksumIEEE([]byte(username))
	emailHash := crc32.ChecksumIEEE([]byte(email))

	// ? Step 2: Check if the bits of (Username, Email) in the Redis
	userExist := db_redis.GetBit(db_redis.Context(), "u:bit", int64(usernameHash)).Val()
	if userExist != 0 {
		return fmt.Errorf("username already exists")
	}
	emailExist := db_redis.GetBit(db_redis.Context(), "u:bit", int64(emailHash)).Val()
	if emailExist != 0 {
		return fmt.Errorf("email already exists")
	}
	return nil
}

func (UserData *UserData) CreateUser(params DTO.Param) (*Model.UserProfile, error) {
	var db = UserData.DB_CONNECTION.Model(&Model.UserProfile{})
	var userprofileModel Model.UserProfile
	
	// * Create new User in the database
	userprofileModel.Username = Utils.ConvertInterface(params["Username"])
	userprofileModel.Email = Utils.ConvertInterface(params["Email"])
	userprofileModel.Password = Utils.ConvertInterface(params["Password"])
	userprofileModel.Salt = Utils.ConvertInterface(params["Salt"])
	userprofileModel.FirstName = Utils.ConvertInterface(params["FirstName"])
	userprofileModel.LastName = Utils.ConvertInterface(params["LastName"])
	userprofileModel.Active = true // ! Temporarily set active to true, in the future it will need email verification
	userprofileModel.Status = "ACTIVE"
	userprofileModel.CreatedDate = Utils.GetCurrentTime()

	if err := db.Create(&userprofileModel).Error; err != nil {
		return nil, err
	}

	// * Create new User in the Redis
	// ? Step 1: Hash the username and email using CRC32 (no duplicate should appear, if it does, we can let user know to recreate new username or email)
	usernameHash := crc32.ChecksumIEEE([]byte(userprofileModel.Username))
	emailHash := crc32.ChecksumIEEE([]byte(userprofileModel.Email))

	// ? Step 2: Set the bits of (Username, Email) in the Redis 
	err := UserData.DB_REDIS.SetBit(UserData.DB_REDIS.Context(), "u:bit", int64(usernameHash), 1).Err()
	if err != nil {
		return nil, err
	}
	err = UserData.DB_REDIS.SetBit(UserData.DB_REDIS.Context(), "u:bit", int64(emailHash), 1).Err()
	if err != nil {
		return nil, err
	}

	return &userprofileModel, nil
}