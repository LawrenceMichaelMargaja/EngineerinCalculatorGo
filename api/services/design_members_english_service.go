package services

import (
	"db/api/domain"
)

type designMembersEnglishService struct {

}

var (
	DesignMembersEnglishService *designMembersEnglishService
)

func init()  {
	DesignMembersEnglishService = &designMembersEnglishService{}
}

func (s *designMembersEnglishService) Get() (*[]domain.DesignMembersEnglish, error) {
	return domain.DesignMembersEnglishDao.Get()
}

//func (s *iShapeEnglishService) Create(iShapeEnglishName string) error {
//	if iShapeEnglishName == "" {
//		return errors.New("english i shape name is invalid")
//	}
//	return domain.IShapeEnglishDao.Create(iShapeEnglishName)
//}
