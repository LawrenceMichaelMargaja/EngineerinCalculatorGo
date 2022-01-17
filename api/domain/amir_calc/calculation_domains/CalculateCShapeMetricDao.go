package calculation_domains

import (
	"db/api/utils"
	"github.com/jmoiron/sqlx"
)

type cShapeMetricDaoInterface interface {
	Create(cShapeEnglishName string) error
	SetClient()
}

type calculateCShapeMetricDao struct {
	client *sqlx.DB
}

var (
	CalculateCShapeMetricDao cShapeMetricDaoInterface
)

func init()  {
	CalculateCShapeMetricDao = &calculateCShapeMetricDao{}
	CalculateCShapeMetricDao.SetClient()
}

func (s *calculateCShapeMetricDao) SetClient()  {
	s.client = utils.GetMYSQLConnection()
}

func (s *calculateCShapeMetricDao) Get() (*[]CalculateCShapeMetric, error) {
	var cShapeEnglishData []CalculateCShapeMetric
	sql := `
			SELECT
				idc_shape_metric AS id,
				calculate_c_shape_metric_name AS calculate_c_shape_metric_name
			FROM
				c_shape_english
			`
	err := s.client.Select(&cShapeEnglishData, sql)
	return &cShapeEnglishData, err
}

func (s *calculateCShapeMetricDao) Create(calculateCShapeMetricName string) error {
	sql := `INSERT INTO calculate_c_shape_metric (calculate_c_shape_metric_name) VALUES(?)`
	_, err := s.client.Exec(sql, calculateCShapeMetricName)
	return err
}