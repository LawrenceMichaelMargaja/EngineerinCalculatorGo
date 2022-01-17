package domain

import (
	"db/api/utils"
	"github.com/jmoiron/sqlx"
)

type twoLShapeEnglishDaoInterface interface {
	Get() (*[]TwoLShapeEnglish, error)
	Create(twoLShapeEnglishName string) error
	SetClient()
}

type twoLShapeEnglishDao struct {
	client *sqlx.DB
}

var (
	TwoLShapeEnglishDao twoLShapeEnglishDaoInterface
)

func init()  {
	TwoLShapeEnglishDao = &twoLShapeEnglishDao{}
	TwoLShapeEnglishDao.SetClient()
}

func (s *twoLShapeEnglishDao) SetClient()  {
	s.client = utils.GetMYSQLConnection()
}

func (s *twoLShapeEnglishDao) Get() (*[]TwoLShapeEnglish, error) {
	var twoLShapeEnglishData []TwoLShapeEnglish
	sql := `
			SELECT
				id2l_shape_english AS id,
				2l_shape_english_name AS two_l_shape_english_name
			FROM
				2l_shape_english
			`
	err := s.client.Select(&twoLShapeEnglishData, sql)
	return &twoLShapeEnglishData, err
}

func (s *twoLShapeEnglishDao) Create(twoLShapeEnglishName string) error {
	sql := `INSERT INTO 2l_shape_english (2l_shape_english_name) VALUES(?)`
	_, err := s.client.Exec(sql, twoLShapeEnglishName)
	return err
}