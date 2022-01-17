package controllers

import (
	"db/api/domain"
	"db/api/services"
	"db/api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type lShapeEnglishController struct {
	
}

var (
	LShapeEnglishController *lShapeEnglishController
)

func init() {
	LShapeEnglishController = &lShapeEnglishController{}
}

func (controller *lShapeEnglishController) GetLShapeEnglish(c *gin.Context) {
	lShapeEnglishData, err := services.LShapeEnglishService.Get()

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "Error when trying to fetch english L shape data : " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, lShapeEnglishData)
}

func (controller *lShapeEnglishController) AddLShapeEnglish(c *gin.Context) {
	var body domain.CreateLShapeEnglish

	if err := c.ShouldBindJSON(&body); err != nil {
		apiErr := &utils.ApplicationError{
			Message: "body must conform to format",
			StatusCode: http.StatusBadRequest,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}

	err := services.LShapeEnglishService.Create(body.LShapeEnglishName)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "Error when trying to insert english L shaped data : " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, "Successfully inserted english L shaped data!")
}