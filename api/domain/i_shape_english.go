package domain

type IShapeEnglish struct {
	Id int `json:"id" db:"id"`
	IShapeEnglishName string `json:"i_shape_english_name" db:"i_shape_english_name"`
}

type CreateIShapeEnglish struct {
	IShapeEnglishName string `json:"i_shape_english_name" db:"i_shape_english_name"`
}
