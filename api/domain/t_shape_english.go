package domain

type TShapeEnglish struct {
	Id int `json:"id" db:"id"`
	TShapeEnglishName string `json:"t_shape_english_name" db:"t_shape_english_name"`
}

type CreateTShapeEnglish struct {
	TShapeEnglishName string `json:"t_shape_english_name" db:"t_shape_english_name"`
}
