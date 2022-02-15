package domain

import (
	"db/api/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type shapeDaoInterface interface {
	GetShapes() (*[]Shapes, error)
	//Create(shapeName string) error
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

//func (s *shapeDao) Create(shapeName string) error {
//	sql := `INSERT INTO shapes (shape_name, db_value) VALUES (?, ?)`
//
//	_, err := s.client.Exec(sql, shapeName)
//	return err
//}

func (s *shapeDao) GetShapes() (*[]Shapes, error) {
	var shapes []Shapes
	sql := `SELECT 
				idshapes AS id,
				shapes_name AS shape_name
			FROM shapes`
	err := s.client.Select(&shapes, sql)
	return &shapes, err
}
