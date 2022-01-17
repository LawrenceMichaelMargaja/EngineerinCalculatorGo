package domain

import (
	"db/api/utils"
	"github.com/jmoiron/sqlx"
)

type sectionDimensionEnglishDaoInterface interface {
	Get() (*[]SectionDimensionEnglish, error)
	Create(sectionDimensionName string, d string, b string, tw string, bf string, tf string, tb string, t string, r string) error
	SetClient()
}

type sectionDimensionEnglishDao struct {
	client *sqlx.DB
}

var (
	SectionDimensionEnglishDao sectionDimensionEnglishDaoInterface
)

func init()  {
	SectionDimensionEnglishDao = &sectionDimensionEnglishDao{}
	SectionDimensionEnglishDao.SetClient()
}

func (s *sectionDimensionEnglishDao) SetClient()  {
	s.client = utils.GetMYSQLConnection()
}

func (s *sectionDimensionEnglishDao) Get() (*[]SectionDimensionEnglish, error) {
	var sectionDimensionEnglishData []SectionDimensionEnglish
	sql := `
			SELECT
				idsection_dimension_english AS id,
				section_dimension_english_name AS section_dimension_english_name,
				section_dimension_english_d AS section_dimension_english_d,
				section_dimension_english_b AS section_dimension_english_b,
				section_dimension_english_tw AS section_dimension_english_tw,
				section_dimension_english_bf AS section_dimension_english_bf,
				section_dimension_english_tf AS section_dimension_english_tf,
				section_dimension_english_tb AS section_dimension_english_tb,
				section_dimension_english_t AS section_dimension_english_t,
				section_dimension_english_r AS section_dimension_english_r
			FROM
				section_dimension_english
			`
	err := s.client.Select(&sectionDimensionEnglishData, sql)
	return &sectionDimensionEnglishData, err
}

func (s *sectionDimensionEnglishDao) Create(sectionDimensionName string, d string, b string, tw string, bf string, tf string, tb string, t string, r string) error {
	sql := `INSERT INTO section_dimension_english (section_dimension_english_name, section_dimension_english_d, section_dimension_english_b, section_dimension_english_tw, section_dimension_english_bf, section_dimension_english_tf, section_dimension_english_tb, section_dimension_english_t, section_dimension_english_r) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := s.client.Exec(sql, sectionDimensionName, d, b, tw, bf, tf, tb, t, r)
	return err
}