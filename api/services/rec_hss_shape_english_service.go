package services

import (
	"db/api/domain"
	"errors"
)

type recHssShapeEnglishService struct {

}

var (
	RecHssShapeEnglishService *recHssShapeEnglishService
)

func init()  {
	RecHssShapeEnglishService = &recHssShapeEnglishService{}
}

func (s *recHssShapeEnglishService) Get() (*[]domain.RecHsShapeEnglish, error) {
	return domain.RecHssShapeEnglishDao.Get()
}

func (s *recHssShapeEnglishService) Create(recHssShapeEnglishName string) error {
	if recHssShapeEnglishName == "" {
		return errors.New("english rec hss shape is invalid")
	}
	return domain.RecHssShapeEnglishDao.Create(recHssShapeEnglishName)
}
