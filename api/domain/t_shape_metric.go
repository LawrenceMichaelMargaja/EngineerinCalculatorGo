package domain

type TShapeMetric struct {
	Id int `json:"id" db:"id"`
	TShapeMetricName string `json:"t_shape_metric_name" db:"t_shape_metric_name"`
}

type CreateTShapeMetric struct {
	TShapeMetricName string `json:"t_shape_metric_name" db:"t_shape_metric_name"`
}