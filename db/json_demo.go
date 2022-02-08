package main

import (
	"encoding/json"
	"fmt"
	"github.com/Baz2013/learning_golang/db/model"
)

//func main() {
//	fmt.Println("json test ...")
//	jsonDemo()
//}

func jsonDemo() {
	author1 := model.Author{
		Name:     "Tom",
		Age:      22,
		Address:  "",
		PostCode: "250100",
	}

	json, err := json.Marshal(author1)
	if err != nil {
		fmt.Println("序列化失败 ")
	}

	fmt.Println(string(json))
}
