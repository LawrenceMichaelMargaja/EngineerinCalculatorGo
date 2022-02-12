package services

import (
	"db/api/domain"
)

type steelArgumentServiceDesignEnglish struct {

}

var (
	SteelArgumentServiceDesignEnglish *steelArgumentServiceDesignEnglish
)

func init()  {
	SteelArgumentServiceDesignEnglish = &steelArgumentServiceDesignEnglish{}
}

func(s *steelArgumentServiceDesignEnglish) GetEnglish(c *domain.CreateSteelDesignArguments) (*domain.CreateSteelDesignArguments, *domain.AdditionalResultsEnglish, *domain.BeamAnalysisDesignResultsEnglish, error) {
	return domain.SteelArgumentsDesignDaoEnglish.GetEnglish(c)
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
