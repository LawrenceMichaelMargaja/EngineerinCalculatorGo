package domain

type LShapeEnglish struct {
	Id int `json:"id" db:"id"`
	LShapeEnglishName string `json:"l_shape_english_name" db:"l_shape_english_name"`
}

type CreateLShapeEnglish struct {
	LShapeEnglishName string `json:"l_shape_english_name" db:"l_shape_english_name"`
}
