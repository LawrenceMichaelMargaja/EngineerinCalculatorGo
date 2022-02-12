package domain

import (
	"db/api/utils"
	"github.com/jmoiron/sqlx"
)

type designMembersMetricDaoInterface interface {
	Get() (*[]DesignMembersMetric, error)
	//Create(iShapeEnglishName string) error
	SetClient()
}

type designMemberMetricDao struct {
	client *sqlx.DB
}

var (
	DesignMembersMetricDao designMembersMetricDaoInterface
)

func init() {
	DesignMembersMetricDao = &designMemberMetricDao{}
	DesignMembersMetricDao.SetClient()
}

func (s *designMemberMetricDao) SetClient() {
	s.client = utils.GetMYSQLConnection()
}

func (s *designMemberMetricDao) Get() (*[]DesignMembersMetric, error) {
	var designMembersMetricData []DesignMembersMetric
	sql := `
			SELECT
				iddesign_member_metric AS id,
				design_member_metric_name AS design_members_metric_name,
				design_member_metric_total_depth AS design_members_metric_total_depth,
			    design_member_metric_weight AS design_members_metric_weight   
			FROM
				design_member_metric
			`
	err := s.client.Select(&designMembersMetricData, sql)
	return &designMembersMetricData, err
}

//func (s *designMemberMetricDao) Create(iShapeEnglishName string) error {
//	sql := `INSERT INTO i_shape_english (i_shape_english_name) VALUES(?)`
//	_, err := s.client.Exec(sql, iShapeEnglishName)
//	return err
//}
