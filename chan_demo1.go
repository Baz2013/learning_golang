package main

import (
	"fmt"
)

func hello(done chan bool) {
	fmt.Println("hello gorutine")
	done <- true
	fmt.Println("data sent") // 数据未被读取的话，会一直阻塞
}

func demo() {
	ch := make(chan bool)
	go hello(ch)

	<-ch
	//time.Sleep(1 * time.Second)

	fmt.Println("Main function")
}

//func main() {
//	demo()
//}
