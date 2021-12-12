package main

import (
	"fmt"
	"time"
)

//func main() {
//demo2()
//}

func demo2() {
	done := make(chan bool)
	fmt.Println("Main going to call hello go goroutine")
	go hello2(done)
	<-done
	fmt.Println("Main received data")
}

func hello2(done chan bool) {
	fmt.Println("hello go routine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("hello go routine awake and going to write to done")
	done <- true
}
