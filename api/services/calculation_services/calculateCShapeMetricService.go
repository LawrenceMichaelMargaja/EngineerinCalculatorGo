package calculation_services

import (
	"db/api/domain/amir_calc/calculation_domains"
	"errors"
)

type calculateCShapeMetricService struct {

}

var (
	CalculateCShapeMetricService *calculateCShapeMetricService
)

func init()  {
	CalculateCShapeMetricService = &calculateCShapeMetricService{}
}

//func (s *cShapeEnglishService) Get() (*[]domain.CShapeEnglish, error) {
//	return domain.CShapeEnglishDao.Get()
//}

func (s *calculateCShapeMetricService) Create(CalculateCShapeMetricName string) error {
	if CalculateCShapeMetricName == "" {
		return errors.New("english C shape name is invalid")
	}
	return calculation_domains.CalculateCShapeMetricDao .Create(CalculateCShapeMetricName)
}
