package Model

import "wan-api-verify-user/AppConfig/Common"

type UserProfile struct {
	UserID      int64  `json:"user_id" gorm:"primaryKey; column:user_id"`
	Username    string `json:"username" gorm:"column:username"`
	Email       string `json:"email" gorm:"column:email"`
	Password    string `json:"password" gorm:"column:password"`
	Salt 	  	string `json:"salt" gorm:"column:salt"`
	FirstName   string `json:"first_name" gorm:"column:first_name"`
	LastName    string `json:"last_name" gorm:"column:last_name"`
	Active      bool   `json:"active" gorm:"column:active"`
	Status      string `json:"status" gorm:"column:status"`
	CreatedBy   string `json:"created_by" gorm:"column:created_by"`
	CreatedDate string `json:"created_date" gorm:"column:created_date"`
	UpdatedBy   string `json:"updated_by" gorm:"column:updated_by"`
	UpdatedDate string `json:"updated_date" gorm:"column:updated_date"`
}

func (UserProfile) TableName() string {
	return Common.TABLE_USERPROFILE
}