package domain

import (
	"db/api/utils"
	"github.com/jmoiron/sqlx"
)

type designMembersEnglishDaoInterface interface {
	Get() (*[]DesignMembersEnglish, error)
	//Create(iShapeEnglishName string) error
	SetClient()
}

type designMemberEnglishDao struct {
	client *sqlx.DB
}

var (
	DesignMembersEnglishDao designMembersEnglishDaoInterface
)

func init() {
	DesignMembersEnglishDao = &designMemberEnglishDao{}
	DesignMembersEnglishDao.SetClient()
}

func (s *designMemberEnglishDao) SetClient() {
	s.client = utils.GetMYSQLConnection()
}

func (s *designMemberEnglishDao) Get() (*[]DesignMembersEnglish, error) {
	var designMembersEnglishData []DesignMembersEnglish
	sql := `
			SELECT
				iddesign_member_english AS id,
				design_member_english_name AS design_members_english_name,
				design_member_english_total_depth AS design_members_english_total_depth,
			    design_member_english_weight AS design_members_english_weight   
			FROM
				design_member_english
			`
	err := s.client.Select(&designMembersEnglishData, sql)
	return &designMembersEnglishData, err
}

//func (s *designMemberEnglishDao) Create(iShapeEnglishName string) error {
//	sql := `INSERT INTO i_shape_english (i_shape_english_name) VALUES(?)`
//	_, err := s.client.Exec(sql, iShapeEnglishName)
//	return err
//}
