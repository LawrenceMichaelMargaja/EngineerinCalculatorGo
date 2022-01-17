package services

import (
	"db/api/domain"
	"errors"
)

type twoLShapeEnglishService struct {

}

var (
	TwoLShapeEnglishService *twoLShapeEnglishService
)

func init()  {
	TwoLShapeEnglishService = &twoLShapeEnglishService{}
}

func (s *twoLShapeEnglishService) Get() (*[]domain.TwoLShapeEnglish, error) {
	return domain.TwoLShapeEnglishDao.Get()
}

func (s *twoLShapeEnglishService) Create(twoLShapeEnglishName string) error {
	if twoLShapeEnglishName == "" {
		return errors.New("english 2L shape name is invalid")
	}
	return domain.TwoLShapeEnglishDao.Create(twoLShapeEnglishName)
}
