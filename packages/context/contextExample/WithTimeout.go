package contextExample

import (
	"context"
	"fmt"
	"time"
)

func StudyWithTimeout() {
	//传递一个带超时的上下文来告诉阻塞函数它
	//应该在超时过后放弃它的工作
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	}
}
