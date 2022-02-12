package services

import (
	"db/api/domain"
)

type steelArgumentServiceDesign struct {

}

var (
	SteelArgumentServiceDesign *steelArgumentServiceDesign
)

func init()  {
	SteelArgumentServiceDesign = &steelArgumentServiceDesign{}
}

func(s *steelArgumentServiceDesign) Get(c *domain.CreateSteelDesignArguments) (*domain.CreateSteelDesignArguments, *domain.AdditionalResults, *domain.BeamAnalysisDesignResults, error) {
	return domain.SteelArgumentsDesignDao.Get(c)
}

//func (s *steelArgumentService) Get() (*[]domain.CreateSteelArguments, error) {
//	return domain.SteelArgumentsDao.Get()
//}
//
//func (s *steelArgumentService) Create(steelTypeEnglishName string, steelTypeEnglishE string, steelTypeEnglishFy string, steelTypeEnglishFu string) error {
//	if steelTypeEnglishName == "" {
//		return errors.New("english steel type name is invalid")
//	}
//	if steelTypeEnglishE == "" {
//		return errors.New("english steel type E is invalid")
//	}
//	if steelTypeEnglishFy == "" {
//		return errors.New("english steel type Fy is invalid")
//	}
//	if steelTypeEnglishFu == "" {
//		return errors.New("english steel type Fu is invalid")
//	}
//	return domain.SteelTypeEnglishDao.Create(steelTypeEnglishName, steelTypeEnglishE, steelTypeEnglishFy, steelTypeEnglishFu)
//}
