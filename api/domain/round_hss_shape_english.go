package domain

type RoundHsShapeEnglish struct {
	Id int `json:"id" db:"id"`
	RoundHsShapeEnglishName string `json:"round_hs_shape_english_name" db:"round_hs_shape_english_name"`
}

type CreateRoundHssShapeEnglish struct {
	RoundHssShapeName string `json:"round_hss_shape_name" db:"round_hss_shape_name"`
}
