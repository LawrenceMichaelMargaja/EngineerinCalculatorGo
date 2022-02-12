package app

import (
	"db/api/controllers"
	"db/api/controllers/calculations"
)

func mapUrls()  {
	//CALCULATE POST REQUEST
	router.POST("/steelArgumentsDesignEnglish", calculations.SteelArgumentsDesignControllerEnglish.GetSteelArgumentsDesignEnglish)
	router.POST("/steelArgumentsDesign", calculations.SteelArgumentsDesignController.GetSteelArgumentsDesign)
	router.POST("/steelArguments", calculations.SteelArgumentsController.GetSteelArguments)

	//CALCULATE GET REQUEST
	router.GET("/steelArguments", calculations.SteelArgumentsController.GetSteelArguments)

	//POST REQUESTS
	router.POST("/shape", controllers.ShapeController.AddShape)
	router.POST("/tshapemetric", controllers.TShapeMetricController.AddTShapeMetric)
	router.POST("/tshapeenglish", controllers.TShapeEnglishController.AddTShapeEnglish)
	router.POST("/steeltypesmetric", controllers.SteelTypesMetricController.AddSteelTypeMetric)
	router.POST("/steeltypesenglish", controllers.SteelTypeEnglishController.AddSteelTypeEnglish)
	router.POST("/sectionpropertiesmetric", controllers.SectionPropertiesMetricController.AddSectionPropertiesMetric)
	router.POST("/sectionpropertiesenglish", controllers.SectionPropertiesEnglishController.AddSectionPropertiesEnglish)
	router.POST("/sectiondimensionmetric", controllers.SectionDimensionMetricController.AddSectionDimensionMetric)
	router.POST("/sectiondimensionenglish", controllers.SectionDimensionEnglishController.AddSectionDimensionEnglish)
	router.POST("/roundhsshapemetric", controllers.RoundHssShapeController.AddRoundHssShapeMetric)
	router.POST("/roundhsshapeenglish", controllers.RoundHssShapeEnglishController.AddRoundHssShapeEnglish)
	router.POST("/rechsshapemetric", controllers.RecHssShapeMetricController.AddRecHssShapeMetric)
	router.POST("/rechsshapeenglish", controllers.RecHssShapeEnglishController.AddRecHssShapeEnglish)
	router.POST("/pipeshapemetric", controllers.PipeShapeMetricController.AddPipeShapeMetric)
	router.POST("/pipeshapeenglish", controllers.PipeShapeEnglishController.AddPipeShapeEnglish)
	router.POST("/lshapemetric", controllers.LShapeMetricController.AddLShapeMetric)
	router.POST("/lshapeenglish", controllers.LShapeEnglishController.AddLShapeEnglish)
	router.POST("/ishapemetric", controllers.IShapeMetricController.AddIShapeMetric)
	router.POST("/ishapeenglish", controllers.IShapeEnglishController.AddIShapeEnglish)
	router.POST("/cshapemetric", controllers.CShapeMetricController.AddCShapeMetric)
	router.POST("/cshapeenglish", controllers.CShapeEnglishController.AddCShapeEnglish)
	router.POST("/2lshapemetric", controllers.TwoLShapeMetricController.AddTwoLShapeMetric)
	router.POST("/2lshapeenglish", controllers.TwoLShapeEnglishController.AddTwoLShapeEnglish)

	//GET REQUESTS
	//router.GET("/steelSections", controllers.SteelSectionsController.GetSteelSections)
	router.GET("/designMembersEnglish", controllers.DesignMembersEnglishController.GetDesignMemberEnglish)
	router.GET("/designMembersMetric", controllers.DesignMembersMetricController.GetDesignMemberMetric)
	router.GET("/shape", controllers.ShapeController.GetShape)
	router.GET("/tshapemetric", controllers.TShapeMetricController.GetTShapeMetricData)
	router.GET("/tshapeenglish", controllers.TShapeEnglishController.GetTShapeEngligh)
	router.GET("/steeltypesmetric", controllers.SteelTypesMetricController.GetSteelTypeMetric)
	router.GET("/steeltypesenglish", controllers.SteelTypeEnglishController.GetSteelTypeEnglish)
	router.GET("/sectionpropertiesmetric", controllers.SectionPropertiesMetricController.GetSectionPropertiesMetric)
	router.GET("/sectionpropertiesenglish", controllers.SectionPropertiesEnglishController.GetSectionPropertiesEnglish)
	router.GET("/sectiondimensionmetric", controllers.SectionDimensionMetricController.GetSectionDimensionMetric)
	router.GET("/sectiondimensionenglish", controllers.SectionDimensionEnglishController.GetSectionDimensionEnglish)
	router.GET("/roundhsshapemetric", controllers.RoundHssShapeController.GetRoundHsShapeMetric)
	router.GET("/roundhsshapeenglish", controllers.RoundHssShapeEnglishController.GetRoundHsShapeEnglish)
	router.GET("/rechsshapemetric", controllers.RecHssShapeMetricController.GetRecHsShapeMetric)
	router.GET("/rechsshapeenglish", controllers.RecHssShapeEnglishController.GetRecHsShapeEnglish)
	router.GET("/pipeshapemetric", controllers.PipeShapeMetricController.GetPipeShapeMetric)
	router.GET("/pipeshapeenglish", controllers.PipeShapeEnglishController.GetPipeShapeEnglish)
	router.GET("/lshapemetric", controllers.LShapeMetricController.GetLShapeMetric)
	router.GET("/lshapeenglish", controllers.LShapeEnglishController.GetLShapeEnglish)
	router.GET("/ishapemetric", controllers.IShapeMetricController.GetIShapeMetric)
	router.GET("/ishapeenglish", controllers.IShapeEnglishController.GetIShapeEnglish)
	router.GET("/cshapemetric", controllers.CShapeMetricController.GetCShapeMetric)
	router.GET("/cshapeenglish", controllers.CShapeEnglishController.GetCSShapeEnglish)
	router.GET("/2lshapemetric", controllers.TwoLShapeMetricController.GetTwoLShapeMetric)
	router.GET("/2lshapeenglish", controllers.TwoLShapeEnglishController.GetTwoLShapeEnglish)
}