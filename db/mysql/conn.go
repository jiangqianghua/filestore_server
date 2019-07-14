package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var db *sql.DB

// 数据库初始化
func init() {
	fmt.Println("connect init...")
	db, _ = sql.Open("mysql", "root:Jiang123!@tcp(180.76.105.202:3306)/filestore?charset=utf8")
	db.SetMaxOpenConns(1000)
	err := db.Ping()
	if err != nil {
		fmt.Println("Failed to connect to mysql, err:" + err.Error())
		os.Exit(1)
	}
}

func DBConn() *sql.DB {
	return db
}
