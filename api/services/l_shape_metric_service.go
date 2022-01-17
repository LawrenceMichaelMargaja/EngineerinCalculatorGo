package services

import (
	"db/api/domain"
	"errors"
)

type lShapeMetricService struct {

}

var (
	LShapeMetricService *lShapeMetricService
)

func init()  {
	LShapeMetricService = &lShapeMetricService{}
}

func (s *lShapeMetricService) Get() (*[]domain.LShapeMetric, error) {
	return domain.LShapeMetricDao.Get()
}

func (s *lShapeMetricService) Create(lShapeMetricName string) error {
	if lShapeMetricName == "" {
		return errors.New("metric l shape name is invalid")
	}
	return domain.LShapeMetricDao.Create(lShapeMetricName)
}
