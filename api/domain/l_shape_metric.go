package domain

type LShapeMetric struct {
	Id int `json:"id" db:"id"`
	LShapeMetricName string `json:"l_shape_metric_name" db:"l_shape_metric_name"`
}

type CreateLShapeMetric struct {
	LShapeMetricName string `json:"l_shape_metric_name" db:"l_shape_metric_name"`
}