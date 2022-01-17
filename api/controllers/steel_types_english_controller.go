package controllers

import (
	"db/api/domain"
	"db/api/services"
	"db/api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type steelTypeEnglishController struct {

}

var (
	SteelTypeEnglishController *steelTypeEnglishController
)

func init()  {
	SteelTypeEnglishController = &steelTypeEnglishController{}
}

func (controller *steelTypeEnglishController) GetSteelTypeEnglish(c *gin.Context) {
	steelTypeEnglishData, err := services.SteelTypeEnglishService.Get()

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "Error when trying to fetch english steel type data : " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, steelTypeEnglishData)
}

func (controller *steelTypeEnglishController) AddSteelTypeEnglish(c *gin.Context) {
	var body domain.CreateSteelTypeEnglish

	if err := c.ShouldBindJSON(&body); err != nil {
		apiErr := &utils.ApplicationError{
			Message: "body must conform to format",
			StatusCode: http.StatusBadRequest,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}

	err := services.SteelTypeEnglishService.Create(body.SteelTypeEnglishName, body.SteelTypeEnglishE, body.SteelTypeEnglishFy, body.SteelTypeEnglishFu)

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "Error when attempting to insert english steel type data : " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, "Successfully inserted english steel type data!")
}