package services

import (
	"db/api/domain"
	"errors"
)

type cShapeMetricService struct {

}

var (
	CShapeMetricService *cShapeMetricService
)

func init()  {
	CShapeMetricService = &cShapeMetricService{}
}

func (s *cShapeMetricService) Get() (*[]domain.CShapeMetric, error) {
	return domain.CShapeMetricDao.Get()
}

func (s *cShapeMetricService) Create(cShapeMetricName string) error {
	if cShapeMetricName == "" {
		return errors.New("metric c shape name is invalid")
	}
	return domain.CShapeMetricDao.Create(cShapeMetricName)
}

