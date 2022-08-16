package contextExample

import (
	"context"
	"fmt"
)

func StudyWithCancel() {
	//gen在一个单独的gor例程中生成整数
	//发送它们到返回的通道。
	//gen的调用者需要取消上下文一次
	//它们使用生成的整数不泄漏
	//由gen. exe启动的内部goroutine。
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return // returning not to leak the goroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}
