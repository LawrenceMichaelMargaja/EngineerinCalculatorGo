package domain

type RecHsShapeMetric struct {
	Id int `json:"id" db:"id"`
	RecHsShapeMetricName string `json:"rec_hs_shape_metric_name" db:"rec_hs_shape_metric_name"`
}

type CreateRecHssShapeMetric struct {
	RecHssShapeMetricName string `json:"rec_hss_shape_metric_name" db:"rec_hss_shape_metric_name"`
}