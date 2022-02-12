package calculations

import (
	"db/api/domain"
	"db/api/services"
	"db/api/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type steelArgumentsDesignControllerEnglish struct {

}

var (
	SteelArgumentsDesignControllerEnglish *steelArgumentsDesignControllerEnglish
)

func init() {
	SteelArgumentsDesignControllerEnglish = &steelArgumentsDesignControllerEnglish{}
}

func (controller *steelArgumentsDesignControllerEnglish) GetSteelArgumentsDesignEnglish(c *gin.Context) {

	var body []domain.CreateSteelDesignArguments
	var results []domain.BeamAnalysisDesignResultsEnglish
	var dataValues []domain.AdditionalResultsEnglish
	//var additionalResults []domain.SteelArgumentsDesignResults

	if err := c.ShouldBindJSON(&body); err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "calculation body must conform to format: " + err.Error(),
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}

	for _, v := range body {
		fmt.Println("The body in the controller ==== ", body)
		_, data, result, err := services.SteelArgumentServiceDesignEnglish.GetEnglish(&v)

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
		dataValues = append(dataValues, *data)
		//fmt.Println("the data values --- ", results)
		//additionalResults = append(additionalResults, *additionalDataResult)
	}
	utils.Respond(c, http.StatusOK, results)
}

//func (controller *steelArgumentsDesignController) GetBeamAnalysisResult(c *gin.Context) {
//	cShapeEnglishData, err := services.CShapeEnglishService.Get()
//
//	if err != nil {
//		apiErr := &utils.ApplicationError{
//			Message: "Error when trying to fetch english C shape data : " + err.Error(),
//			StatusCode: http.StatusInternalServerError,
//			Code: "bad_request",
//		}
//		utils.RespondError(c, apiErr)
//		return
//	}
//	utils.Respond(c, http.StatusOK, cShapeEnglishData)
//}
