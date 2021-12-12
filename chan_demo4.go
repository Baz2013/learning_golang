package main

import "fmt"

//func main() {
//	demo4()
//}

func demo4() {
	ch := make(chan int)
	go producer(ch)
	//for {
	//	i, ok := <- ch
	//	if ok == false {
	//		break
	//	}
	//
	//	fmt.Println("Received: ", i , ok)
	//}

	for v := range ch {
		fmt.Println("Received:", v)
	}
}

func producer(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}
