package controllers

import (
	"db/api/domain"
	"db/api/services"
	"db/api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type pipeShapeMetricController struct {

}

var (
	PipeShapeMetricController *pipeShapeMetricController
)

func init()  {
	PipeShapeMetricController = &pipeShapeMetricController{}
}

func (controller *pipeShapeMetricController) GetPipeShapeMetric(c *gin.Context) {
	pipeShapeMetricData, err := services.PipeShapeMetricService.Get()

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "Error when trying to fetch metric pipe shape data : " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, pipeShapeMetricData)
}

func (controller *pipeShapeMetricController) AddPipeShapeMetric(c *gin.Context) {
	var body domain.CreatePipeShapeMetric

	if err := c.ShouldBindJSON(&body); err != nil {
		apiErr := &utils.ApplicationError{
			Message: "body must conform to format",
			StatusCode: http.StatusBadRequest,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}

	err := services.PipeShapeMetricService.Create(body.PipeShapeMetricName)

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "Error when trying to insert metric pipe shape data : " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, "Successfully inserted metric pipe shape data!")
}