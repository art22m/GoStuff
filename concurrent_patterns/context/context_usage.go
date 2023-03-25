package context

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func testTwoContext() {
	ctx := context.Background()
	another_process(ctx)
}

func another_process(ctx context.Context) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("Operation A: ", operationA(ctx))
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Operation B: ", operationB(ctx))
	}()

	wg.Wait()
}

func operationA(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	return serviceA(ctx)
}

func serviceA(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		case <-time.After(5 * time.Second):
			return nil
		}
	}
}

func operationB(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	return serviceB(ctx)
}

func serviceB(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		case <-time.After(time.Second):
			return nil
		}
	}
}
