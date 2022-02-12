package controllers

import (
	"db/api/services"
	"db/api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type designMembersMetricController struct {

}

var (
	DesignMembersMetricController *designMembersMetricController
)

func init()  {
	DesignMembersMetricController = &designMembersMetricController{}
}

func (controller *designMembersMetricController) GetDesignMemberMetric(c *gin.Context) {
	designMembersMetricData, err := services.DesignMembersMetricService.Get()

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "Error when trying to fetch english Design Members Metric data : " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, designMembersMetricData)
}

//func (controller *iShapeEnglishController) AddIShapeEnglish(c *gin.Context) {
//	var body domain.CreateIShapeEnglish
//
//	if err := c.ShouldBindJSON(&body); err != nil {
//		apiErr := &utils.ApplicationError{
//			Message: "body must conform to format",
//			StatusCode: http.StatusBadRequest,
//			Code: "bad_request",
//		}
//		utils.RespondError(c, apiErr)
//		return
//	}
//
//	err := services.IShapeEnglishService.Create(body.IShapeEnglishName)
//	if err != nil {
//		apiErr := &utils.ApplicationError{
//			Message: "Error when trying to insert english I shape data : " + err.Error(),
//			StatusCode: http.StatusInternalServerError,
//			Code: "bad_request",
//		}
//		utils.RespondError(c, apiErr)
//		return
//	}
//	utils.Respond(c, http.StatusOK, "Successfully inserted english I shape data!")
//}
