package services

import (
	"db/api/domain"
	"errors"
)

type roundHssShapeMetricService struct {
}

var (
	RoundHssShapeMetricService *roundHssShapeMetricService
)

func init() {
	RoundHssShapeMetricService = &roundHssShapeMetricService{}
}

func (s *roundHssShapeMetricService) Get() (*[]domain.RoundHSShapeMetric, error) {
	return domain.RoundHssShapeMetricDao.Get()
}

func (s *roundHssShapeMetricService) Create(roundHssShapeName string) error {
	if roundHssShapeName == "" {
		return errors.New("the round hss name is invalid")
	}
	return domain.RoundHssShapeMetricDao.Create(roundHssShapeName)
}