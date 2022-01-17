package services

import (
	"db/api/domain"
	"errors"
)

type lShapeEnglishService struct {

}

var (
	LShapeEnglishService *lShapeEnglishService
)

func init()  {
	LShapeEnglishService = &lShapeEnglishService{}
}

func (s *lShapeEnglishService) Get() (*[]domain.LShapeEnglish, error) {
	return domain.LShapeEnglishDao.Get()
}

func (s *lShapeEnglishService) Create(lShapeEnglishName string) error {
	if lShapeEnglishName == "" {
		return errors.New("english l shaped name is invalid")
	}
	return domain.LShapeEnglishDao.Create(lShapeEnglishName)
}
