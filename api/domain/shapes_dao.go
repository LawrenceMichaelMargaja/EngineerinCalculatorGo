package domain

import (
	"db/api/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type shapeDaoInterface interface {
	GetShapes() (*[]Shapes, error)
	Create(shapeName string, DbValue string) error
	SetClient()
}

type shapeDao struct {
	client *sqlx.DB
}

var (
	ShapeDao shapeDaoInterface
)

func init() {
	ShapeDao = &shapeDao{}
	ShapeDao.SetClient()
}

func (s *shapeDao) SetClient() {
	s.client = utils.GetMYSQLConnection()
}

func (s *shapeDao) Create(shapeName string, DbValue string) error {
	sql := `INSERT INTO shapes (shapeName, db_value) VALUES (?, ?)`

	_, err := s.client.Exec(sql, shapeName, DbValue)
	return err
}

func (s *shapeDao) GetShapes() (*[]Shapes, error) {
	var shapes []Shapes
	sql := `SELECT 
				idshapes AS id,
				shapeName AS shape_name,
       			db_value AS db_value
			FROM shapes`
	err := s.client.Select(&shapes, sql)
	return &shapes, err
}
