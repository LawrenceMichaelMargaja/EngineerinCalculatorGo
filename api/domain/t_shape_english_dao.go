package domain

import (
	"db/api/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type tShapeEnglishInterface interface {
	Get() (*[]TShapeEnglish, error)
	Create(tShapeEnglishName string) error
	SetClient()
}

type tShapeEnglishDao struct {
	client *sqlx.DB
}

var (
	TShapeEnglishDao tShapeEnglishInterface
)

func init()  {
	TShapeEnglishDao = &tShapeEnglishDao{}
	TShapeEnglishDao.SetClient()
}

func (s *tShapeEnglishDao) SetClient() {
	s.client = utils.GetMYSQLConnection()
}

func (s *tShapeEnglishDao) Get() (*[]TShapeEnglish, error) {
	var tShapeEnglishData []TShapeEnglish
	sql := `SELECT idt_shape_english AS id, t_shape_english_name AS t_shape_english_name FROM t_shape_english`
	err := s.client.Select(&tShapeEnglishData, sql)
	return &tShapeEnglishData, err
}

func (s *tShapeEnglishDao) Create(tShapeEnglishName string) error {
	sql := `INSERT INTO t_shape_english (t_shape_english_name) VALUES (?)`
	_, err := s.client.Exec(sql, tShapeEnglishName)
	return err
}