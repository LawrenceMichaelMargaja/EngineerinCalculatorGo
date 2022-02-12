package services

import (
	"db/api/domain"
)

type designMembersMetricService struct {

}

var (
	DesignMembersMetricService *designMembersMetricService
)

func init()  {
	DesignMembersMetricService = &designMembersMetricService{}
}

func (s *designMembersMetricService) Get() (*[]domain.DesignMembersMetric, error) {
	return domain.DesignMembersMetricDao.Get()
}

//func (s *iShapeEnglishService) Create(iShapeEnglishName string) error {
//	if iShapeEnglishName == "" {
//		return errors.New("english i shape name is invalid")
//	}
//	return domain.IShapeEnglishDao.Create(iShapeEnglishName)
//}
