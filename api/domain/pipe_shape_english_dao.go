package domain

import (
	"db/api/utils"
	"github.com/jmoiron/sqlx"
)

type pipeShapeEnglishDaoInterface interface {
	Get() (*[]PipeShapeEnglish, error)
	Create(pipeShapeEnglishName string) error
	SetClient()
}

type pipeShapeEnglishDao struct {
	client *sqlx.DB
}

var (
	PipeShapeEnglishDao pipeShapeEnglishDaoInterface
)

func init()  {
	PipeShapeEnglishDao = &pipeShapeEnglishDao{}
	PipeShapeEnglishDao.SetClient()
}

func (s *pipeShapeEnglishDao) SetClient() {
	s.client = utils.GetMYSQLConnection()
}

func (s *pipeShapeEnglishDao) Get() (*[]PipeShapeEnglish, error) {
	var pipeShapeEnglishData []PipeShapeEnglish
	sql := `
			SELECT
				idpipe_shape_english AS id,
				pipe_shape_english_name AS pipe_shape_english_name
			FROM
				pipe_shape_english
			`
	err := s.client.Select(&pipeShapeEnglishData, sql)
	return &pipeShapeEnglishData, err
}

func (s *pipeShapeEnglishDao) Create(pipeShapeEnglishName string) error {
	sql := `INSERT INTO pipe_shape_english (pipe_shape_english_name) VALUES(?)`
	_, err := s.client.Exec(sql, pipeShapeEnglishName)
	return err
}