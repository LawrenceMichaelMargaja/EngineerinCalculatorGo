package controllers

import (
	"db/api/domain"
	"db/api/services"
	"db/api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type sectionDimensionMetricController struct {

}

var (
	SectionDimensionMetricController *sectionDimensionMetricController
)

func init()  {
	SectionDimensionMetricController = &sectionDimensionMetricController{}
}

func (controller *sectionDimensionMetricController) GetSectionDimensionMetric(c *gin.Context) {
	sectionDimensionMetricData, err := services.SectionDimensionMetricService.Get()

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "Error when trying to fetch metric section dimension data : " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, sectionDimensionMetricData)
}

func (controller *sectionDimensionMetricController) AddSectionDimensionMetric(c *gin.Context) {
	var body domain.CreateSectionDimensionMetric

	if err := c.ShouldBindJSON(&body); err != nil {
		apiErr := &utils.ApplicationError{
			Message: "body must conform to format",
			StatusCode: http.StatusBadRequest,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
	}

	err := services.SectionDimensionMetricService.Create(body.SectionDimensionName, body.SectionDimensionD, body.SectionDimensionB, body.SectionDimensionTW, body.SectionDimensionBF, body.SectionDimensionTF, body.SectionDimensionTB, body.SectionDimensionT, body.SectionDimensionR)

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "Error when attempting to insert metric section dimension data : " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, "Successfully inserted metric section dimension data!")
}