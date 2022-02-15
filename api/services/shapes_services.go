package services

import (
	"db/api/domain"
)

type shapesService struct {

}

var (
	ShapeService *shapesService
)

func init()  {
	ShapeService = &shapesService{}
}

//func (s *shapesService) Create(shapeName string) error {
//	if shapeName == "" {
//		return errors.New("shape name is invalid")
//	}
//	return domain.ShapeDao.Create(shapeName)
//}

func (s *shapesService) GetShapes() (*[]domain.Shapes, error) {
	return domain.ShapeDao.GetShapes()
}