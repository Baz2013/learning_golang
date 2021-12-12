package main

import (
	"fmt"
	"sync"
	"time"
)

//func main() {
//	demo6()
//}

func demo6() {
	no := 3
	var wg sync.WaitGroup

	for i := 0; i < no; i++ {
		wg.Add(1)
		// 注意！！！！ 如果没有传递 wg 的地址，那么每个 Go 协程将会得到一个 WaitGroup 值的拷贝，因而当它们执行结束时，main 函数并不会知道
		go process1(i, &wg)
	}

	wg.Wait()
}

func process1(i int, s *sync.WaitGroup) {
	fmt.Println("Started gorutine", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	s.Done()
}
