package main

import "fmt"

type Describer2 interface {
	Describe2()
}
type Person2 struct {
	name string
	age  int
}

func (p Person2) Describe2() { // 使用值接受者实现
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

type Address struct {
	state   string
	country string
}

func (a *Address) Describe2() { // 使用指针接受者实现
	fmt.Printf("State %s Country %s", a.state, a.country)
}

func demo11() {
	var d1 Describer2
	p1 := Person2{"Sam", 25}
	d1 = p1
	d1.Describe2()
	p2 := Person2{"James", 32}
	d1 = &p2
	d1.Describe2()

	var d2 Describer2
	a := Address{"Washington", "USA"}

	/* 如果下面一行取消注释会导致编译错误：
	   cannot use a (type Address) as type Describer
	   in assignment: Address does not implement
	   Describer (Describe method has pointer
	   receiver)
	*/
	//d2 = a

	d2 = &a // 这是合法的
	// 因为在第 22 行，Address 类型的指针实现了 Describer 接口
	d2.Describe2()
}

func main() {
	demo11()

}
