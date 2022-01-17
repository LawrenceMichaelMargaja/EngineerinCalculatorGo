package domain

import (
	"db/api/utils"
	"github.com/jmoiron/sqlx"
)

type steelArgumentsDaoInterface interface {
	Get(ID int) (*CreateSteelArguments, error)
	SetClient()
}

type steelArgumentsDao struct {
	client *sqlx.DB
}

var (
	SteelArgumentsDao steelArgumentsDaoInterface
)

func init()  {
	SteelArgumentsDao = &steelArgumentsDao{}
	SteelArgumentsDao.SetClient()
}

func (s *steelArgumentsDao) SetClient() {
	s.client = utils.GetMYSQLConnection()
}

func (s *steelArgumentsDao) Get(ID int) (*CreateSteelArguments, error) {
	var steelArgumentsData CreateSteelArguments
	sql := `
			SELECT
				steel_sections_t_f AS steel_sections_t_f,
				steel_sections_w AS steel_sections_w,
				steel_sections_a AS steel_sections_a,
				steel_sections_d AS steel_sections_d,
				steel_sections_ddet AS steel_sections_ddet,
				steel_sections_ht AS steel_sections_ht,
				steel_sections_h AS steel_sections_h,
				steel_sections_od AS steel_sections_od,
				steel_sections_bf AS steel_sections_bf,
				steel_sections_w AS steel_sections_w,
				steel_sections_bhss AS steel_sections_bhss,
				steel_sections_b AS steel_sections_b,
				steel_sections_id AS steel_sections_id,
				steel_sections_tw AS steel_sections_tw,
				steel_sections_twdet AS steel_sections_twdet,
				steel_sections_twdet_2 AS steel_sections_twdet_2,
				steel_sections_tf AS steel_sections_tf,
				steel_sections_tfdet AS steel_sections_tfdet,
				steel_sections_th AS steel_sections_th,
				steel_sections_tnom AS steel_sections_tnom,
				steel_sections_tdes AS steel_sections_tdes,
				steel_sections_kdes AS steel_sections_kdes,
				steel_sections_kdet AS steel_sections_kdet,
				steel_sections_k1 AS steel_sections_k1,
			    steel_sections_x AS steel_sections_x,
			    steel_sections_y AS steel_sections_y,
			    steel_sections_eo AS steel_sections_eo,
			    steel_sections_xp AS steel_sections_xp,
			    steel_sections_yp AS steel_sections_yp,
			    steel_sections_bf_2tf AS steel_sections_bf_2_tf,
			    steel_sections_b_t AS steel_sections_b_t,
			    steel_sections_b_tdes AS steel_sections_b_tdes,
			    steel_sections_h_tw AS steel_sections_h_tw,
			    steel_sections_h_tdes AS steel_sections_h_tdes,
			    steel_sections_d_t AS steel_sections_d_t,
			    steel_sections_ix AS steel_sections_ix,
			    steel_sections_zx AS steel_sections_zx,
			    steel_sections_sx AS steel_sections_sx,
			    steel_sections_rx AS steel_sections_rx,
			    steel_sections_iy AS steel_sections_iy,
			    steel_sections_zy AS steel_sections_zy,
			    steel_sections_sy AS steel_sections_sy,
			    steel_sections_ry AS steel_sections_ry,
			    steel_sections_iz AS steel_sections_iz,
			    steel_sections_rz AS steel_sections_rz,
			    steel_sections_sz AS steel_sections_sz,
			    steel_sections_j AS steel_sections_j,
			    steel_sections_cw AS steel_sections_cw,
			    steel_sections_c AS steel_sections_c,
			    steel_sections_wno AS steel_sections_wno,
			    steel_sections_sw1 AS steel_sections_sw_1,
			    steel_sections_sw2 AS steel_sections_sw_2,
			    steel_sections_sw3 AS steel_sections_sw_3,
			    steel_sections_qf AS steel_sections_qf,
			    steel_sections_qw AS steel_sections_qw,
			    steel_sections_ro AS steel_sections_ro,
			    steel_sections_hfc AS steel_sections_hfc,
			    steel_sections_tan_a AS steel_sections_tan_a,
			    steel_sections_iw AS steel_sections_iw,
			    steel_sections_za AS steel_sections_za,
				steel_sections_zb AS steel_sections_zb,
			    steel_sections_zc AS steel_sections_zc,
			    steel_sections_wa AS steel_sections_wa,
			    steel_sections_wb AS steel_sections_wb,
			    steel_sections_wc AS steel_sections_wc,
			    steel_sections_swa AS steel_sections_swa,
			    steel_sections_swb AS steel_sections_swb,
			    steel_sections_swc AS steel_sections_swc,
			    steel_sections_sza AS steel_sections_sza,
				steel_sections_szb AS steel_sections_szb,
			    steel_sections_szc AS steel_sections_szc,
			    steel_sections_rts AS steel_sections_rts,
			    steel_sections_ho AS steel_sections_ho,
			    steel_sections_pa AS steel_sections_pa,
			    steel_sections_pa2 AS steel_sections_pa_2,
			    steel_sections_pb AS steel_sections_pb,
			    steel_sections_pc AS steel_sections_pc,
			    steel_sections_pd AS steel_sections_pd,
				steel_sections_t AS steel_sections_t,
			    steel_sections_wgi AS steel_sections_wgi,
			    steel_sections_wgo AS steel_sections_wgo
			FROM
				steel_sections
			WHERE 
				idsteel_sections = ?
			`
	err := s.client.Get(&steelArgumentsData, sql, ID)
	return &steelArgumentsData, err
}

func (s *steelArgumentsDao) Create(steelTypeEnglishName string, steelTypeEnglishE string, steelTypeEnglishFy string, steelTypeEnglishFu string) error {
	sql := `INSERT INTO steel_types_english (steel_types_english_name, steel_types_english_E, steel_types_english_Fy, steel_types_english_Fu) VALUES(?,?,?,?)`
	_, err := s.client.Exec(sql, steelTypeEnglishName, steelTypeEnglishE, steelTypeEnglishFy, steelTypeEnglishFu)
	return err
}