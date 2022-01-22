package services

import (
	"db/api/domain"
)

type steelArgumentService struct {

}

var (
	SteelArgumentService *steelArgumentService
)

func init()  {
	SteelArgumentService = &steelArgumentService{}
}

func(s *steelArgumentService) Get(c *domain.CreateSteelArguments) (*domain.CreateSteelArguments, *domain.BeamAnalysisResults, error) {
	return domain.SteelArgumentsDao.Get(c)
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
