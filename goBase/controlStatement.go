package goBase

import "fmt"

func StudyIfFor() {
	if true {
		fmt.Println("我是if")
	} else if true {
		fmt.Println("我是if else")
	}
	for i := 0; i <= 10; i++ {
		fmt.Printf("hello")
	}
	var str string
	switch str {
	case "":
		fmt.Println("")
	case "hi":
		fmt.Println("hi")
	case "hello":
		fmt.Println("hello")
	}
	arr := [6]int{1, 2, 3, 4, 5, 6}
	//for range 对arr、slice、string返回索引和值
	//          map返回键和值
	//			channel只返回通道内的值
	for in, v := range arr {
		fmt.Printf("arr[%v] 对应的值：%v", in, v)
	}

}
