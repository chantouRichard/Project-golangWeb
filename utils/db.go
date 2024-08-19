package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("mysql", "root:123456@tcp(localhost:3306)/theater")
	if err != nil {
		log.Fatal(err)
	}

	// 检查数据库连接
	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Database connected!")
}
