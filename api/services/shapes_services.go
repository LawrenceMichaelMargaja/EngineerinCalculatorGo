package services

import (
	"db/api/domain"
	"errors"
)

type shapesService struct {

}

var (
	ShapeService *shapesService
)

func init()  {
	ShapeService = &shapesService{}
}

func (s *shapesService) Create(shapeName string, DbValue string) error {
	if shapeName == "" {
		return errors.New("shape name is invalid")
	}
	if DbValue == "" {
		return errors.New("DB value is invalid")
	}
	return domain.ShapeDao.Create(shapeName, DbValue)
}

func (s *shapesService) GetShapes() (*[]domain.Shapes, error) {
	return domain.ShapeDao.GetShapes()
}