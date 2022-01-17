package controllers

import (
	"db/api/domain"
	"db/api/services"
	"db/api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type pipeShapeEnglishController struct {

}

var (
	PipeShapeEnglishController *pipeShapeEnglishController
)

func init()  {
	PipeShapeEnglishController = &pipeShapeEnglishController{}
}

func (controller *pipeShapeEnglishController) GetPipeShapeEnglish(c *gin.Context) {
	pipeShapeEnglishData, err := services.PipeShapeEnglishService.Get()

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "Error when trying to fetch english pipe shape data : " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, pipeShapeEnglishData)
}

func (controller *pipeShapeEnglishController) AddPipeShapeEnglish(c *gin.Context) {
	var body domain.CreatePipeShapeEnglish

	if err := c.ShouldBindJSON(&body); err != nil {
		apiErr := &utils.ApplicationError{
			Message: "body must conform to format",
			StatusCode: http.StatusBadRequest,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}

	err := services.PipeShapeEnglishService.Create(body.PipeShapeNameEnglish)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "Error when attempting to insert english pipe shape data : " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, "Successfully inserted english pipe shape data!")
}
