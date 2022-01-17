package services

import (
	"db/api/domain"
	"errors"
)

type twoLShapeMetricService struct {

}

var (
	TwoLShapeMetricService *twoLShapeMetricService
)

func init()  {
	TwoLShapeMetricService = &twoLShapeMetricService{}
}

func (s *twoLShapeMetricService) Get() (*[]domain.TwoLShapeMetric, error) {
	return domain.TwoLMetricShapeDao.Get()
}

func (s *twoLShapeMetricService) Create(twoLShapeMetricName string) error {
	if twoLShapeMetricName == "" {
		return errors.New("metric 2 L shape is invalid")
	}
	return domain.TwoLMetricShapeDao.Create(twoLShapeMetricName)
}
