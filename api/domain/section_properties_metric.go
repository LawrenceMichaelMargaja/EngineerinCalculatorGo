package domain

type SectionPropertiesMetric struct {
	Id int `json:"id" db:"id"`
	SectionPropertiesMetricName string `json:"section_properties_metric_name" db:"section_properties_metric_name"`
	SectionPropertiesMetricA string `json:"section_properties_metric_a" db:"section_properties_metric_a"`
	SectionPropertiesMetricJ string `json:"section_properties_metric_j" db:"section_properties_metric_j"`
	SectionPropertiesMetricIxp string `json:"section_properties_metric_ixp" db:"section_properties_metric_ixp"`
	SectionPropertiesMetricIyp string `json:"section_properties_metric_iyp" db:"section_properties_metric_iyp"`
	SectionPropertiesMetricIw string `json:"section_properties_metric_iw" db:"section_properties_metric_iw"`
	SectionPropertiesMetricSxp string `json:"section_properties_metric_sxp" db:"section_properties_metric_sxp"`
	SectionPropertiesMetricSyp string `json:"section_properties_metric_syp" db:"section_properties_metric_syp"`
}

type CreateSectionPropertiesMetric struct {
	SectionName string `json:"section_name" db:"section_name"`
	A           string `json:"a" db:"a"`
	J           string `json:"j" db:"j"`
	Ixp         string `json:"ixp"db:"ixp"'`
	Iyp         string `json:"iyp" db:"iyp"`
	Iw          string `json:"iw" db:"iw"`
	Sxp         string `json:"sxp" db:"sxp"`
	Syp         string `json:"syp" db:"syp"`
}
