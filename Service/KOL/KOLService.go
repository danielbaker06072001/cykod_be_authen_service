package Service

import (
	"context"
	"wan-api-verify-user/DTO"
	"wan-api-verify-user/Model"
	"wan-api-verify-user/Service/KOL/Interface"
	"wan-api-verify-user/Utils"
)

type KOLService struct {
	KOLDL Interface.IKOLDataLayer
}

func NewKOLService(KolDataLayer Interface.IKOLDataLayer) Interface.IKOLService {
	return &KOLService{
		KOLDL: KolDataLayer,
	}
}

func (service *KOLService) UpdateKol(params DTO.AddedParam) (*DTO.KolDTO, error) {

	// KolID := params.KolID
	// VerificationStatus := params.VerificationStatus
	var KolId = Utils.StringToInt64(params["KolID"])
	KolModel, err := service.KOLDL.GetKolByID(KolId)
	if err != nil {
		return nil, err
	}

	for key, value := range params {
		switch key {
		case "IDFrontURL":
			if value != "" {
				KolModel.IDFrontURL = &value
			}
		case "IDBackURL":
			if value != "" {
				KolModel.IDBackURL = &value
			}
		case "PortraitURL":
			if value != "" {
				KolModel.PortraitURL = &value
			}
		case "PortraitRightURL":
			if value != "" {
				KolModel.PortraitRightURL = &value
			}
		case "PortraitLeftURL":
			if value != "" {
				KolModel.PortraitLeftURL = &value
			}
		}

	}

	updatedKolModel, err := service.KOLDL.UpdateKol(KolId, KolModel)
	if err != nil {
		return nil, err
	}

	KolDto := ConvertToKolDto(*updatedKolModel)
	return KolDto, nil
}

func (service *KOLService) UpdateClient(ctx context.Context) {

	return
}

func ConvertToKolDto(data Model.KOL) *DTO.KolDTO {
	var dto DTO.KolDTO
	dto.KolID = data.KolID
	dto.UserProfileID = data.UserProfileID
	dto.Language = data.Language
	dto.Education = data.Education
	dto.ExpectedSalary = data.ExpectedSalary
	dto.ExpectedSalaryEnable = data.ExpectedSalaryEnable
	dto.ChannelSettingTypeID = data.ChannelSettingTypeID
	dto.IDFrontURL = data.IDFrontURL
	dto.IDBackURL = data.IDBackURL
	dto.PortraitURL = data.PortraitURL
	dto.RewardID = data.RewardID
	dto.PaymentMethodID = data.PaymentMethodID
	dto.TestimonialsID = data.TestimonialsID
	dto.VerificationStatus = data.VerificationStatus
	dto.Enabled = data.Enabled
	dto.ActiveDate = data.ActiveDate
	dto.Active = data.Active
	dto.CreatedBy = data.CreatedBy
	dto.CreatedDate = data.CreatedDate
	dto.ModifiedBy = data.ModifiedBy
	dto.ModifiedDate = data.ModifiedDate
	dto.IsRemove = data.IsRemove
	dto.IsOnBoarding = data.IsOnBoarding
	dto.Code = data.Code
	dto.PortraitRightURL = data.PortraitRightURL
	dto.PortraitLeftURL = data.PortraitLeftURL
	return &dto
}
