package domain

import (
	"db/api/utils"
	"github.com/jmoiron/sqlx"
)

type lShapeEnglishDaoInterface interface {
	Get() (*[]LShapeEnglish, error)
	Create(lShapeEnglishName string) error
	SetClient()
}

type lShapeEnglishDao struct {
	client *sqlx.DB
}

var (
	LShapeEnglishDao lShapeEnglishDaoInterface
)

func init()  {
	LShapeEnglishDao = &lShapeEnglishDao{}
	LShapeEnglishDao.SetClient()
}

func (s *lShapeEnglishDao) SetClient()  {
	s.client = utils.GetMYSQLConnection()
}

func (s *lShapeEnglishDao) Get() (*[]LShapeEnglish, error) {
	var lShapeEnglishData []LShapeEnglish
	sql := `
			SELECT
				idl_shape_english AS id,
				l_shape_english_name AS l_shape_english_name
			FROM
				l_shape_english
			`
	err := s.client.Select(&lShapeEnglishData, sql)
	return &lShapeEnglishData, err
}

func (s *lShapeEnglishDao) Create(lShapeEnglishName string) error {
	sql := `INSERT INTO l_shape_english (l_shape_english_name) VALUES(?)`
	_, err := s.client.Exec(sql, lShapeEnglishName)
	return err
}