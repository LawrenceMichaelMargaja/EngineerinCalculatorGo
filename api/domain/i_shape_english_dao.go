package domain

import (
	"db/api/utils"
	"github.com/jmoiron/sqlx"
)

type iShapeEnglishDaoInterface interface {
	Get() (*[]IShapeEnglish, error)
	Create(iShapeEnglishName string) error
	SetClient()
}

type iShapeEnglishDao struct {
	client *sqlx.DB
}

var (
	IShapeEnglishDao iShapeEnglishDaoInterface
)

func init()  {
	IShapeEnglishDao = &iShapeEnglishDao{}
	IShapeEnglishDao.SetClient()
}

func (s *iShapeEnglishDao) SetClient()  {
	s.client = utils.GetMYSQLConnection()
}

func (s *iShapeEnglishDao) Get() (*[]IShapeEnglish, error) {
	var iShapeEnglishData []IShapeEnglish
	sql := `
			SELECT
				idi_shape_english AS id,
				i_shape_english_name AS i_shape_english_name
			FROM
				i_shape_english
			`
	err := s.client.Select(&iShapeEnglishData, sql)
	return &iShapeEnglishData, err
}

func (s *iShapeEnglishDao) Create(iShapeEnglishName string) error {
	sql := `INSERT INTO i_shape_english (i_shape_english_name) VALUES(?)`
	_, err := s.client.Exec(sql, iShapeEnglishName)
	return err
}
