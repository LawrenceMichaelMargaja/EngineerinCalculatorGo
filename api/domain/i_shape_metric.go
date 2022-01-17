package domain

type IShapeMetric struct {
	Id int `json:"id" db:"id"`
	IShapeMetricName string `json:"i_shape_metric_name" db:"i_shape_metric_name"`
}

type CreateIShapeMetric struct {
	IShapeMetricName string `json:"i_shape_metric_name" db:"i_shape_metric_name"`
}
