package main

import (
	"context"
	"fmt"
)

//func main() {
//	demo13()

//headers := make(map[string]string)
//headers["Date"] = "Tue, 21 Dec 2021 02:41:17 GMT"
//headers["AUTHORIZATION"] = "Sign f00f98d5-56a5-4047-a420-01b4aebc158a:ZDMxNmNlZDdjY2VjMDIyMDRiYjY2MDIwMWMzOGM3NDQ="
//// loop over elements of slice
//for k, v := range headers {
//	fmt.Println(k, "value is:", v)
//}
//}

func demo13() {
	// gen generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the context once
	// they are done consuming generated integers not to leak
	// the internal goroutine started by gen.
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

	fmt.Printf("%T\n", gen) // func(context.Context) <-chan int

	//for n := range gen(ctx) {
	//	fmt.Println(n)
	//	if n == 15 {
	//		break
	//	}
	//}
	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
	ch := gen(ctx) // 调用函数，其返回值是channel
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
