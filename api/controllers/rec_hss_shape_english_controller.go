package controllers

import (
	"db/api/domain"
	"db/api/services"
	"db/api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type recHssShapeEnglishController struct {

}

var (
	RecHssShapeEnglishController *recHssShapeEnglishController
)

func init()  {
	RecHssShapeEnglishController = &recHssShapeEnglishController{}
}

func (controller *recHssShapeEnglishController) GetRecHsShapeEnglish(c *gin.Context) {
	recHsShapeEnglishData, err := services.RecHssShapeEnglishService.Get()

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "Error when trying to fetch english rec hs shape data : " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, recHsShapeEnglishData)
}

func (controller *recHssShapeEnglishController) AddRecHssShapeEnglish(c *gin.Context) {
	var body domain.CreateRecHssShapeEnglish

	if err := c.ShouldBindJSON(&body); err != nil {
		apiErr := &utils.ApplicationError{
			Message: "body must conform to format",
			StatusCode: http.StatusBadRequest,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}

	err := services.RecHssShapeEnglishService.Create(body.RecHssShapeEnglishName)

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "Error when attempting to insert english recHSS shape data : " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, "Successfully inserted english rec hss shape data!")
}
