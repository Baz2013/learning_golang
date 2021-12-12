package main

import (
	"time"
)

/**
func main() {

	ch := make(chan string)
	go process(ch)

	for  {
		time.Sleep(1000*time.Millisecond)
		select {
		case s1 := <-ch:
			fmt.Println("received value", s1)
			return
		default:
			fmt.Println("no value received")
		}
	}

}

*/

func process(ch chan string) {
	time.Sleep(10500 * time.Millisecond)
	ch <- "process successful"
}
