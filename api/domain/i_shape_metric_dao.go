package domain

import (
	"db/api/utils"
	"github.com/jmoiron/sqlx"
)

type iShapeMetricDaoInterface interface {
	Get() (*[]IShapeMetric, error)
	Create(iShapeMetricName string) error
	SetClient()
}

type iShapeMetricDao struct {
	client *sqlx.DB
}

var (
	IShapeMetricDao iShapeMetricDaoInterface
)

func init()  {
	IShapeMetricDao = &iShapeMetricDao{}
	IShapeMetricDao.SetClient()
}

func (s *iShapeMetricDao) SetClient()  {
	s.client = utils.GetMYSQLConnection()
}

func (s *iShapeMetricDao) Get() (*[]IShapeMetric, error) {
	var iShapeMetricdata []IShapeMetric
	sql := `
			SELECT
				idi_shape_metric AS id,
				i_shape_metric_name AS i_shape_metric_name
			FROM
				i_shape_metric
			`
	err := s.client.Select(&iShapeMetricdata, sql)
	return &iShapeMetricdata, err
}

func (s *iShapeMetricDao) Create(iShapeMetricName string) error {
	sql := `INSERT INTO i_shape_metric (i_shape_metric_name) VALUES(?)`
	_, err := s.client.Exec(sql, iShapeMetricName)
	return err
}