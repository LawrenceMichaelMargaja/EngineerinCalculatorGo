package domain

type SteelTypeEnglish struct {
	Id                   int    `json:"id" db:"id"`
	SteelTypeEnglishName string `json:"steel_type_english_name" db:"steel_type_english_name"`
	SteelTypeEnglishE    string `json:"steel_type_english_e" db:"steel_type_english_e"`
	SteelTypeEnglishFy   string `json:"steel_type_english_fy" db:"steel_type_english_fy"`
	SteelTypeEnglishFu   string `json:"steel_type_english_fu" db:"steel_type_english_fu"`
}

type CreateSteelTypeEnglish struct {
	SteelTypeEnglishName string `json:"steel_type_english_name" db:"steel_type_english_name"`
	SteelTypeEnglishE    string `json:"steel_type_english_e" db:"steel_type_english_e"`
	SteelTypeEnglishFy   string `json:"steel_type_english_fy" db:"steel_type_english_fy"`
	SteelTypeEnglishFu   string `json:"steel_type_english_fu" db:"steel_type_english_fu"`
}
