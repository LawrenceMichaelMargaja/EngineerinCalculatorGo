package services

import (
	"db/api/domain"
	"errors"
)

type tShapeEnglishService struct {

}

var (
	TShapeEnglishService *tShapeEnglishService
)

func init()  {
	TShapeEnglishService = &tShapeEnglishService{}
}

func (s *tShapeEnglishService) Get() (*[]domain.TShapeEnglish, error) {
	return domain.TShapeEnglishDao.Get()
}

func (s *tShapeEnglishService) Create(tShapeEnglishName string) error {
	if tShapeEnglishName == "" {
		return errors.New("T shape name is invalid")
	}
	return domain.TShapeEnglishDao.Create(tShapeEnglishName)
}