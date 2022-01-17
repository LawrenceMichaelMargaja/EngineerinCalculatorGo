package services

import (
	"db/api/domain"
	"errors"
)

type sectionDimensionEnglishService struct {

}

var (
	SectionDimensionEnglishService *sectionDimensionEnglishService
)

func init()  {
	SectionDimensionEnglishService = &sectionDimensionEnglishService{}
}

func (s *sectionDimensionEnglishService) Get() (*[]domain.SectionDimensionEnglish, error) {
	return domain.SectionDimensionEnglishDao.Get()
}

func (s *sectionDimensionEnglishService) Create(sectionDimensionName string, d string, b string, tw string, bf string, tf string, tb string, t string, r string) error {
	if sectionDimensionName == ""{
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
	return domain.SectionDimensionEnglishDao.Create(sectionDimensionName, d, b, tw, bf, tf, tb, t, r)
}