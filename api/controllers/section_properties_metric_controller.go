package controllers

import (
	"db/api/domain"
	"db/api/services"
	"db/api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type sectionPropertiesMetricController struct {

}

var (
	SectionPropertiesMetricController *sectionPropertiesMetricController
)

func init()  {
	SectionPropertiesMetricController = &sectionPropertiesMetricController{}
}

func (controller *sectionPropertiesMetricController) GetSectionPropertiesMetric(c *gin.Context) {
	sectionPropertiesMetricData, err := services.SectionPropertiesMetricService.Get()

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "Error when trying to fetch metric section properties data : " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, sectionPropertiesMetricData)
}

func (controller *sectionPropertiesMetricController) AddSectionPropertiesMetric(c *gin.Context) {
	var body domain.CreateSectionPropertiesMetric

	if err := c.ShouldBindJSON(&body); err != nil {
		apiErr := &utils.ApplicationError{
			Message: "body must conform to format",
			StatusCode: http.StatusBadRequest,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}

	err := services.SectionPropertiesMetricService.Create(body.SectionName, body.A, body.J, body.Ixp, body.Iyp, body.Iw, body.Sxp, body.Syp)

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "Error when trying to insert metric section properties data : " + err.Error(),
			StatusCode: http.StatusBadRequest,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, "Successfully inserted metric section properties data!")
}