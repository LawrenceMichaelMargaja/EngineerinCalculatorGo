package services

import (
	"db/api/domain"
	"errors"
)

type cShapeEnglishService struct {

}

var (
	CShapeEnglishService *cShapeEnglishService
)

func init()  {
	CShapeEnglishService = &cShapeEnglishService{}
}

func (s *cShapeEnglishService) Get() (*[]domain.CShapeEnglish, error) {
	return domain.CShapeEnglishDao.Get()
}

func (s *cShapeEnglishService) Create(cShapeEnglishName string) error {
	if cShapeEnglishName == "" {
		return errors.New("english C shape name is invalid")
	}
	return domain.CShapeEnglishDao.Create(cShapeEnglishName)
}
