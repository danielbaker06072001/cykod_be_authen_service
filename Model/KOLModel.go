package Model

import (
	"time"
	"wan-api-verify-user/AppConfig/Common"
)

type KOL struct {
	KolID                int64      `json:"kolID"  gorm:"primaryKey;  column:KolID"`
	UserProfileID        int64      `json:"userProfileID" gorm:"column:UserProfileID"`
	Language             *string    `json:"language" gorm:"column:Language"`
	Education            *string    `json:"education" gorm:"column:Education"`
	ExpectedSalary       *int64     `json:"expectedSalary" gorm:"column:ExpectedSalary"`
	ExpectedSalaryEnable *bool      `json:"expectedSalaryEnable" gorm:"column:ExpectedSalaryEnable"`
	ChannelSettingTypeID *int64     `json:"channelSettingTypeID" gorm:"column:ChannelSettingTypeID"`
	IDFrontURL           *string    `json:"iDFrontURL" gorm:"column:IDFrontURL"`
	IDBackURL            *string    `json:"iDBackURL" gorm:"column:IDBackURL"`
	PortraitURL          *string    `json:"portraitURL" gorm:"column:PortraitURL"`
	RewardID             *int64     `json:"rewardID" gorm:"column:RewardID"`
	PaymentMethodID      *int64     `json:"paymentMethodID" gorm:"column:PaymentMethodID"`
	TestimonialsID       *int64     `json:"testimonialsID" gorm:"column:TestimonialsID"`
	VerificationStatus   *bool      `json:"verificationStatus" gorm:"column:VerificationStatus"`
	Enabled              bool       `json:"enabled" gorm:"column:Enabled"`
	ActiveDate           *time.Time `json:"activeDate" gorm:"column:ActiveDate"`
	Active               *bool      `json:"active" gorm:"column:Active"`
	CreatedBy            *string    `json:"createdBy" gorm:"column:CreatedBy"`
	CreatedDate          *time.Time `json:"createdDate" gorm:"column:CreatedDate"`
	ModifiedBy           *string    `json:"modifiedBy" gorm:"column:ModifiedBy"`
	ModifiedDate         *time.Time `json:"modifiedDate" gorm:"column:ModifiedDate"`
	IsRemove             bool       `json:"isRemove" gorm:"column:IsRemove"`
	IsOnBoarding         bool       `json:"isOnBoarding" gorm:"column:IsOnBoarding"`
	Code                 *string    `json:"code" gorm:"column:Code"`
	PortraitRightURL     *string    `json:"portraitRightURL" gorm:"column:PortraitRightURL"`
	PortraitLeftURL      *string    `json:"portraitLeftURL" gorm:"column:PortraitLeftURL"`
}

func (KOL) TableName() string {
	return Common.TABLE_KOL
}
