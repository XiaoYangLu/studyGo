package main

import (
	"fmt"
	"studyGo/goBase"
)


func main() {
	//数据类型
	goBase.StudyDataType()
	fmt.Println("---------------------")

	//结构体
	goBase.StudyStruct()
	fmt.Println("---------------------")

	//interface
	p1 := goBase.Works{"搬砖",10,14}
	p2 := goBase.Works{"卖报纸",10,13.5}
	p3 := goBase.BrushVideo{"抖音极速版",2,5}
	ic := []goBase.Income{p1,p2,p3}
	fmt.Println(goBase.CalIncome(ic))
	fmt.Println("---------------------")

	//goroutime
	goBase.StudyGoroutine()
	//fmt.Println("---------------------")

	//channel
	goBase.StudyChannel()


	goBase.Wg.Wait()
}
