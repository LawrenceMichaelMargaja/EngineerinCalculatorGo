package domain

import (
	"db/api/utils"
	"github.com/jmoiron/sqlx"
)

type recHssShapeEnglishDaoInterface interface {
    Get() (*[]RecHsShapeEnglish, error)
	Create(recHssShapeEnglishName string) error
	SetClient()
}

type recHssShapeEnglishDao struct {
	client *sqlx.DB
}

var (
	RecHssShapeEnglishDao recHssShapeEnglishDaoInterface
)

func init()  {
	RecHssShapeEnglishDao = &recHssShapeEnglishDao{}
	RecHssShapeEnglishDao.SetClient()
}

func (s *recHssShapeEnglishDao) SetClient()  {
	s.client = utils.GetMYSQLConnection()
}

func (s *recHssShapeEnglishDao) Get() (*[]RecHsShapeEnglish, error) {
	var recHsShapeEnglishData []RecHsShapeEnglish
	sql := `
			SELECT
				idrec_hs_shape_english AS id,
				rec_hs_shape_english_name AS rec_hs_shape_english_name
			FROM
				rec_hs_shape_english
			`
	err := s.client.Select(&recHsShapeEnglishData, sql)
	return &recHsShapeEnglishData, err
}

func (s *recHssShapeEnglishDao) Create(recHssShapeEnglishName string) error {
	sql := `INSERT INTO rechss_shape_english(rechss_shape_english_name) VALUES(?)`
	_, err := s.client.Exec(sql, recHssShapeEnglishName)
	return err
}