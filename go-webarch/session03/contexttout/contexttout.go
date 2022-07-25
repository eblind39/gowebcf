package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancelIf := context.WithTimeout(ctx, 500*time.Microsecond)
	defer cancelIf()

	time.Sleep(100 * time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("work didn't finish")
	default:
		fmt.Println("work finished")
	}
}
