package controllers

import (
	"db/api/domain"
	"db/api/services"
	"db/api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type twoLShapeEnglishController struct {

}

var (
	TwoLShapeEnglishController *twoLShapeEnglishController
)

func init()  {
	TwoLShapeEnglishController = &twoLShapeEnglishController{}
}

func (controller *twoLShapeEnglishController) GetTwoLShapeEnglish(c *gin.Context) {
	twoLShapeEnglishData, err := services.TwoLShapeEnglishService.Get()

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "Error when trying to fetch english 2L shape data : " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, twoLShapeEnglishData)
}

func (controller *twoLShapeEnglishController) AddTwoLShapeEnglish(c *gin.Context) {
	var body domain.CreateTwoLShapeEnglish

	if err := c.ShouldBindJSON(&body); err != nil {
		apiErr := &utils.ApplicationError{
			Message: "body must conform to format",
			StatusCode: http.StatusBadRequest,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}

	err := services.TwoLShapeEnglishService.Create(body.TwoLShapeEnglishName)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "Error when trying to insert english 2L shape data : " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, "Successfully inserted english 2L shape data!")
}
