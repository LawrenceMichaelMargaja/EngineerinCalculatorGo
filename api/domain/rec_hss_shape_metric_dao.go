package domain

import (
	"db/api/utils"
	"github.com/jmoiron/sqlx"
)

type recHssShapeMetricDaoInterface interface {
	Get() (*[]RecHsShapeMetric, error)
	Create(recHssShapeMetricName string) error
	SetClient()
}

type recHssShapeMetricDao struct {
	client *sqlx.DB
}

var (
	RecHssShapeMetricDao recHssShapeMetricDaoInterface
)

func init()  {
	RecHssShapeMetricDao = &recHssShapeMetricDao{}
	RecHssShapeMetricDao.SetClient()
}

func (s *recHssShapeMetricDao) SetClient() {
	s.client = utils.GetMYSQLConnection()
}

func (s *recHssShapeMetricDao) Get() (*[]RecHsShapeMetric, error) {
	var recHsShapeMetricData []RecHsShapeMetric
	sql := `
			SELECT
				idrec_hs_shape_metric AS id,
				rec_hs_shape_metric_name AS rec_hs_shape_metric_name
			FROM
				rec_hs_shape_metric
			`
	err := s.client.Select(&recHsShapeMetricData, sql)
	return &recHsShapeMetricData, err
}

func (s *recHssShapeMetricDao) Create(recHssShapeMetricName string) error {
	sql := `INSERT INTO rechss_shape_metric(recHSS_shape_metric_name) VALUES(?)`
	_, err := s.client.Exec(sql, recHssShapeMetricName)
	return err
}


