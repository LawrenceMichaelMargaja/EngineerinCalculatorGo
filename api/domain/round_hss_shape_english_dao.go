package domain

import (
	"db/api/utils"
	"github.com/jmoiron/sqlx"
)

type roundHssShapeEnglishDaoInterface interface {
	Get() (*[]RoundHsShapeEnglish, error)
	Create(roundHssShapeEnglishName string) error
	SetClient()
}

type roundHssShapeEnglishDao struct {
	client *sqlx.DB
}

var (
	RoundHssShapeEnglishDao roundHssShapeEnglishDaoInterface
)

func init() {
	RoundHssShapeEnglishDao = &roundHssShapeEnglishDao{}
	RoundHssShapeEnglishDao.SetClient()
}

func (s *roundHssShapeEnglishDao) SetClient() {
	s.client = utils.GetMYSQLConnection()
}

func (s *roundHssShapeEnglishDao) Get() (*[]RoundHsShapeEnglish, error) {
	var roundHsShapeEnglishData []RoundHsShapeEnglish
	sql := `
			SELECT
				idround_hs_shape_english AS id,
				round_hs_shape_english_name AS round_hs_shape_english_name
			FROM
				round_hs_shape_english
			`
	err := s.client.Select(&roundHsShapeEnglishData, sql)
	return &roundHsShapeEnglishData, err
}

func (s *roundHssShapeEnglishDao) Create(roundHssShapeEnglishName string) error {
	sql := `INSERT INTO roundhss_shape_english(roundhss_shape_english_name) VALUES(?)`
	_, err := s.client.Exec(sql, roundHssShapeEnglishName)
	return err
}