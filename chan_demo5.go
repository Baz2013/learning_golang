package main

import "fmt"

func main() {
	demo5()
}

func demo5() {
	squareCh := make(chan int)
	cubeCh := make(chan int)

	number := 589
	go calcSquare(number, squareCh)
	go calcCube(number, cubeCh)

	n1, n2 := <-squareCh, <-cubeCh
	fmt.Println("Sum: ", n1+n2)
}

func calcCube(number int, ch chan int) {
	digitCh := make(chan int)
	sum := 0

	go getDigit(number, digitCh)
	for v := range digitCh {
		sum += v * v
	}

	ch <- sum
}

func getDigit(number int, ch chan int) {
	for number != 0 {
		digit := number % 10
		ch <- digit
		number /= 10
	}

	close(ch) // for range 读取信道时，要关闭信道，法则会引发死锁 panic
}

func calcSquare(number int, ch chan int) {
	digitCh := make(chan int)
	sum := 0

	go getDigit(number, digitCh)
	for v := range digitCh {
		sum += v * v * v
	}

	ch <- sum
}
