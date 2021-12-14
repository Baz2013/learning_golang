package main

import "fmt"

type Describer interface {
	Describe()
}
type Person struct {
	name string
	age  int
}

func (p Person) Describe() {
	fmt.Printf("%s is %d years old", p.name, p.age)
}

func findType2(i interface{}) {
	switch v := i.(type) {
	case Describer:
		v.Describe()
	//case string:
	//	fmt.Println("string type")
	default:
		fmt.Printf("unknown type\n")
	}
}

func demo10() {
	findType2("Naveen")
	p := Person{
		name: "Naveen R",
		age:  25,
	}
	findType2(p)
}

//func main() {
//	demo10()
//}
