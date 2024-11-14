package DTO

import (
	"time"
)

type KolDTO struct {
	KolID                int64      `json:"kolID"`
	UserProfileID        int64      `json:"userProfileID"`
	Language             *string    `json:"language"`
	Education            *string    `json:"education"`
	ExpectedSalary       *int64     `json:"expectedSalary"`
	ExpectedSalaryEnable *bool      `json:"expectedSalaryEnable"`
	ChannelSettingTypeID *int64     `json:"channelSettingTypeID"`
	IDFrontURL           *string    `json:"iDFrontURL"`
	IDBackURL            *string    `json:"iDBackURL"`
	PortraitURL          *string    `json:"portraitURL"`
	RewardID             *int64     `json:"rewardID"`
	PaymentMethodID      *int64     `json:"paymentMethodID"`
	TestimonialsID       *int64     `json:"testimonialsID"`
	VerificationStatus   *bool      `json:"verificationStatus"`
	Enabled              bool       `json:"enabled"`
	ActiveDate           *time.Time `json:"activeDate"`
	Active               *bool      `json:"active"`
	CreatedBy            *string    `json:"createdBy"`
	CreatedDate          *time.Time `json:"createdDate"`
	ModifiedBy           *string    `json:"modifiedBy"`
	ModifiedDate         *time.Time `json:"modifiedDate"`
	IsRemove             bool       `json:"isRemove"`
	IsOnBoarding         bool       `json:"isOnBoarding"`
	Code                 *string    `json:"code"`
	PortraitRightURL     *string    `json:"portraitRightURL"`
	PortraitLeftURL      *string    `json:"portraitLeftURL"`
}

type KOLInputDTO struct {
	KolID              string           `json:"kolID"`
	VerificationStatus *string          `json:"verificationStatus"`
	ImageUrl           *[]ImageInputDTO `json:"imageUrl"`
}

type ImageInputDTO struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
