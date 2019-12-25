package dbconfig

import (
	"fmt"

	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

func newConnection(username, password, host string) *gorm.DB {
	DB, err := gorm.Open("mysql", username+":"+password+"@tcp("+host+":3306)/")

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Database connected")
	//defer db.Close()
	//err = DB.Ping()
	//if err != nil {
	//	panic(err.Error())
	//}
	DB.Exec("CREATE DATABASE IF NOT EXISTS NotesApp;")
	DB.Exec("use NotesApp;")
	DB.Exec("CREATE TABLE IF NOT EXISTS Notes (id VARCHAR(30) NOT NULL, title VARCHAR(30) NOT NULL, content VARCHAR(100), created_at TIMESTAMP, updated_at TIMESTAMP);")

	return DB
}
