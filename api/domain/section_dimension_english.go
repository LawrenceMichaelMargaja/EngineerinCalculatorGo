package domain

type SectionDimensionEnglish struct {
	Id                          int    `json:"id" db:"id"`
	SectionDimensionEnglishName string `json:"section_dimension_english_name" db:"section_dimension_english_name"`
	SectionDimensionEnglishD    string `json:"section_dimension_english_d" db:"section_dimension_english_d"`
	SectionDimensionEnglishB    string `json:"section_dimension_english_b" db:"section_dimension_english_b"`
	SectionDimensionEnglishTW   string `json:"section_dimension_english_tw" db:"section_dimension_english_tw"`
	SectionDimensionEnglishBF   string `json:"section_dimension_english_bf" db:"section_dimension_english_bf"`
	SectionDimensionEnglishTF   string `json:"section_dimension_english_tf" db:"section_dimension_english_tf"`
	SectionDimensionEnglishTB   string `json:"section_dimension_english_tb" db:"section_dimension_english_tb"`
	SectionDimensionEnglishT    string `json:"section_dimension_english_t" db:"section_dimension_english_t"`
	SectionDimensionEnglishR    string `json:"section_dimension_english_r" db:"section_dimension_english_r"`
}

type CreateSectionDimensionEnglish struct {
	SectionDimensionEnglishName string `json:"section_dimension_english_name" db:"section_dimension_english_name"`
	SectionDimensionEnglishD    string `json:"section_dimension_english_d" db:"section_dimension_english_d"`
	SectionDimensionEnglishB    string `json:"section_dimension_english_b" db:"section_dimension_english_b"`
	SectionDimensionEnglishTW   string `json:"section_dimension_english_tw" db:"section_dimension_english_tw"`
	SectionDimensionEnglishBF   string `json:"section_dimension_english_bf" db:"section_dimension_english_bf"`
	SectionDimensionEnglishTF   string `json:"section_dimension_english_tf" db:"section_dimension_english_tf"`
	SectionDimensionEnglishTB   string `json:"section_dimension_english_tb" db:"section_dimension_english_tb"`
	SectionDimensionEnglishT    string `json:"section_dimension_english_t" db:"section_dimension_english_t"`
	SectionDimensionEnglishR    string `json:"section_dimension_english_r" db:"section_dimension_english_r"`
}
