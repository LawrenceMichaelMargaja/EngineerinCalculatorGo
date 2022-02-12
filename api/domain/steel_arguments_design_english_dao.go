package domain

import (
	"db/api/domain/amir_calc"
	//"db/api/domain/amir_calc/calculation_domains"
	"db/api/utils"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type steelArgumentsDesignDaoEnglishInterfaceEnglish interface {
	//GetAdditionalDatabaseValues(c *SteelArgumentsDesignResults) (*SteelArgumentsDesignResults, error)
	GetEnglish(c *CreateSteelDesignArguments) (*CreateSteelDesignArguments, *AdditionalResultsEnglish, *BeamAnalysisDesignResultsEnglish, error)
	SetClient()
}

type steelArgumentsDesignDaoEnglish struct {
	client *sqlx.DB
}

type BeamAnalysisDesignResultsEnglish struct {
	Id                  int     `json:"id"`
	Name                string  `json:"name"`
	OverallDepth        float64 `json:"overall_depth"`
	Weight              float64 `json:"weight"`
	CriticalDesignRatio float64 `json:"critical_design_ratio"`
	KlrValue            float64 `json:"klr_value"`
	Pt                  float64 `json:"pt"`
	Pc                  float64 `json:"pc"`
	Mcx                 float64 `json:"mcx"`
	Mcy                 float64 `json:"mcy"`
	Vcx                 float64 `json:"vcx"`
	Vcy                 float64 `json:"vcy"`
	Pratio              float64 `json:"pratio"`
	MxRatio             float64 `json:"mx_ratio"`
	MyRatio             float64 `json:"my_ratio"`
	VxRatio             float64 `json:"vx_ratio"`
	VyRatio             float64 `json:"vy_ratio"`
	Combined            float64 `json:"combined"`
	KLr                 float64 `json:"k_lr"`
}

type AdditionalResultsEnglish struct {
	Id                  int     `json:"id" db:"id"`
	Name                string  `json:"name" db:"name"`
	OverallDepth        float64 `json:"overall_depth" db:"overall_depth"`
	Weight              float64 `json:"weight" db:"weight"`
	CriticalDesignRatio float64 `json:"critical_design_ratio" db:"critical_design_ratio"`
	KlrValue            float64 `json:"klr_value" db:"klr_value"`
	Pt                  float64 `json:"pt"`
	Pc                  float64 `json:"pc"`
	Mcx                 float64 `json:"mcx"`
	Mcy                 float64 `json:"mcy"`
	Vcx                 float64 `json:"vcx"`
	Vcy                 float64 `json:"vcy"`
	Pratio              float64 `json:"pratio"`
	MxRatio             float64 `json:"mx_ratio"`
	MyRatio             float64 `json:"my_ratio"`
	VxRatio             float64 `json:"vx_ratio"`
	VyRatio             float64 `json:"vy_ratio"`
	Combined            float64 `json:"combined"`
	KLr                 float64 `json:"k_lr"`
}

//func (c *BeamAnalysisDesignResultsEnglish) DisplayDesignData() {
//	fmt.Println("c.Pt", c.Pt)
//	fmt.Println("c.Pc", c.Pc)
//	fmt.Println("c.Mcx", c.Mcx)
//	fmt.Println("c.Mcy", c.Mcy)
//	fmt.Println("c.Vcx", c.Vcx)
//	fmt.Println("c.Vcy", c.Vcy)
//	fmt.Println("c.Pratio", c.Pratio)
//	fmt.Println("c.MxRatio", c.MxRatio)
//	fmt.Println("c.MyRatio", c.MyRatio)
//	fmt.Println("c.VxRatio", c.VxRatio)
//	fmt.Println("c.VyRatio", c.VyRatio)
//	fmt.Println("c.Combined", c.Combined)
//	fmt.Println("c.KLr", c.KLr)
//}

var (
	SteelArgumentsDesignDaoEnglish steelArgumentsDesignDaoEnglishInterfaceEnglish
)

func init() {
	SteelArgumentsDesignDaoEnglish = &steelArgumentsDesignDaoEnglish{}
	SteelArgumentsDesignDaoEnglish.SetClient()
}

func (s *steelArgumentsDesignDaoEnglish) SetClient() {
	s.client = utils.GetMYSQLConnection()
}

func (s *steelArgumentsDesignDaoEnglish) GetEnglish(c *CreateSteelDesignArguments) (*CreateSteelDesignArguments, *AdditionalResultsEnglish, *BeamAnalysisDesignResultsEnglish, error) {
	var steelArgumentsData CreateSteelDesignArguments
	//var theAdditionalResultsEnglish SteelArgumentsDesignResults
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
				steel_sections_aisc_manual_label_english = ?
			`
	err := s.client.Get(&steelArgumentsData, sql, c.Name)

	SteelSections_w, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_w, 64)
	SteelSections_a, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_a, 64)
	SteelSections_d, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_d, 64)
	SteelSections_ddet, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_ddet, 64)
	SteelSections_ht, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_ht, 64)
	SteelSections_h, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_h, 64)
	SteelSections_od, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_od, 64)
	SteelSections_bf, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_bf, 64)
	SteelSections_bfdet, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_bfdet, 64)
	SteelSections_bhss, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_bhss, 64)
	SteelSections_b, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_b, 64)
	SteelSections_id, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_id, 64)
	SteelSections_tw, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_tw, 64)
	SteelSections_twdet, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_twdet, 64)
	SteelSections_twdet_2, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_twdet_2, 64)
	SteelSections_tf, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_tf, 64)
	SteelSections_tfdet, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_tfdet, 64)
	SteelSections_th, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_th, 64)
	SteelSections_tnom, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_tnom, 64)
	SteelSections_tdes, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_tdes, 64)
	SteelSections_kdes, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_kdes, 64)
	SteelSections_kdet, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_kdet, 64)
	SteelSections_k1, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_k1, 64)
	SteelSections_x, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_x, 64)
	SteelSections_y, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_y, 64)
	SteelSections_eo, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_eo, 64)
	SteelSections_xp, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_xp, 64)
	SteelSections_yp, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_yp, 64)
	SteelSections_bf_2tf, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_bf_2tf, 64)
	SteelSections_b_t, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_b_t, 64)
	SteelSections_b_tdes, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_b_tdes, 64)
	SteelSections_h_tw, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_h_tw, 64)
	SteelSections_h_tdes, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_h_tdes, 64)
	SteelSections_d_t, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_d_t, 64)
	SteelSections_ix, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_ix, 64)
	SteelSections_zx, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_zx, 64)
	SteelSections_sx, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_sx, 64)
	SteelSections_rx, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_rx, 64)
	SteelSections_iy, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_iy, 64)
	SteelSections_zy, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_zy, 64)
	SteelSections_sy, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_sy, 64)
	SteelSections_ry, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_ry, 64)
	SteelSections_iz, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_iz, 64)
	SteelSections_rz, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_rz, 64)
	SteelSections_sz, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_sz, 64)
	SteelSections_j, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_j, 64)
	SteelSections_cw, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_cw, 64)
	SteelSections_c, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_c, 64)
	SteelSections_wno, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_wno, 64)
	SteelSections_sw1, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_sw1, 64)
	SteelSections_sw2, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_sw2, 64)
	SteelSections_sw3, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_sw3, 64)
	SteelSections_qf, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_qf, 64)
	SteelSections_qw, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_qw, 64)
	SteelSections_ro, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_ro, 64)
	SteelSections_hfc, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_hfc, 64)
	SteelSections_tan_a, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_tan_a, 64)
	SteelSections_iw, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_iw, 64)
	SteelSections_za, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_za, 64)
	SteelSections_zb, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_zb, 64)
	SteelSections_zc, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_zc, 64)
	SteelSections_wa, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_wa, 64)
	SteelSections_wb, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_wb, 64)
	SteelSections_wc, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_wc, 64)
	SteelSections_swa, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_swa, 64)
	SteelSections_swb, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_swb, 64)
	SteelSections_swc, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_swc, 64)
	SteelSections_sza, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_sza, 64)
	SteelSections_szb, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_szb, 64)
	SteelSections_szc, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_szc, 64)
	SteelSections_rts, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_rts, 64)
	SteelSections_ho, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_ho, 64)
	SteelSections_pa, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_pa, 64)
	SteelSections_pa2, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_pa2, 64)
	SteelSections_pb, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_pb, 64)
	SteelSections_pc, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_pc, 64)
	SteelSections_pd, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_pd, 64)
	SteelSections_t, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_t, 64)
	SteelSections_wgi, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_wgi, 64)
	SteelSections_wgo, _ := strconv.ParseFloat(steelArgumentsData.SteelSections_wgo, 64)

	//fmt.Println("steelArgumentsData.Analysis", c.Analysis, reflect.ValueOf(c.Analysis).Kind())
	//fmt.Println("steelArgumentsData.Method", c.Method, reflect.ValueOf(c.Method).Kind())
	//fmt.Println("steelArgumentsData", utils.ToJSON(steelArgumentsData))
	Pt, Pc, Mcx, Mcy, Vcx, Vcy, Pratio, MxRatio, MyRatio, VxRatio, VyRatio, Combined, KLr := amir_calc.BeamColumnAnalysis(
		c.Analysis,
		c.Analysis,
		c.Units,
		c.TensileFactor,
		c.CompressFactor,
		c.BendingFactor,
		c.ShearFactor,
		c.Vx,
		c.Vy,
		c.Mx,
		c.My,
		c.P,
		c.Lb,
		c.Lx,
		c.Kx,
		c.Ly,
		c.Ky,
		c.LLT,
		c.Cbx,
		c.Cby,
		c.SRc,
		c.SRt,
		c.ModE,
		c.YieldStr,
		c.UltStr,
		c.Shape,
		steelArgumentsData.SteelSections_t_f,
		SteelSections_w,
		SteelSections_a,
		SteelSections_d,
		SteelSections_ddet,
		SteelSections_ht,
		SteelSections_h,
		SteelSections_od,
		SteelSections_bf,
		SteelSections_bfdet,
		SteelSections_bhss,
		SteelSections_b,
		SteelSections_id,
		SteelSections_tw,
		SteelSections_twdet,
		SteelSections_twdet_2,
		SteelSections_tf,
		SteelSections_tfdet,
		SteelSections_th,
		SteelSections_tnom,
		SteelSections_tdes,
		SteelSections_kdes,
		SteelSections_kdet,
		SteelSections_k1,
		SteelSections_x,
		SteelSections_y,
		SteelSections_eo,
		SteelSections_xp,
		SteelSections_yp,
		SteelSections_bf_2tf,
		SteelSections_b_t,
		SteelSections_b_tdes,
		SteelSections_h_tw,
		SteelSections_h_tdes,
		SteelSections_d_t,
		SteelSections_ix,
		SteelSections_zx,
		SteelSections_sx,
		SteelSections_rx,
		SteelSections_iy,
		SteelSections_zy,
		SteelSections_sy,
		SteelSections_ry,
		SteelSections_iz,
		SteelSections_rz,
		SteelSections_sz,
		SteelSections_j,
		SteelSections_cw,
		SteelSections_c,
		SteelSections_wno,
		SteelSections_sw1,
		SteelSections_sw2,
		SteelSections_sw3,
		SteelSections_qf,
		SteelSections_qw,
		SteelSections_ro,
		SteelSections_hfc,
		SteelSections_tan_a,
		SteelSections_iw,
		SteelSections_za,
		SteelSections_zb,
		SteelSections_zc,
		SteelSections_wa,
		SteelSections_wb,
		SteelSections_wc,
		SteelSections_swa,
		SteelSections_swb,
		SteelSections_swc,
		SteelSections_sza,
		SteelSections_szb,
		SteelSections_szc,
		SteelSections_rts,
		SteelSections_ho,
		SteelSections_pa,
		SteelSections_pa2,
		SteelSections_pb,
		SteelSections_pc,
		SteelSections_pd,
		SteelSections_t,
		SteelSections_wgi,
		SteelSections_wgo)

	calculationResult := &BeamAnalysisDesignResultsEnglish{
		Pt:       Pt,
		Pc:       Pc,
		Mcx:      Mcx,
		Mcy:      Mcy,
		Vcx:      Vcx,
		Vcy:      Vcy,
		Pratio:   Pratio,
		MxRatio:  MxRatio,
		MyRatio:  MyRatio,
		VxRatio:  VxRatio,
		VyRatio:  VyRatio,
		Combined: Combined,
		KLr:      KLr,
	}

	//Data, err := calculation_domains.GetTest(c.Name)
	data, testResult, _ := s.GetDesignResultEnglish(c, calculationResult)

	//NewResult := &SteelArgumentsDesignResults{
	//	id:                  id,
	//	name:                name,
	//	overallDepth:        overallDepth,
	//	weight:              weight,
	//	criticalDesignRatio: criticalDesignRatio,
	//	klrValue:            klrValue,
	//	Pt:                  Pt,
	//	Pc:                  Pc,
	//	Mcx:                 Mcx,
	//	Mcy:                 Mcy,
	//	Vcx:                 Vcx,
	//	Vcy:                 Vcy,
	//	Pratio:              Pratio,
	//	MxRatio:             MxRatio,
	//	MyRatio:             MyRatio,
	//	VxRatio:             VxRatio,
	//	VyRatio:             VyRatio,
	//	Combined:            Combined,
	//	KLr:                 KLr,
	//}
	//fmt.Println("newErr === ", newErr)
	//fmt.Println("data === ", data)
	fmt.Println("the test results === ", testResult)
	return &steelArgumentsData, data, testResult, err
}

func (s *steelArgumentsDesignDaoEnglish) GetDesignResultEnglish(c *CreateSteelDesignArguments, calculatedResult *BeamAnalysisDesignResultsEnglish) (*AdditionalResultsEnglish, *BeamAnalysisDesignResultsEnglish, error) {
	var theAdditionalResultsEnglish AdditionalResultsEnglish
	sqlNew := `SELECT
					iddesign_member_english AS id,
					design_member_english_name AS name,
					design_member_english_total_depth AS overall_depth,
					design_member_english_weight AS weight
				FROM
					design_member_english
				WHERE
					design_member_english_name = ?`

	errNew := s.client.Get(&theAdditionalResultsEnglish, sqlNew, c.Name)

	promisedResult := &BeamAnalysisDesignResultsEnglish{
		Id:                  theAdditionalResultsEnglish.Id,
		Name:                theAdditionalResultsEnglish.Name,
		OverallDepth:        theAdditionalResultsEnglish.OverallDepth,
		Weight:              theAdditionalResultsEnglish.Weight,
		CriticalDesignRatio: theAdditionalResultsEnglish.CriticalDesignRatio,
		KlrValue:            theAdditionalResultsEnglish.KLr,
		Pt:                  calculatedResult.Pt,
		Pc:                  calculatedResult.Pc,
		Mcx:                 calculatedResult.Mcx,
		Mcy:                 calculatedResult.Mcy,
		Vcx:                 calculatedResult.Vcx,
		Vcy:                 calculatedResult.Vcy,
		Pratio:              calculatedResult.Pratio,
		MxRatio:             calculatedResult.MxRatio,
		MyRatio:             calculatedResult.MyRatio,
		VxRatio:             calculatedResult.VxRatio,
		VyRatio:             calculatedResult.VyRatio,
		Combined:            calculatedResult.Combined,
		KLr:                 calculatedResult.KLr,
	}
	fmt.Println("the new promisedResult", promisedResult)
	fmt.Println("the new body", theAdditionalResultsEnglish)
	return &theAdditionalResultsEnglish, promisedResult, errNew
}

//func (s *steelArgumentsDesignDaoEnglish) GetAdditionalDatabaseValues(c *SteelArgumentsDesignResults) (*SteelArgumentsDesignResults, error) {
//	var additionalDesignData SteelArgumentsDesignResults
//	sql := `SELECT idt_shape_metric AS id, i_shape_metric_name AS i_shape_english_name FROM i_shape_english  WHERE steel_sections_aisc_manual_label_metric = ?`
//	err := s.client.Get(&additionalDesignData, sql, c.name)
//	return &additionalDesignData, err
//}
