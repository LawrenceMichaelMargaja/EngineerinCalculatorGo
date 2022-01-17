package services

import (
	"db/api/domain"
	"errors"
)

type pipeShapeEnglishService struct {

}

var (
	PipeShapeEnglishService *pipeShapeEnglishService
)

func init()  {
	PipeShapeEnglishService = &pipeShapeEnglishService{}
}

func (s *pipeShapeEnglishService) Get() (*[]domain.PipeShapeEnglish, error) {
	return domain.PipeShapeEnglishDao.Get()
}

func (s *pipeShapeEnglishService) Create(pipeShapeEnglishName string) error {
	if pipeShapeEnglishName == "" {
		return errors.New("english pipe shape name is invalid")
	}
	return domain.PipeShapeEnglishDao.Create(pipeShapeEnglishName)
}