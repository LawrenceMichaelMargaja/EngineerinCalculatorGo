package services

import (
	"db/api/domain"
	"errors"
)

type iShapeEnglishService struct {

}

var (
	IShapeEnglishService *iShapeEnglishService
)

func init()  {
	IShapeEnglishService = &iShapeEnglishService{}
}

func (s *iShapeEnglishService) Get() (*[]domain.IShapeEnglish, error) {
	return domain.IShapeEnglishDao.Get()
}

func (s *iShapeEnglishService) Create(iShapeEnglishName string) error {
	if iShapeEnglishName == "" {
		return errors.New("english i shape name is invalid")
	}
	return domain.IShapeEnglishDao.Create(iShapeEnglishName)
}
