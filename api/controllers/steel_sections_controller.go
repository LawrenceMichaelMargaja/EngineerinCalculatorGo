package controllers

type steelSectionsController struct {

}

var (
	SteelSectionsController *steelSectionsController
)

func init()  {
	SteelSectionsController = &steelSectionsController{}
}

//func (controller *steelSectionsController) GetSteelSections(c *gin.Context) {
//	steelSectionsData, err := services.SteelSectionService.Get()
//
//	if err != nil {
//		apiErr := &utils.ApplicationError{
//			Message: "Error when trying to fetch steel sections data : " + err.Error(),
//			StatusCode: http.StatusInternalServerError,
//			Code: "bad_request",
//		}
//		utils.RespondError(c, apiErr)
//		return
//	}
//	utils.Respond(c, http.StatusOK, steelSectionsData)
//}
//
//func (controller *steelSectionsController) AddSteelSections(c *gin.Context) {
//	var body domain.CreateSteelSection
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
////
//	err := services.SteelSectionService.Create(body.SteelTypeEnglishName, body.SteelTypeEnglishE, body.SteelTypeEnglishFy, body.SteelTypeEnglishFu)
//
//	if err != nil {
//		apiErr := &utils.ApplicationError{
//			Message: "Error when attempting to insert english steel type data : " + err.Error(),
//			StatusCode: http.StatusInternalServerError,
//			Code: "bad_request",
//		}
//		utils.RespondError(c, apiErr)
//		return
//	}
//	utils.Respond(c, http.StatusOK, "Successfully inserted english steel type data!")
//}