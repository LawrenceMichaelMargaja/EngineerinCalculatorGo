package domain

type SectionDimensionMetric struct {
	Id int `json:"id" db:"id"`
	SectionDimensionMetricName string `json:"section_dimension_metric_name" db:"section_dimension_metric_name"`
	SectionDimensionMetricD string `json:"section_dimension_metric_d" db:"section_dimension_metric_d"`
	SectionDimensionMetricB string `json:"section_dimension_metric_b" db:"section_dimension_metric_b"`
	SectionDimensionMetricTW string `json:"section_dimension_metric_tw" db:"section_dimension_metric_tw"`
	SectionDimensionMetricBF string `json:"section_dimension_metric_bf" db:"section_dimension_metric_bf"`
	SectionDimensionMetricTF string `json:"section_dimension_metric_tf" db:"section_dimension_metric_tf"`
	SectionDimensionMetricTB string `json:"section_dimension_metric_tb" db:"section_dimension_metric_tb"`
	SectionDimensionMetricT string `json:"section_dimension_metric_t" db:"section_dimension_metric_t"`
	SectionDimensionMetricR string `json:"section_dimension_metric_r" db:"section_dimension_metric_r"`
}

type CreateSectionDimensionMetric struct {
	SectionDimensionName string `json:"section_dimension_name" db:"section_dimension_name"`
	SectionDimensionD    string `json:"section_dimension_d" db:"section_dimension_d"`
	SectionDimensionB    string `json:"section_dimension_b" db:"section_dimension_b"`
	SectionDimensionTW   string `json:"section_dimension_tw" db:"section_dimension_tw"`
	SectionDimensionBF   string `json:"section_dimension_bf" db:"section_dimension_bf"`
	SectionDimensionTF   string `json:"section_dimension_tf" db:"section_dimension_tf"`
	SectionDimensionTB   string `json:"section_dimension_tb" db:"section_dimension_tb"`
	SectionDimensionT    string `json:"section_dimension_t" db:"section_dimension_t"`
	SectionDimensionR    string `json:"section_dimension_r" db:"section_dimension_r"`
}
