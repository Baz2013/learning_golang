package main

import (
	"fmt"
	//并不需要使用其API，只需要执行该包的init方法（加载MySQL的驱动程序）
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//参考 ： https://bbs.huaweicloud.com/blogs/detail/289053

type Person struct {
	// 对应name表字段
	Name string `db:"name"`
	// 对应age表字段
	Age int `db:"age"`
	// 对应rmb表字段
	Money float64 `db:"salary"`
}

//func main() {
//	demo_crud()
//}

func demo_crud() {
	db, _ := sqlx.Open("mysql", "root:root@tcp(localhost:23339)/test_orm")
	defer db.Close()

	insertData(db)
	//预定义Person切片用于接收查询结果
	var ps []Person
	//执行查询，得到Perosn对象的集合，丢入预定义的ps地址
	e := db.Select(&ps, "select name,age,salary from person where name like ?;", "%Tom%")
	if e != nil {
		fmt.Println("err=", e)
	}
	fmt.Println("查询成功", ps)
}

func insertData(db *sqlx.DB) {
	result, err := db.Exec("insert into person(name,age,salary) values(?,?,?);", "Tom", 21, 8888)
	if err != nil {
		fmt.Printf("插入数据错误， Error message: %s \n", err)
	} else {
		rowsAffected, _ := result.RowsAffected()
		fmt.Printf("插入数据库成功, 影响行数：%d \n", rowsAffected)
	}
}
