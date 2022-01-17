package services

import (
	"db/api/domain"
	"errors"
)

type steelTypeEnglishService struct {

}

var (
	SteelTypeEnglishService *steelTypeEnglishService
)

func init()  {
	SteelTypeEnglishService = &steelTypeEnglishService{}
}

func (s *steelTypeEnglishService) Get() (*[]domain.SteelTypeEnglish, error) {
	return domain.SteelTypeEnglishDao.Get()
}

func (s *steelTypeEnglishService) Create(steelTypeEnglishName string, steelTypeEnglishE string, steelTypeEnglishFy string, steelTypeEnglishFu string) error {
	if steelTypeEnglishName == "" {
		return errors.New("english steel type name is invalid")
	}
	if steelTypeEnglishE == "" {
		return errors.New("english steel type E is invalid")
	}
	if steelTypeEnglishFy == "" {
		return errors.New("english steel type Fy is invalid")
	}
	if steelTypeEnglishFu == "" {
		return errors.New("english steel type Fu is invalid")
	}
	return domain.SteelTypeEnglishDao.Create(steelTypeEnglishName, steelTypeEnglishE, steelTypeEnglishFy, steelTypeEnglishFu)
}
