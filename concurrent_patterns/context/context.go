package context

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func testContext() {
	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		create(ctx)
	}()

	time.Sleep(time.Second)
	cancel()

	wg.Wait()
}

func create(ctx context.Context) {
	process(ctx)
}

func process(ctx context.Context) {
	set(ctx)
}

func set(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
		return

	case <-time.After(10 * time.Second):
	}
}
