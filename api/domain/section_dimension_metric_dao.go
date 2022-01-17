package domain

import (
	"db/api/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type sectionDimensionMetricDaoInterface interface {
	Get() (*[]SectionDimensionMetric, error)
	Create(sectionName string, d string, b string, tw string, bf string, tf string, tb string, t string, r string) error
	SetClient()
}

type sectionDimensionMetricDao struct {
	client *sqlx.DB
}

var (
	SectionDimensionMetricDao sectionDimensionMetricDaoInterface
)

func init()  {
	SectionDimensionMetricDao = &sectionDimensionMetricDao{}
	SectionDimensionMetricDao.SetClient()
}

func (s *sectionDimensionMetricDao) SetClient() {
	s.client = utils.GetMYSQLConnection()
}

func (s *sectionDimensionMetricDao) Get() (*[]SectionDimensionMetric, error) {
	var sectionDimensionMetricData []SectionDimensionMetric
	sql := `
			SELECT
				idsection_dimension_metric AS id,
				section_dimension_metric_name AS section_dimension_metric_name,
				section_dimension_metric_d AS section_dimension_metric_d,
				section_dimension_metric_b AS section_dimension_metric_b,
				section_dimension_metric_tw AS section_dimension_metric_tw,
				section_dimension_metric_bf AS section_dimension_metric_bf,
				section_dimension_metric_tf AS section_dimension_metric_tf,
				section_dimension_metric_tb AS section_dimension_metric_tb,
				section_dimension_metric_t AS section_dimension_metric_t,
				section_dimension_metric_r AS section_dimension_metric_r
			FROM
				section_dimension_metric
			`
	err := s.client.Select(&sectionDimensionMetricData, sql)
	return &sectionDimensionMetricData, err
}

func (s *sectionDimensionMetricDao) Create(sectionName string, d string, b string, tw string, bf string, tf string, tb string, t string, r string) error {
	sql := `INSERT INTO section_dimension_metric (section_dimension_metric_name, section_dimension_metric_d, section_dimension_metric_b, section_dimension_metric_tw, section_dimension_metric_bf, section_dimension_metric_tf, section_dimension_metric_tb, section_dimension_metric_t, section_dimension_metric_r) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := s.client.Exec(sql, sectionName, d, b, tw, bf, tf, tb, t, r)
	return err
}