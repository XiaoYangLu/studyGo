package goBase

import "fmt"


	//结构体
	//type name struct {
	//	//属性1  type1
	//	//属性2  type2
	//	//属性3，属性4  type3
	//}
	type name1 struct {
		name string
		age int
		addr string
	}
	//聚合关系：一个类作为另一个类的属性
	type name2 struct {
		name1 interface{} //接口可接受任何数据类型
		sex string
	}
	//聚合关系；一个类作为另一个类的子类
	type name3 struct {
		uid string
		name1
	}
/*
 * 学习结构体
 */
func StudyStruct()  {


	//聚合关系，聚合关系的实现
	n1 := name1{
		name: "张三",
		age:  18,
		addr: "上海",
	}
	n2 := name2{
		name1: n1,
		sex:   "男",
	}
	n3 := name3{
		uid:   "3306",
		name1: name1{
			name: "李四",
			age:  18,
			addr: "北京",
		},
	}
	fmt.Printf("n1类型%T,n1的值：%v \n n2类型%T,n2的值：%v \n n3类型%T,n3的值：%v \n ",n1,n1,n2,n2,n3,n3)

}
