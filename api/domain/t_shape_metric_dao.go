package domain

import (
	"db/api/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type tShapeMetricInterface interface {
	Get() (*[]TShapeMetric, error)
	Create(tShapeMetricName string) error
	SetClient()
}

type tShapeMetricDao struct {
	client *sqlx.DB
}

var (
	TShapeMetricDao tShapeMetricInterface
)

func init() {
	TShapeMetricDao = &tShapeMetricDao{}
	TShapeMetricDao.SetClient()
}

func (s *tShapeMetricDao) SetClient()  {
	s.client = utils.GetMYSQLConnection()
}

func (s *tShapeMetricDao) Get() (*[]TShapeMetric, error) {
	var tShapeMetricData []TShapeMetric
	sql := `SELECT
				idt_shape_metric AS id,
				t_shape_metric_name AS t_shape_metric_name
			FROM t_shape_metric`
	err := s.client.Select(&tShapeMetricData, sql)
	return &tShapeMetricData, err
}

func (s *tShapeMetricDao) Create(tShapeMetricName string) error {
	sql := `INSERT INTO t_shape_metric (t_shape_metric_name) VALUES(?)`

	_, err := s.client.Exec(sql, tShapeMetricName)
	return err
}