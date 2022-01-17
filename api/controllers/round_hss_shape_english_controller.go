package controllers

import (
	"db/api/domain"
	"db/api/services"
	"db/api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type roundHssShapeEnglishController struct {

}

var (
	RoundHssShapeEnglishController *roundHssShapeEnglishController
)

func init()  {
	RoundHssShapeEnglishController = &roundHssShapeEnglishController{}
}

func (controller *roundHssShapeEnglishController) GetRoundHsShapeEnglish(c *gin.Context) {
	roundHsShapeEnglishData, err := services.RoundHssShapeEnglishService.Get()

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "Error when trying to fetch english round hs shape data : " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, roundHsShapeEnglishData)
}

func (controller *roundHssShapeEnglishController) AddRoundHssShapeEnglish(c *gin.Context) {
	var body domain.CreateRoundHssShapeEnglish

	if err := c.ShouldBindJSON(&body); err != nil {
		apiErr := &utils.ApplicationError{
			Message: "body must conform to format",
			StatusCode: http.StatusBadRequest,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}

	err := services.RoundHssShapeEnglishService.Create(body.RoundHssShapeName)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "Error when attempting to insert english round hss shape date : " + err.Error(),
			StatusCode: http.StatusBadRequest,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
	}
	utils.Respond(c, http.StatusOK, "Successfully inserted english round hss shape data!")
}
