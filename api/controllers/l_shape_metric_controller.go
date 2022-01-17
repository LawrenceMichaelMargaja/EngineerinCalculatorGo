package controllers

import (
	"db/api/domain"
	"db/api/services"
	"db/api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type lShapeMetricController struct {

}

var (
	LShapeMetricController *lShapeMetricController
)

func init()  {
	LShapeMetricController = &lShapeMetricController{}
}

func (controller *lShapeMetricController) GetLShapeMetric(c *gin.Context) {
	lShapeMetricData, err := services.LShapeMetricService.Get()

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "Error when trying to fetch metric L shape data : " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, lShapeMetricData)
}

func (controller *lShapeMetricController) AddLShapeMetric(c *gin.Context) {
	var body domain.CreateLShapeMetric

	if err := c.ShouldBindJSON(&body); err != nil {
		apiErr := &utils.ApplicationError{
			Message: "body should conform to format",
			StatusCode: http.StatusBadRequest,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}

	err := services.LShapeMetricService.Create(body.LShapeMetricName)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "Error when trying to insert metric L shape data : " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, "Successfully inserted metric L shaped data!")
}
