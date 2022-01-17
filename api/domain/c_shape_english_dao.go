package domain

import (
	"db/api/utils"
	"github.com/jmoiron/sqlx"
)

type cShapeEnglishDaoInterface interface {
	Get() (*[]CShapeEnglish, error)
	Create(cShapeEnglishName string) error
	SetClient()
}

type cShapeEnglishDao struct {
	client *sqlx.DB
}

var (
	CShapeEnglishDao cShapeEnglishDaoInterface
)

func init()  {
	CShapeEnglishDao = &cShapeEnglishDao{}
	CShapeEnglishDao.SetClient()
}

func (s *cShapeEnglishDao) SetClient()  {
	s.client = utils.GetMYSQLConnection()
}

func (s *cShapeEnglishDao) Get() (*[]CShapeEnglish, error) {
	var cShapeEnglishData []CShapeEnglish
	sql := `
			SELECT
				idc_shape_english AS id,
				c_shape_english_name AS c_shape_english_name
			FROM
				c_shape_english
			`
	err := s.client.Select(&cShapeEnglishData, sql)
	return &cShapeEnglishData, err
}

func (s *cShapeEnglishDao) Create(cShapeEnglishName string) error {
	sql := `INSERT INTO c_shape_english (c_shape_english_name) VALUES(?)`
	_, err := s.client.Exec(sql, cShapeEnglishName)
	return err
}
