package controllers

import (
	"db/api/domain"
	"db/api/services"
	"db/api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type tShapeEnglishController struct {
}

var (
	TShapeEnglishController *tShapeEnglishController
)

func init() {
	TShapeEnglishController = &tShapeEnglishController{}
}

func (controller *tShapeEnglishController) GetTShapeEngligh(c *gin.Context) {
	tShapeEnglishData, err := services.TShapeEnglishService.Get()

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "Error when trying to fetch english T-shape data : " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, tShapeEnglishData)
}

func (controller *tShapeEnglishController) AddTShapeEnglish(c *gin.Context) {
	var body domain.CreateTShapeEnglish
	if err := c.ShouldBindJSON(&body); err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "body must conform to format",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}

	err := services.TShapeEnglishService.Create(body.TShapeEnglishName)

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "Error when attempting to insert t-shape data : " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, "Successfully inserted t-shape data!")
}
