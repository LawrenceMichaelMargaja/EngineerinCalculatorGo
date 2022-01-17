package controllers

import (
	"db/api/domain"
	"db/api/services"
	"db/api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type twoLShapeMetricController struct {

}

var (
	TwoLShapeMetricController *twoLShapeMetricController
)

func init()  {
	TwoLShapeMetricController = &twoLShapeMetricController{}
}

func (controller *twoLShapeMetricController) GetTwoLShapeMetric(c *gin.Context) {
	twoLShapeMetricData, err := services.TwoLShapeMetricService.Get()

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "Error when trying to fetch metric 2L shape data : " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, twoLShapeMetricData)
}

func (controller *twoLShapeMetricController) AddTwoLShapeMetric(c *gin.Context) {
	var body domain.CreateTwoLShapeMetric

	if err := c.ShouldBindJSON(&body); err != nil {
		apiErr := &utils.ApplicationError{
			Message: "body must conform to format",
			StatusCode: http.StatusBadRequest,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}

	err := services.TwoLShapeMetricService.Create(body.TwoLShapeMetricName)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "Error when trying to insert metric 2L shape data : " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, "Successfully inserted metric 2L shape data!")
}
