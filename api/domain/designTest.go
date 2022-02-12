package domain

//type steelArgumentsDesignDaoInterface interface {
//	//GetAdditionalDatabaseValues(c *SteelArgumentsDesignResults) (*SteelArgumentsDesignResults, error)
//	GetTest(c *CreateSteelDesignArguments) (*OverallResults, error)
//	SetClient()
//}
//
//type steelArgumentsDesignDao struct {
//	client *sqlx.DB
//}
//
//type BeamAnalysisDesignResults struct {
//	Pt       float64 `json:"pt"`
//	Pc       float64 `json:"pc"`
//	Mcx      float64 `json:"mcx"`
//	Mcy      float64 `json:"mcy"`
//	Vcx      float64 `json:"vcx"`
//	Vcy      float64 `json:"vcy"`
//	Pratio   float64 `json:"pratio"`
//	MxRatio  float64 `json:"mx_ratio"`
//	MyRatio  float64 `json:"my_ratio"`
//	VxRatio  float64 `json:"vx_ratio"`
//	VyRatio  float64 `json:"vy_ratio"`
//	Combined float64 `json:"combined"`
//	KLr      float64 `json:"k_lr"`
//}

//type AdditionalResults struct {
//	id                  int     `json:"id" db:"id"`
//	name                string  `json:"name" db:"name"`
//	overallDepth        float64 `json:"overall_depth" db:"overall_depth"`
//	weight              float64 `json:"weight" db:"weight"`
//	criticalDesignRatio float64 `json:"critical_design_ratio" db:"critical_design_ratio"`
//	Pt                  float64 `json:"pt"`
//	Pc                  float64 `json:"pc"`
//	Mcx                 float64 `json:"mcx"`
//	Mcy                 float64 `json:"mcy"`
//	Vcx                 float64 `json:"vcx"`
//	Vcy                 float64 `json:"vcy"`
//	Pratio              float64 `json:"pratio"`
//	MxRatio             float64 `json:"mx_ratio"`
//	MyRatio             float64 `json:"my_ratio"`
//	VxRatio             float64 `json:"vx_ratio"`
//	VyRatio             float64 `json:"vy_ratio"`
//	Combined            float64 `json:"combined"`
//	KLr                 float64 `json:"k_lr"`
//}
//
//func (c *BeamAnalysisDesignResults) DisplayDesignData() {
//	fmt.Println("c.Pt", c.Pt)
//	fmt.Println("c.Pc", c.Pc)
//	fmt.Println("c.Mcx", c.Mcx)
//	fmt.Println("c.Mcy", c.Mcy)
//	fmt.Println("c.Vcx", c.Vcx)
//	fmt.Println("c.Vcy", c.Vcy)
//	fmt.Println("c.Pratio", c.Pratio)
//	fmt.Println("c.MxRatio", c.MxRatio)
//	fmt.Println("c.MyRatio", c.MyRatio)
//	fmt.Println("c.VxRatio", c.VxRatio)
//	fmt.Println("c.VyRatio", c.VyRatio)
//	fmt.Println("c.Combined", c.Combined)
//	fmt.Println("c.KLr", c.KLr)
//}
//
//var (
//	SteelArgumentsDesignDao steelArgumentsDesignDaoInterface
//)
//
//func init() {
//	SteelArgumentsDesignDao = &steelArgumentsDesignDao{}
//	SteelArgumentsDesignDao.SetClient()
//}
//
//func (s *steelArgumentsDesignDao) SetClient() {
//	s.client = utils.GetMYSQLConnection()
//}


