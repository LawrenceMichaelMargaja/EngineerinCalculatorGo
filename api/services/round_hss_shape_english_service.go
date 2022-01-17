package services

import (
	"db/api/domain"
	"errors"
)

type roundHssShapeEnglishService struct {

}

var (
	RoundHssShapeEnglishService *roundHssShapeEnglishService
)

func init()  {
	RoundHssShapeEnglishService = &roundHssShapeEnglishService{}
}

func (s *roundHssShapeEnglishService) Get() (*[]domain.RoundHsShapeEnglish, error) {
	return domain.RoundHssShapeEnglishDao.Get()
}

func (s *roundHssShapeEnglishService) Create(roundHssShapeEnglishName string) error {
	if roundHssShapeEnglishName == "" {
		return errors.New("english round hss name is invalid")
	}
	return domain.RoundHssShapeEnglishDao.Create(roundHssShapeEnglishName)
}
