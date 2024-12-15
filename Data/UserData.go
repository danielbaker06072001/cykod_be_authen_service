package Data

import (
	"fmt"
	"hash/crc32"
	"strconv"
	"time"
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

func (UserData *UserData) CheckUserExistsActive(username string) (error, string ,string) {
	var db_redis = UserData.DB_REDIS
	var thresholdTime = float64(time.Now().Unix() - 2592000) // 30 days in seconds
	
	// ? Step 1: Hash the username using CRC32 (no duplicate should appear, if it does, we can let user know to recreate new username)
	usernamehash := crc32.ChecksumIEEE([]byte(username))
	userkey := "active_user" 
	member := strconv.FormatUint(uint64(usernamehash), 10)
	score, err := db_redis.ZScore(db_redis.Context(), userkey, member).Result()
	if err == redis.Nil {
		return fmt.Errorf("username does not exist"), "", ""
	} else if err != nil {
		return err, "", ""
	}

	// ? Step 2: Check if the user is active
	if score > thresholdTime { 
		data, err := db_redis.HGetAll(db_redis.Context(), "u:"+member).Result()
		if err != nil {
			return err, "", ""
		}

		// Retrieve passhash and salt from Redis
		passhash := data["passhash"]
		salt := data["salt"]
		if passhash == "" && salt == "" {
			return fmt.Errorf("error retrieving user data"), "",""
		}
		return nil, passhash, salt
	} 
	
	
	return fmt.Errorf("user is not active"), "", ""
}

/*
	* Set the user to active in the Redis after they logged in
	? Active user will be set with time to live (TTL) of 30 days
*/
func (UserData *UserData) SetUserActive(username string, passHash string, salt string) (error) {
	var db_redis = UserData.DB_REDIS
	usernamehash := crc32.ChecksumIEEE([]byte(username))
	userkey := "active_user"
	member := strconv.FormatUint(uint64(usernamehash), 10)
	
	err := db_redis.ZAdd(db_redis.Context(), userkey, &redis.Z{
		Score: float64(time.Now().Unix()), 
		Member: member,
	}).Err()
	if err != nil {
		return err
	}

	// ? Store passhass and salt in Redis a key prefix
	err = db_redis.HMSet(db_redis.Context(), "u:"+member, map[string]interface{}{
		"passhash": passHash,
		"salt": salt,
	}).Err()
	if err != nil {
		return err
	}

	// ! this is cron job to remove the user from the active_user set (passhash)
	// err = db_redis.Expire(ctx, "user:"+member, 30*24*time.Hour).Err()
	// if err != nil {
	// 	return err
	// }

	// ! this is cron job to remove the user from the zAdd 
	// err = db_redis.Expire(ctx, "active_user", 30*24*time.Hour).Err()
	// if err != nil {
	// 	return err
	// }
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