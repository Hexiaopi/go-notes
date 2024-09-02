package context

import (
	"context"
	"fmt"
	"time"
)

func ExampleContextTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("timeout")
			return
		case <-time.After(3 * time.Second):
			fmt.Println("3s done")
			return
		}
	}
	// Output:
	// timeout
}

func ExampleContextValue() {
	ctx := context.WithValue(context.Background(), "key", "value")
	// 定义一个函数，接收一个 context.Context 类型的参数
	f := func(ctx context.Context) {
		fmt.Println(ctx.Value("key"))
	}
	f(ctx)
	// Output:
	// value
}

func ExampleContextCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("done")
				return
			default:
				fmt.Println("running")
				time.Sleep(1 * time.Second)
			}
		}
	}(ctx)
	time.Sleep(3 * time.Second)
	// 主动取消
	cancel()
	time.Sleep(1 * time.Second)
	// Output:
	// running
	// running
	// running
	// done
}

func ExampleContextAfterFunc() {
	ctx, cancel := context.WithCancel(context.Background())
	stop := context.AfterFunc(ctx, func() {
		fmt.Println("stop")
	})
	fmt.Println("cancel")
	cancel()
	if stop() {
		fmt.Println("done")
	}
	time.Sleep(1 * time.Second)
	// Output:
	// cancel
	// stop
}
