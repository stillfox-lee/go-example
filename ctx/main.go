package main

import (
	"context"
	"fmt"
	"time"
)

func todoSomething(ctx context.Context) {
	fmt.Println("todoSomething")
}

func backgroundSomething(ctx context.Context) {
	fmt.Println("backgroundSomething")
}

func valueSomething(ctx context.Context) {
	fmt.Printf("valueSomething: %v\n", ctx.Value("key"))
}

func nestedValueSomething(ctx context.Context) {
	fmt.Printf("originValue: %v\n", ctx.Value("key"))
	nestedCtx := context.WithValue(ctx, "key", "nestedValue")
	valueSomething(nestedCtx)
	fmt.Printf("originValue: %v\n", ctx.Value("key"))
}

func cancelSomething(ctx context.Context) {
	ctx, cancelFn := context.WithCancel(ctx)
	printCh := make(chan int)
	go printWithCancel(ctx, printCh)
	for num := 1; num <= 3; num++ {
		printCh <- num
	}

	cancelFn()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("cancelSomething done")
}

func printWithCancel(ctx context.Context, printCh <-chan int) {
	for {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				fmt.Printf("printWithCancel error: %v\n", err)
			} else {
				fmt.Println("printWithCancel done")
			}
			return
		case num := <-printCh:
			fmt.Printf("print number: %d\n", num)
		}
	}
}

func main() {
	todoCtx := context.TODO()
	todoSomething(todoCtx)
	backgroudCtx := context.Background()
	backgroundSomething(backgroudCtx)

	// base case for context.WithValue
	valueCtx := context.WithValue(backgroudCtx, "key", "value")
	valueSomething(valueCtx)

	// nested case for context.WithValue
	originValue := context.WithValue(backgroudCtx, "key", "originValue")
	nestedValueSomething(originValue)

	// cancel case for context.WithCancel
	cancelSomething(backgroudCtx)

}
