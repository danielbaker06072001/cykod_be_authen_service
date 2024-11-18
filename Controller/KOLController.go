package Controller

import (
	"net/http"
	"wan-api-verify-user/DTO"
	"wan-api-verify-user/Service/KOL/Interface"
	"wan-api-verify-user/ViewModel"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type KOLController struct {
	KOLService Interface.IKOLService
}

func NewKOLController(context *gin.Engine, KOLServiceObject Interface.IKOLService) {

	KOLControllerObject := &KOLController{
		KOLService: KOLServiceObject,
	}

	context.GET("/kol/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	VerifyUserGroup := context.Group("/verify-user")
	{
		VerifyUserGroup.POST("/kol", func(c *gin.Context) {
			KOLControllerObject.UpdateKol(c)
		})
		VerifyUserGroup.POST("/client", func(c *gin.Context) {
			KOLControllerObject.UpdateClient(c)
		})
	}
}

func (KolController *KOLController) UpdateKol(echoCtx *gin.Context) error {
	var KolVM ViewModel.UpdateKolViewModel

	Guid := uuid.New().String()

	var input DTO.KOLInputDTO
	if err := echoCtx.Bind(&input); err != nil {
		KolVM.CommonUpdateResponse.Result = "Failed"
		KolVM.CommonUpdateResponse.Guid = Guid
		echoCtx.JSON(http.StatusBadRequest, KolVM)
		return nil
	}

	var params DTO.AddedParam = make(DTO.AddedParam)
	params["KolID"] = input.KolID
	params["VerificationStatus"] = *input.VerificationStatus
	if input.ImageUrl != nil {
		for _, p := range *input.ImageUrl {
			params[p.Key] = p.Value
		}
	}

	KolDto, err := KolController.KOLService.UpdateKol(params)
	if err != nil {
		KolVM.CommonUpdateResponse.Result = "Failed"
		KolVM.CommonUpdateResponse.ErrorMessage = err.Error()
		KolVM.CommonUpdateResponse.Guid = Guid
		echoCtx.JSON(http.StatusBadRequest, KolVM)
		return nil
	}

	// * Update successfully
	KolVM.CommonUpdateResponse.Result = "Success"
	KolVM.CommonUpdateResponse.Guid = Guid
	KolVM.Kol = KolDto
	echoCtx.JSON(http.StatusOK, KolVM)
	return nil
}

func (KolController *KOLController) UpdateClient(context *gin.Context) error {
	return nil
}
