package main

import (
	"fmt"
)

// 由于空接口没有方法，因此所有类型都实现了空接口
func describe(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

//func main() {
//	demo7()
//}

func demo7() {
	s := "Hello World"
	describe(s)
	i := 55
	describe(i)
	strt := struct {
		name string
	}{
		name: "Naveen R",
	}
	describe(strt)
}
