package domain

type TwoLShapeMetric struct {
	Id int `json:"id" db:"id"`
	TwoLShapeMetricName string `json:"two_l_shape_metric_name" db:"two_l_shape_metric_name"`
}

type CreateTwoLShapeMetric struct {
	TwoLShapeMetricName string `json:"two_l_shape_metric_name" db:"two_l_shape_metric_name"`
}
