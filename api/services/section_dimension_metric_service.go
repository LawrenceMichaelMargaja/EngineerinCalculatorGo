package services

import (
	"db/api/domain"
	"errors"
)

type sectionDimensionMetricService struct {

}

var (
	SectionDimensionMetricService *sectionDimensionMetricService
)

func init()  {
	SectionDimensionMetricService = &sectionDimensionMetricService{}
}

func (s *sectionDimensionMetricService) Get() (*[]domain.SectionDimensionMetric, error) {
	return domain.SectionDimensionMetricDao.Get()
}

func (s *sectionDimensionMetricService) Create(sectionDimensionName string, d string, b string, tw string, bf string, tf string, tb string, t string, r string) error {
	if sectionDimensionName == "" {
		return errors.New("section dimension name is invalid")
	}
	if d == "" {
		return errors.New("d is invalid")
	}
	if b == "" {
		return errors.New("b is invalid")
	}
	if tw == "" {
		return errors.New("tw is invalid")
	}
	if bf == "" {
		return errors.New("bf is invalid")
	}
	if tf == "" {
		return errors.New("tf is invalid")
	}
	if tb == "" {
		return errors.New("tb is invalid")
	}
	if t == "" {
		return errors.New("t is invalid")
	}
	if r == "" {
		return errors.New("r is invalid")
	}
	return domain.SectionDimensionMetricDao.Create(sectionDimensionName, d, b, tw, bf, tf, tb, t, r)
}