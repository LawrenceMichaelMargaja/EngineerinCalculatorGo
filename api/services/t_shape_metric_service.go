package services

import (
	"db/api/domain"
	"errors"
)

type tShapeMetricService struct {

}

var(
	TShapeMetricService *tShapeMetricService
)

func init()  {
	TShapeMetricService = &tShapeMetricService{}
}

func (s *tShapeMetricService) Get() (*[]domain.TShapeMetric, error) {
	return domain.TShapeMetricDao.Get()
}

func (s *tShapeMetricService) Create(tShapeMetricName string) error {
	if tShapeMetricName == "" {
		return errors.New("t shape name is invalid")
	}
	return domain.TShapeMetricDao.Create(tShapeMetricName)
}
