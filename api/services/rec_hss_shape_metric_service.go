package services

import (
	"db/api/domain"
	"errors"
)

type recHssShapeMetricService struct {

}

var (
	RecHssShapeMetricService *recHssShapeMetricService
)

func init()  {
	RecHssShapeMetricService = &recHssShapeMetricService{}
}

func (s *recHssShapeMetricService) Get() (*[]domain.RecHsShapeMetric, error) {
	return domain.RecHssShapeMetricDao.Get()
}

func (s *recHssShapeMetricService) Create(recHssShapeMetricName string) error {
	if recHssShapeMetricName == "" {
		return errors.New("metric rec hss name is invalid")
	}
	return domain.RecHssShapeMetricDao.Create(recHssShapeMetricName)
}