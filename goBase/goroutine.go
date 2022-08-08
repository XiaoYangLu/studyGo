package goBase

import (
	"fmt"
	"sync"
)

var Wg sync.WaitGroup  //定义全局变量 wg

func StudyGoroutine()  {
	//WaitGroup：它能够一直等到所有的goroutine执行完成，并且阻塞主线程的执行，直到所有的goroutine执行完成。
	//WaitGroup总共有三个方法：Add(delta int),Done(),Wait()。
		//	Add:添加或者减少等待goroutine的数量
		//	Done:相当于Add(-1)
		//	Wait:执行阻塞，直到所有的WaitGroup数量变成0


	//串行：当前任务执行完成，才能执行下一个任务
	//并行；每个任务分配给每个处理器独立运行，即多个任务时运行
	//并发：同时处理多个任务，微观层面，不是同时执行
	//进程：即某个程序运行时产生
	//线程：轻量级进程，通常一个进程包含多个线程
	//协程：微线程，一个线程可以拥有多个协程，编译器级别
	//goroutine并发执行，coroutine顺序执行
	//goroutine可用关键字go创建，其没有返回值，可用goroutine之间的通信通信机制channel
	Wg.Add(3)
	go testGoroutine()
	//多个goroutine时，可用runtime.Gosched()交出控制权


	//线程池中的线程与CPU核心数量的对应关系，可用runtime.COMAXPROCS()
	//runtime.COMAXPROCS(逻辑CPU数量)
	//  <1    不修改任何数值
	//  =1     单核执行
	//  >1    多核并发执行

	go testGoroutine1()
	go testGoroutine2()


	//两个任务只有一个逻辑核心，此时是做完一个任务再做另一个任务。
	//将逻辑核心数设为2，此时两个任务并行执行
	//runtime.GOMAXPROCS(1)
	//runtime.GOMAXPROCS(2)
	//go testGoroutine1()
	//go testGoroutine2()

}

func testGoroutine() {
	defer Wg.Done()
	fmt.Println("hello,我是goroutine")
}

//多个goroutime
func testGoroutine1()  {
	defer Wg.Done()
	for i := 0;i <= 3;i++{
		fmt.Println("test1:",i)
	}
}
func testGoroutine2()  {
	defer Wg.Done()
	for i := 0;i <= 3;i++{
		fmt.Printf("test2: %c \n",i) //%c格式化打印出字符（ACCIS）
	}
}