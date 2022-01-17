package domain

import (
	"db/api/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type steelTypeEnglishDaoInterface interface {
	Get() (*[]SteelTypeEnglish, error)
	Create(steelTypeEnglishName string, steelTypeEnglishE string, steelTypeEnglishFy string, steelTypeEnglishFu string) error
	SetClient()
}

type steelTypeEnglishDao struct {
	client *sqlx.DB
}

var (
	SteelTypeEnglishDao steelTypeEnglishDaoInterface
)

func init()  {
	SteelTypeEnglishDao = &steelTypeEnglishDao{}
	SteelTypeEnglishDao.SetClient()
}

func (s *steelTypeEnglishDao) SetClient() {
	s.client = utils.GetMYSQLConnection()
}

func (s *steelTypeEnglishDao) Get() (*[]SteelTypeEnglish, error) {
	var steelTypeEnglishData []SteelTypeEnglish
	sql := `
			SELECT
				idsteel_types_english AS id,
				steel_types_english_name AS steel_type_english_name,
				steel_types_english_E AS steel_type_english_e,
				steel_types_english_Fy AS steel_type_english_fy,
				steel_types_english_Fu AS steel_type_english_fu
			FROM
				steel_types_english
			`
	err := s.client.Select(&steelTypeEnglishData, sql)
	return &steelTypeEnglishData, err
}

func (s *steelTypeEnglishDao) Create(steelTypeEnglishName string, steelTypeEnglishE string, steelTypeEnglishFy string, steelTypeEnglishFu string) error {
	sql := `INSERT INTO steel_types_english (steel_types_english_name, steel_types_english_E, steel_types_english_Fy, steel_types_english_Fu) VALUES(?,?,?,?)`
	_, err := s.client.Exec(sql, steelTypeEnglishName, steelTypeEnglishE, steelTypeEnglishFy, steelTypeEnglishFu)
	return err
}