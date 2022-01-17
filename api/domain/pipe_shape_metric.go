package domain

type PipeShapeMetric struct {
	Id int `json:"id" db:"id"`
	PipeShapeMetricName string `json:"pipe_shape_metric_name" db:"pipe_shape_metric_name"`
}

type CreatePipeShapeMetric struct {
	PipeShapeMetricName string `json:"pipe_shape_metric_name" db:"pipe_shape_metric_name"`
}
