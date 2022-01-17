package domain

import (
	"db/api/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type steelTypeMetricInterface interface {
	Get() (*[]SteelTypeMetric, error)
	Create(steelTypeMetricName string, steelTypeMetricE string, steelTypeMetricFy string, steelTypeMetricFu string) error
	SetClient()
}

type steelTypeMetricDao struct {
	client *sqlx.DB
}

var (
	SteelTypeMetricDao steelTypeMetricInterface
)

func init() {
	SteelTypeMetricDao = &steelTypeMetricDao{}
	SteelTypeMetricDao.SetClient()
}

func (s *steelTypeMetricDao) SetClient() {
	s.client = utils.GetMYSQLConnection()
}

func (s *steelTypeMetricDao) Get() (*[]SteelTypeMetric, error) {
	var steelTypeMetricData []SteelTypeMetric
	sql := `SELECT 
				idsteel_types_metric AS id, 
				steel_types_metric_name AS steel_type_metric_name, 
				steel_types_metric_E AS steel_type_metric_e, 
				steel_types_metric_Fy AS steel_type_metric_fy,
				steel_types_metric_Fu AS steel_type_metric_fu
			FROM
				steel_types_metric
			`
	err := s.client.Select(&steelTypeMetricData, sql)
	return &steelTypeMetricData, err
}

func (s *steelTypeMetricDao) Create(steelTypeMetricName string, steelTypeMetricE string, steelTypeMetricFy string, steelTypeMetricFu string) error {
	sql := `INSERT INTO steel_types_metric (steel_types_metric_name, steel_types_metric_E, steel_types_metric_Fy, steel_types_metric_Fu) VALUES(?, ?, ?, ?)`
	_, err := s.client.Exec(sql, steelTypeMetricName, steelTypeMetricE, steelTypeMetricFy, steelTypeMetricFu)
	return err
}
