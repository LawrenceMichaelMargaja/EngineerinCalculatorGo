package controllers

import (
	"db/api/domain"
	"db/api/services"
	"db/api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type iShapeMetricController struct {

}

var (
	IShapeMetricController *iShapeMetricController
)

func init()  {
	IShapeMetricController = &iShapeMetricController{}
}

func (controller *iShapeMetricController) GetIShapeMetric(c *gin.Context) {
	iShapeMetricData, err := services.IShapeMetricService.Get()

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "Error when trying to fetch metric I shape data : " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, iShapeMetricData)
}

func (controller *iShapeMetricController) AddIShapeMetric(c *gin.Context) {
	var body domain.CreateIShapeMetric

	if err := c.ShouldBindJSON(&body); err != nil {
		apiErr := &utils.ApplicationError{
			Message: "body must conform to format",
			StatusCode: http.StatusBadRequest,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
	}

	err := services.IShapeMetricService.Create(body.IShapeMetricName)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "Error when trying to insert metric i shape data : " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, "Successfully inserted metric i shape data!")
}