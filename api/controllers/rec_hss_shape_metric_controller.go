package controllers

import (
	"db/api/domain"
	"db/api/services"
	"db/api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type recHssShapeMetricController struct {

}

var (
	RecHssShapeMetricController *recHssShapeMetricController
)

func init()  {
	RecHssShapeMetricController = &recHssShapeMetricController{}
}

func (controller *recHssShapeMetricController) GetRecHsShapeMetric(c *gin.Context) {
	recHsShapeMetricData, err := services.RecHssShapeMetricService.Get()

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "Error when trying to fetch metric rec hs shape data : " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, recHsShapeMetricData)
}

func (controller *recHssShapeMetricController) AddRecHssShapeMetric(c *gin.Context) {
	var body domain.CreateRecHssShapeMetric

	if err := c.ShouldBindJSON(&body); err != nil {
		apiErr := &utils.ApplicationError{
			Message: "body should conform to format",
			StatusCode: http.StatusBadRequest,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}

	err := services.RecHssShapeMetricService.Create(body.RecHssShapeMetricName)

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "Error when trying to insert metric rec hss shape data : " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, "Successfully inserted metric rec hss shape data!")
}