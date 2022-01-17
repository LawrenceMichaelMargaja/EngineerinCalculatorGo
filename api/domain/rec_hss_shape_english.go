package domain

type RecHsShapeEnglish struct {
	Id int `json:"id" db:"id"`
	RecHsShapeEnglishName string `json:"rec_hs_shape_english_name" db:"rec_hs_shape_english_name"`
}

type CreateRecHssShapeEnglish struct {
	RecHssShapeEnglishName string `json:"rec_hss_shape_english_name" db:"rec_hss_shape_english_name"`
}
