package goBase



//接口
	//type name interface {
	//	//方法体
	//}
	//type  name struct {
	//	//结构体
	//}
	//实现接口
	//func （变量名 结构体类型） 方法（[参数列表]） [返回参数类别] {
	//	//方法体
	//}

	
//多态	
type Income interface {
	source() string //来源
	calculatetotal() float64 //总收入
}

//工作收入
type Works struct {
	Work string
	WorkHours int //工作时间
	WageHours float64 //每小时收入
}

func (e Works) source() string  {
	return e.Work
}
func (e Works) calculatetotal() float64  {
	return float64(e.WorkHours) * e.WageHours
}
//刷视频
type BrushVideo struct {
	Video string
	Frequency int //次数
	LookVideo float64 //看多少
}

func (b BrushVideo) source() string {
	return b.Video
}
func (b BrushVideo) calculatetotal() float64  {
	return float64(b.Frequency) * b.LookVideo
}

//计算总收入
func CalIncome(ic []Income) float64  {
	var netIncome float64
	for _, v := range ic {
		netIncome += v.calculatetotal()
	}
	return netIncome
}




