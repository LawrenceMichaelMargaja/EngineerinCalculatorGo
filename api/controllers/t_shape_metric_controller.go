package controllers

import (
	"db/api/domain"
	"db/api/services"
	"db/api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type tShapeMetricController struct {
}

var (
	TShapeMetricController *tShapeMetricController
)

func init() {
	TShapeMetricController = &tShapeMetricController{}
}

func (controller *tShapeMetricController) GetTShapeMetricData(c *gin.Context) {
	tShapeMetricData, err := services.TShapeMetricService.Get()

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "Error when attempting to fetch the T-shape metric data : " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, tShapeMetricData)
}

func (controller *tShapeMetricController) AddTShapeMetric(c *gin.Context) {
	var body domain.CreateTShapeMetric
	if err := c.ShouldBindJSON(&body); err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "body must conform to format",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}

	err := services.TShapeMetricService.Create(body.TShapeMetricName)

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "Error when attempting to insert t-shape data : " + err.Error(),
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}

	utils.Respond(c, http.StatusOK, "Successfully added the t-shape")
}
