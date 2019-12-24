package dbconfig

import (
	"database/sql"
	"fmt"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func NewConnection(username, password, host string) {
	DB, err := sql.Open("mysql", username+":"+password+"@tcp("+host+":3306)/NotesApp")

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Database connected")
	//defer db.Close()
	err = DB.Ping()
	if err != nil {
		panic(err.Error())
	}

	_, err = DB.Exec("CREATE DATABASE IF NOT EXISTS NotesApp;")
	if err != nil {
		panic(err)
	}

	_, err = DB.Exec("use NotesApp;")
	if err != nil {
		panic(err)
	}

	_, err = DB.Exec("create table IF NOT EXISTS Notes (Id VARCHAR(30) NOT NULL, Title VARCHAR(30) NOT NULL, Content VARCHAR(100), CreatedAt TIMESTAMP, UpdatedAt TIMESTAMP);")
	if err != nil {
		panic(err)
	}
}
