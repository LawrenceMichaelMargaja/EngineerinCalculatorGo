package domain

import (
	"db/api/utils"
	"github.com/jmoiron/sqlx"
)

type twoLMetricShapeDaoInterface interface {
	Get() (*[]TwoLShapeMetric, error)
	Create(twoLShapeMetricName string) error
	SetClient()
}

type twoLMetricShapeDao struct {
	client *sqlx.DB
}

var (
	TwoLMetricShapeDao twoLMetricShapeDaoInterface
)

func init()  {
	TwoLMetricShapeDao = &twoLMetricShapeDao{}
	TwoLMetricShapeDao.SetClient()
}

func (s *twoLMetricShapeDao) SetClient()  {
	s.client = utils.GetMYSQLConnection()
}

func (s *twoLMetricShapeDao) Get() (*[]TwoLShapeMetric, error) {
	var twoLShapeMetricData []TwoLShapeMetric
	sql := `
			SELECT
				id2l_shape_metric AS id,
				2l_shape_metric_name AS two_l_shape_metric_name
			FROM
				2l_shape_metric
			`
	err := s.client.Select(&twoLShapeMetricData, sql)
	return &twoLShapeMetricData, err
}

func (s *twoLMetricShapeDao) Create(twoLShapeMetricName string) error {
	sql := `INSERT INTO 2l_shape_metric (2l_shape_metric_name) VALUES(?)`
	_, err := s.client.Exec(sql, twoLShapeMetricName)
	return err
}
