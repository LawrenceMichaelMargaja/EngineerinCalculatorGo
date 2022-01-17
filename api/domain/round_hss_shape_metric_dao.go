package domain

import (
	"db/api/utils"
	"github.com/jmoiron/sqlx"
)

type roundHssShapeMetricInterface interface {
	Get() (*[]RoundHSShapeMetric, error)
	Create(roundHssShapeName string) error
	SetClient()
}

type roundHssShapeMetricDao struct {
	client *sqlx.DB
}

var (
	RoundHssShapeMetricDao roundHssShapeMetricInterface
)

func init()  {
	RoundHssShapeMetricDao = &roundHssShapeMetricDao{}
	RoundHssShapeMetricDao.SetClient()
}

func (s *roundHssShapeMetricDao) SetClient()  {
	s.client = utils.GetMYSQLConnection()
}

func (s *roundHssShapeMetricDao) Get() (*[]RoundHSShapeMetric, error) {
	var roundHsShapeMetricData []RoundHSShapeMetric
	sql := `
			SELECT
				idround_hs_shape_metric AS id,
				round_hs_shape_metric_name AS round_hs_shape_name
			FROM
				round_hs_shape_metric
			`
	err := s.client.Select(&roundHsShapeMetricData, sql)
	return &roundHsShapeMetricData, err
}

func (s *roundHssShapeMetricDao) Create(roundHssShapeName string) error {
	sql := `INSERT INTO roundhss_shape_metric(roundhss_shape_metric_name) VALUES(?)`
	_, err := s.client.Exec(sql, roundHssShapeName)
	return err
}
