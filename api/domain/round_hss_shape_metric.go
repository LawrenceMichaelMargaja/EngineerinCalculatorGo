package domain

type RoundHSShapeMetric struct {
	Id int `json:"id" db:"id"`
	RoundHsShapeName string `json:"round_hs_shape_name" db:"round_hs_shape_name"`
}

type CreateRoundHssShape struct {
	RoundHssShapeName string `json:"round_hss_shape_name" db:"round_hss_shape_name"`
}
