package services

import (
	"db/api/domain"
	"errors"
)

type sectionPropertiesEnglishService struct {

}

var (
	SectionPropertiesEnglishService *sectionPropertiesEnglishService
)

func init()  {
	SectionPropertiesEnglishService = &sectionPropertiesEnglishService{}
}

func (s *sectionPropertiesEnglishService) Get() (*[]domain.SectionPropertiesEnglish, error) {
	return domain.SectionPropertiesEnglishDao.Get()
}

func (s *sectionPropertiesEnglishService) Create(sectionName string, a string, j string, ixp string, iyp string, iw string, sxp string, syp string) error {
	if sectionName == "" {
		return errors.New("section name is invalid")
	}
	if a == "" {
		return errors.New("a is invalid")
	}
	if j == "" {
		return errors.New("j is invalid")
	}
	if ixp == "" {
		return errors.New("Ixp is invalid")
	}
	if iyp == "" {
		return errors.New("Iyp is invalid")
	}
	if iw == "" {
		return errors.New("Iw is invalid")
	}
	if sxp == "" {
		return errors.New("Sxp is invalid")
	}
	if syp == "" {
		return errors.New("Syp is invalid")
	}
	return domain.SectionPropertiesEnglishDao.Create(sectionName, a, j, ixp, iyp, iw, sxp,syp)
}