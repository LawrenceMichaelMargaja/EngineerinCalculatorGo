package calculations

import (
	"db/api/domain"
	"db/api/domain/amir_calc"
	"db/api/services"
	"db/api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type steelArgumentsController struct {
}

var (
	SteelArgumentsController *steelArgumentsController
)

func init() {
	SteelArgumentsController = &steelArgumentsController{}
}

func (controller *steelArgumentsController) GetSteelArguments(c *gin.Context) {

	var body domain.CreateSteelArguments

	if err := c.ShouldBindJSON(&body); err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "calculation body must conform to format",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}

	data, err := services.SteelArgumentService.Get(body.ID)

	SteelSections_w, _ := strconv.ParseFloat(body.SteelSections_w, 64)
	SteelSections_a, _ := strconv.ParseFloat(body.SteelSections_a, 64)
	SteelSections_d, _ := strconv.ParseFloat(body.SteelSections_d, 64)
	SteelSections_ddet, _ := strconv.ParseFloat(body.SteelSections_ddet, 64)
	SteelSections_ht, _ := strconv.ParseFloat(body.SteelSections_ht, 64)
	SteelSections_h, _ := strconv.ParseFloat(body.SteelSections_h, 64)
	SteelSections_od, _ := strconv.ParseFloat(body.SteelSections_od, 64)
	SteelSections_bf, _ := strconv.ParseFloat(body.SteelSections_bf, 64)
	SteelSections_bfdet, _ := strconv.ParseFloat(body.SteelSections_bfdet, 64)
	SteelSections_bhss, _ := strconv.ParseFloat(body.SteelSections_bhss, 64)
	SteelSections_b, _ := strconv.ParseFloat(body.SteelSections_b, 64)
	SteelSections_id, _ := strconv.ParseFloat(body.SteelSections_id, 64)
	SteelSections_tw, _ := strconv.ParseFloat(body.SteelSections_tw, 64)
	SteelSections_twdet, _ := strconv.ParseFloat(body.SteelSections_twdet, 64)
	SteelSections_twdet_2, _ := strconv.ParseFloat(body.SteelSections_twdet_2, 64)
	SteelSections_tf, _ := strconv.ParseFloat(body.SteelSections_tf, 64)
	SteelSections_tfdet, _ := strconv.ParseFloat(body.SteelSections_tfdet, 64)
	SteelSections_th, _ := strconv.ParseFloat(body.SteelSections_th, 64)
	SteelSections_tnom, _ := strconv.ParseFloat(body.SteelSections_tnom, 64)
	SteelSections_tdes, _ := strconv.ParseFloat(body.SteelSections_tdes, 64)
	SteelSections_kdes, _ := strconv.ParseFloat(body.SteelSections_kdes, 64)
	SteelSections_kdet, _ := strconv.ParseFloat(body.SteelSections_kdet, 64)
	SteelSections_k1, _ := strconv.ParseFloat(body.SteelSections_k1, 64)
	SteelSections_x, _ := strconv.ParseFloat(body.SteelSections_x, 64)
	SteelSections_y, _ := strconv.ParseFloat(body.SteelSections_y, 64)
	SteelSections_eo, _ := strconv.ParseFloat(body.SteelSections_eo, 64)
	SteelSections_xp, _ := strconv.ParseFloat(body.SteelSections_xp, 64)
	SteelSections_yp, _ := strconv.ParseFloat(body.SteelSections_yp, 64)
	SteelSections_bf_2tf, _ := strconv.ParseFloat(body.SteelSections_bf_2tf, 64)
	SteelSections_b_t, _ := strconv.ParseFloat(body.SteelSections_b_t, 64)
	SteelSections_b_tdes, _ := strconv.ParseFloat(body.SteelSections_b_tdes, 64)
	SteelSections_h_tw, _ := strconv.ParseFloat(body.SteelSections_h_tw, 64)
	SteelSections_h_tdes, _ := strconv.ParseFloat(body.SteelSections_h_tdes, 64)
	SteelSections_d_t, _ := strconv.ParseFloat(body.SteelSections_d_t, 64)
	SteelSections_ix, _ := strconv.ParseFloat(body.SteelSections_ix, 64)
	SteelSections_zx, _ := strconv.ParseFloat(body.SteelSections_zx, 64)
	SteelSections_sx, _ := strconv.ParseFloat(body.SteelSections_sx, 64)
	SteelSections_rx, _ := strconv.ParseFloat(body.SteelSections_rx, 64)
	SteelSections_iy, _ := strconv.ParseFloat(body.SteelSections_iy, 64)
	SteelSections_zy, _ := strconv.ParseFloat(body.SteelSections_zy, 64)
	SteelSections_sy, _ := strconv.ParseFloat(body.SteelSections_sy, 64)
	SteelSections_ry, _ := strconv.ParseFloat(body.SteelSections_ry, 64)
	SteelSections_iz, _ := strconv.ParseFloat(body.SteelSections_iz, 64)
	SteelSections_rz, _ := strconv.ParseFloat(body.SteelSections_rz, 64)
	SteelSections_sz, _ := strconv.ParseFloat(body.SteelSections_sz, 64)
	SteelSections_j, _ := strconv.ParseFloat(body.SteelSections_j, 64)
	SteelSections_cw, _ := strconv.ParseFloat(body.SteelSections_cw, 64)
	SteelSections_c, _ := strconv.ParseFloat(body.SteelSections_c, 64)
	SteelSections_wno, _ := strconv.ParseFloat(body.SteelSections_wno, 64)
	SteelSections_sw1, _ := strconv.ParseFloat(body.SteelSections_sw1, 64)
	SteelSections_sw2, _ := strconv.ParseFloat(body.SteelSections_sw2, 64)
	SteelSections_sw3, _ := strconv.ParseFloat(body.SteelSections_sw3, 64)
	SteelSections_qf, _ := strconv.ParseFloat(body.SteelSections_qf, 64)
	SteelSections_qw, _ := strconv.ParseFloat(body.SteelSections_qw, 64)
	SteelSections_ro, _ := strconv.ParseFloat(body.SteelSections_ro, 64)
	SteelSections_hfc, _ := strconv.ParseFloat(body.SteelSections_hfc, 64)
	SteelSections_tan_a, _ := strconv.ParseFloat(body.SteelSections_tan_a, 64)
	SteelSections_iw, _ := strconv.ParseFloat(body.SteelSections_iw, 64)
	SteelSections_za, _ := strconv.ParseFloat(body.SteelSections_za, 64)
	SteelSections_zb, _ := strconv.ParseFloat(body.SteelSections_zb, 64)
	SteelSections_zc, _ := strconv.ParseFloat(body.SteelSections_zc, 64)
	SteelSections_wa, _ := strconv.ParseFloat(body.SteelSections_wa, 64)
	SteelSections_wb, _ := strconv.ParseFloat(body.SteelSections_wb, 64)
	SteelSections_wc, _ := strconv.ParseFloat(body.SteelSections_wc, 64)
	SteelSections_swa, _ := strconv.ParseFloat(body.SteelSections_swa, 64)
	SteelSections_swb, _ := strconv.ParseFloat(body.SteelSections_swb, 64)
	SteelSections_swc, _ := strconv.ParseFloat(body.SteelSections_swc, 64)
	SteelSections_sza, _ := strconv.ParseFloat(body.SteelSections_sza, 64)
	SteelSections_szb, _ := strconv.ParseFloat(body.SteelSections_szb, 64)
	SteelSections_szc, _ := strconv.ParseFloat(body.SteelSections_szc, 64)
	SteelSections_rts, _ := strconv.ParseFloat(body.SteelSections_rts, 64)
	SteelSections_ho, _ := strconv.ParseFloat(body.SteelSections_ho, 64)
	SteelSections_pa, _ := strconv.ParseFloat(body.SteelSections_pa, 64)
	SteelSections_pa2, _ := strconv.ParseFloat(body.SteelSections_pa2, 64)
	SteelSections_pb, _ := strconv.ParseFloat(body.SteelSections_pb, 64)
	SteelSections_pc, _ := strconv.ParseFloat(body.SteelSections_pc, 64)
	SteelSections_pd,_ := strconv.ParseFloat(body.SteelSections_pd, 64)
	SteelSections_t, _ := strconv.ParseFloat(body.SteelSections_t, 64)
	SteelSections_wgi, _ := strconv.ParseFloat(body.SteelSections_wgi, 64)
	SteelSections_wgo, _ := strconv.ParseFloat(body.SteelSections_wgo, 64)


	Pt, Pc, Mcx, Mcy, Vcx, Vcy, Pratio, MxRatio, MyRatio, VxRatio, VyRatio, Combined, KLr := amir_calc.BeamColumnAnalysis(body.Analysis, body.Method, body.Units, body.TensileFactor, body.CompressFactor, body.BendingFactor, body.ShearFactor, body.Vx, body.Vy, body.Mx, body.My, body.P, body.Lb, body.Lx, body.Kx, body.Ly, body.Ky, body.LLT, body.Cbx, body.Cby, body.SRc, body.SRt, body.ModE, body.YieldStr, body.UltStr, body.Shape, body.SteelSections_t_f, SteelSections_w, SteelSections_a, SteelSections_d, SteelSections_ddet, SteelSections_ht, SteelSections_h, SteelSections_od, SteelSections_bf, SteelSections_bfdet, SteelSections_bhss, SteelSections_b, SteelSections_id, SteelSections_tw, SteelSections_twdet, SteelSections_twdet_2, SteelSections_tf, SteelSections_tfdet, SteelSections_th, SteelSections_tnom, SteelSections_tdes, SteelSections_kdes, SteelSections_kdet, SteelSections_k1, SteelSections_x, SteelSections_y, SteelSections_eo, SteelSections_xp, SteelSections_yp, SteelSections_bf_2tf, SteelSections_b_t, SteelSections_b_tdes, SteelSections_h_tw, SteelSections_h_tdes, SteelSections_d_t, SteelSections_ix, SteelSections_zx, SteelSections_sx, SteelSections_rx, SteelSections_iy, SteelSections_zy, SteelSections_sy, SteelSections_ry, SteelSections_iz, SteelSections_rz, SteelSections_sz, SteelSections_j, SteelSections_cw, SteelSections_c, SteelSections_wno, SteelSections_sw1, SteelSections_sw2, SteelSections_sw3, SteelSections_qf, SteelSections_qw, SteelSections_ro, SteelSections_hfc, SteelSections_tan_a, SteelSections_iw, SteelSections_za, SteelSections_zb, SteelSections_zc, SteelSections_wa, SteelSections_wb, SteelSections_wc, SteelSections_swa, SteelSections_swb, SteelSections_swc, SteelSections_sza, SteelSections_szb, SteelSections_szc, SteelSections_rts, SteelSections_ho, SteelSections_pa, SteelSections_pa2, SteelSections_pb, SteelSections_pc, SteelSections_pd, SteelSections_t, SteelSections_wgi, SteelSections_wgo)

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "Calculation Error when trying to insert UI data : " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code:       "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, data)
}

func (controller *steelArgumentsController) GetBeamAnalysisResult(c *gin.Context) {
	cShapeEnglishData, err := services.CShapeEnglishService.Get()

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "Error when trying to fetch english C shape data : " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, cShapeEnglishData)
}
