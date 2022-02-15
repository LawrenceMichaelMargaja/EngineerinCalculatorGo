package domain

type Shapes struct {
	Id int `json:"id" db:"id"`
	ShapeName string `json:"shape_name" db:"shape_name"`
}

//type CreateShapes struct {
//	ShapeName string `json:"shape_name" db:"shape_name"`
//	DbValue   string `json:"db_value" db:"db_value"`
//}
