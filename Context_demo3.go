package main

import (
	"context"
	"fmt"
	"time"
)

const shortDuration = 1 * time.Millisecond // a reasonable duration to block in an example

//func main() {
//	/////////////    WithDeadline()和WithTimeout()差不多，只不过两者传的参数不一样，一个传的是截止时间，一个是时间间隔。
//	//ExampleWithDeadline()
//	ExampleWithTimeout()
//	//ExampleWithValue()
//}

// ExampleWithDeadline This example passes a context with an arbitrary deadline to tell a blocking
// function that it should abandon its work as soon as it gets to it.
func ExampleWithDeadline() {
	d := time.Now().Add(shortDuration)
	//fmt.Printf("%T \n", d)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// Even though ctx will be expired, it is good practice to call its
	// cancellation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}

	// Output:
	// context deadline exceeded
}

// ExampleWithTimeout This example passes a context with a timeout to tell a blocking function that
// it should abandon its work after the timeout elapses.
func ExampleWithTimeout() {
	// Pass a context with a timeout to tell a blocking function that it
	// should abandon its work after the timeout elapses.
	ctx, cancel := context.WithTimeout(context.Background(), shortDuration)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		d, x := ctx.Deadline()
		fmt.Println("dealine is: ", d)
		fmt.Println("deadline : ", x)
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	}

	// Output:
	// context deadline exceeded
}

// ExampleWithValue This example demonstrates how a value can be passed to the context
// and also how to retrieve it if it exists.
func ExampleWithValue() {
	type favContextKey string

	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}

	k := favContextKey("language")
	ctx := context.WithValue(context.Background(), k, "Go")

	f(ctx, k)
	f(ctx, favContextKey("color"))

	// Output:
	// found value: Go
	// key not found: color
}
