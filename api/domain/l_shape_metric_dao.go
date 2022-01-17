package domain

import (
	"db/api/utils"
	"github.com/jmoiron/sqlx"
)

type lShapeMetricDaoInterface interface {
	Get() (*[]LShapeMetric, error)
	Create(lShapeMetricName string) error
	SetClient()
}

type lShapeMetricDao struct {
	client *sqlx.DB
}

var (
	LShapeMetricDao lShapeMetricDaoInterface
)

func init()  {
	LShapeMetricDao = &lShapeMetricDao{}
	LShapeMetricDao.SetClient()
}

func (s *lShapeMetricDao) SetClient() {
	s.client = utils.GetMYSQLConnection()
}

func (s *lShapeMetricDao) Get() (*[]LShapeMetric, error) {
	var lShapeMetricData []LShapeMetric
	sql := `
			SELECT
				idl_shape_metric AS id,
				l_shape_metric_name AS l_shape_metric_name
			FROM
				l_shape_metric
			`
	err := s.client.Select(&lShapeMetricData, sql)
	return &lShapeMetricData, err
}

func (s *lShapeMetricDao) Create(lShapeMetricName string) error {
	sql := `INSERT INTO l_shape_metric (L_shape_metric_name) VALUES(?)`
	_,err := s.client.Exec(sql, lShapeMetricName)
	return err
}
