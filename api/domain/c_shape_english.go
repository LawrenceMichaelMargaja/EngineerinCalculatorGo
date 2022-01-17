package domain

type CShapeEnglish struct {
	Id                int    `json:"id" db:"id"`
	CShapeEnglishName string `json:"c_shape_english_name" db:"c_shape_english_name"`
}

type CreateCShapeEnglish struct {
	CShapeEnglishName string `json:"c_shape_english_name" db:"c_shape_english_name"`
}