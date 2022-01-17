package services

import (
	"db/api/domain"
	"errors"
)

type iShapeMetricService struct {

}

var (
	IShapeMetricService *iShapeMetricService
)

func init()  {
	IShapeMetricService = &iShapeMetricService{}
}

func (s *iShapeMetricService) Get() (*[]domain.IShapeMetric, error) {
	return domain.IShapeMetricDao.Get()
}

func (s *iShapeMetricService) Create(iShapeMetricName string) error {
	if iShapeMetricName == "" {
		return errors.New("metric i shape is invalid")
	}
	return domain.IShapeMetricDao.Create(iShapeMetricName)
}
