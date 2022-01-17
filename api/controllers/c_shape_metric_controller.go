package controllers

import (
	"db/api/domain"
	"db/api/services"
	"db/api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type cShapeMetricController struct {

}

var (
	CShapeMetricController *cShapeMetricController
)

func init()  {
	CShapeMetricController = &cShapeMetricController{}
}

func (controller *cShapeMetricController) GetCShapeMetric(c *gin.Context) {
	cShapeMetricData, err := services.CShapeMetricService.Get()

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "Error when trying to fetch metric C shape data : " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, cShapeMetricData)
}

func (controller *cShapeMetricController) AddCShapeMetric(c *gin.Context) {
	var body domain.CreateCShapeMetric

	if err := c.ShouldBindJSON(&body); err != nil {
		apiErr := &utils.ApplicationError{
			Message: "body must conform to the format",
			StatusCode: http.StatusBadRequest,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}

	err := services.CShapeMetricService.Create(body.CShapeMetricName)

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "Error when trying to insert metric C shape data : " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, "Successfully inserted metric C shape data!")
}
