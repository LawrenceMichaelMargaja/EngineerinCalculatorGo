package domain

type TwoLShapeEnglish struct {
	Id int `json:"id" db:"id"`
	TwoLShapeEnglishName string `json:"two_l_shape_english_name" db:"two_l_shape_english_name"`
}

type CreateTwoLShapeEnglish struct {
	TwoLShapeEnglishName string `json:"two_l_shape_english_name" db:"two_l_shape_english_name"`
}
