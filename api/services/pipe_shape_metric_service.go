package services

import (
	"db/api/domain"
	"errors"
)

type pipeShapeMetricService struct {

}

var (
	PipeShapeMetricService *pipeShapeMetricService
)

func init()  {
	PipeShapeMetricService = &pipeShapeMetricService{}
}

func (s *pipeShapeMetricService) Get() (*[]domain.PipeShapeMetric, error) {
	return domain.PipeShapeMetricDao.Get()
}

func (s *pipeShapeMetricService) Create(pipeShapeMetricName string) error {
	if pipeShapeMetricName == "" {
		return errors.New("metric pipe shape name is invalid")
	}
	return domain.PipeShapeMetricDao.Create(pipeShapeMetricName)
}