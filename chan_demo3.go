package main

import "fmt"

//func main() {
//	demo3()
//}

func demo3() {
	number := 589
	squareCh := make(chan int)
	cubeCh := make(chan int)
	go callCalculateSquares(number, squareCh)
	go callCalculateCubes(number, cubeCh)
	n1, n2 := <-squareCh, <-cubeCh

	fmt.Println("sum:", n1+n2)
}

func callCalculateCubes(number int, ch chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}

	fmt.Println("square sum:", sum)
	ch <- sum
}

func callCalculateSquares(number int, ch chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}

	fmt.Println("cube sum:", sum)
	ch <- sum
}
