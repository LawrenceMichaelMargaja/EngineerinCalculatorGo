package controllers

import (
	"db/api/domain"
	"db/api/services"
	"db/api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type steelTypesMetricController struct {

}

var (
	SteelTypesMetricController *steelTypesMetricController
)

func init()  {
	SteelTypesMetricController = &steelTypesMetricController{}
}

func (controller *steelTypesMetricController) GetSteelTypeMetric(c *gin.Context) {
	steelTypeMetricData, err := services.SteelTypeMetricService.Get()

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "Error when trying to fetch metric steel type data : " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, steelTypeMetricData)
}

func (controller *steelTypesMetricController) AddSteelTypeMetric(c *gin.Context) {
	var body domain.CreateSteelTypeMetric

	if err := c.ShouldBindJSON(&body); err != nil {
		apiErr := &utils.ApplicationError{
			Message: "body must conform to format",
			StatusCode: http.StatusBadRequest,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}

	err := services.SteelTypeMetricService.Create(body.SteelTypeMetricName, body.SteelTypeMetricE, body.SteelTypeMetricFy, body.SteelTypeMetricFu)

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "Error when attempting to insert metric steel type data",
			StatusCode: http.StatusInternalServerError,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}

	utils.Respond(c, http.StatusOK, "Successfully inserted metric steel type data!")
}