package services

import (
	"db/api/domain"
	"errors"
)

type steelTypeMetricService struct {

}

var (
	SteelTypeMetricService *steelTypeMetricService
)

func init()  {
	SteelTypeMetricService = &steelTypeMetricService{}
}

func (s *steelTypeMetricService) Get() (*[]domain.SteelTypeMetric, error) {
	return domain.SteelTypeMetricDao.Get()
}

func (s *steelTypeMetricService) Create(steelTypeMetricName string, steelTypeMetricE string, steelTypeMetricFy string, steelTypeMetricFu string) error {
	if steelTypeMetricName == "" {
		return errors.New("steel type name is invalid")
	}
	if steelTypeMetricE == "" {
		return errors.New("steel type E is invalid")
	}
	if steelTypeMetricFy == "" {
		return errors.New("steel type Fy is invalid")
	}
	if steelTypeMetricFu == "" {
		return errors.New("steel type Fu is invalid")
	}
	return domain.SteelTypeMetricDao.Create(steelTypeMetricName, steelTypeMetricE, steelTypeMetricFy, steelTypeMetricFu)
}