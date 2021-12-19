package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB //声明一个全局的 db 变量

// 初始化 MySQL 函数
func initMySQL() (err error) {
	dsn := "root:root@tcp(127.0.0.1:23339)/uss_1029"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}

	// 判断是否连接成功
	err = db.Ping()
	if err != nil {
		return
	}
	return
}

//func main() {
//	connDBDemo()
//}

func connDBDemo() {
	// 初始化 MySQL
	err := initMySQL()
	if err != nil {
		panic(err)
	}
	defer db.Close()
}

/**
Go 语言操作 MySQL 之 SQLX 包 https://segmentfault.com/a/1190000023113675?utm_source=sf-similar-article
Go 语言操作 MySQL 之 CURD 操作 https://segmentfault.com/a/1190000023067651
*/
