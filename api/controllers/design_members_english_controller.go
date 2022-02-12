package controllers

import (
	"db/api/services"
	"db/api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type designMembersEnglishController struct {

}

var (
	DesignMembersEnglishController *designMembersEnglishController
)

func init()  {
	DesignMembersEnglishController = &designMembersEnglishController{}
}

func (controller *designMembersEnglishController) GetDesignMemberEnglish(c *gin.Context) {
	designMembersEnglishData, err := services.DesignMembersEnglishService.Get()

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "Error when trying to fetch english Design Members English data : " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, designMembersEnglishData)
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
