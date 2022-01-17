package domain

import (
	"db/api/utils"
	"github.com/jmoiron/sqlx"
)

type pipeShapeMetricDaoInterface interface {
	Get() (*[]PipeShapeMetric, error)
	Create(pipeShapeMetricName string) error
	SetClient()
}

type pipeShapeMetricDao struct {
	client *sqlx.DB
}

var (
	PipeShapeMetricDao pipeShapeMetricDaoInterface
)

func init()  {
	PipeShapeMetricDao = &pipeShapeMetricDao{}
	PipeShapeMetricDao.SetClient()
}

func (s *pipeShapeMetricDao) SetClient() {
	s.client = utils.GetMYSQLConnection()
}

func (s *pipeShapeMetricDao) Get() (*[]PipeShapeMetric, error) {
	var pipeShapeMetricData []PipeShapeMetric
	sql := `
			SELECT
				idpipe_shape_metric AS id,
				pipe_shape_metric_name AS pipe_shape_metric_name
			FROM
				pipe_shape_metric
			`
	err := s.client.Select(&pipeShapeMetricData, sql)
	return &pipeShapeMetricData, err
}

func (s *pipeShapeMetricDao) Create(pipeShapeMetricName string) error {
	sql := `INSERT INTO pipe_shape_metric (pipe_shape_metric_name) VALUES(?)`
	_, err := s.client.Exec(sql, pipeShapeMetricName)
	return err
}