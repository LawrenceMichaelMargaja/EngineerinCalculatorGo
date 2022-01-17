package domain

import (
	"db/api/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type sectionPropertiesEnglishDaoInterface interface {
	Get() (*[]SectionPropertiesEnglish, error)
	Create(sectionName string, a string, j string, ixp string, iyp string, iw string, sxp string, syp string) error
	SetClient()
}

type sectionPropertiesEnglishDao struct {
	client *sqlx.DB
}

var (
	SectionPropertiesEnglishDao sectionPropertiesEnglishDaoInterface
)

func init()  {
	SectionPropertiesEnglishDao = &sectionPropertiesEnglishDao{}
	SectionPropertiesEnglishDao.SetClient()
}

func (s *sectionPropertiesEnglishDao) SetClient() {
	s.client = utils.GetMYSQLConnection()
}

func (s *sectionPropertiesEnglishDao) Get() (*[]SectionPropertiesEnglish, error) {
	var sectionPropertiesEnglishData []SectionPropertiesEnglish
	sql := `
			SELECT
				idsection_properties_english AS id,
				section_properties_english_name AS section_properties_english_name,
				section_properties_english_A AS section_properties_english_a,
				section_properties_english_J AS section_properties_english_j,
				section_properties_english_Ixp AS section_properties_english_ixp,
				section_properties_english_Iyp AS section_properties_english_iyp,
				section_properties_english_Iw AS section_properties_english_iw,
				section_properties_english_Sxp AS section_properties_english_sxp,
				section_properties_english_Syp AS section_properties_english_syp
			FROM
				section_properties_english
			`
	err := s.client.Select(&sectionPropertiesEnglishData, sql)
	return &sectionPropertiesEnglishData, err
}

func (s *sectionPropertiesEnglishDao) Create(sectionName string, a string, j string, ixp string, iyp string, iw string, sxp string, syp string) error {
	sql := `INSERT INTO section_properties_english (section_properties_english_name, A_english, J_english, Ixp_english, Iyp_english, Iw_english, Sxp_english, Syp_english) VALUES(?,?,?,?,?,?,?,?)`
	_, err := s.client.Exec(sql, sectionName, a, j, ixp, iyp, iw, sxp, syp)
	return err
}