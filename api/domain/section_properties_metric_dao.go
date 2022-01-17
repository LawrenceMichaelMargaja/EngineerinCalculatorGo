package domain

import (
	"db/api/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type sectionPropertiesMetricDaoInterface interface {
	Get() (*[]SectionPropertiesMetric, error)
	Create(sectionName string, a string, j string, ixp string, iyp string, iw string, sxp string, syp string) error
	SetClient()
}

type sectionPropertiesMetricDao struct {
	client *sqlx.DB
}

var (
	SectionPropertiesMetricDao sectionPropertiesMetricDaoInterface
)

func init()  {
	SectionPropertiesMetricDao = &sectionPropertiesMetricDao{}
	SectionPropertiesMetricDao.SetClient()
}

func (s *sectionPropertiesMetricDao) SetClient() {
	s.client = utils.GetMYSQLConnection()
}

func (s *sectionPropertiesMetricDao) Get() (*[]SectionPropertiesMetric, error) {
	var sectionPropertiesMetricData []SectionPropertiesMetric
	sql := `
			SELECT 
				idsection_properties_metric AS id,
				section_properties_metric_name AS section_properties_metric_name,
				section_properties_metric_A AS section_properties_metric_a,
				section_properties_metric_J AS section_properties_metric_j,
				section_properties_metric_Ixp AS section_properties_metric_ixp,
				section_properties_metric_Iyp AS section_properties_metric_iyp,
				section_properties_metric_Iw AS section_properties_metric_iw,
				section_properties_metric_Sxp AS section_properties_metric_sxp,
				section_properties_metric_Syp AS section_properties_metric_syp
			FROM
				section_properties_metric
			`
	err := s.client.Select(&sectionPropertiesMetricData, sql)
	return &sectionPropertiesMetricData, err
}

func (s *sectionPropertiesMetricDao) Create(sectionName string, a string, j string, ixp string, iyp string, iw string, sxp string, syp string) error {
	sql := `INSERT INTO section_properties_metric (sectionName, A, J, Ixp, Iyp, Iw, Sxp, Syp) VALUES(?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := s.client.Exec(sql, sectionName, a, j, ixp, iyp, iw, sxp, syp)
	return err
}