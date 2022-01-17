package domain

import (
	"db/api/utils"
	"github.com/jmoiron/sqlx"
)

type cShapeMetricDaoInterface interface {
	Get() (*[]CShapeMetric, error)
	Create(cShapeMetricName string) error
	SetClient()
}

type cShapeMetricDao struct {
	client *sqlx.DB
}

var (
	CShapeMetricDao cShapeMetricDaoInterface
)

func init()  {
	CShapeMetricDao = &cShapeMetricDao{}
	CShapeMetricDao.SetClient()
}

func (s *cShapeMetricDao) SetClient()  {
	s.client = utils.GetMYSQLConnection()
}

func (s *cShapeMetricDao) Get() (*[]CShapeMetric, error) {
	var cShapeMetricData []CShapeMetric
	sql := `
			SELECT
				idc_shape_metric AS id,
				c_shape_metric_name AS c_shape_metric_name
			FROM
				c_shape_metric
			`
	err := s.client.Select(&cShapeMetricData, sql)
	return &cShapeMetricData, err
}

func (s *cShapeMetricDao) Create(cShapeMetricName string) error {
	sql := `INSERT INTO c_shape_metric (c_shape_metric_name) VALUES(?)`
	_, err := s.client.Exec(sql, cShapeMetricName)
	return err
}
