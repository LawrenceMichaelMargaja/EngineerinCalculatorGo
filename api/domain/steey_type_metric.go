package domain

type SteelTypeMetric struct {
	Id int `json:"id" db:"id"`
	SteelTypeMetricName string `json:"steel_type_metric_name" db:"steel_type_metric_name"`
	SteelTypeMetricE string `json:"steel_type_metric_e" db:"steel_type_metric_e"`
	SteelTypeMetricFy string `json:"steel_type_metric_fy" db:"steel_type_metric_fy"`
	SteelTypeMetricFu string `json:"steel_type_metric_fu" db:"steel_type_metric_fu"`
}

type CreateSteelTypeMetric struct {
	SteelTypeMetricName string `json:"steel_type_metric_name" db:"steel_type_metric_name"`
	SteelTypeMetricE    string `json:"steel_type_metric_e" db:"steel_type_metric_e"`
	SteelTypeMetricFy   string `json:"steel_type_metric_fy" db:"steel_type_metric_fy"`
	SteelTypeMetricFu   string `json:"steel_type_metric_fu" db:"steel_type_metric_fu"`
}
