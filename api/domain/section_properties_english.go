package domain

type SectionPropertiesEnglish struct {
	Id int `json:"id" db:"id"`
	SectionPropertiesEnglishName string `json:"section_properties_english_name" db:"section_properties_english_name"`
	SectionPropertiesEnglishA string `json:"section_properties_english_a" db:"section_properties_english_a"`
	SectionPropertiesEnglishJ string `json:"section_properties_english_j" db:"section_properties_english_j"`
	SectionPropertiesEnglishIxp string `json:"section_properties_english_ixp" db:"section_properties_english_ixp"`
	SectionPropertiesEnglishIyp string `json:"section_properties_english_iyp" db:"section_properties_english_iyp"`
	SectionPropertiesEnglishIw string `json:"section_properties_english_iw" db:"section_properties_english_iw"`
	SectionPropertiesEnglishSxp string `json:"section_properties_english_sxp" db:"section_properties_english_sxp"`
	SectionPropertiesEnglishSyp string `json:"section_properties_english_syp" db:"section_properties_english_syp"`
}

type CreateSectionPropertiesEnglish struct {
	SectionName string `json:"section_name" db:"section_name"`
	A           string `json:"a" db:"a"`
	J           string `json:"j" db:"j"`
	Ixp         string `json:"ixp" db:"ixp"`
	Iyp         string `json:"iyp" db:"iyp"`
	Iw          string `json:"iw" db:"iw"`
	Sxp         string `json:"sxp" db:"sxp"`
	Syp         string `json:"syp" db:"syp"`
}
