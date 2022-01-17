package domain

type CShapeMetric struct {
	Id int `json:"id" db:"id"`
	CShapeMetricName string `json:"c_shape_metric_name" db:"c_shape_metric_name"`
}

type CreateCShapeMetric struct {
	CShapeMetricName string `json:"c_shape_metric_name" db:"c_shape_metric_name"`
}
