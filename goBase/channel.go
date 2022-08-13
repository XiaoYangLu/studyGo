package goBase

import (
	"fmt"
)

func StudyChannel() {
	//1、阻塞接受数据
	// data := <- ch
	//2、完整写法
	// data, ok := <- ch
	//3、忽略接收数据
	// <- ch
	//4、循环接受数据
	// for ... range

	ch := make(chan string, 5) //缓冲通道
	//Wg.Add(2)  //此处使用 Wg 如下:
	//fatal error: all goroutines are asleep - deadlock!
	//因为main在等goroutine，而此时所有的goroutine已经结束
	//可研究35行的select中的case
	go testChannel(ch)
	for v := range ch {
		fmt.Println(v)
	}

	ch1 := make(chan int)
	go func() {
		//defer Wg.Done()
		ch1 <- 10
	}()
	//select与Switch相似，随机可通信case执行，如果所有的case都没有数据到达，
	//则执行default，如果没有default，则会阻塞，直到case接受到数据
	select {
	case data := <-ch1:
		fmt.Println("ch1的数据：", data)
	case data := <-ch1:
		fmt.Println("ch2的数据：", data)
	default: //注：case随机效果不明显，可去掉default多次尝试
		fmt.Println("执行了default...")
	}
}

func testChannel(ch chan string) {
	//defer Wg.Done()
	defer close(ch)
	for i := 0; i <= 100; i++ {
		ch <- fmt.Sprint(i)

	}
	//strings.Split()
}
