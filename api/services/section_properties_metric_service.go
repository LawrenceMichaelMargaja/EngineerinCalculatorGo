package services

import (
	"db/api/domain"
	"errors"
)

type sectionPropertiesMetricService struct {
}

var (
	SectionPropertiesMetricService *sectionPropertiesMetricService
)

func init() {
	SectionPropertiesMetricService = &sectionPropertiesMetricService{}
}

func (s *sectionPropertiesMetricService) Get() (*[]domain.SectionPropertiesMetric, error) {
	return domain.SectionPropertiesMetricDao.Get()
}

func (s *sectionPropertiesMetricService) Create(sectionName string, a string, j string, ixp string, iyp string, iw string, sxp string, syp string) error {
	if sectionName == "" {
		return errors.New("section name is invalid")
	}
	if a == "" {
		return errors.New("A is invalid")
	}
	if j == "" {
		return errors.New("j is invalid")
	}
	if ixp == "" {
		return errors.New("ixp is invalid")
	}
	if iyp == "" {
		return errors.New("iyp is invalid")
	}
	if iw == "" {
		return errors.New("iw is invalid")
	}
	if sxp == "" {
		return errors.New("sxp is invalid")
	}
	if syp == "" {
		return errors.New("syp is invalid")
	}
	return domain.SectionPropertiesMetricDao.Create(sectionName, a, j, ixp, iyp, iw, sxp, syp)
}
