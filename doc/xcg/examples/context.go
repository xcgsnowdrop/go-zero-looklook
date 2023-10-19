package main

import (
	"context"
	"fmt"
	"time"
)

func doWork(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Work canceled")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	parentContext := context.Background()
	ctx, cancel := context.WithCancel(parentContext)

	go doWork(ctx)

	time.Sleep(3 * time.Second)
	cancel() // 取消工作

	time.Sleep(1 * time.Second)
	fmt.Println("Main function completed")
}
