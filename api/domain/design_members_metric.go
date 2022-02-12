package domain

type DesignMembersMetric struct {
	ID                            int    `json:"id" db:"id"`
	DesignMembersMetricName       string `json:"design_members_metric_name" db:"design_members_metric_name"`
	DesignMembersMetricTotalDepth string `json:"design_members_metric_total_depth" db:"design_members_metric_total_depth"`
	DesignMembersMetricWeight     string `json:"design_members_metric_weight" db:"design_members_metric_weight"`
}
