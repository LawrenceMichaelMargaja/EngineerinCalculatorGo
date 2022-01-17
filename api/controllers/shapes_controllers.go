package controllers

import (
	"db/api/domain"
	"db/api/services"
	"db/api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type shapeController struct {
}

var (
	ShapeController *shapeController
)

func init() {
	ShapeController = &shapeController{}
}

func (controller *shapeController) GetShape(c *gin.Context) {
	shapes, err := services.ShapeService.GetShapes()

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "Error when attempting to fetch the shapes data : " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, shapes)
}


func (controller *shapeController) AddShape(c *gin.Context) {

	var body domain.CreateShapes

	if err := c.ShouldBindJSON(&body); err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "body must conform to format",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}

	err := services.ShapeService.Create(body.ShapeName, body.DbValue)

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "Error wen attempting to insert shape data",
			StatusCode: http.StatusInternalServerError,
			Code:       "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, "Successfully Added Shape")
	return
}
