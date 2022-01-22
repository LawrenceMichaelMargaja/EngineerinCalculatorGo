package calculations

import (
	"db/api/domain"
	"db/api/services"
	"db/api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type steelArgumentsController struct {
}

var (
	SteelArgumentsController *steelArgumentsController
)

func init() {
	SteelArgumentsController = &steelArgumentsController{}
}

func (controller *steelArgumentsController) GetSteelArguments(c *gin.Context) {

	var body []domain.CreateSteelArguments
	var results []domain.BeamAnalysisResults

	if err := c.ShouldBindJSON(&body); err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "calculation body must conform to format",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	for _, v := range body {
		_, result, err := services.SteelArgumentService.Get(&v)

		if err != nil {
			apiErr := &utils.ApplicationError{
				Message:    "Calculation Error when trying to insert UI data : " + err.Error(),
				StatusCode: http.StatusInternalServerError,
				Code:       "bad_request",
			}
			utils.RespondError(c, apiErr)
			return
		}
		results = append(results, *result)
	}
	utils.Respond(c, http.StatusOK, results)
}

func (controller *steelArgumentsController) GetBeamAnalysisResult(c *gin.Context) {
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
