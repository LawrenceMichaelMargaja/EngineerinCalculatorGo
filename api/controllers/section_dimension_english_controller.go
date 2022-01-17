package controllers

import (
	"db/api/domain"
	"db/api/services"
	"db/api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type sectionDimensionEnglishController struct {

}

var (
	SectionDimensionEnglishController *sectionDimensionEnglishController
)

func init()  {
	SectionDimensionEnglishController = &sectionDimensionEnglishController{}
}

func (controller *sectionDimensionEnglishController) GetSectionDimensionEnglish(c *gin.Context) {
	sectionDimensionEnglishData, err := services.SectionDimensionEnglishService.Get()

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "Error when trying to fetch english section dimension data : " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, sectionDimensionEnglishData)
}

func (controller *sectionDimensionEnglishController) AddSectionDimensionEnglish(c *gin.Context) {
	var body domain.CreateSectionDimensionEnglish

	if err := c.ShouldBindJSON(&body); err != nil {
		apiErr := &utils.ApplicationError{
			Message: "body must conform to format",
			StatusCode: http.StatusBadRequest,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}

	err := services.SectionDimensionEnglishService.Create(body.SectionDimensionEnglishName, body.SectionDimensionEnglishD, body.SectionDimensionEnglishB, body.SectionDimensionEnglishTW, body.SectionDimensionEnglishBF, body.SectionDimensionEnglishTF, body.SectionDimensionEnglishTB, body.SectionDimensionEnglishT, body.SectionDimensionEnglishR)

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "Error when attempting to insert english section dimension data : " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, "Successfully inserted english section dimension data!")
}