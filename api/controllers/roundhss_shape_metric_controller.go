package controllers

import (
	"db/api/domain"
	"db/api/services"
	"db/api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type roundHssShapeMetricController struct {

}

var (
	RoundHssShapeController *roundHssShapeMetricController
)

func init()  {
	RoundHssShapeController = &roundHssShapeMetricController{}
}

func (controller *roundHssShapeMetricController) GetRoundHsShapeMetric(c *gin.Context) {
	roundHsShapeMetricData, err := services.RoundHssShapeMetricService.Get()

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "Error when trying to fetch metric round hs shape data : " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, roundHsShapeMetricData)
}

func (controller *roundHssShapeMetricController) AddRoundHssShapeMetric(c *gin.Context) {
	var body domain.CreateRoundHssShape

	if err := c.ShouldBindJSON(&body); err != nil {
		apiErr := &utils.ApplicationError{
			Message: "body must conform to format",
			StatusCode: http.StatusBadRequest,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}

	err := services.RoundHssShapeMetricService.Create(body.RoundHssShapeName)

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "Error when trying to insert round hss shape data : " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, "Successfully inserted metric round hss shape!")
}