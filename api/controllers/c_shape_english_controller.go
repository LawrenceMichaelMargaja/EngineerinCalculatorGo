package controllers

import (
	"db/api/domain"
	"db/api/services"
	"db/api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type cShapeEnglishController struct {

}

var (
	CShapeEnglishController *cShapeEnglishController
)

func init()  {
	CShapeEnglishController = &cShapeEnglishController{}
}

func (controller *cShapeEnglishController) GetCSShapeEnglish(c *gin.Context) {
	cShapeEnglishData, err := services.CShapeEnglishService.Get()

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "Error when trying to fetch english C shape data : " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, cShapeEnglishData)
}

func (controller *cShapeEnglishController) AddCShapeEnglish(c *gin.Context) {
	var body domain.CreateCShapeEnglish

	if err := c.ShouldBindJSON(&body); err != nil {
		apiErr := &utils.ApplicationError{
			Message: "body must conform to format",
			StatusCode: http.StatusBadRequest,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}

	err := services.CShapeEnglishService.Create(body.CShapeEnglishName)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "Error when trying to insert english C shape data : " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, "Successfully inserted english C shape data!")
}
