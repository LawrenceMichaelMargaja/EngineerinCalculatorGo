package controllers

import (
	"db/api/domain"
	"db/api/services"
	"db/api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type sectionPropertiesEnglishController struct {

}

var (
	SectionPropertiesEnglishController *sectionPropertiesEnglishController
)

func init()  {
	SectionPropertiesEnglishController = &sectionPropertiesEnglishController{}
}

func (controller *sectionPropertiesEnglishController) GetSectionPropertiesEnglish(c *gin.Context) {
	sectionPropertiesEnglishData, err := services.SectionPropertiesEnglishService.Get()

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "Error when trying to fetch english section properties data : " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, sectionPropertiesEnglishData)
}

func (controller *sectionPropertiesEnglishController) AddSectionPropertiesEnglish(c *gin.Context) {
	var body domain.CreateSectionPropertiesEnglish

	if err := c.ShouldBindJSON(&body); err != nil {
		apiErr := &utils.ApplicationError{
			Message: "body must conform to format",
			StatusCode: http.StatusBadRequest,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}

	err := services.SectionPropertiesEnglishService.Create(body.SectionName, body.A, body.J, body.Ixp, body.Iyp, body.Iw, body.Sxp, body.Syp)

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "Error when attempting to insert english section properties data : " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, "Successfully inserted english section properties data!")
}