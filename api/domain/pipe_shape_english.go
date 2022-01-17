package domain

type PipeShapeEnglish struct {
	Id int `json:"id" db:"id"`
	PipeShapeEnglishName string `json:"pipe_shape_english_name" db:"pipe_shape_english_name"`
}

type CreatePipeShapeEnglish struct {
	PipeShapeNameEnglish string `json:"pipe_shape_name_english" db:"pipe_shape_name_english"`
}
