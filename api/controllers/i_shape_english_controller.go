package controllers

import (
	"db/api/domain"
	"db/api/services"
	"db/api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type iShapeEnglishController struct {

}

var (
	IShapeEnglishController *iShapeEnglishController
)

func init()  {
	IShapeEnglishController = &iShapeEnglishController{}
}

func (controller *iShapeEnglishController) GetIShapeEnglish(c *gin.Context) {
	iShapeEnglishData, err := services.IShapeEnglishService.Get()

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "Error when trying to fetch english I shape data : " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, iShapeEnglishData)
}

func (controller *iShapeEnglishController) AddIShapeEnglish(c *gin.Context) {
	var body domain.CreateIShapeEnglish

	if err := c.ShouldBindJSON(&body); err != nil {
		apiErr := &utils.ApplicationError{
			Message: "body must conform to format",
			StatusCode: http.StatusBadRequest,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}

	err := services.IShapeEnglishService.Create(body.IShapeEnglishName)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "Error when trying to insert english I shape data : " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, "Successfully inserted english I shape data!")
}
